# 🎮 Project-0: Gameplay & Technical Specification

This document outlines the core gameplay loop, technical architecture for dynamic systems, and the integration of AI-driven visual consistency.

---

## 1. Core Gameplay Loop: The Extraction Cycle

The game follows a high-stakes "Extraction RPG" loop centered around **The Bastion** and **The Void**.

### Phase 1: The Bastion (Strategic Hub)
*   **Energy Management:** Players allocate the Bastion's limited power core between systems (Shields, Warp Drive, Auto-Repair).
*   **Anatomical Equipment System:** A visual mapping interface (Silhouette Map) where players install modules onto specific anatomical slots (HEAD, CORE, ARM_L, ARM_R, LEGS).
*   **Bastion Operations:** Maintenance of Vehicles and Pilot Gear. Repairing items consumes resources gathered from The Void.
*   **Combat Power (CP) Calculation:** A standardized formula to measure the synergy of the vehicle and its parts: `(Total ATK * 3) + (Total DEF * 2) + (Total HP / 5)`.
*   **Pilot Resonance:** Training and neural synchronization to increase the Pilot's base stats.

### Phase 2: The Warp (Deployment)
*   **Navigation:** Players select coordinates on the Universe Map.
*   **Vehicle Selection:** Choosing the right archetype (VEHICLE, TANK, SHIP, etc.) and equipping modules for the mission.
*   **Cost:** Deployment consumes Fuel and O2 based on distance and sector danger level.

### Phase 3: The Void (Timeline Exploration)
*   **Timeline View:** Exploration is represented as a linear timeline with sequential nodes.
*   **Infiltration Layer:** Certain nodes or sub-paths allow the Pilot to **Eject** and use an **Exosuit** for stealth or precision tasks.
*   **Transformation:** Vehicles equipped with T-Modules can switch modes (e.g., Tank to Robot) to overcome environmental hazards.
*   **Dominance & Annihilation:** If the player's CP significantly exceeds the node's difficulty, they can choose to **Annihilate** threats, bypassing combat for instant rewards.
*   **Pre-Encounter Strategy (The Approach):** Before entering a node, players can choose an approach based on their Bastion's modules:
    *   **Passive Scan:** Low energy cost. Reveals basic node type.
    *   **Deep Analysis:** High energy cost. Reveals potential rewards and enemy stats.
    *   **Stealth Approach:** Increases success chance for [Stealth] options but consumes more Fuel.
*   **Node Types & Strategies:**
    *   **Standard Node:** Balanced mix of combat and loot. Strategy: *Steady Advance*.
    *   **Resource Node:** High chance of materials, low combat. Strategy: *Deep Scan* vs *Quick Scavenge*.
    *   **Combat Node:** High-intensity battles, rare equipment drops. Strategy: *Full Assault* vs *Tactical Flank*.
    *   **Anomaly Node:** High risk, high reward. Potential for "Void-Touched" items. Strategy: *Scientific Study* vs *Containment*.
    *   **Outpost Node:** Rare safe zones for trading and minor repairs. Strategy: *Rest & Refit*.
*   **Strategic Choice:** Branching paths requiring stat-based decisions.

### Phase 4: Extraction (Risk vs Reward)
*   **Safe Extraction:** Reaching a Warp Gate node secures all loot.
*   **Emergency Warp:** Immediate retreat with a 50% chance of cargo loss and high stress on vehicle durability.
*   **Total Failure:** If durability reaches zero, all loot is lost.
*   **Stage Change (V2O):** Players can choose to "Mint" their manifested assets (items in the database) to on-chain NFTs on Base L2.

---

## 2. Design Philosophy: Power vs. Strategy

The game's progression is built on a dual-track system that balances raw power with tactical depth. This is our core "Unique Selling Point" (USP) to prevent the "CP Pumping" trap.

### 2.1 The Power Fantasy: Combat Power (CP) & Annihilation
*   **CP as a Milestone:** Combat Power is the primary measure of a player's growth. It represents the synergy between Pilot Resonance, Vehicle Stats, and Equipment.
*   **Annihilation (ถล่มยับ):** When a player's CP significantly exceeds a node's difficulty, they unlock the ability to **Annihilate**. This allows them to bypass the encounter and claim rewards instantly.

### 2.2 The Tactical Counterweight: The "Anti-Pump" Mechanisms

#### 2.2.1 Suitability System (ความเหมาะสมของพื้นที่)
Every node and sub-sector is assigned a **Terrain Type**. Vehicles have a `SuitabilityRating` (0-100) for each terrain.

*   **Terrain Types:**
    *   **URBAN:** High density, tight corners. Favors MECH/EXOSUIT.
    *   **ISLANDS:** Water-heavy. Favors SHIP/SPEEDER.
    *   **SKY:** High altitude. Favors JET/SPEEDER.
    *   **DESERT:** Open, harsh. Favors TANK/MECH.
    *   **VOID:** Zero-G, neural interference. Favors specialized VOID-TYPE vehicles.
