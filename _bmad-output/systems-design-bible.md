# Systems Design Bible: Project-0 (The Mechanics)

**Goal:** Define the core gameplay systems, including materials, items, skills, and the research tree to ensure deep progression and clear utility.

## 1. Exploration Mechanics: The "Expedition and Encounters" System
Exploration is governed by the interaction between the Pilot's resources and the AI-generated narrative timeline:

*Implementation Note: The logic for generating these encounters is driven by the **YAML Blueprint System**, allowing for decoupled content management.*

1.  **The Expedition (Narrative Spine):** A sequence of fixed goals (e.g., "Locate the Signal Source").
2.  **The Encounters (Procedural Events):** Each "Click" or "Move" on the timeline generates an Encounter.
    *   **Generation Logic:** The AI checks the Pilot's current **O2** and **Fuel**.
    *   **Low Resources:** Triggers "Resource Encounters" (e.g., finding an O2 tank) or "High-Stakes Narrative Encounters" (e.g., a desperate gamble).
    *   **High Resources:** Triggers "Combat Encounters" or "Lore Encounters" to progress the story.
3.  **Visual DNA Synthesis:** Every Encounter generates a unique visual prompt combining the Vehicle's parts and the current environment, providing a "Visual Reveal" for the player.

## 2. Data-Driven Engine: The Blueprint System
To ensure scalability and rapid iteration, the game engine uses a YAML-based Blueprint system to define all game entities:
- **Node Blueprints (`nodes.yaml`):** Defines the metadata, strategic choices, success chances, and reward pools for every exploration node.
- **Enemy Blueprints (`enemies.yaml`):** Defines the stats (HP, ATK, DEF, SPD), rarity, and Combat Rating (CR) for all NPC factions.
- **Future Expansion:** This system will be expanded to cover **Modular Parts** and **Loot Tables**, moving all balance-related data out of the core Go logic.

## 3. Materials & Resource Economy

| Material | Rarity | Primary Source | Primary Use |
| :--- | :--- | :--- | :--- |
| **Scrap Metal** | Common | Derelict Wrecks | Repairs, Basic Crafting, Radar Restoration |
| **Fuel Isotopes** | Common | Asteroid Belts | Mothership Travel, Vehicle Energy |
| **O2 Crystals** | Common | Frozen Moons | Refilling Pilot O2 Tanks |
| **Neural Links** | Uncommon | AI Drones | AI Synthesis, Pilot Skill Upgrades |
| **Void Shards** | Rare | Void Cult Vaults | **The Singularity** Research, V2O Bridging |
| **Ancient Tech** | Legendary | **The Singularity** Ruins | Crafting Legendary Blueprints |
| **Nexus Cores** | Mythic | The Singularity (Location) | Final Evolution of Vehicles |

## 2. Item Classification
Items are divided into three categories based on their persistence and utility:

### 2.1 Modular Parts (NFTs)
- **Nature:** Permanent assets (once minted).
- **Slots:** Chassis, Arms, Legs, Weapons, Engines.
- **Utility:** Determines the stats and "Visual DNA" of your vehicles.

### 2.2 Consumables (Non-NFT)
- **Nature:** Single-use items.
- **Examples:** 
    - **Repair Nanites:** Instant 25% HP repair during a mission.
    - **O2 Stim:** Instant 50% O2 refill for Pilot EVA.
    - **Signal Flare:** Reveals all nodes in a small radius for 1 scan.

### 2.3 Blueprints & Data Cores
- **Nature:** Crafting requirements.
- **Utility:** A Blueprint + specific Materials = A new Modular Part.
- **Data Cores:** Can be "Decrypted" in the Lab to reveal new Research nodes.

## 3. The Research Tree (Mothership Progression)
Progression is gated by the Mothership's facilities. You must invest Materials to unlock new tiers.

### Tier 1: The Scavenger (Early Game)
- **Radar Level 1:** Unlocks basic mission nodes in Sol Gate.
- **Bastion Level 1:** Allows repair of Vehicles and basic Modules.
- **Basic Refiner:** Converts raw Scrap into "Refined Alloy."

### Tier 2: The Operative (Mid Game)
- **Scanner Level 1:** Unlocks "Deep Scan" to find hidden high-value nodes.
- **Lab Level 1:** Unlocks "Blueprint Decryption" and Pilot Skill training.
- **Factory Level 1:** Allows the assembly of Rare and Epic modular parts.

### Tier 3: The Commander (End Game)
- **Quantum Gate:** Unlocks Teleportation (bypassing Atmospheric Entry).
- **Singularity Archive:** Unlocks research into Ancient Tech and Legendary parts.
- **V2O Bridge:** Unlocks the ability to mint Virtual Assets into On-chain NFTs.

## 4. Skill System: Character vs. Vehicle
Skills provide the tactical edge needed to survive high-risk missions. All progression is tied to the **Active Character**.

