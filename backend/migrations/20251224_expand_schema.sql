-- Migration: Expand Schema for Players, Items, and Gacha
-- Date: 2025-12-24

-- 1. Update Rarity Enum
ALTER TYPE rarity_tier ADD VALUE IF NOT EXISTS 'REFINED';
ALTER TYPE rarity_tier ADD VALUE IF NOT EXISTS 'PROTOTYPE';
ALTER TYPE rarity_tier ADD VALUE IF NOT EXISTS 'RELIC';
ALTER TYPE rarity_tier ADD VALUE IF NOT EXISTS 'SINGULARITY';

-- 2. Update Mechs Table
ALTER TABLE mechs ADD COLUMN IF NOT EXISTS tier INTEGER DEFAULT 1;
ALTER TABLE mechs ADD COLUMN IF NOT EXISTS is_void_touched BOOLEAN DEFAULT FALSE;
-- is_minted is already handled by 'status' enum but let's add a explicit flag if needed
-- status enum has PENDING, MINTED, BURNED. That's enough.

-- 3. Update Parts Table
ALTER TABLE parts ADD COLUMN IF NOT EXISTS tier INTEGER DEFAULT 1;
ALTER TABLE parts ADD COLUMN IF NOT EXISTS is_void_touched BOOLEAN DEFAULT FALSE;
ALTER TABLE parts ADD COLUMN IF NOT EXISTS is_minted BOOLEAN DEFAULT FALSE;

-- 4. Update Pilot Stats
ALTER TABLE pilot_stats ADD COLUMN IF NOT EXISTS xp INTEGER DEFAULT 0;
ALTER TABLE pilot_stats ADD COLUMN IF NOT EXISTS rank INTEGER DEFAULT 1;

-- 5. Gacha Pity System
CREATE TABLE IF NOT EXISTS gacha_stats (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    pity_relic_count INTEGER DEFAULT 0,
    pity_singularity_count INTEGER DEFAULT 0,
    total_pulls INTEGER DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 6. Gacha History
CREATE TABLE IF NOT EXISTS gacha_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    item_id UUID, -- Can be Mech or Part ID
    item_type TEXT NOT NULL, -- 'MECH', 'PART'
    pull_type TEXT NOT NULL, -- 'STANDARD_SIGNAL', 'VOID_SIGNAL'
    rarity rarity_tier NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 7. Initialize Gacha Stats for existing users
INSERT INTO gacha_stats (user_id)
SELECT id FROM users
ON CONFLICT (user_id) DO NOTHING;
