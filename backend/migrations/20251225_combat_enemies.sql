-- Add enemy_id to encounters table
ALTER TABLE encounters ADD COLUMN IF NOT EXISTS enemy_id UUID REFERENCES mechs(id);

-- Create a System User for NPC Mechs if not exists
INSERT INTO users (id, wallet_address, username)
VALUES ('00000000-0000-0000-0000-000000000000', '0x0000000000000000000000000000000000000000', 'SYSTEM_NPC')
ON CONFLICT (id) DO NOTHING;

-- Seed some Enemy Mechs
INSERT INTO mechs (id, owner_id, vehicle_type, class, rarity, stats, status)
VALUES 
('e0000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000000', 'MECH', 'STRIKER', 'COMMON', '{"hp": 80, "attack": 15, "defense": 5, "speed": 60}', 'MINTED'),
('e0000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000000', 'TANK', 'GUARDIAN', 'RARE', '{"hp": 200, "attack": 10, "defense": 20, "speed": 30}', 'MINTED'),
('e0000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000000', 'SHIP', 'SCOUT', 'COMMON', '{"hp": 60, "attack": 12, "defense": 3, "speed": 100}', 'MINTED')
ON CONFLICT (id) DO NOTHING;
