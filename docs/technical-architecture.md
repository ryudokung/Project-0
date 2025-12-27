# Technical Architecture Summary

**Date:** December 25, 2025
**Status:** Approved

## 1. Frontend Architecture: "Decoupled Game Engine"

The frontend has been refactored from a standard React state-management approach to a decoupled "Game Engine" architecture. This pattern is designed to handle complex game logic independently of the UI rendering.

### Core Components:
- **Event Bus (`EventBus.ts`)**: A global, type-safe event emitter that serves as the central nervous system of the game. Systems and UI components communicate via events (e.g., `GAME_EVENTS.NOTIFICATION`, `GAME_EVENTS.STATS_UPDATE`).
- **Singleton Systems**: Logic is encapsulated into standalone classes (Systems) that manage specific game domains.
    - `ExplorationSystem.ts`: Manages mission starts, timeline advancement, and narrative state.
    - `CombatSystem.ts` (Planned): Will manage turn-based battle logic.
- **XState Machine (`gameMachine.ts`)**: Controls the high-level "Stages" of the application:
    - `initializing` -> `landing` -> `characterCreation` -> `bastion` -> `exploration` / `combat`.
- **React UI Layer**: Purely functional components that listen to the Event Bus for updates and send commands to the Systems.

## 2. Authentication & Onboarding Flow

The project follows a "Web2.5" onboarding strategy to minimize friction for new players.

### Authentication Rules:
1.  **No Initial Wallet Requirement**: Users are **NOT** required to connect a wallet to start playing.
2.  **Multi-Provider Login**:
    - **Guest Login**: Instant entry using a unique ID stored in `localStorage`.
    - **Social Login**: Email or Google login via **Privy**.
3.  **Pilot Registration**: Every new account must create a "Pilot" (Character). This process:
    - Creates a record in the `characters` table.
    - Initializes `pilot_stats` (O2, Fuel, XP).
    - Initializes a "Starter Pack" (Manifested Assets) in the database.
4.  **Wallet Linking (Late-Stage)**: Wallet connection is an optional step used only when the player decides to "Mint" their manifested assets into on-chain NFTs or interact with the Base L2 blockchain.

## 3. Backend & Database Structure

### Backend:
- **Go (Modular Monolith)**: Organized into domain-specific packages (`auth`, `vehicle`, `exploration`, `game`).
- **YAML Blueprint System**: A data-driven engine that separates game content from logic.
    - **Registry**: The `game.BlueprintRegistry` loads YAML files from `backend/blueprints/` on startup.
    - **Content Types**: Currently supports `nodes.yaml` (Exploration nodes and choices) and `enemies.yaml` (NPC stats and classes).
    - **Extensibility**: Designed to be expanded for `parts.yaml` (Vehicle modules) and `loot.yaml` (Reward tables).
- **JWT Middleware**: All protected routes require a valid JWT. The middleware extracts the `user_id` and injects it into the request context.
- **Context Keys**: Shared constants are used for context keys to prevent circular dependencies between packages.

### Database (PostgreSQL):
- **Relational Integrity**: Uses UUIDs for all primary and foreign keys.
- **Nullable Vehicles**: The `expeditions` table allows `vehicle_id` to be `NULL` to support "Pilot Only" (on-foot) exploration modes.
- **Atomic Updates**: Resource consumption (O2/Fuel) and Gacha pulls use atomic SQL updates to prevent race conditions.
