# Vehicle & Item Bible: Project-0 (The Asset Catalog)

**Goal:** Define the specific modular parts, rarity tiers, and "Visual DNA" traits for all vehicle types (Mechs, Tanks, Ships) to ensure a consistent asset ecosystem.

## 1. Modular Vehicle Anatomy (Functional Hardpoints)
All vehicles are composed of modular NFT slots. The **Bastion** serves as a central **Strategic Hub** where players can view their collection and manage their fleet. Every new character begins with a **Starter Vehicle** to facilitate initial exploration. Actual vehicle selection occurs at the **Deployment Phase** based on mission requirements.

### 1.1 Ship (Atmospheric/Space)
- **Role:** Interception, Long-range Scouting, Mothership Defense.
- **Transformation:** Can switch to **Robot Mode** for precision landing and surface combat.
- **Hardpoints:** Fuselage, Wing/Pylon, Engine, Transformation Module (T-Module).

### 1.2 Mech (Bipedal/Multipedal)
- **Role:** Heavy Combat, Extraction, All-terrain.
- **Hardpoints:** Chassis, Left Arm, Right Arm, Legs.

### 1.3 Tank (Treaded/Wheeled)
- **Role:** Siege, Defensive Escort, High-threat Zones.
- **Transformation:** Can switch to **Robot Mode** for terrain adaptability and verticality.
- **Hardpoints:** Hull, Turret, Sponson, Transformation Module (T-Module).

### 1.4 Speeder (Hoverbike/Swoop)
- **Role:** High-speed Scouting, Time-sensitive Salvage, Stealth.
- **Hardpoints:** Frame, High-output Engine, Handlebar/Sensors, Side-mount (Light Weapon).

### 1.5 Pilot & Exosuit (The Infiltration Layer)
- **Role:** Indoor Salvage, Boarding Actions, Stealth Infiltration.
- **Deployment:** Used when the Pilot **Ejects** from a vehicle to enter areas inaccessible to heavy assets.
- **Hardpoints:** Exosuit Chassis (Body), Neural Interface (Head), Utility Belt, Personal Weapon.

### 1.6 Heavy Hauler (Industrial)
- **Role:** Resource Transport, Salvaging Massive Wrecks.
- **Hardpoints:** Hull, Crane/Grapple, Cargo Bay, Defense Turrets.

## 2. Progression & Performance Metrics

### 2.1 Combat Rating (CR)
The **Combat Rating (CR)** is the primary indicator of a vehicle's power. It is calculated dynamically based on:
- **Base Stats:** HP, Attack, Defense, Speed.
- **Pilot Resonance:** The synergy between the Pilot and the Vehicle.
- **Equipment Quality:** The rarity and tier of equipped parts.
- **Void-Touch Status:** A multiplier for assets corrupted by the void.

### 2.2 Suitability System
Vehicles possess **Suitability Tags** (e.g., `urban`, `desert`, `void`, `high-gravity`) that determine their effectiveness in specific environments.
- **Matching Tags:** Grant bonuses to Speed and Defense.
- **Mismatched Tags:** Impose penalties to Energy Consumption and Accuracy.

## 3. Item Rarity & "Options" (Sub-stats)
Items are minted with a base stat and 1-4 "Options" based on rarity:
- **Common (White):** 0 Options.
- **Rare (Blue):** 1 Option.
- **Epic (Purple):** 2 Options.
- **Legendary (Gold):** 3 Options + 1 Unique Perk.
- **Singularity (Red):** 4 Options + Unique Visual Effect (e.g., "Energy Wings", "Plasma Trail").

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

### 4.3 Transformation Modules (T-Modules)
- **TM-01: "V-Shift Core":** Standard transformation logic for Ships.
- **TM-02: "Siege-to-Stride":** Heavy-duty transformation for Tanks.

## 5. Pilot Gear (Exosuit Focus)
- **Exosuit Chassis:** O2 Capacity, Agility, Armor.
- **Neural Interface:** Resonance Level, Accuracy.
- **Utility Gadget:** Hacking tools, Scanners, Stealth Field.
- **Personal Weapon:** Sidearm for CQC combat.

## 6. Economy & Salvage
- **Scrap:** Common currency for repairs.
- **Cores:** Rare material for "Overclocking" (Upgrading stats).
- **Blueprints:** Required to mint new NFTs from salvage.
