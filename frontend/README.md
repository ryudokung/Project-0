# Project-0 Frontend

This is the frontend for **Project-0**, a Tactical Noir AI-Powered Web Game. Built with Next.js 15, it uses a decoupled architecture with singleton systems and an EventBus to manage game logic.

## 🚀 Getting Started

1.  **Install Dependencies**:
    ```bash
    npm install
    ```

2.  **Environment Setup**:
    Create a `.env.local` file (refer to `.env.example` if available) with your API endpoints and keys.

3.  **Run Development Server**:
    ```bash
    npm run dev
    ```

4.  **Build for Production**:
    ```bash
    npm run build
    ```

## 🏗 Architecture

-   **Systems**: Located in `src/systems/`. These are singleton classes (e.g., `BastionSystem`, `ExplorationSystem`) that handle core game logic.
-   **EventBus**: A global event emitter (`src/systems/EventBus.ts`) used for communication between systems and UI components.
-   **Game Machine**: An XState state machine (`src/machines/gameMachine.ts`) that governs the high-level game states (Bastion, Map, Exploration, Combat).
-   **Components**: React components in `src/components/game/` are designed to be "thin" view layers that react to system events.

## 🛠 Tech Stack

-   **Framework**: Next.js 15 (App Router)
-   **State Management**: XState, React Context
-   **Styling**: Tailwind CSS, Lucide React (Icons)
-   **3D/Visuals**: React Three Fiber (R3F), Three.js
-   **Communication**: Axios (API), EventBus (Internal)

---
Part of the [Project-0](../README.md) ecosystem.