*   **The Penalty Logic:**
    *   `ECP = Base_CP * (Suitability_Rating / 100)`
    *   If `Suitability < 40%`, Fuel consumption increases by 2x.
    *   If `Suitability < 20%`, the vehicle suffers "Structural Stress" (HP decay per click).

#### 2.2.2 Infiltration & Detection (ด่านลอบเร้น)
Nodes have a `DetectionThreshold` (Signature Limit).

*   **Signature Calculation:**
    *   `Signature = (Vehicle_Size_Factor * 100) + (Current_CP / 10)`
    *   Tanks/Ships have high `Size_Factor` (3.0+). Pilots/Exosuits have low (0.5).
*   **Detection States:**
    *   **STEALTH:** `Signature < Threshold`. Normal encounter rates.
    *   **CAUTION:** `Signature` within 20% of `Threshold`. Enemy detection range increases.
    *   **ALARM:** `Signature > Threshold`. Triggers "Reinforcement Waves" (Enemies +200% Stats, 0% Loot drop).
*   **Tactical Choice:** Players can "Deploy Pilot" to enter a node in EVA mode to sabotage sensors, lowering the `Threshold` for the main Vehicle.

#### 2.2.3 Resonance & Synergy (เพดานความซิงโคร)
The bond between Pilot and Machine.

*   **Resonance Level (RL):** A value from 1 to 100, unique to each Pilot-Vehicle pair.
*   **The Cap Formula:**
    *   `Effective_Stats = Base_Stats * Min(1.0, (RL / Vehicle_Tier_Requirement))`
    *   A Tier 5 Vehicle might require RL 80 to function at 100% power.
*   **Resonance Growth:** Increases through "Sync Actions" (Perfect dodges, critical hits, and successful mission completions).

#### 2.2.4 Transformation Strategy (การเปลี่ยนร่าง)
Vehicles with the "Transformable" trait can switch modes mid-mission.

*   **Modes:** e.g., **CRUISER MODE** (High Speed, Low ATK) vs. **ASSAULT MODE** (Low Speed, High ATK).
*   **Transformation Cost:** Consumes 10% of Max Fuel and 1 turn.
*   **Strategic Use:** Switch to Cruiser mode to bypass a high-detection zone, then transform to Assault mode once the combat encounter is triggered.

## 3. Sector Archetypes & Gameplay Patterns

To support diverse gameplay, we use **Archetypes** that define the default behavior, hazards, and rewards for a Sector.

### 3.1 Archetype: The Iron Graveyard (ซากปรักหักพัง)
*   **Primary Terrain:** `VOID` or `URBAN` (Debris fields act as tight urban corridors).
*   **Gameplay Pattern:**
    *   **High Infiltration:** Many nodes have low `DetectionThreshold`. Players are encouraged to use **Pilot EVA** or **Exosuits** to navigate between wrecks.
    *   **Scavenger's Paradise:** High drop rates for `Scrap Metal` and `Neural Links`.
    *   **Hazard:** `Neural Static` (Stress +2 per click) from malfunctioning ancient AI cores.
*   **Best Fit:** Mechs with high maneuverability or Pilots with "Stealth Training".

### 3.2 Archetype: The Frontier System (ระบบดาวเคราะห์)
*   **Primary Terrain:** `DESERT`, `ISLANDS`, or `SKY`.
*   **Gameplay Pattern:**
    *   **Suitability Test:** Each planet in the system has a dominant terrain. Players must switch vehicles between sub-sectors.
    *   **Resource Extraction:** Primary source of `Fuel Isotopes` and `O2 Crystals`.
    *   **Hazard:** Environmental hazards like `Sandstorms` or `Turbulence`.
*   **Best Fit:** Heavy Vehicles (Tanks/Ships) that match the planet's terrain.

### 3.3 Archetype: The Void Rift (รอยแยกมิติ)
*   **Primary Terrain:** `VOID` or `SPACE`.
*   **Gameplay Pattern:**
    *   **Resonance Focus:** Nodes have high difficulty but grant massive `Resonance EXP`.
    *   **High Risk/High Reward:** Rare drops like `Void Shards` and `Ancient Tech`.
    *   **Hazard:** `Neural Static` is extreme. Pilot **Stress** management is the primary challenge.
*   **Best Fit:** High-tier Vehicles with legendary Pilots who have high Resonance potential.

## 4. Technical Implementation Pattern

### 4.1 Data Inheritance
1.  **Sector Level:** Defines the `Archetype` and `Default_Terrain`.
2.  **Sub-Sector Level:** Inherits from Sector but can override (e.g., a Desert Planet in a Frontier System).
3.  **Node Level:** Inherits `Terrain` and `Hazard` from Sub-Sector.

