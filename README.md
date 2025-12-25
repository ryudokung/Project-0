# Project-0: Tactical Noir AI-Powered Web Game

Project-0 is a high-stakes, 1:1 scale Sci-Fi Exploration and Combat game on **Base L2**. It features a hybrid economy and AI-driven innovation, where players operate modular vehicles in an intense, evolving universe.

## 🌌 The Vision: "Tactical Noir"
Project-0 is not just a game; it's an evolving universe. Players take on the role of **Operatives** in a dark, industrial future where every mission is a chapter in a gripping narrative. The core focus is on **Social Prestige**, **Collection Pride**, and **High-Stakes Survival**.

## 🚀 Core Gameplay Loop (MVP1 Focus)
1.  **The Bastion (Preparation):** Wake up in your derelict **Bastion**. Manage energy, install modules, and repair your Vehicles using salvaged resources.
2.  **The Warp (Deployment):** Scan the Universe Map for signals and select a destination. Warp consumes Fuel and O2.
3.  **The Timeline (Exploration):** Navigate through a linear timeline of events. Encounter combat, salvage opportunities, and anomalies.
4.  **Strategic Choices:** Make critical decisions (Stealth, Assault, Bypass) based on your Pilot's stats and Vehicle's condition.
5.  **The Extraction (Risk vs Reward):** Reach a Warp Gate to secure loot, or risk an Emergency Warp if your **Durability** is critical.
6.  **The Salvage (Economy):** Refine scrap, research blueprints, and choose when to perform a **Stage Change (Mint to NFT)**.

## 🚜 Vehicle Archetypes & Transformation
- **Mechs:** Balanced bipedal combat units. All-terrain capability.
- **Tanks:** Heavy-armored siege units. High defense. Can **Transform** into Robot mode for terrain adaptability.
- **Ships:** High-speed aerial/space units. Can **Transform** into Robot mode for precision landing and combat.
- **Speeders:** Hoverbikes for high-velocity scouting and stealth.
- **Haulers:** Industrial units for bulk transport and massive salvage.
- **Pilot & Exosuits:** The operative's "Infiltration Layer." Exosuits provide power armor for indoor salvage, boarding actions, and stealth missions where vehicles cannot enter.

## ✨ Unique Selling Points
- **AI-Generated Visual DNA:** Vehicles are synthesized by AI (FLUX.1) based on their specific NFT parts and stats.
- **Resonance & Dominance:** High-level Pilot Resonance and gear quality allow players to "Annihilate" lower-tier threats, focusing on high-stakes progression.
- **Deep Durability System (DDS):** Items aren't just "broken" or "fixed." They have 5 thresholds (Pristine to Broken) affecting stats and visuals.
- **Security-First Architecture:** Every game mechanic is validated server-side. No client-side trust. Atomic database operations prevent race conditions and double-spending.
- **Hybrid Narrative Architecture:** A unique blend of Director-fixed plot points ("Narrative Anchors") and AI-generated dynamic content, ensuring every player has a unique but cohesive story.
- **Node-Based Exploration:** A "Tactical Noir" radar interface where you click to discover the universe.
- **Showcase Engine:** A high-fidelity 3D Bastion for "flexing" your unique assets to the community.
- **Stage Change Model (V2O):** Play for free, find rare items, and choose to bridge them to the blockchain as NFTs while maintaining their in-game utility and durability.
- **Captain's Log:** A personal, AI-narrated history of your greatest victories and "First Discoveries."

## 🛠 Tech Stack
- **Frontend**: Next.js 15+, **Decoupled Systems (EventBus + Singletons)**, **XState (Finite State Machine)**, **WebGPU**, **React Three Fiber (R3F)**, Tailwind CSS.
- **Backend**: Go (Modular Monolith), Clean Architecture, Saga Pattern, **JWT Security Middleware**, **Atomic Concurrency Control**.
- **Blockchain**: Base L2, **ERC-6551 (Token Bound Accounts)**, ERC-721.
- **AI**: FLUX.1 via Fal.ai/Replicate, Structured Output.

## 🏗 Architecture Principles
- **Security-First**: Every game mechanic is validated server-side. No client-side trust. Atomic database operations prevent race conditions and double-spending.
- **Decoupled Systems**: The frontend is built like a game engine. Logic is separated into **Singleton Systems** that communicate via a global **EventBus**, ensuring the UI remains a pure view layer.
- **State-Driven UI**: The frontend uses **XState** to manage the high-level game loop. This ensures the UI is always in a valid state and simplifies complex transitions between Bastion, Map, Exploration, and Combat.

## 📖 The Technical Bibles
The project is governed by core design and technical documents:
0.  [Gameplay & Technical Spec](_bmad-output/gameplay-technical-spec.md): **The Master Blueprint** for Loop, DDS, and AI.
1.  [Level Design Bible](_bmad-output/level-design-bible.md): Sectors, Hazards, and Zones.
2.  [Vehicle & Item Bible](_bmad-output/vehicle-item-bible.md): Modular Parts and Visual DNA.
3.  [Combat Design Bible](_bmad-output/combat-design-bible.md): Formulas, Status Effects, and Logic.
4.  [Bastion Bible](_bmad-output/bastion-bible.md): The Command Center, Facilities, and DDS.
5.  [Narrative & Lore Bible](_bmad-output/narrative-lore-bible.md): History, Factions, and Hooks.
6.  [Economy Bible](_bmad-output/economy-bible.md): Dual-Currency and Stage Change Model.
7.  [UI/UX Bible](_bmad-output/ui-ux-bible.md): Tactical Noir HUD and User Journey.
8.  [Creator Studio Bible](_bmad-output/creator-studio-bible.md): God Mode and Context Patching.
9.  [Project Journey](_bmad-output/project-journey.md): The evolution of Tech & Design.

---
*Developed using the BMad Method with specialized Game Dev Agents.*
