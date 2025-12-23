# Project-0: AI-Powered Seasonal NFT Web Game

Project-0 is a sustainable Crypto Web Game that bridges traditional gaming mechanics with AI-driven innovation. It features a hybrid economy designed to eliminate "FOMO" cycles, utilizing AI to generate unique, seasonal NFT assets (Mechs, Tanks, Ships) with genuine aesthetic and functional rarity.

## 🚀 Core Features

- **Modular NFT Assembly:** Every part (Railgun, Shield, Pilot Suit) is an individual NFT using **ERC-6551 (Token Bound Accounts)**. AI dynamically synthesizes these parts into a single visual representation.
- **Multi-Stage Exploration:** Seamless 3D transitions (Mothership -> Mech/Aircraft -> Pilot EVA) using **WebGPU + React Three Fiber**.
- **Risk-Based Gameplay:** Radar Scan & Risk Assessment with **Web2 Backend Fog of War** and AI-generated accidents via **Model Context Protocol (MCP)**.
- **Personal Scale (1:1):** Focus on the bond between one Pilot and their unique Mech/Aircraft.
- **Closed-Loop Economy:** Sustainable USDT-in model with Bitcoin-style difficulty adjustment.
- **Web3 Integration:** Base L2 blockchain, Social Login, and Embedded Wallets for seamless onboarding.

## 🛠 Tech Stack

- **Frontend:** Next.js 15+ (App Router), **WebGPU**, **React Three Fiber**, Tailwind CSS, Zustand, [UI/UX Pro Max](https://github.com/nextlevelbuilder/ui-ux-pro-max-skill/).
- **Backend:** Go (Modular Monolith), Clean Architecture, Saga Pattern, **MCP (Model Context Protocol)**.
- **Blockchain:** Base L2, Solidity (**ERC-6551**, ERC-721), Privy/Dynamic (Auth).
- **AI:** FLUX.1 via Fal.ai/Replicate, Structured Output (JSON Schema).
- **Database:** PostgreSQL (ACID compliant for Saga).
- **Infrastructure:** Docker & Docker Compose.

## 📂 Project Structure

- `_bmad/`: BMad Method configuration and agent manifests.
- `_bmad-output/`: Project documentation (PRD, Architecture, Research).
- `backend/`: (Pending) Go backend source code.
- `frontend/`: (Pending) Next.js frontend source code.

## 📖 Documentation

- [Product Requirements Document (PRD)](_bmad-output/prd.md)
- [Architecture Decisions](_bmad-output/architecture.md)

## 🚀 Getting Started

### Prerequisites
- Docker & Docker Compose installed.

### Running the Project
1. Clone the repository.
2. Run the following command:
   ```bash
   docker-compose up --build
   ```
3. Access the services:
   - Frontend: `http://localhost:3000`
   - Backend API: `http://localhost:8080`
   - Database: `localhost:5432`

---
*Developed using the BMad Method.*
