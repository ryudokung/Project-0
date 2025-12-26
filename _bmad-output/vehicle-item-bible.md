# Vehicle & Item Bible: Project-0 (The Asset Catalog)

**Goal:** Define the specific modular parts, rarity tiers, and "Visual DNA" traits for all vehicle types (Mechs, Tanks, Ships) to ensure a consistent asset ecosystem.

## 1. Modular Vehicle Anatomy (Anatomical Equipment System)
All vehicles are composed of modular slots mapped to a visual silhouette. The **Bastion** serves as a central **Strategic Hub** where players can view their collection and manage their fleet using the **Visual Equipment Map**.

### 1.1 Anatomical Slots
- **HEAD:** Sensors, Neural Interfaces, and Targeting Systems.
- **CORE:** Power Source, Armor Plating, and Life Support.
- **ARM_L / ARM_R:** Weapons, Shields, and Utility Tools.
- **LEGS:** Mobility Systems, Thrusters, and Stability Anchors.

### 1.2 Vehicle Archetypes
- **Vehicle (General):** Balanced bipedal or multipedal units.
- **Tank:** Heavy-armored siege units. High defense.
- **Ship:** High-speed aerial/space units.
- **Speeder:** Hoverbikes for high-velocity scouting.
- **Pilot & Exosuit:** The operative's "Infiltration Layer." Exosuits are **Character Gear** worn by the Pilot, providing a base layer of protection and utility.

### 1.3 Layered Equipment System
- **Inner Layer (Pilot):** The Pilot wears an **Exosuit**. This gear is bound to the character and provides bonuses even when inside a vehicle.
- **Outer Layer (Vehicle):** The Pilot operates a **Vehicle**. This is the primary asset for heavy combat and long-distance travel.
- **Set Synergy:** Equipping an Exosuit and a Vehicle from the same **Series** (e.g., "Void-Walker") unlocks unique **Set Passives**.

## 2. Progression & Performance Metrics

### 2.1 Combat Power (CP)
The **Combat Power (CP)** is the primary indicator of a vehicle's power. It is calculated dynamically based on the synergy of the vehicle and its equipped parts:
- **Formula:** `(Total ATK * 2) + (Total DEF * 2) + (Total HP / 10)`
- **Effective CP (ECP):** `(Vehicle_CP + Exosuit_CP) * Suitability_Mod * Resonance_Sync * (1 - Fatigue_Penalty) * Synergy_Mod`.
- **Base Stats:** HP, Attack, Defense, Speed.
- **Pilot Resonance:** A multiplier based on the Pilot's rank and synchronization.
- **Equipment Quality:** The rarity and tier of equipped parts.
- **Void-Touch Status:** A multiplier for assets corrupted by the void.
- **Set Synergy:** +15% ECP bonus when Exosuit and Vehicle share the same **Series** metadata.

### 2.2 Suitability System
Vehicles possess **Suitability Tags** (e.g., `urban`, `desert`, `void`, `high-gravity`) that determine their effectiveness in specific environments.
- **Matching Tags:** Grant bonuses to Speed and Defense.
- **Mismatched Tags:** Impose penalties to Energy Consumption and Accuracy.

## 3. Item Rarity & "Options" (Sub-stats)
Items are manifested with a base stat and 1-4 "Options" based on rarity:
- **Common (White):** 0 Options.
- **Rare (Blue):** 1 Option.
- **Epic (Purple):** 2 Options.
- **Legendary (Gold):** 3 Options + 1 Unique Perk.
- **Singularity (Red):** 4 Options + Unique Visual Effect (e.g., "Energy Wings", "Plasma Trail").

## 4. Visual DNA Framework (AI Synthesis)
The AI (FLUX.1) uses "Visual Keywords" derived from metadata to synthesize a cohesive image:
- **The Synthesis Process:**
    1. **Metadata Collection:** Backend gathers all part IDs, faction styles, and rarity traits.
    2. **Prompt Engineering:** A structured prompt is generated (e.g., "A brutalist Iron Syndicate Vehicle with rusted steel plating and a glowing green core, tactical noir style, high-fidelity 3D render").
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
Exosuits use the same **Anatomical Mapping** (HEAD, CORE, ARMS, LEGS) as vehicles but at a personal scale.
- **HEAD (Neural Interface):** Resonance Level, Accuracy, HUD enhancements.
- **CORE (Exosuit Chassis):** O2 Capacity, Armor, Life Support.
- **ARMS (Power Grips):** Melee damage, Hacking speed, Recoil control.
- **LEGS (Kinetic Actuators):** Agility, Jump height, Stealth movement.
- **Utility Gadget:** Specialized tools like Cloaking Fields or Scanners.
- **Personal Weapon:** Sidearm for CQC combat when outside the vehicle.

## 6. Economy & Salvage
- **Scrap:** Common currency for repairs.
- **Cores:** Rare material for "Overclocking" (Upgrading stats).
- **Blueprints:** Required to mint new NFTs from salvage.
