-- Migration: Character System and Starter Ship Pivot
-- Date: 2025-12-24

-- 1. Create Characters Table
CREATE TABLE IF NOT EXISTS characters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    name VARCHAR(50) NOT NULL,
    gender VARCHAR(10), -- 'MALE', 'FEMALE', 'NON_BINARY'
    face_index INTEGER DEFAULT 0,
    hair_index INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. Add active_character_id to users
ALTER TABLE users ADD COLUMN IF NOT EXISTS active_character_id UUID REFERENCES characters(id);

-- 3. Update pilot_stats to link to characters
ALTER TABLE pilot_stats ADD COLUMN IF NOT EXISTS character_id UUID REFERENCES characters(id);

-- 4. Update mechs to link to characters
ALTER TABLE mechs ADD COLUMN IF NOT EXISTS character_id UUID REFERENCES characters(id);

-- 5. Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_characters_user_id ON characters(user_id);
CREATE INDEX IF NOT EXISTS idx_pilot_stats_character_id ON pilot_stats(character_id);
CREATE INDEX IF NOT EXISTS idx_mechs_character_id ON mechs(character_id);
