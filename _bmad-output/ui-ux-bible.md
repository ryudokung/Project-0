# UI/UX & Immersive HUD Bible: Project-0 (The Experience)

**Goal:** Define the visual language, HUD elements, and user journey to ensure maximum immersion and "Flex" potential.

## 1. Visual Language: "Tactical Noir"
- **Color Palette:** Deep Blacks, Neon Greens (HUD), Warning Oranges, and Industrial Grays.
- **Typography:** Monospaced fonts (for that "Terminal" feel) and bold, italicized headers for impact.
- **Effects:** Scanlines, slight chromatic aberration, and HUD glitches when taking damage.
- **Dynamic Background:** A global 3D space environment rendered with R3F, featuring a rotating Earth-like planet, floating space debris, and a deep starfield to maintain immersion across all application states.

## 2. The Immersive Cockpit & Bridge HUD (3D)
Rendered using **WebGPU + R3F**, the HUD is part of the 3D world:
- **The Bridge View:** The primary hub interface. A wide-angle viewport showing the Bastion's exterior and the current star system.
- **Center:** Crosshair and Target Lock-on.
- **Left Panel:** System Integrity (HP of each modular part) with **DDS Condition Indicators**.
- **Right Panel:** Radar (3D sphere showing enemy signatures).
- **Bottom:** O2 Levels (for EVA) and Fuel Gauge (for The Bastion).
- **Side Monitors:** Real-time AI Narrative logs and Pilot status.

## 3. The Showcase Engine (The "Flex" View)
The Bastion is the primary social and management space:
- **Visual Equipment Map:** A silhouette-based interface for mapping modules to anatomical slots (HEAD, CORE, ARMS, LEGS).
- **Dynamic Lighting:** High-fidelity shadows and reflections to make the AI-generated textures pop.
- **Photo Mode:** A dedicated UI to take high-res snapshots of the vehicle with custom filters.
- **Social Sharing:** One-click sharing to X (Twitter) or Discord with the vehicle's stats, CP, and rarity.

## 4. The Narrative Timeline: "Expedition and Encounters" UX
The exploration interface is designed around the "Expedition and Encounters" model, focusing on the feeling of a journey unfolding in real-time:

### 4.1 The Timeline Tracker (The Expedition)
- **Visual:** A glowing, data-driven line that runs across the top of the HUD.
- **Nodes:** As the player clicks "Proceed," the current node is highlighted, and the next node is revealed.
- **Color Coding:** 
    - **Pink Nodes:** Combat encounters.
    - **Blue Nodes:** Resource/Salvage points.
    - **White Nodes:** Standard/Narrative fragments.
    - **Yellow Nodes:** Outposts/Safe zones.

### 4.2 The Visual DNA Synthesis (AI Image Display)
- **Presentation:** The AI-generated image (based on Visual DNA) appears as a "Main Viewport" in the center of the screen.
- **Transition:** When a new node is reached, the image "synthesizes" into view with a scanline effect.
- **Metadata Overlay:** Small, monospaced text at the corner of the image showing the "DNA Keywords" used to generate it (e.g., `VISUAL DNA SYNTHESIS: BRUTALIST | INDUSTRIAL_STEEL`).

### 4.3 Resource-Driven Interaction
- **The "Proceed" Button:** The primary interaction. It appears once a node is resolved (e.g., after combat or a choice).
- **DDS Alerts:** When an item enters "Damaged" or "Critical" status, the HUD triggers visual glitches, flickering lights, and warning sirens.

## 5. User Journey: The "Single-Page Game Loop"

### 5.1 Onboarding: "The Operative Recruitment"
- **Login:** Minimalist screen with "Enter Project-0" via Google/Email. No wallet popup.
- **Character Creation:** Immediate transition to the **Pilot Registration** screen.
- **Selection:** Visual choice between **Male** or **Female** pilot using high-fidelity AI-generated portraits.
- **UI Element:** Interactive cards that highlight on selection, with a grayscale-to-color transition effect.
- **Transition:** Cinematic zoom-out from the pilot's profile to the full Bastion view.
- **First Contact:** After registration, the player enters the Bastion where their **Starter Vehicle** is waiting.
- **Late Binding:** A subtle "Link External Wallet" button in the Bastion settings or profile, framed as "Securing Assets to the Void-Chain."

