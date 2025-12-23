# Gameplay Architecture: Project-0

## 1. Overview
This document defines the technical implementation of core gameplay pillars, ensuring they align with the Modular Monolith and Saga Pattern architecture.

## 2. Core Systems

### 2.1 Colony Management
- **Module:** `internal/colony`
- **State:** Managed in PostgreSQL (`colonies` table).
- **Mobility:** Colony position affects fuel consumption for exploration missions. Moving the colony requires a `ColonyMoveSaga`.
- **Progression:** Upgrading facilities (Hangar, Lab, Refinery) unlocks higher-tier Mechs and faster resource processing.

### 2.2 Exploration Stages
- **Module:** `internal/exploration`
- **Frontend Tech:** React Three Fiber (Three.js) for seamless 3D Cockpit and EVA sequences.
- **Stages:**
    - **Near-Space (Debris Field):** Mothership only. Collect salvage from space debris. Low risk.
    - **Deep-Space (Star Systems):** Mothership + Mechs/Aircraft. High risk, requires advanced Ships.
    - **Planetary Surface:** Mothership (Atmospheric Entry) + Mechs/Aircraft + Pilot EVA.
- **Vehicle Roles:**
    - **Mech:** Best for heavy resource extraction and high-durability ground combat.
    - **Aircraft:** Best for rapid scouting, aerial superiority, and high-evasion hit-and-run tactics.
- **Saga:** `ExplorationMissionSaga`
    1. **Radar Scan & Risk Assessment:** Use Mothership Scanner to probe target. Calculate success probability and Threat Level (Low/Med/High/Extreme).
    2. **Pre-flight Briefing:** Present fuzzy radar data, risk %, and warnings to the player. Wait for "Proceed" or "Abort" command.
    3. **Departure:** Deduct fuel and lock participating units.
    4. **Event Resolution:** 
        - Roll for random accidents based on risk %.
        - If an accident occurs, trigger **AI Event Generation** (context-aware narrative).
        - Calculate combat/loot/exploration results (including hidden elite threats).
    5. **Return:** Update unit durability (including accident damage), distribute loot, and generate AI Combat Log.
    6. **Visualization:** Trigger 3D sequence: Mothership Transit -> Accident/Entry Sequence -> Cockpit Zoom -> Pilot EVA.

### 2.3 Salvage & Research
- **Module:** `internal/salvage`
- **Modular Salvage:** Players extract individual **NFT Components** from defeated enemies.
- **Pilot Requirement:** High-tier salvage in tight spaces requires a **Pilot EVA** with specific gear (O2, Tools).
- **Actions:**
    - **Mint Part:** Convert salvaged component into a new NFT part.
    - **Research:** Consume part to progress Tech Tree.
    - **Scrap:** Break down part into resources.
- **Saga:** `SalvageOperationSaga` ensures atomic processing of the chosen action.

### 2.4 Pilot & Gear
- **Module:** `internal/pilot`
- **State:** Managed in PostgreSQL (`pilots` table).
- **Gear:** O2 Tanks, Swords, Guns, and Suits.
- **Progression:** Upgrading Pilot gear is essential for surviving the final stage of high-value missions.

### 2.4 Story Mode
- **Module:** `internal/story`
- **Fixed Items:** Unlike AI-generated seasonal NFTs, Story Mode rewards are "Fixed" to ensure balanced progression.
- **Narrative Integration:** Acts as a trigger for the `StoryProgressionSaga`, which unlocks new star systems and colony facilities.

## 3. AI Integration
- **Combat Logs:** Every mission result is sent to the AI Service to generate a unique narrative summary.
- **Modular NFT Assembly (FLUX.1):**
    - **Visual Synthesis:** The AI Service "reads" the visual traits of individual NFT components (Railgun, Shield, Pilot Suit) and synthesizes them into a single image.
    - **Prompt Engineering:** The AI Service constructs prompts based on:
        - **Equipped NFT Parts:** Specific visual identifiers from each NFT's metadata (e.g., *"Chrome-plated Railgun NFT #402"*).
        - **Durability State:** Visual damage descriptors applied to specific parts.
        - **Event Context:** Background context from the mission (e.g., *"during a high-speed atmospheric entry"*).
    - **Consistency:** Ensures the final "Master Image" is a 1:1 representation of the combined NFT parts and their current condition.
- **Visual Wear & Tear:** 3D models use procedural shaders for real-time damage, while AI generates the high-fidelity "Master Image" for the NFT.

## 4. Consistency & Idempotency
- All gameplay actions that modify player assets or currency must use **Idempotency Keys**.
- Failed steps in any gameplay Saga must trigger compensating transactions to prevent state desync.
