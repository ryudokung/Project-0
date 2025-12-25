# Bastion Bible: Project-0 (The Eternal Command Center)

**Goal:** Define the progression, modular components, and NFT integration of the player's mobile base, "The Bastion."

## 1. The Bastion: The Command Hub
The Bastion is the player's primary hub, experienced through the **Bridge View**. It is a living entity that requires maintenance and upgrades to survive the deep void.

### 1.1 Modular Hardpoints (The Slots)
The Bastion has specific slots where components can be installed:
- **Warp Drive Slot:** Determines travel distance and speed.
- **Shield Projector Slot:** Provides defense during atmospheric entry and space encounters.
- **Weapon Hardpoints (x2):** Defensive turrets visible from the Bridge.
- **Utility Slot:** Scanners, Labs, or Resource Refineries.

### 1.2 Bastion Facilities (The Tech Tree)
Upgrading facilities unlocks new gameplay mechanics:

| Facility | Level 1 (Starter) | Level 5 (Advanced) | Level 10 (Elite) |
| :--- | :--- | :--- | :--- |
| **Bastion Ops** | Basic Repairs (Slow) | Auto-Repair (Fast) | Modular Refit (Instant) |
| **Scanner/Radar** | **Broken (Requires Repair)** | Deep Scan (Sector) | Void Probe (Inter-Sector) |
| **Lab** | Basic Scrap Refining | Core Synthesis | Blueprint Overclocking |
| **Bridge** | Manual Navigation | AI Autopilot | Tactical Command (Fleet) |

## 2. Deployment Technologies (The "Entry" Choice)
How the player reaches the surface determines the risk and cost of the mission.

### 2.1 Atmospheric Entry (Standard)
- **Mechanic:** The vehicle is launched in a drop-pod or flies down manually.
- **Requirement:** Vehicle must have high **Heat Shielding** and **Armor**.
- **Risk:** High. Random "Entry Events" (Turbulence, AA Fire, Heat Damage).
- **Cost:** Low (Fuel only).

### 2.2 Quantum Teleportation (Advanced Upgrade)
- **Mechanic:** Instant deployment of the Pilot or a Light Vehicle (Speeder/Exosuit) to the surface.
- **Requirement:** Bastion Upgrade: **"Quantum Gate" (Level 8 Bridge)**.
- **Risk:** Zero (Bypasses the atmospheric entry phase).
- **Cost:** High (Requires "Core Fragments" or massive Energy consumption).
- **Limitation:** Cannot teleport Heavy Vehicles (Tanks/Heavy Mechs) until Level 10.

## 3. The Maintenance Loop: Repair & Refit
Items in Project-0 (Virtual or NFT) are tools that wear down with use. The Bastion provides the facilities to keep them operational.

### 3.1 Durability & Repair
- **Usage:** Every mission, warp, or combat encounter reduces the durability of installed components.
- **The Repair Station:** Located in the Bastion Ops. Players use gathered resources (Scrap, Energy) to restore durability.
- **NFT Status:** Minting an item to NFT doesn't make it indestructible; it just changes its "Stage" to be tradeable and verified. It still needs to be repaired like any other tool.

## 4. Item Lifecycle: The "Stage Change" Model
1. **Discovery:** Explore planets to find Blueprints and raw materials.
2. **Crafting:** Create a Virtual Component in the Lab. Stats are randomized.
3. **Maintenance:** Use the item, repair it when low, and keep it in top shape.
4. **Minting (Stage Change):** If the item is valuable, the player "Mints" it to show it off on the blockchain or sell it. It remains fully functional and repairable in-game.

## 5. Bridge Immersion
- **Visual DNA:** The Bridge reflects the current state of the Bastion.
- **Maintenance Alerts:** When a component needs repair, the Bridge UI shows warning lights or holographic alerts.
- **The Viewport:** A window into the current star system, showing the Bastion's active weapons and shields.

## 6. Resource Management (Click-Based)
- **Fuel:** Required for all travel and node transitions.
- **Energy:** Required for Scanning and Teleportation.
- **O2 Reserves:** Refills Pilot EVA tanks.
- **Radar Status:** If Radar is broken (Level 0), no mission nodes are visible. Repairing it is the first objective.

## 7. Technical Implementation
- **Unified Database:** A single `items` table tracks all components, including their `current_durability` and `max_durability`.
- **Repair API:** `POST /api/v1/bastion/repair` - Consumes resources to reset an item's durability.
- **Ownership Sync:** If an NFT is sold, the database record (including its current durability state) is transferred to the new owner.