### 4.1 Character Skills (Passive/Permanent)
*Trained in the Mothership Lab using Neural Links.*
- **O2 Discipline:** Reduces O2 consumption during EVA by 10/20/30%.
- **Scavenger's Eye:** Increases the chance of finding Rare materials by 5/10/15%.
- **Stealth Training:** Reduces detection range from enemy drones.

### 4.2 Neural Resonance (The "Awakening" System)
Inspired by the "Newtype" concept, this is a latent power that triggers when a **Character's** intent aligns with their Vehicle's core.

- **Resonance Potential:** A hidden stat for every **Character**. Some are born with high potential, others develop it through high-stakes survival.
- **The Trigger (Willpower):** Triggers during "Critical Moments" (e.g., O2 < 5%, Hull Integrity < 10%, or after 5 consecutive High-Risk successes).
- **Awakened State: "The Ghost Shift":**
    - **Visual:** The Vehicle's Visual DNA changes in real-time (AI-generated glowing vents, "Ghost" after-images, or shifting armor plates).
    - **Gameplay:** 
        - **Omniscience:** All hidden nodes in the sector are temporarily revealed.
        - **Zero-Latency:** The next 3 clicks cost 0 O2/Fuel.
        - **Perfect Strike:** The next attack ignores all enemy defense.
- **The Cost:** After the state ends, the Pilot suffers "Neural Strain," requiring a rest period in the Bastion.

### 4.3 Vehicle Skills (Active/Part-Based)
*Linked to specific Modular Parts (Weapons/Engines).*
- **Overdrive (Engine):** Increases Speed by 50% for 3 clicks (consumes extra Fuel).
- **Shield Pulse (Chassis):** Negates the next instance of damage.
- **Tactical Scan (Sensors):** Reveals the contents of a node before clicking it.

## 5. The "Loki" Context: Dynamic Research
As the Director, you can release **"Context Patches"** that add temporary or permanent branches to the Research Tree (e.g., "A new Void Cult technology has been discovered—unlock the 'Void Shield' research node this week only").

## 6. Deep Gameplay Mechanics: The "Anti-Pump" System

### 6.1 Effective Combat Power (ECP) Model
ECP is the final value used for combat calculations and "Annihilation" checks. It ensures that raw stats are filtered through tactical context.

**Formula:**
$$ECP = (Base\_CP \times Suitability\_Mod) \times Resonance\_Sync \times (1 - Fatigue\_Penalty)$$

*   **Suitability_Mod:**
    *   `1.2x` (Perfect): Vehicle Type matches Terrain Tag.
    *   `1.0x` (Neutral): Standard compatibility.
    *   `0.5x` (Incompatible): e.g., Tank in Islands.
*   **Resonance_Sync:** `Min(1.0, Pilot_Resonance / Vehicle_Tier_Requirement)`.
*   **Fatigue_Penalty:** `Pilot_Stress / 200` (Max 50% penalty at 100 Stress).

### 6.2 Terrain & Environmental Hazards
Terrains are not just stat modifiers; they introduce active gameplay risks that require specific counter-measures.

| Terrain | Hazard Name | Effect | Counter-Module |
| :--- | :--- | :--- | :--- |
| **DESERT** | Sandstorm | -30% Accuracy, +10% Fuel Cost | Advanced Air Filters |
| **ISLANDS** | Corrosive Salt | -5% HP per 3 clicks | Anti-Corrosion Coating |
| **VOID** | Neural Static | +2 Stress per click | Neural Dampener |
| **URBAN** | High Signal | +20% Detection Signature | Stealth Plating |
| **SKY** | Turbulence | -20% Evasion | Gyro-Stabilizers |

### 6.3 Pilot Stress & Fatigue System
To prevent "Single-Pilot Dominance" and encourage roster rotation, pilots accumulate mental strain.

*   **Stress Accumulation:**
    *   +1 per click in standard nodes.
    *   +5 per Combat Encounter.
    *   +10 if Vehicle HP falls below 20%.
*   **Stress Thresholds:**
    *   **0-30 (Relaxed):** No penalties.
    *   **31-70 (Stressed):** -10% Resonance Sync.
    *   **71-99 (Burnout):** -30% Resonance Sync, cannot use "Signature Skills".
    *   **100 (Breakdown):** Pilot becomes "Inactive" for 24 hours (real-time).
*   **Recovery:**
    *   **Passive:** -5 Stress per hour while in Bastion.
    *   **Active:** Use "Neural Stim" consumable (-50 Stress instantly).

### 6.4 Transformation Logic (Tactical Shift)
Transformable vehicles have two distinct "States" (A and B) to adapt to the Expedition's Timeline.

*   **State A (Travel/Stealth):** High Speed, Low Fuel Cost, Low Signature.
*   **State B (Assault/Combat):** Low Speed, High ATK/DEF, High Signature.
*   **Switching:** Can be done mid-expedition. Each switch costs 5 Fuel and 1 "Click" on the timeline. This allows a player to bypass a detection gate in State A and then switch to State B for the boss fight.
