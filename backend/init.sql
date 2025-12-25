-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enum Types (Idempotent)
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vehicle_type') THEN
        CREATE TYPE vehicle_type AS ENUM ('MECH', 'TANK', 'SHIP');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vehicle_class') THEN
        CREATE TYPE vehicle_class AS ENUM ('STRIKER', 'GUARDIAN', 'SCOUT', 'ARTILLERY');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'rarity_tier') THEN
        CREATE TYPE rarity_tier AS ENUM ('COMMON', 'RARE', 'LEGENDARY', 'REFINED', 'PROTOTYPE', 'RELIC', 'SINGULARITY');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'saga_status') THEN
        CREATE TYPE saga_status AS ENUM ('STARTED', 'COMPLETED', 'FAILED', 'COMPENSATING');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vehicle_status') THEN
        CREATE TYPE vehicle_status AS ENUM ('PENDING', 'MINTED', 'BURNED');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'item_type') THEN
        CREATE TYPE item_type AS ENUM ('VEHICLE', 'PART', 'BASTION_MODULE', 'CONSUMABLE');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'item_condition') THEN
        CREATE TYPE item_condition AS ENUM ('PRISTINE', 'WORN', 'DAMAGED', 'CRITICAL', 'BROKEN');
    END IF;
END $$;

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    privy_did VARCHAR(255) UNIQUE,
    wallet_address VARCHAR(42) UNIQUE,
    username VARCHAR(50),
    email VARCHAR(255) UNIQUE,
    password_hash TEXT,
    guest_id VARCHAR(255) UNIQUE,
    auth_type VARCHAR(20) DEFAULT 'SOCIAL',
    credits DECIMAL(18, 8) DEFAULT 0,
    active_character_id UUID, -- Set after character creation
    last_login TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Characters Table
CREATE TABLE IF NOT EXISTS characters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    gender VARCHAR(20),
    face_index INTEGER DEFAULT 0,
    hair_index INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Add circular reference back to users
ALTER TABLE users ADD CONSTRAINT fk_active_character FOREIGN KEY (active_character_id) REFERENCES characters(id);

-- Vehicles (NFTs) Table
CREATE TABLE IF NOT EXISTS vehicles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    token_id NUMERIC(78, 0) UNIQUE, -- Supports uint256
    owner_id UUID REFERENCES users(id),
    character_id UUID REFERENCES characters(id),
    vehicle_type vehicle_type NOT NULL,
    class vehicle_class NOT NULL,
    image_url TEXT,
    stats JSONB NOT NULL DEFAULT '{}', -- Base stats: HP, Attack, Defense, etc.
    cr INTEGER DEFAULT 0, -- Combat Rating
    suitability_tags JSONB NOT NULL DEFAULT '[]', -- Terrain suitability (e.g., ["AERIAL", "AQUATIC"])
    rarity rarity_tier NOT NULL DEFAULT 'COMMON',
    tier INTEGER DEFAULT 1,
    is_void_touched BOOLEAN DEFAULT FALSE,
    season VARCHAR(50),
    status vehicle_status DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Unified Items Table (Deep Durability System & Stage Change)
CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID REFERENCES users(id),
    character_id UUID REFERENCES characters(id),
    
    -- Identity
    name VARCHAR(100) NOT NULL,
    item_type item_type NOT NULL,
    rarity rarity_tier NOT NULL DEFAULT 'COMMON',
    tier INTEGER DEFAULT 1,
    slot VARCHAR(50), -- e.g., 'ARM_L', 'WARP_DRIVE', 'HEAD'
    
    -- Stage Change (NFT) Support
    is_nft BOOLEAN DEFAULT FALSE,
    token_id NUMERIC(78, 0) UNIQUE,
    
    -- Deep Durability System (DDS)
    durability INTEGER NOT NULL DEFAULT 1000,
    max_durability INTEGER NOT NULL DEFAULT 1000,
    condition item_condition NOT NULL DEFAULT 'PRISTINE',
    
    -- Dynamic Data
    stats JSONB NOT NULL DEFAULT '{}', -- HP, ATK, DEF, Energy, etc.
    visual_dna JSONB NOT NULL DEFAULT '{}', -- AI Keywords, Glitch Intensity, Smoke Level
    metadata JSONB NOT NULL DEFAULT '{}', -- Season, Crafting Info, Origin
    
    -- State
    is_equipped BOOLEAN DEFAULT FALSE,
    parent_item_id UUID REFERENCES items(id), -- For parts attached to a Vehicle or Bastion
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Parts (Legacy/Modular Items) Table - Keeping for compatibility or specific part logic
CREATE TABLE IF NOT EXISTS parts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID REFERENCES users(id),
    vehicle_id UUID REFERENCES vehicles(id), -- NULL if not equipped
    character_id UUID REFERENCES characters(id), -- For Pilot Gear (NULL if Vehicle part)
    slot VARCHAR(50) NOT NULL, -- CORE, ARMOR, WEAPON, T_MODULE (Vehicle) | HEAD, BODY, UTILITY (Pilot)
    name VARCHAR(100) NOT NULL,
    rarity rarity_tier NOT NULL DEFAULT 'COMMON',
    tier INTEGER DEFAULT 1,
    is_void_touched BOOLEAN DEFAULT FALSE,
    is_minted BOOLEAN DEFAULT FALSE,
    stats JSONB NOT NULL DEFAULT '{}', -- Bonus stats: +HP, +Crit, etc.
    visual_dna JSONB NOT NULL DEFAULT '{}', -- AI Keywords for FLUX.1
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Pilot Stats (Neural Resonance & Resources)
CREATE TABLE IF NOT EXISTS pilot_stats (
    user_id UUID REFERENCES users(id),
    character_id UUID PRIMARY KEY REFERENCES characters(id),
    resonance_level INTEGER DEFAULT 0,
    resonance_exp INTEGER DEFAULT 0,
    xp INTEGER DEFAULT 0,
    rank INTEGER DEFAULT 1,
    current_o2 DECIMAL(5, 2) DEFAULT 100.00,
    current_fuel DECIMAL(5, 2) DEFAULT 100.00,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Stars (Universe) Table
CREATE TABLE IF NOT EXISTS stars (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    difficulty_level INTEGER DEFAULT 1,
    discovered_by UUID REFERENCES users(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Nodes (Encounters on an Expedition)
CREATE TYPE node_type AS ENUM ('STANDARD', 'RESOURCE', 'COMBAT', 'ANOMALY', 'OUTPOST', 'NARRATIVE', 'REST', 'BOSS');

-- Universe Map Structure
CREATE TABLE IF NOT EXISTS sectors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    difficulty VARCHAR(20), -- LOW, MEDIUM, HIGH, EXTREME
    coordinates_x INTEGER,
    coordinates_y INTEGER,
    color VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sub_sectors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sector_id UUID REFERENCES sectors(id),
    type VARCHAR(20), -- PLANET, STATION, WRECK, ANOMALY
    name VARCHAR(100) NOT NULL,
    description TEXT,
    rewards TEXT[],
    requirements TEXT[],
    allowed_modes TEXT[], -- PILOT, VEHICLE, TANK, etc.
    requires_atmosphere BOOLEAN DEFAULT FALSE,
    suitability_pilot INTEGER DEFAULT 50,
    suitability_vehicle INTEGER DEFAULT 50,
    coordinates_x INTEGER,
    coordinates_y INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS planet_locations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sub_sector_id UUID REFERENCES sub_sectors(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    rewards TEXT[],
    requirements TEXT[],
    allowed_modes TEXT[],
    requires_atmosphere BOOLEAN DEFAULT FALSE,
    suitability_pilot INTEGER DEFAULT 50,
    suitability_vehicle INTEGER DEFAULT 50,
    coordinates_x INTEGER,
    coordinates_y INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Expeditions & Encounters (Exploration Timeline)
CREATE TABLE IF NOT EXISTS expeditions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    sub_sector_id UUID REFERENCES sub_sectors(id),
    planet_location_id UUID REFERENCES planet_locations(id),
    vehicle_id UUID REFERENCES vehicles(id),
    title VARCHAR(200),
    description TEXT,
    goal TEXT,
    status VARCHAR(20) DEFAULT 'ACTIVE', -- ACTIVE, COMPLETED, FAILED
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS encounters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    expedition_id UUID REFERENCES expeditions(id),
    type VARCHAR(20), -- COMBAT, RESOURCE, NARRATIVE, ANCHOR
    title VARCHAR(200),
    description TEXT,
    visual_prompt TEXT,
    image_url TEXT,
    enemy_id UUID, -- References vehicles(id) for NPC enemies
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS nodes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    star_id UUID REFERENCES stars(id),
    name VARCHAR(100) NOT NULL,
    type node_type NOT NULL,
    environment_description TEXT, -- e.g., "Acid Rain Desert", "Neon Slums"
    difficulty_multiplier DECIMAL(3, 2) DEFAULT 1.0,
    position_index INTEGER NOT NULL, -- Order on the "Expedition"
    metadata JSONB DEFAULT '{}', -- Enemy types, loot tables, etc.
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Exploration Sessions
CREATE TABLE IF NOT EXISTS exploration_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    vehicle_id UUID REFERENCES vehicles(id),
    current_node_id UUID REFERENCES nodes(id),
    status VARCHAR(20) DEFAULT 'ACTIVE', -- ACTIVE, COMPLETED, FAILED
    logs JSONB DEFAULT '[]',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Saga Transactions (Orchestration Log)
CREATE TABLE IF NOT EXISTS saga_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    type VARCHAR(50) NOT NULL, -- VEHICLE_ASSEMBLY, STAR_DISCOVERY, etc.
    current_step VARCHAR(100),
    status saga_status DEFAULT 'STARTED',
    payload JSONB DEFAULT '{}',
    idempotency_key VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Combat Logs & Battle Records
CREATE TABLE IF NOT EXISTS combat_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    mission_id UUID REFERENCES saga_transactions(id),
    battle_data JSONB NOT NULL,
    image_url TEXT,
    is_permanent BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP WITH TIME ZONE,
    saved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Gacha Pity System
CREATE TABLE IF NOT EXISTS gacha_stats (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    pity_relic_count INTEGER DEFAULT 0,
    pity_singularity_count INTEGER DEFAULT 0,
    total_pulls INTEGER DEFAULT 0,
    last_free_pull_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Gacha History
CREATE TABLE IF NOT EXISTS gacha_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    item_id UUID, -- Can be Vehicle or Part ID
    item_type TEXT NOT NULL, -- 'VEHICLE', 'PART'
    pull_type TEXT NOT NULL, -- 'STANDARD_SIGNAL', 'VOID_SIGNAL'
    rarity rarity_tier NOT NULL,
    seed BIGINT,
    pity_relic_before INTEGER,
    pity_relic_after INTEGER,
    pity_singularity_before INTEGER,
    pity_singularity_after INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Mothership Upgrades table
CREATE TABLE IF NOT EXISTS mothership_upgrades (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    node_id VARCHAR(100) NOT NULL, -- e.g., 'SIGNAL_BOOSTER_1'
    path VARCHAR(50) NOT NULL, -- 'TELEPORT' or 'ENTRY'
    unlocked_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, node_id)
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_vehicles_owner ON vehicles(owner_id);
CREATE INDEX IF NOT EXISTS idx_saga_user ON saga_transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_saga_idempotency ON saga_transactions(idempotency_key);
CREATE INDEX IF NOT EXISTS idx_combat_user ON combat_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_combat_expires ON combat_logs(expires_at) WHERE is_permanent = FALSE;
CREATE INDEX IF NOT EXISTS idx_items_owner ON items(owner_id);
CREATE INDEX IF NOT EXISTS idx_items_character ON items(character_id);
CREATE INDEX IF NOT EXISTS idx_items_type ON items(item_type);
CREATE INDEX IF NOT EXISTS idx_items_parent ON items(parent_item_id);
CREATE INDEX IF NOT EXISTS idx_characters_user_id ON characters(user_id);

-- Triggers for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_items_updated_at BEFORE UPDATE ON items FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_characters_updated_at BEFORE UPDATE ON characters FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_exploration_sessions_updated_at BEFORE UPDATE ON exploration_sessions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

--------------------------------------------------------------------------------
-- INITIAL SEED DATA (Single Source of Truth)
--------------------------------------------------------------------------------

-- 1. SYSTEM_NPC User
INSERT INTO users (id, wallet_address, username)
VALUES ('00000000-0000-0000-0000-000000000000', '0x0000000000000000000000000000000000000000', 'SYSTEM_NPC')
ON CONFLICT (id) DO NOTHING;

-- 2. NPC Vehicles (Enemies)
INSERT INTO vehicles (id, owner_id, vehicle_type, class, rarity, stats, status, cr)
VALUES 
('e0000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000000', 'MECH', 'STRIKER', 'COMMON', '{"hp": 80, "attack": 15, "defense": 5, "speed": 60}', 'MINTED', 150),
('e0000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000000', 'TANK', 'GUARDIAN', 'RARE', '{"hp": 200, "attack": 10, "defense": 20, "speed": 30}', 'MINTED', 300),
('e0000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000000', 'SHIP', 'SCOUT', 'COMMON', '{"hp": 60, "attack": 12, "defense": 3, "speed": 100}', 'MINTED', 120)
ON CONFLICT (id) DO NOTHING;

-- 3. Sample Universe Data (SOL GATE)
INSERT INTO sectors (id, name, description, difficulty, coordinates_x, coordinates_y, color) 
VALUES ('s0000000-0000-0000-0000-000000000001', 'SOL GATE', 'The industrial gateway to the system. Relatively safe but heavily monitored.', 'LOW', 15, 20, 'blue')
ON CONFLICT (id) DO NOTHING;

INSERT INTO sub_sectors (id, sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
VALUES ('ss000000-0000-0000-0000-000000000001', 's0000000-0000-0000-0000-000000000001', 'STATION', 'Outpost 01', 'A standard refueling station for independent scavengers.', '{"Scrap Metal", "Fuel Isotopes"}', '{}', '{"PILOT", "VEHICLE"}', FALSE, 100, 20, 40, 30)
ON CONFLICT (id) DO NOTHING;

-- 4. Sample Narrative Expedition
INSERT INTO expeditions (id, title, description, goal) 
VALUES ('00000000-0000-0000-0000-000000000001', 'The Silent Signal', 'A mysterious signal is emanating from the Iron Nebula.', 'Locate the source of the signal and decrypt it.')
ON CONFLICT (id) DO NOTHING;

INSERT INTO encounters (id, expedition_id, type, title, description, visual_prompt) 
VALUES ('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', 'NARRATIVE', 'The Awakening', 'You wake up in the cold silence of your cockpit. The radar is flickering.', 'Tactical Noir style, a pilot waking up in a dark cockpit, flickering neon lights, high detail')
ON CONFLICT (id) DO NOTHING;
