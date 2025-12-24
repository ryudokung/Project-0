# Vehicle & Item Bible: Project-0 (The Asset Catalog)

**Goal:** Define the specific modular parts, rarity tiers, and "Visual DNA" traits for all vehicle types (Mechs, Tanks, Ships) to ensure a consistent asset ecosystem.

## 1. Modular Vehicle Anatomy (Functional Hardpoints)
All vehicles are composed of 4 core NFT slots. The Hangar serves as a central **Asset Inventory** where players can view their collection. Every new character begins with a **Starter Ship** to facilitate initial exploration. Actual vehicle selection occurs at the **Deployment Phase** based on mission requirements.

### 1.1 Ship (Atmospheric/Space) - **Starter Vehicle**
- **Role:** Interception, Long-range Scouting, Mothership Defense.
- **Deployment:** Space combat, atmospheric entry escort.
- **Hardpoints:** Fuselage, Wing/Pylon, Engine, Avionics.
- **Starter Model:** `FS-01: "Void Runner"` (Standard Fuselage).

### 1.2 Mech (Bipedal/Multipedal)
- **Role:** Heavy Combat, Extraction, All-terrain.
- **Deployment:** Surface missions, high-gravity zones.
- **Hardpoints:** Chassis, Left Arm, Right Arm, Legs.

### 1.3 Tank (Treaded/Wheeled)
- **Role:** Siege, Defensive Escort, High-threat Zones.
- **Deployment:** Open surface combat, defensive missions.
- **Hardpoints:** Hull, Turret, Sponson, Drive System.

### 1.4 Speeder (Hoverbike/Swoop)
- **Role:** High-speed Scouting, Time-sensitive Salvage, Stealth.
- **Deployment:** Recon, low-threat salvage, stealth infiltration.
- **Hardpoints:** Frame, High-output Engine, Handlebar/Sensors, Side-mount (Light Weapon).

### 1.5 Exosuit (Power Armor)
- **Role:** Indoor Salvage, Boarding Actions, Close-quarters Combat (CQC).
- **Deployment:** Derelict ship interiors, bunkers, narrow caves.
- **Hardpoints:** Core, Arms (Melee/Small arms), Back-pack (O2/Thrusters), Boots.

### 1.6 Heavy Hauler (Industrial)
- **Role:** Resource Transport, Salvaging Massive Wrecks, Base Building.
- **Deployment:** Bulk resource extraction, wreck towing.
- **Hardpoints:** Hull, Crane/Grapple, Cargo Bay, Defense Turrets.

## 2. Item Rarity & "Options" (Sub-stats)
Items are minted with a base stat and 1-4 "Options" based on rarity:
- **Common (White):** 0 Options.
- **Rare (Blue):** 1 Option.
- **Epic (Purple):** 2 Options.
- **Legendary (Gold):** 3 Options + 1 Unique Perk.
- **Seasonal (Red):** 4 Options + Unique Visual Effect (e.g., "Energy Wings", "Plasma Trail").

## 3. Visual DNA Framework (AI Synthesis)
The AI (FLUX.1) uses "Visual Keywords" derived from metadata to synthesize a cohesive image:
- **The Synthesis Process:**
    1. **Metadata Collection:** Backend gathers all part IDs, faction styles, and rarity traits.
    2. **Prompt Engineering:** A structured prompt is generated (e.g., "A brutalist Iron Syndicate Mech with rusted steel plating and a glowing green core, tactical noir style, high-fidelity 3D render").
    3. **AI Generation:** FLUX.1 generates the unique visual representation.
    4. **Caching:** The image is stored as the "Visual DNA" of that specific asset.
- **Materials:** `industrial_steel`, `carbon_fiber`, `weathered_copper`, `obsidian_glass`.
- **Styles:** `brutalist`, `sleek_aerodynamic`, `gothic_mechanical`, `salvaged_junk`.
- **Conditions:** `pristine`, `battle_scarred`, `rusted`, `overheated`.

## 4. Item Catalog (Sample Items)

### 4.1 Chassis/Hull/Fuselage
- **CH-01: "The Bastion" (Mech Chassis):** Brutalist, Heavy Armor.
- **HL-01: "Iron Tusk" (Tank Hull):** Weathered, Sloped Armor.
- **FS-01: "Void Runner" (Ship Fuselage):** Sleek, Carbon Fiber.

### 4.2 Weapons
- **W-01: "Railgun-X" (Kinetic):** Long Barrel, High Range.
- **W-02: "Plasma Cutter" (Energy):** Glowing Blue, Short Range.
- **W-03: "Hellfire Rack" (Explosive):** Missile Pods, Area Damage.

## 5. Pilot Gear (EVA Focus)
- **Suit:** O2 Capacity, Agility.
- **Sidearm:** Melee/Ranged damage for EVA combat.
- **Gadget:** Hacking tools, Scanners, Jetpacks.

## 6. Economy & Salvage
- **Scrap:** Common currency for repairs.
- **Cores:** Rare material for "Overclocking" (Upgrading stats).
- **Blueprints:** Required to mint new NFTs from salvage.