### 5.2 The Unified Play Page (The Loop)
The entire game experience is contained within a single route (`/`), managed by a **Decoupled Systems Architecture**:
- **XState v5 Orchestration:** Controls the high-level game stages (Landing -> Bastion -> Map -> Exploration -> Combat -> Debrief).
- **EventBus Communication:** UI components are "dumb" views that subscribe to events from the **ExplorationSystem**, **CombatSystem**, and **BastionSystem**.
- **Seamless Transitions:** Using Framer Motion `AnimatePresence`, the UI morphs between stages without page reloads.
- **Persistent HUD:** The Pilot ID Badge, O2 levels, and Fuel reserves remain visible or transition smoothly between stages, maintaining the "Cockpit" feel.
- **Stage Flow:**
    1. **The Bastion:** Manage Vehicle, Equipment (Visual Map), and Repairs.
    2. **Universe Map:** Select Sector.
    3. **Location Scan:** Identify POIs.
    4. **Exploration Loop:** Advance through nodes on the timeline.
    5. **Combat:** Tactical engagement.
    6. **Debrief:** Return to The Bastion.

### 5.3 The Pilot ID Badge (Hangar UI)
- **Visual:** A glassmorphic "ID Card" displayed in the Hangar.
- **Details:** Shows Character Name, Sync Rate, Gender, and a portrait based on the selected appearance.
- **Functionality:** Serves as the primary entry point for Character Stats and Skill progression.

## 6. Rarity & Tier Visuals (The "Flex" Hierarchy)
To distinguish between items and allow players to "show off," the UI uses a strict color and effect hierarchy:

### 6.1 Rarity Color Palette
- **Standard (White):** `#FFFFFF` - Clean, industrial look.
- **Refined (Green):** `#4ADE80` - Stable, functional glow.
- **Prototype (Blue):** `#60A5FA` - Experimental, pulsing energy.
- **Relic (Purple):** `#A855F7` - Ancient, deep resonance. High-fidelity textures.
- **Singularity (Gold):** `#FBBF24` - God-tier. Includes "Void Glow" particle effects and animated UI borders.

### 6.2 Premium vs. Standard Distinction
- **Void-Touched (Premium):** Items obtained via Gacha or Seasonal Wormholes have a **"Void Signature"**—a subtle, dark smoke or glitch effect that surrounds the item's icon and 3D model.
- **Mintable Badge:** A small, glowing "M" icon on the item card indicates it can be bridged to the blockchain.

## 7. The Gacha UI: "Void Signals"
The Gacha experience is designed to be high-tension and visually rewarding:
- **The Pull:** The player activates a "Signal Decoder." The screen goes dark, and a 3D radar sweep begins.
- **The Reveal:** A beam of light (colored by rarity) shoots from the center. 
- **The "Gasp" Moment:** For **Relic** or **Singularity** items, a full-screen splash art (AI-generated) is revealed with a dramatic sound effect and a "DNA Sequence" animation.
- **Pity Tracker:** A subtle "Signal Strength" bar at the bottom shows how close the player is to a guaranteed high-tier drop.

## 9. The Bastion Systems HUD
The Bastion interface must clearly distinguish between the two independent engineering paths:

### 9.1 Teleport Interface (Dimensional)
- **Visuals:** Distorted, "glitchy" blue/purple UI elements.
- **Metrics:** Stability %, Energy Drain, Gravity Interference.
- **Feedback:** A "Dimensional Tear" animation when jumping.

### 9.2 Atmospheric Entry Interface (Structural)
- **Visuals:** Solid, industrial orange/red UI elements.
- **Metrics:** Hull Temperature, G-Force, Structural Integrity.
- **Feedback:** Heat haze and vibration effects during descent.

### 9.3 Engineering Profile
Ships are categorized by their "Engineering Bias":
- **Warp-Bias:** High Teleport stability, fragile hull (Risky landing).
- **Entry-Bias:** Reinforced hull (Safe landing), no Teleport capacity.
- **Hybrid (End-game):** Balanced metrics for both systems.
