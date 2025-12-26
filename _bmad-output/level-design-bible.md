# Level Design Bible: Project-0 (The Star Map)

**Goal:** Define the specific structure of the game world, sectors, and environmental hazards to ensure a consistent exploration experience.

## 1. World Structure: The "Expedition and Encounters" Narrative Timeline
The game world is explored through a dynamic timeline model that balances fixed narrative goals with procedural variety:

1.  **The Expedition (Narrative Anchors):** Fixed milestones defined by the Director (e.g., "Reach the Sol Gate", "Repair the Radar"). These are the "Anchors" that provide the context for the journey.
2.  **The Encounters (Procedural Events):** AI-generated sub-events strung onto the Expedition. These include:
    *   **Combat Encounters:** Tactical encounters with enemy factions.
    *   **Resource Encounters:** Salvage opportunities or fuel/O2 replenishment.
    *   **Narrative Encounters:** Lore fragments, pilot logs, or environmental storytelling.
3.  **Discovery Nodes:** The physical locations on the Star Map where Expeditions are initiated.
4.  **Encounter Scenes:** The visual result of an "Encounter" being activated, featuring AI-generated Visual DNA prompts.

## 2. Sector Archetypes
Each sector has a specific "Theme" and "Difficulty Tier":

| Sector Name | Tier | Theme | Hazards | Loot Focus |
| :--- | :--- | :--- | :--- | :--- |
| **Sol Gate** | 1 | Industrial / Safe | Low Radiation | Basic Scrap, Fuel |
| **Iron Nebula** | 2 | Mining / Debris | High Gravity, Debris | Kinetic Parts, Ores |
| **Neon Abyss** | 3 | Cyber / High-Tech | EMP Storms, Hackers | Energy Parts, Data |
| **The Dead Rim** | 4 | Ancient / Eldritch | Dark Matter, Ghost Signals | Legendary Blueprints |

## 3. Zone Types & Gameplay Patterns
Zones define the "Scene" and the specific click-based interaction:

### 3.1 Orbital Nodes (Mothership Focus)
- **Gameplay:** Click to Scan, Click to Travel, Click to Manage Fuel.
- **Visuals:** Static or subtly animated starfield with clickable UI markers.
- **Mechanic:** "Radar Repair" unlocks the ability to see more distant nodes.

### 3.2 Surface & EVA Nodes (Action Focus)
- **Gameplay:** Click-based decision making (e.g., "Enter Room", "Search Crate", "Fire Pistol").
- **Visuals:** High-fidelity 3D scenes (R3F) that update based on clicks.
- **Mechanic:** "O2 Management" - each click or action consumes a small amount of O2.
- **Vehicle:** Light Vehicles or Exosuits are the only vehicles allowed in narrow EVA zones.

### 3.3 High-Speed Corridors (Speeder Focus)
- **Gameplay:** Time-trials, Chases, Rapid Scouting via click-timing.
- **Visuals:** Canyons, Tunnels, Dense Forests.
- **Mechanic:** "Heat Signature" (Speed reduces detection).
- **Vehicle:** Speeders are required for these high-velocity missions.

## 4. Environmental Hazards (The "Narrative" Layer)
Hazards are not just visual; they affect stats and trigger AI events:
- **EMP Storms:** Disables Energy weapons, reduces Radar range.
- **Corrosive Rain:** Constant Durability damage to Vehicle Chassis.
- **Solar Flares:** Increases Heat levels, potentially causing "System Overheat" (Stun).
- **Void Echoes:** Triggers "Psychological Stress" for the Pilot (reduces Accuracy).

## 5. Level Design Framework (The "Scene" Template)
Every new node added to the game must define:
1. **ID & Name:** (e.g., `NODE-K7-01: The Frozen Spire`)
2. **Type:** (Surface / EVA / Orbital)
3. **Hazard Level:** (0-10)
4. **Primary Enemy Faction:** (e.g., "The Scavengers")
5. **Loot Table ID:** (Reference to Item Bible)
6. **AI Narrative Prompt:** (The context provided to the AI for event generation)
