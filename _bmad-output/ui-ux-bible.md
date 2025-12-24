# UI/UX & Immersive HUD Bible: Project-0 (The Experience)

**Goal:** Define the visual language, HUD elements, and user journey to ensure maximum immersion and "Flex" potential.

## 1. Visual Language: "Tactical Noir"
- **Color Palette:** Deep Blacks, Neon Greens (HUD), Warning Oranges, and Industrial Grays.
- **Typography:** Monospaced fonts (for that "Terminal" feel) and bold, italicized headers for impact.
- **Effects:** Scanlines, slight chromatic aberration, and HUD glitches when taking damage.
- **Dynamic Background:** A global 3D space environment rendered with R3F, featuring a rotating Earth-like planet, floating space debris, and a deep starfield to maintain immersion across all application states.

## 2. The Immersive Cockpit HUD (3D)
Rendered using **WebGPU + R3F**, the HUD is part of the 3D world:
- **Center:** Crosshair and Target Lock-on.
- **Left Panel:** System Integrity (HP of each modular part).
- **Right Panel:** Radar (3D sphere showing enemy signatures).
- **Bottom:** O2 Levels (for EVA) and Fuel Gauge (for Mothership).
- **Side Monitors:** Real-time AI Narrative logs and Pilot status.

## 3. The Showcase Engine (The "Flex" View)
The Hangar is the primary social space:
- **Dynamic Lighting:** High-fidelity shadows and reflections to make the AI-generated textures pop.
- **Photo Mode:** A dedicated UI to take high-res snapshots of the vehicle with custom filters.
- **Social Sharing:** One-click sharing to X (Twitter) or Discord with the vehicle's stats and rarity.

## 4. The Narrative Timeline: "Expedition and Encounters" UX
The exploration interface is designed around the "Expedition and Encounters" model, focusing on the feeling of a journey unfolding in real-time:

### 4.1 The Timeline String (The Expedition)
- **Visual:** A glowing, data-driven line that runs across the top or side of the HUD.
- **Encounters:** As the player clicks "Advance," new nodes (Encounters) appear on the string.
- **Color Coding:** 
    - **Red Encounters:** Combat encounters.
    - **Blue Encounters:** Resource/Salvage points.
    - **White Encounters:** Narrative/Lore fragments.
    - **Gold Encounters:** Major Anchors (Fixed plot points).

### 4.2 The Visual Reveal (AI Image Display)
- **Presentation:** The AI-generated image (based on Visual DNA) appears as a "Main Viewport" or a "Tactical Feed" in the center of the cockpit.
- **Transition:** When a new Encounter is generated, the image "glitches" or "scans" into view, emphasizing the AI-generated nature of the world.
- **Metadata Overlay:** Small, monospaced text at the corner of the image showing the "DNA Keywords" used to generate it (e.g., `FACTION: IRON_SYNDICATE | STYLE: BRUTALIST`).

### 4.3 Resource-Driven Interaction
- **The "Advance" Button:** The primary interaction. It displays the O2/Fuel cost for the next step.
- **Dynamic Warnings:** If O2 is low, the HUD turns red, and the "Advance" button text changes to "DESPERATE MOVE," signaling that the AI is now more likely to generate high-stakes or resource-focused Encounters.

## 5. User Journey: The "Click-to-Explore" Loop

### 5.1 Onboarding: "The Operative Recruitment"
- **Login:** Minimalist screen with "Enter Project-0" via Google/Email. No wallet popup.
- **Character Creation:** Immediate transition to the **Pilot Registration** screen.
    - **UI Style:** Bento Grid layout for selecting Gender, Face, and Hair.
    - **Visuals:** Real-time preview of the character's appearance.
- **First Contact:** After registration, the player enters the Hangar where their **Starter Ship** is waiting.
- **Late Binding:** A subtle "Link External Wallet" button in the Hangar settings or profile, framed as "Securing Assets to the Void-Chain."

### 5.2 The Pilot ID Badge (Hangar UI)
- **Visual:** A glassmorphic "ID Card" displayed in the Hangar.
- **Details:** Shows Character Name, Rank, Gender, and a portrait based on the selected appearance.
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

## 9. Mothership Systems HUD
The Mothership interface must clearly distinguish between the two independent engineering paths:

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
