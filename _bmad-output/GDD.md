# Game Design Document: Project-0

**Author:** Samus Shepard (Game Designer)
**Date:** 2025-12-23
**Version:** 0.1 (Draft)

## 1. Executive Summary

### 1.1 Game Concept
Project-0 is a high-stakes, 1:1 scale Sci-Fi Exploration and Combat game. Players begin their journey by creating a unique **Pilot Character**, starting with a **Starter Pack** to navigate the void. As they progress, they operate modular **Vehicles** (Mechs, Aircraft, Tanks), all supported by **The Bastion**, their mobile command center. The game blends deep customization (Modular NFTs) with a high-fidelity 3D experience (WebGPU) and a creator-driven evolving universe.

### 1.2 Design Pillars (The "Samus" Touch)
- **Intense Narrative (The Core):** A deep, dark, and gripping story that drives every mission. You aren't just exploring; you're surviving a complex plot.
- **Social Prestige & Envy:** Vehicles are high-value assets. It's about having the rarest, most powerful Vehicle to dominate and show off to the community.
- **High-Stakes Exploration:** Risk vs. Reward is at the heart of every mission. Pilot death is meaningful and impacts your progression.
- **Visual Pride (The Hook):** AI-generated assets that look so cool players *must* show them off.
- **Evolving Universe:** A world that breathes and changes based on the Creator's "Context Patches."

---

## 2. Gameplay Systems

### 2.1 The Core Loop (The "Heartbeat")
The game utilizes a **Decoupled Systems Architecture** (EventBus + Singletons) to manage transitions between states without page reloads.
1.  **Onboarding (Identity):** Create a Pilot Character (Name, Gender: Male/Female).
2.  **The Bastion (Bridge View):** The primary hub for navigation, maintenance, and fleet management.
3.  **Preparation:** Customize Vehicle, equip Modules via the **Visual Equipment Map**, repair damage using the **Deep Durability System (DDS)**, and refuel.
4.  **Scanner (Risk Assessment):** Probe star systems for loot and threats.
5.  **Transit (The Journey):** Travel to the target, surviving atmospheric entry.
6.  **Combat/Salvage (The Action):** 1:1 Vehicle combat or Pilot EVA for rare tech.
7.  **Return (The Debrief):** Refine salvaged scrap and upgrade The Bastion.

### 2.2 Multi-Stage Exploration
- **Stage 1: The Bastion (Strategic):** Managing fuel, scanner range, and modular hardpoints.
- **Stage 2: Vehicle (Tactical):** Surface combat and aerial superiority.
- **Stage 3: Pilot EVA (Precision):** High-risk, high-reward salvage on foot.

---

## 3. Core Mechanics

### 3.1 Modular NFT Assembly (ERC-6551) & Stage Change Model
- Vehicles are composed of **Chassis, Modules (HEAD, CORE, ARM_L, ARM_R, LEGS)**.
- **Stage Change:** Items start as **Manifested Assets** in the database. Players can "Mint" them to become NFTs, making them tradeable while retaining all in-game stats and durability.
- Each part is an individual NFT with AI-generated stats and visual traits.
- **Visual Synthesis:** AI combines all equipped parts into a single, cohesive 3D/2D representation.

### 3.2 Deep Durability System (DDS)
- Items have 5 condition thresholds: **Pristine, Worn, Damaged, Critical, Broken**.
- **Impact:** Lower durability reduces stats (Attack/Defense) and introduces visual glitches/smoke in the UI.
- **Maintenance:** Players must use Scrap and Energy at The Bastion's Repair Station to restore items.

### 3.3 Combat Mechanics
- **Asynchronous Auto-Resolve with 3D Visualization:** Strategic depth meets cinematic flair.
- **Combat Power (CP):** A weighted sum of stats: `(ATK*3) + (DEF*2) + (HP/5)`.
- **Real Combat Integration:** Combat encounters are now driven by real backend data. When a player enters a `COMBAT` node during exploration, the system identifies a specific `enemy_id` linked to a seeded NPC Vehicle.
- **NPC Enemy Seeding:** The backend maintains a pool of NPC Vehicles (e.g., "Iron Syndicate Striker", "Void Guardian") with unique stats (HP, Attack, Defense, Speed).
- **Elemental Triangle:** Kinetic > Energy > Explosive > Kinetic.
- **Visual Wear & Tear:** Real-time damage visualization on specific anatomical slots.

### 3.3 Pilot EVA & Death
- **EVA Mode:** Pilot exits the Vehicle to enter narrow spaces.
- **Death Penalty:** Loss of all loot from the current run + critical Vehicle damage. This creates the "Sweaty Palms" feeling.

### 3.4 Character Creation & Identity
- **Pilot Registration:** New players must register their first character before entering the game world.
- **Customization Options:**
    - **Name:** Unique callsign for the pilot.
    - **Gender:** Male / Female / Non-binary.
    - **Appearance:** Selection of Face types and Hair styles.
- **Starter Asset:** Upon completion, the character is granted a **Starter Pack** (Ship + Modules) and basic **Pilot Stats**.
- **Character Instances:** A single user account can host multiple characters (acquired via Gacha or progression).

---

## 4. Progression Systems

### 4.1 Pilot Rank & Mastery
- **XP:** Earned from successful missions and star discoveries.
- **Licenses:** Unlock higher-tier vehicles (Motherships, Vehicles).
- **Skill Tree:** Focus on Agility, O2 Efficiency, or Weapon Mastery.

### 4.2 Vehicle Sync Rate
- Rewards players for sticking with their favorite gear.
- High Sync = Passive buffs (Accuracy, Evasion).

### 4.3 The Bastion Tech Tree
- Upgrade Preparation Area (Repair speed), Lab (Crafting quality), and Scanner (Range).
- Install modular components into Bastion Hardpoints (Warp Drive, Shields, Turrets).

---

## 5. UI/UX Design (Pro Max Standards)
- **Decoupled UI Layer:** React components subscribe to a global **EventBus**, ensuring a responsive and "Unity-like" experience.
- **Bridge View:** An immersive 3D/2D viewport showing the Bastion's current location and status.
- **HUD/Cockpit:** Glassmorphism for an immersive "Inside the Helmet" feel.
- **Dashboard:** Bento Grid layout for clear resource management.
- **Visual DNA:** Seasonal themes (e.g., Cyber-Samurai) that define the aesthetic.

---

## 6. Technical Requirements (Cloud Dragonborn's Domain)
- **Architecture:** Decoupled Systems (EventBus + Singleton Systems).
- **Engine:** WebGPU + React Three Fiber for seamless 3D.
- **Backend:** Go (Modular Monolith) with Saga Pattern for Web3 consistency.
- **AI:** FLUX.1 for visual synthesis + MCP for real-time narrative.
- **Blockchain:** Base L2 with ERC-6551.

---

## 7. Narrative & Lore
- **Intense Storytelling:** The universe is filled with dark secrets, political intrigue, and high-stakes drama. Every "Context Patch" expands a gripping narrative that keeps players on the edge of their seats.
- **Social Validation:** The game is built for "Flexing." Your profile, your hangar, and your battle history are public, designed to trigger social envy and prestige.
- **The Evolving Universe:** Lore is delivered via "Context Patches" that change the world state permanently.
- **Hall of Fame:** Immortalizing players who make "First Discoveries" or dominate the seasonal leaderboards.

---
*Drafted by Samus Shepard - Let's GOOO!*
