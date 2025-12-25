# Systems Design Bible: Project-0 (The Mechanics)

**Goal:** Define the core gameplay systems, including materials, items, skills, and the research tree to ensure deep progression and clear utility.

## 1. Exploration Mechanics: The "Expedition and Encounters" System
Exploration is governed by the interaction between the Pilot's resources and the AI-generated narrative timeline:

1.  **The Expedition (Narrative Spine):** A sequence of fixed goals (e.g., "Locate the Signal Source").
2.  **The Encounters (Procedural Events):** Each "Click" or "Move" on the timeline generates an Encounter.
    *   **Generation Logic:** The AI checks the Pilot's current **O2** and **Fuel**.
    *   **Low Resources:** Triggers "Resource Encounters" (e.g., finding an O2 tank) or "High-Stakes Narrative Encounters" (e.g., a desperate gamble).
    *   **High Resources:** Triggers "Combat Encounters" or "Lore Encounters" to progress the story.
3.  **Visual DNA Synthesis:** Every Encounter generates a unique visual prompt combining the Vehicle's parts and the current environment, providing a "Visual Reveal" for the player.

## 2. Materials & Resource Economy

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
- **Hangar Level 1:** Allows repair of Exosuits and basic Mechs.
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
- **The Cost:** After the state ends, the Pilot suffers "Neural Strain," requiring a rest period in the Mothership.

### 4.3 Vehicle Skills (Active/Part-Based)
*Linked to specific Modular Parts (Weapons/Engines).*
- **Overdrive (Engine):** Increases Speed by 50% for 3 clicks (consumes extra Fuel).
- **Shield Pulse (Chassis):** Negates the next instance of damage.
- **Tactical Scan (Sensors):** Reveals the contents of a node before clicking it.

## 5. The "Loki" Context: Dynamic Research
As the Director, you can release **"Context Patches"** that add temporary or permanent branches to the Research Tree (e.g., "A new Void Cult technology has been discovered—unlock the 'Void Shield' research node this week only").
