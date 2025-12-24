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
        CREATE TYPE rarity_tier AS ENUM ('COMMON', 'RARE', 'LEGENDARY');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'saga_status') THEN
        CREATE TYPE saga_status AS ENUM ('STARTED', 'COMPLETED', 'FAILED', 'COMPENSATING');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mech_status') THEN
        CREATE TYPE mech_status AS ENUM ('PENDING', 'MINTED', 'BURNED');
    END IF;
END $$;

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    wallet_address VARCHAR(42) UNIQUE NOT NULL,
    username VARCHAR(50),
    credits DECIMAL(18, 8) DEFAULT 0,
    last_login TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Mechs (NFTs) Table
CREATE TABLE IF NOT EXISTS mechs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    token_id NUMERIC(78, 0) UNIQUE, -- Supports uint256
    owner_id UUID REFERENCES users(id),
    vehicle_type vehicle_type NOT NULL,
    class vehicle_class NOT NULL,
    image_url TEXT,
    stats JSONB NOT NULL DEFAULT '{}', -- Base stats: HP, Attack, Defense, etc.
    rarity rarity_tier NOT NULL DEFAULT 'COMMON',
    season VARCHAR(50),
    status mech_status DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Parts (Modular Items) Table
CREATE TABLE IF NOT EXISTS parts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID REFERENCES users(id),
    mech_id UUID REFERENCES mechs(id), -- NULL if not equipped
    slot VARCHAR(50) NOT NULL, -- CHASSIS, ARM_L, ARM_R, LEGS, etc.
    name VARCHAR(100) NOT NULL,
    rarity rarity_tier NOT NULL DEFAULT 'COMMON',
    stats JSONB NOT NULL DEFAULT '{}', -- Bonus stats: +HP, +Crit, etc.
    visual_dna JSONB NOT NULL DEFAULT '{}', -- AI Keywords for FLUX.1
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Pilot Stats (Neural Resonance & Resources)
CREATE TABLE IF NOT EXISTS pilot_stats (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    resonance_level INTEGER DEFAULT 0,
    resonance_exp INTEGER DEFAULT 0,
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
CREATE TYPE node_type AS ENUM ('COMBAT', 'RESOURCE', 'NARRATIVE', 'REST', 'BOSS');

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
    allowed_modes TEXT[], -- PILOT, MECH, TANK, etc.
    requires_atmosphere BOOLEAN DEFAULT FALSE,
    suitability_pilot INTEGER DEFAULT 50,
    suitability_mech INTEGER DEFAULT 50,
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
    suitability_mech INTEGER DEFAULT 50,
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
    vehicle_id UUID REFERENCES mechs(id),
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
    mech_id UUID REFERENCES mechs(id),
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
    type VARCHAR(50) NOT NULL, -- MECH_ASSEMBLY, STAR_DISCOVERY, etc.
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

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_mechs_owner ON mechs(owner_id);
CREATE INDEX IF NOT EXISTS idx_saga_user ON saga_transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_saga_idempotency ON saga_transactions(idempotency_key);
CREATE INDEX IF NOT EXISTS idx_combat_user ON combat_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_combat_expires ON combat_logs(expires_at) WHERE is_permanent = FALSE;

-- Sample Data for Narrative Timeline
INSERT INTO expeditions (id, title, description, goal) VALUES 
('00000000-0000-0000-0000-000000000001', 'The Silent Signal', 'A mysterious signal is emanating from the Iron Nebula.', 'Locate the source of the signal and decrypt it.')
ON CONFLICT (id) DO NOTHING;

INSERT INTO encounters (id, expedition_id, type, title, description, visual_prompt) VALUES 
('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', 'NARRATIVE', 'The Awakening', 'You wake up in the cold silence of your cockpit. The radar is flickering.', 'Tactical Noir style, a pilot waking up in a dark cockpit, flickering neon lights, high detail')
ON CONFLICT (id) DO NOTHING;
