# 🎮 Project-0: Gameplay & Technical Specification

This document outlines the core gameplay loop, technical architecture for dynamic systems, and the integration of AI-driven visual consistency.

---

## 1. Core Gameplay Loop: The Extraction Cycle

The game follows a high-stakes "Extraction RPG" loop centered around **The Bastion** and **The Void**.

### Phase 1: The Bastion (Strategic Hub)
*   **Energy Management:** Players allocate the Bastion's limited power core between systems (Shields, Warp Drive, Auto-Repair).
*   **Modular Hardpoints:** The Bastion has slots for modules (Scanners, Scavenger Arms, Defense Turrets) that unlock new exploration capabilities.
*   **Bastion Operations:** Maintenance of Vehicles and Pilot Gear. Repairing items consumes resources gathered from The Void.
*   **Pilot Resonance:** Training and neural synchronization to increase the Pilot's base Combat Rating (CR).

### Phase 2: The Warp (Deployment)
*   **Navigation:** Players select coordinates on the Universe Map.
*   **Vehicle Selection:** Choosing the right archetype (MECH, TANK, SHIP, etc.) and equipping Transformation Modules for the mission.
*   **Cost:** Deployment consumes Fuel and O2 based on distance and sector danger level.

### Phase 3: The Void (Timeline Exploration)
*   **Timeline View:** Exploration is represented as a linear timeline with sequential nodes.
*   **Infiltration Layer:** Certain nodes or sub-paths allow the Pilot to **Eject** and use an **Exosuit** for stealth or precision tasks.
*   **Transformation:** Vehicles equipped with T-Modules can switch modes (e.g., Tank to Robot) to overcome environmental hazards.
*   **Dominance & Annihilation:** If the player's CR significantly exceeds the node's difficulty, they can choose to **Annihilate** threats, bypassing combat for instant rewards.
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

---

## 2. Design Philosophy: Power vs. Strategy

The game's progression is built on a dual-track system that balances raw power with tactical depth. This is our core "Unique Selling Point" (USP).

### The Power Fantasy: Combat Rating (CR) & Annihilation
*   **CR as a Milestone:** Combat Rating is the primary measure of a player's growth. It represents the synergy between Pilot Resonance, Vehicle Stats, and Equipment.
*   **Annihilation (ถล่มยับ):** When a player's CR significantly exceeds a node's difficulty, they unlock the ability to **Annihilate**. This allows them to bypass the encounter and claim rewards instantly.
*   **Respecting Time:** This mechanic rewards high-level players by making "farming" efficient and satisfying, reinforcing the feeling of becoming a "Legend of the Void."

### The Tactical Counterweight: Suitability & Infiltration
*   **Horizontal Progression:** Power alone cannot solve every problem. The **Suitability System** ensures that a high-CR Tank cannot easily dominate a high-altitude or aquatic mission. Players must build a diverse "Fleet" of vehicles.
*   **Infiltration Layer:** High-CR vehicles are often "loud" and easily detected. Stealth-focused missions or restricted areas require players to switch to **Pilot/Exosuit** modes, where strategy and skill outweigh raw stats.
*   **Transformation Strategy:** The ability to **Transform** (e.g., Tank to Robot) provides mid-mission adaptability. Choosing the right form for the right terrain or enemy type can overcome a CR deficit.

---

## 3. Deep Durability System (DDS)

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

## 3. Strategic Encounter Engine

Encounters are generated dynamically based on player stats, vehicle type, and current condition.

### Decision Logic
*   **Stat-Based Choices:** Options like [Stealth], [Assault], or [Bypass] use a combination of Pilot stats (Resonance), Vehicle class, and equipped Modules to calculate success probability.
*   **Archetype Synergy:** Certain choices are only available to specific archetypes (e.g., [Infiltrate] requires Pilot/Exosuit, [Heavy Siege] requires TANK).
*   **DDS Integration:** Damaged parts provide negative modifiers to relevant choices (e.g., a damaged sensor reduces Stealth success).

---

## 4. Dynamic Visual AI Generation

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
