-- Migration: Add Daily Signal and Enhanced Gacha Logging
-- Date: 2025-12-24

-- Add last_free_pull_at to gacha_stats
ALTER TABLE gacha_stats ADD COLUMN IF NOT EXISTS last_free_pull_at TIMESTAMP WITH TIME ZONE;

-- Add enhanced logging columns to gacha_history
ALTER TABLE gacha_history ADD COLUMN IF NOT EXISTS seed BIGINT;
ALTER TABLE gacha_history ADD COLUMN IF NOT EXISTS pity_relic_before INTEGER;
ALTER TABLE gacha_history ADD COLUMN IF NOT EXISTS pity_relic_after INTEGER;
ALTER TABLE gacha_history ADD COLUMN IF NOT EXISTS pity_singularity_before INTEGER;
ALTER TABLE gacha_history ADD COLUMN IF NOT EXISTS pity_singularity_after INTEGER;

-- Create Mothership Upgrades table
CREATE TABLE IF NOT EXISTS mothership_upgrades (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    node_id VARCHAR(100) NOT NULL, -- e.g., 'SIGNAL_BOOSTER_1'
    path VARCHAR(50) NOT NULL, -- 'TELEPORT' or 'ENTRY'
    unlocked_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, node_id)
);

-- Seed Starter Gear (Optional, but good for reference)
-- We'll handle starter gear assignment in the code, but we can define the items here if needed.
