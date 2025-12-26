# Bastion Bible: Project-0 (The Eternal Command Center)

**Goal:** Define the progression, modular components, and NFT integration of the player's mobile base, "The Bastion."

## 1. The Bastion: The Command Hub
The Bastion is the player's primary hub, experienced through the **Bridge View**. It is a living entity that requires maintenance and upgrades to survive the deep void.

### 1.1 Visual Equipment Map (Anatomical Slots)
The Bastion features a strategic interface for managing vehicle equipment. Players map items to specific anatomical slots on a vehicle's silhouette:
- **HEAD:** Sensors, Neural Interfaces, and Targeting Systems.
- **CORE:** Power Source, Armor Plating, and Life Support.
- **ARM_L / ARM_R:** Weapons, Shields, and Utility Tools.
- **LEGS:** Mobility Systems, Thrusters, and Stability Anchors.

### 1.2 Bastion Facilities (The Tech Tree)
Upgrading facilities unlocks new gameplay mechanics and global bonuses:

| Facility | Level 1 (Starter) | Level 5 (Advanced) | Level 10 (Elite) |
| :--- | :--- | :--- | :--- |
| **Radar** | -20% Detection Threshold | -100% Detection Threshold | -200% Detection Threshold |
| **Lab** | +10% Reward / +15% XP | +50% Reward / +75% XP | +100% Reward / +150% XP |
| **Warp Drive** | -10% Fuel Cost | -50% Fuel Cost | -100% Fuel Cost |

## 2. Bastion Modular System (Global Support Layer)
The Bastion is not just a static hub; it is a modular entity. Players can upgrade **Bastion Modules** to provide global buffs and unlock advanced features across all game systems.

### 2.1 Systemic Integration & Bonuses
| Module | Implementation | Core Bonus |
| :--- | :--- | :--- |
| **Warp Drive** | `ExplorationService` | Reduces Fuel cost per node transition (10% reduction per level). |
| **Radar** | `TimelineGenerator` | Reduces **Detection Threshold** (20% reduction per level), making stealth easier. |
| **Lab** | `RewardSystem` | Increases **Scrap Metal** (+10%/lvl) and **XP** (+15%/lvl) gains from all nodes. |

### 2.2 Layered Equipment & Set Synergy
The Bastion manages the "Triple-Layer" gear system:
1.  **Pilot Layer (Exosuit):** Worn by the pilot. Provides base CP and unique passives.
2.  **Vehicle Layer (Chassis/Parts):** The primary combat unit.
3.  **Bastion Layer (Modules):** Global buffs affecting the entire expedition.

**Set Synergy Bonus:**
When the **Exosuit** and **Vehicle** belong to the same **Series** (e.g., *Void-Walker*), the system grants a **+15% ECP (Effective Combat Power) bonus**.

### 2.3 Emergency Retrieval Protocol (Fail-Safe)
A system that auto-warps the pilot back to the Bastion when Fuel or O2 reaches 0.
- **Penalty:** +50 Stress, 50% Reward loss, and a "Critical Fatigue" flag (reduces ECP by 50% in the next mission).

### 2.4 Neural Overdrive (Active Skills)
Tactical skills powered by **Neural Energy (NE)** (Max 100, +10 per node).
- **Overclock:** +30% ECP for 1 node (Cost: 50 NE).
- **Emergency Repair:** Restore 30% Vehicle HP (Cost: 40 NE).

### 2.5 Module Synergy & Energy Management
- **Energy Allocation:** Each equipped module consumes a portion of the Bastion's **Core Energy**. Players must balance which modules are active based on their current mission goals.
- **Fleet-Wide Buffs:** Unlike Vehicle parts, Bastion Modules provide **Global Passives** that affect every Pilot and Vehicle deployed from that Bastion.
- **Architecture Support:** These modules interact directly with the backend services (e.g., the `ExplorationService` checks for Radar modules to determine node visibility).

## 3. Deployment Technologies (The "Entry" Choice)

## 3. The Maintenance Loop: Repair & Refit
Items in Project-0 (Virtual or NFT) are tools that wear down with use. The Bastion provides the facilities to keep them operational.

### 3.1 Durability & Repair
- **Usage:** Every mission, warp, or combat encounter reduces the durability of installed components.
- **The Repair Station:** Located in the Bastion Ops. Players use gathered resources (Scrap, Energy) to restore durability.
- **NFT Status:** Minting an item to NFT doesn't make it indestructible; it just changes its "Stage" to be tradeable and verified. It still needs to be repaired like any other tool.

## 4. Item Lifecycle: The "Stage Change" Model (V2O)
1. **Discovery:** Explore planets to find Blueprints and raw materials.
2. **Crafting:** Create a "Manifested Asset" (Virtual Component) in the Lab. Stats are randomized.
3. **Maintenance:** Use the item, repair it when low, and keep it in top shape.
4. **Minting (Stage Change):** If the item is valuable, the player "Mints" it to the blockchain as an NFT. It remains fully functional and repairable in-game.
    - **Minting Rules:** Epic+ rarity, 10+ Expeditions completed, and >80% Durability.

## 5. Bridge Immersion
- **Visual DNA:** The Bridge reflects the current state of the Bastion.
- **Maintenance Alerts:** When a component needs repair, the Bridge UI shows warning lights or holographic alerts.
- **The Viewport:** A window into the current star system, showing the Bastion's active weapons and shields.

## 6. Resource Management (Click-Based)
- **Fuel:** Required for all travel and node transitions.
- **Energy:** Required for Scanning and Teleportation.
- **O2 Reserves:** Refills Pilot EVA tanks.
- **Neural Energy (NE):** Required for Active Skills.
- **Radar Status:** If Radar is broken (Level 0), no mission nodes are visible. Repairing it is the first objective.

## 7. Technical Implementation
- **Unified Database:** A single `items` table tracks all components, including their `current_durability` and `max_durability`.
- **Repair API:** `POST /api/v1/bastion/repair` - Consumes resources to reset an item's durability.
- **Ownership Sync:** If an NFT is sold, the database record (including its current durability state) is transferred to the new owner.
- **Metadata Storage:** Bastion Module levels are stored in the `pilot_stats.metadata` JSONB field.