### 4.2 The "Transformation" Trigger
When a player encounters a Node with a `Terrain` that results in `Suitability < 40%`, the UI triggers a **"Tactical Recommendation"**:
*   "Current Vehicle unsuitable for [Terrain]. Transformation to [Mode B] recommended."
*   "Cost: 10 Fuel. ECP will increase by 40%."

---

## 5. Deep Durability System (DDS)

Items (Vehicles, Exosuits, Parts, Modules) have a granular durability system that affects both gameplay and visuals.

### Condition Thresholds
| Condition | Durability % | Gameplay Impact | Visual DNA Impact |
| :--- | :--- | :--- | :--- |
| **PRISTINE** | 80% - 100% | Full Stats | Clean, no effects |
| **WORN** | 50% - 79% | Standard Stats | Light scratches |
| **DAMAGED** | 20% - 49% | -10% Stats | Smoke Level: 0.1 - 0.5 |
| **CRITICAL** | 1% - 19% | -30% Stats, Speed -50% | Smoke: 1.0, Glitch: 0.5, Sparks |
| **BROKEN** | 0% | Item Inoperable | Heavy Glitch, Sparks, Fire |

---

## 4. Strategic Encounter Engine

Encounters are generated dynamically based on player stats, vehicle type, and current condition.

### Decision Logic
*   **Stat-Based Choices:** Options like [Stealth], [Assault], or [Bypass] use a combination of Pilot stats (Resonance), Vehicle class, and equipped Modules to calculate success probability.
*   **Archetype Synergy:** Certain choices are only available to specific archetypes (e.g., [Infiltrate] requires Pilot/Exosuit, [Heavy Siege] requires TANK).
*   **DDS Integration:** Damaged parts provide negative modifiers to relevant choices (e.g., a damaged sensor reduces Stealth success).

---

## 5. Dynamic Visual AI Generation

The system generates unique, consistent images for every encounter using AI (FLUX.1 + IP-Adapter).

### Visual Consistency Pillars
1.  **IP-Adapter (Identity):** Uses the original "Mint Image" of the vehicle or pilot as a reference to maintain the same look, colors, and silhouette across all generated scenes.
2.  **ControlNet (Structure):** Uses depth maps or edge detection to ensure the vehicle/pilot is in the correct pose (e.g., firing, retreating, transforming) and that equipped parts are in their correct slots.
3.  **Layered Prompting:**
    *   **Subject:** Vehicle/Pilot Class + Specific Equipment (from DB).
    *   **Condition:** Visual DNA values (Smoke, Glitch, Sparks).
    *   **Environment:** Scenario context (Void base, asteroid field).
    *   **Pilot:** Character features (Hair, eyes) visible in the cockpit or through the Exosuit visor.

---

## 6. Technical Architecture (Frontend)

The frontend is built as a decoupled game engine to handle high complexity and state-driven UI.

### Core Principles
*   **XState v5:** Manages the high-level game loop (Landing -> Bastion -> Map -> Exploration -> Combat -> Debrief).
*   **Singleton Systems:** Standalone logic classes (e.g., `ExplorationSystem`, `BastionSystem`, `CombatSystem`) that manage game state and API interactions.
*   **Event Bus:** A centralized, type-safe event system for cross-component communication.
*   **Tactical Noir HUD:** A high-fidelity, monochromatic UI with scanlines and glitch effects, designed to feel like a neural link.
*   **Framer Motion:** Used for all UI transitions and animations to ensure a premium, responsive feel.

---

## 5. Stage Change (NFT Digital Twin)

The bridge between the virtual game world and the blockchain.

*   **Virtual-to-NFT:** Players can "Stage Change" a high-value or "Battle-Hardened" vehicle/item into a tradeable NFT.
*   **Digital Twin Persistence:** The NFT retains the exact stats, durability, and **Visual DNA (Scars)** it had at the moment of minting.
*   **Metadata Storytelling:** The history of the item (encounters survived, damage taken) is recorded in the NFT's metadata, creating unique value for "veteran" items.

---

## 6. Technical Architecture (Backend)

### Database Schema: `items` Table
*   `id`: UUID
*   `owner_id`: UUID
*   `item_type`: VEHICLE, PART, BASTION_MODULE, EXOSUIT
*   `is_nft`: Boolean
*   `durability`: Integer (0-1000)
*   `condition`: Enum (PRISTINE...BROKEN)
*   `visual_dna`: JSONB (smoke_level, glitch_intensity, sparks_enabled)
*   `stats`: JSONB (HP, ATK, DEF, Energy)
*   `metadata`: JSONB (Prompt keywords, Season, Origin)

### API Endpoints
*   `GET /api/v1/items`: List player inventory with DDS status.
*   `POST /api/v1/exploration/encounter`: Generate a strategic choice node.
*   `POST /api/v1/exploration/resolve`: Process player choice and update DDS/Loot.
*   `POST /api/v1/items/repair`: Restore durability using resources.
