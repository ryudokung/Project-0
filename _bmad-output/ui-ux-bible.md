# UI/UX & Immersive HUD Bible: Project-0 (The Experience)

**Goal:** Define the visual language, HUD elements, and user journey to ensure maximum immersion and "Flex" potential.

## 1. Visual Language: "Tactical Noir"
- **Color Palette:** Deep Blacks, Neon Greens (HUD), Warning Oranges, and Industrial Grays.
- **Typography:** Monospaced fonts (for that "Terminal" feel) and bold, italicized headers for impact.
- **Effects:** Scanlines, slight chromatic aberration, and HUD glitches when taking damage.

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

## 5. Interaction Model: "Tactical Clicks"
- **Node Discovery:** Clicking the "Scan" button reveals new clickable nodes on the Radar.
- **Resource Cost:** Every click in a mission (moving, searching) has a cost (O2 for Pilot, Fuel for Mothership).
- **Visual Feedback:** Every click triggers a high-quality 3D animation or camera shift to keep the experience immersive.
