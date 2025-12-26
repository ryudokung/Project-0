package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/project0?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Seeding Exploration Data...")

	// Clear existing data
	_, err = db.Exec(`TRUNCATE sectors, sub_sectors, planet_locations CASCADE`)
	if err != nil {
		log.Fatal(err)
	}

	// 1. SOL GATE
	solGateID := uuid.New()
	_, err = db.Exec(`INSERT INTO sectors (id, name, description, difficulty, coordinates_x, coordinates_y, color) 
		VALUES ($1, 'SOL GATE', 'The industrial gateway to the system. Relatively safe but heavily monitored.', 'LOW', 15, 20, 'blue')`,
		solGateID)
	if err != nil {
		log.Fatal(err)
	}

	// SubSectors for SOL GATE
	_, err = db.Exec(`INSERT INTO sub_sectors (sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'STATION', 'Outpost 01', 'A standard refueling station for independent scavengers.', '{"Scrap Metal", "Fuel Isotopes"}', '{}', '{"PILOT", "SPEEDER"}', FALSE, 100, 20, 40, 30)`,
		solGateID)
	
	_, err = db.Exec(`INSERT INTO sub_sectors (sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'WRECK', 'Old Freighter', 'A derelict cargo ship drifting near the gate.', '{"Scrap Metal", "O2 Crystals"}', '{}', '{"PILOT", "EXOSUIT"}', FALSE, 90, 10, 60, 50)`,
		solGateID)

	// 2. IRON NEBULA
	ironNebulaID := uuid.New()
	_, err = db.Exec(`INSERT INTO sectors (id, name, description, difficulty, coordinates_x, coordinates_y, color) 
		VALUES ($1, 'IRON NEBULA', 'A dense cloud of metallic dust and derelict warships. High gravity zones.', 'MEDIUM', 45, 40, 'pink')`,
		ironNebulaID)

	// SubSectors for IRON NEBULA
	_, err = db.Exec(`INSERT INTO sub_sectors (sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'WRECK', 'Scrap Graveyard', 'A massive cluster of destroyed freighter hulls.', '{"Scrap Metal", "Neural Links"}', '{}', '{"PILOT", "SPEEDER", "EXOSUIT"}', FALSE, 80, 40, 30, 40)`,
		ironNebulaID)

	kriosPrimeID := uuid.New()
	_, err = db.Exec(`INSERT INTO sub_sectors (id, sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, $2, 'PLANET', 'Krios Prime', 'A frozen planetoid with hidden Syndicate bunkers.', '{"Fuel Isotopes", "Neural Links"}', '{}', '{"MECH", "TANK", "HAULER"}', TRUE, 20, 90, 70, 60)`,
		kriosPrimeID, ironNebulaID)

	// Locations for Krios Prime
	_, err = db.Exec(`INSERT INTO planet_locations (sub_sector_id, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'Bunker Alpha', 'Deep underground storage facility.', '{"Neural Links", "Scrap Metal"}', '{"Hacking Module"}', '{"PILOT", "EXOSUIT"}', TRUE, 90, 10, 20, 30)`,
		kriosPrimeID)

	_, err = db.Exec(`INSERT INTO planet_locations (sub_sector_id, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'Mining Rig 7', 'Automated extraction site on the surface.', '{"Scrap Metal", "Fuel Isotopes"}', '{"Mining Drill"}', '{"MECH", "HAULER"}', TRUE, 10, 100, 60, 70)`,
		kriosPrimeID)

	// 3. NEON ABYSS
	neonAbyssID := uuid.New()
	_, err = db.Exec(`INSERT INTO sectors (id, name, description, difficulty, coordinates_x, coordinates_y, color) 
		VALUES ($1, 'NEON ABYSS', 'A high-tech sector plagued by EMP storms and rogue AI signals.', 'HIGH', 75, 30, 'red')`,
		neonAbyssID)

	_, err = db.Exec(`INSERT INTO sub_sectors (sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'STATION', 'Data Hive', 'A massive server farm drifting in a nebula.', '{"Neural Links", "Void Shards"}', '{"Hacking Module"}', '{"PILOT", "EXOSUIT"}', FALSE, 100, 0, 50, 50)`,
		neonAbyssID)

	// 4. THE DEAD RIM
	deadRimID := uuid.New()
	_, err = db.Exec(`INSERT INTO sectors (id, name, description, difficulty, coordinates_x, coordinates_y, color) 
		VALUES ($1, 'THE DEAD RIM', 'The edge of known space. Ancient ruins and ghost signals.', 'EXTREME', 50, 80, 'red')`,
		deadRimID)

	vulcanisID := uuid.New()
	_, err = db.Exec(`INSERT INTO sub_sectors (id, sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, $2, 'PLANET', 'Vulcanis', 'A high-gravity mining planet near a dying star.', '{"Void Shards", "Ancient Tech"}', '{}', '{"MECH", "TANK", "SHIP"}', TRUE, 5, 95, 40, 30)`,
		vulcanisID, deadRimID)

	_, err = db.Exec(`INSERT INTO planet_locations (sub_sector_id, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y) 
		VALUES ($1, 'Magma Chamber', 'Extreme heat zone with rare mineral deposits.', '{"Ancient Tech", "Nexus Cores"}', '{"Mining Drill"}', '{"MECH", "TANK"}', TRUE, 0, 100, 50, 50)`,
		vulcanisID)

	// 5. Initialize Mock Pilot
	mockUserID := "a58aa13f-f715-4137-bdf3-6ee44dd244ba"
	mockCharacterID := "c58aa13f-f715-4137-bdf3-6ee44dd244ba"

	_, err = db.Exec(`INSERT INTO users (id, wallet_address, username) VALUES ($1, '0x1234567890123456789012345678901234567890', 'PILOT_0') ON CONFLICT DO NOTHING`, mockUserID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO characters (id, user_id, name, gender) VALUES ($1, $2, 'PILOT_0', 'MALE') ON CONFLICT DO NOTHING`, mockCharacterID, mockUserID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO pilot_stats (user_id, character_id, current_o2, current_fuel) VALUES ($1, $2, 100.0, 100.0) ON CONFLICT (character_id) DO UPDATE SET current_o2 = 100.0, current_fuel = 100.0`, mockUserID, mockCharacterID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding Complete!")
}
