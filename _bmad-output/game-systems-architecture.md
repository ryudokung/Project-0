# Game Systems Architecture: Project-0

**Author:** Cloud Dragonborn (Game Architect)
**Date:** 2025-12-23
**Version:** 1.0

## 1. Architectural Philosophy
To support a high-fidelity, real-time 3D experience with complex modular assets, we adopt a **Hybrid ECS-Service Architecture**. This ensures high performance for the 3D frontend while maintaining strict consistency for the Web3/Backend state.

## 2. Core Game Loop & State Management

### 2.1 Single-Page Game Loop (Unified Controller)
- **Architecture:** The entire game experience is managed by a single React entry point (`/`), acting as a **Unified Game Controller**.
- **State Machine:** Uses **XState v5** to manage `GameStage` transitions (Bastion -> Map -> Exploration -> Combat).
- **Immersion:** Eliminates page reloads, allowing for persistent background audio, seamless visual transitions (Framer Motion), and consistent state sharing (O2, Fuel, Pilot Stats) across all stages.
- **Frontend (WebGPU + R3F):** Uses a lightweight **ECS (Entity Component System)** pattern for managing 3D entities (Vehicles, VFX, UI) within the unified scene.
- **State Synchronization:** The frontend maintains a "Local Mirror" of the game state, updated via WebSockets or REST from the Go Backend.

### 2.2 Backend Game Logic (Go)
- **Tick-less Simulation:** Since combat is asynchronous auto-resolve, the backend runs a discrete simulation rather than a continuous tick loop.
- **Deterministic Simulation:** Ensures that the same inputs (Stats, RNG Seed) always produce the same outcome, allowing the frontend to "Replay" the battle with 100% accuracy.

---

## 3. Modular NFT Assembly Pipeline (The "Showcase" Engine)

### 3.1 Visual Synthesis Flow
1. **Metadata Extraction:** Backend reads the `ERC-6551` Token Bound Account to list all equipped NFT parts.
2. **Visual DNA Mapping:** Each NFT part has a "Visual DNA" string in its metadata (e.g., `V-DNA: CHROME_GLOW_01`).
3. **Prompt Construction:** The AI Integration Service assembles a master prompt:
   - `Base: [Vehicle Chassis/Fuselage V-DNA]`
   - `Addon: [Weapon/Engine V-DNA] + [Shield/Wing V-DNA]`
   - `Context: [Environment] + [Damage State]`
4. **AI Generation (FLUX.1):** Generates the high-fidelity "Master Image."
5. **3D Asset Mapping:** The frontend maps the V-DNA to specific 3D models and shaders in the React Three Fiber scene.

### 3.2 Showcase System (Bastion & Profile)
- **Dynamic Lighting:** Uses WebGPU's advanced compute shaders for real-time reflections on the vehicle's surface.
- **Visual Equipment Map:** A silhouette-based interface for mapping modules to anatomical slots (HEAD, CORE, ARMS, LEGS).
- **Photo Mode Engine:** A dedicated R3F scene with adjustable camera parameters (FOV, Aperture, Bloom) to capture the "Perfect Flex."
- **Public API:** Allows external sites (or the Hall of Fame) to fetch and render the 3D model of a player's vehicle.

---

## 4. Technical Systems

### 4.1 Radar & Fog of War (Web2 Backend)
- **Server-Side Validation:** The "Map" is stored on the server. The client only receives data for entities within the Mothership's `Scanner Range`.
- **Anti-Cheat:** Prevents map-hacking by never sending hidden entity data to the client.

### 4.2 V2O Bridge (Virtual-to-On-chain)
- **Lazy Minting:** Assets are created as database records first.
- **Minting Trigger:** User initiates a "Mint" transaction on Base L2.
- **Saga Pattern:** Ensures the database record is marked as "On-chain" only after the transaction is confirmed.

### 4.3 Hybrid Energy System
- **Standard Energy:** Managed in Redis for fast read/write during gameplay.
- **Premium Energy:** Managed as an ERC-20 token or USDT balance, requiring on-chain verification for "Overclock" actions.

---

## 5. Scalability & Performance
- **Asset Streaming:** 3D models are sharded and loaded on-demand to keep initial load times under 2 seconds.
- **GPU Load Balancing:** AI generation requests are queued and processed based on priority (Premium vs. Standard).

---
*Architecture approved by Cloud Dragonborn - Build for tomorrow, ship today.*
