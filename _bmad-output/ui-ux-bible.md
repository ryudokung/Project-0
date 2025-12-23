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

## 4. User Journey: The "Click-to-Explore" Loop
1. **Login:** Terminal-style boot sequence.
2. **Hangar:** View current fleet, repair, and customize.
3. **Radar Interface:** A node-based map where new locations appear as "Blips". Click a blip to view mission details.
4. **Action Scene:** A high-fidelity 3D view of the mission. Interaction is click-based (e.g., clicking a door to enter, clicking a crate to salvage).
5. **Debrief:** Summary of loot and narrative progress.

## 5. Interaction Model: "Tactical Clicks"
- **Node Discovery:** Clicking the "Scan" button reveals new clickable nodes on the Radar.
- **Resource Cost:** Every click in a mission (moving, searching) has a cost (O2 for Pilot, Fuel for Mothership).
- **Visual Feedback:** Every click triggers a high-quality 3D animation or camera shift to keep the experience immersive.
