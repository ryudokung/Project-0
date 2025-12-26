# Gameplay Architecture: Project-0

## 1. Overview
This document defines the technical implementation of core gameplay pillars, ensuring they align with the Modular Monolith and Saga Pattern architecture.

## 2. Core Systems

### 2.1 Bastion Management
- **Module:** `internal/bastion`
- **State:** Managed in PostgreSQL (`bastions` table).
- **Mobility:** Bastion position affects fuel consumption for exploration missions. Moving the bastion requires a `BastionMoveSaga`.
- **Progression:** Upgrading facilities (Preparation Area, Lab, Refinery) unlocks higher-tier Vehicles and faster resource processing.

### 2.2 Exploration Stages
- **Module:** `internal/exploration`
- **Frontend Tech:** **Single-Page Game Loop** with Framer Motion for stage transitions. React Three Fiber (Three.js) for the Bastion Showcase and Cockpit HUD.
- **Stages (State-Driven):**
    - **BASTION:** Pilot and Vehicle management. Showcase engine for "Flexing" assets via the **Visual Equipment Map**.
    - **MAP (Universe Map):** Sector-level navigation.
    - **LOCATION_SCAN:** Sub-sector scanning for Points of Interest (Wrecks, Stations, Planets).
    - **PLANET_SURFACE:** Tactical objective selection for planetary missions.
    - **EXPLORATION (The Loop):** Core encounter progression, O2/Fuel management, and AI narrative generation.
    - **COMBAT:** Turn-based tactical simulation.
    - **DEBRIEF:** Mission summary and loot distribution.
- **Vehicle Roles:**
    - **Vehicle:** Versatile modular units (Mechs, Tanks, Ships) for heavy resource extraction and combat.
- **Saga:** `ExplorationMissionSaga`
    1. **Radar Scan & Risk Assessment:** Use Bastion Scanner to probe target. Calculate success probability and Threat Level (Low/Med/High/Extreme).
    2. **Pre-flight Briefing:** Present fuzzy radar data, risk %, and warnings to the player. Wait for "Proceed" or "Abort" command.
    3. **Departure:** Deduct fuel and lock participating units.
    4. **Event Resolution:** 
        - Roll for random accidents based on risk %.
        - If an accident occurs, trigger **AI Event Generation** (context-aware narrative).
        - Calculate combat/loot/exploration results (including hidden elite threats).
    5. **Return:** Update unit durability (including accident damage), distribute loot, and generate AI Combat Log.
    6. **Visualization:** Trigger **Seamless Stage Transitions**: Bastion -> Map -> Scan -> Exploration Loop. Cockpit HUD remains persistent to maintain immersion.

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
- **Combat Power (CP):** Calculated as `(ATK*3) + (DEF*2) + (HP/5)` to provide a quick strength assessment.

### 2.4 Story Mode
- **Module:** `internal/story`
- **Fixed Items:** Unlike AI-generated seasonal NFTs, Story Mode rewards are "Fixed" to ensure balanced progression.
- **Narrative Integration:** Acts as a trigger for the `StoryProgressionSaga`, which unlocks new star systems and Bastion facilities.

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
