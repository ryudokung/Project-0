# Game Design Document: Project-0

**Author:** Samus Shepard (Game Designer)
**Date:** 2025-12-23
**Version:** 0.1 (Draft)

## 1. Executive Summary

### 1.1 Game Concept
Project-0 is a high-stakes, 1:1 scale Sci-Fi Exploration and Combat game. Players begin their journey by creating a unique **Pilot Character**, starting with a **Starter Ship** to navigate the void. As they progress, they operate modular Mechs, Aircraft, and Tanks, all supported by a massive Mothership. The game blends deep customization (Modular NFTs) with a high-fidelity 3D experience (WebGPU) and a creator-driven evolving universe.

### 1.2 Design Pillars (The "Samus" Touch)
- **Intense Narrative (The Core):** A deep, dark, and gripping story that drives every mission. You aren't just exploring; you're surviving a complex plot.
- **Social Prestige & Envy:** Vehicles are high-value assets. It's about having the rarest, most powerful Mech to dominate and show off to the community.
- **High-Stakes Exploration:** Risk vs. Reward is at the heart of every mission. Pilot death is meaningful and impacts your progression.
- **Visual Pride (The Hook):** AI-generated assets that look so cool players *must* show them off.
- **Evolving Universe:** A world that breathes and changes based on the Creator's "Context Patches."

---

## 2. Gameplay Systems

### 2.1 The Core Loop (The "Heartbeat")
The game utilizes a **Unified Game Controller** (Single Page Architecture) to manage transitions between states without page reloads.
1.  **Onboarding (Identity):** Create a Pilot Character (Name, Gender: Male/Female).
2.  **Hangar (Preparation):** Customize Ship/Mech, repair damage, and refuel.
3.  **Scanner (Risk Assessment):** Probe star systems for loot and threats.
4.  **Transit (The Journey):** Travel to the target, surviving atmospheric entry.
5.  **Combat/Salvage (The Action):** 1:1 Mech/Ship combat or Pilot EVA for rare tech.
6.  **Return (The Debrief):** Refine salvaged scrap and upgrade the Mothership.

### 2.2 Multi-Stage Exploration
- **Stage 1: Mothership (Strategic):** Managing fuel and scanner range.
- **Stage 2: Mech/Aircraft (Tactical):** Surface combat and aerial superiority.
- **Stage 3: Pilot EVA (Precision):** High-risk, high-reward salvage on foot.

---

## 3. Core Mechanics

### 3.1 Modular NFT Assembly (ERC-6551)
- Mechs are composed of **Chassis, Arms (Weapons), and Shields**.
- Each part is an individual NFT with AI-generated stats and visual traits.
- **Visual Synthesis:** AI combines all equipped parts into a single, cohesive 3D/2D representation.

### 3.2 Combat Mechanics
- **Asynchronous Auto-Resolve with 3D Visualization:** Strategic depth meets cinematic flair.
- **Real Combat Integration:** Combat encounters are now driven by real backend data. When a player enters a `COMBAT` node during exploration, the system identifies a specific `enemy_id` linked to a seeded NPC Mech.
- **NPC Enemy Seeding:** The backend maintains a pool of NPC Mechs (e.g., "Iron Syndicate Striker", "Void Guardian") with unique stats (HP, Attack, Defense, Speed).
- **Elemental Triangle:** Kinetic > Energy > Explosive > Kinetic.
- **Visual Wear & Tear:** Real-time damage visualization on specific NFT parts.

### 3.3 Pilot EVA & Death
- **EVA Mode:** Pilot exits the Mech to enter narrow spaces.
- **Death Penalty:** Loss of all loot from the current run + critical Mech damage. This creates the "Sweaty Palms" feeling.

### 3.4 Character Creation & Identity
- **Pilot Registration:** New players must register their first character before entering the game world.
- **Customization Options:**
    - **Name:** Unique callsign for the pilot.
    - **Gender:** Male / Female / Non-binary.
    - **Appearance:** Selection of Face types and Hair styles.
- **Starter Asset:** Upon completion, the character is granted a **Starter Ship** (Type: Ship) and basic **Pilot Stats**.
- **Character Instances:** A single user account can host multiple characters (acquired via Gacha or progression).

---

## 4. Progression Systems

### 4.1 Pilot Rank & Mastery
- **XP:** Earned from successful missions and star discoveries.
- **Licenses:** Unlock higher-tier vehicles (Motherships, Mechs).
- **Skill Tree:** Focus on Agility, O2 Efficiency, or Weapon Mastery.

### 4.2 Mech Sync Rate
- Rewards players for sticking with their favorite gear.
- High Sync = Passive buffs (Accuracy, Evasion).

### 4.3 Mothership Tech Tree
- Upgrade Hangar (Repair speed), Lab (Crafting quality), and Scanner (Range).

---

## 5. UI/UX Design (Pro Max Standards)
- **Unified Game Controller:** A single-page architecture that manages `GameStage` (HANGAR, MAP, EXPLORATION, COMBAT) with fluid Framer Motion transitions.
- **HUD/Cockpit:** Glassmorphism for an immersive "Inside the Helmet" feel.
- **Dashboard:** Bento Grid layout for clear resource management.
- **Visual DNA:** Seasonal themes (e.g., Cyber-Samurai) that define the aesthetic.

---

## 6. Technical Requirements (Cloud Dragonborn's Domain)
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
