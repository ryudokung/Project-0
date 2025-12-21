# Project Context: Project-0

## Project Overview
Project-0 is a Crypto Web Game featuring AI-generated seasonal NFTs (Mechs, Tanks, Ships) on Base L2, with a hybrid economy and a "No USDT Out" revenue model.

## Technical Stack & Infrastructure
- **Frontend:** Next.js 14+ (App Router), Tailwind CSS, Zustand.
- **Backend:** Go (Modular Monolith), Clean Architecture, Saga Pattern.
- **Blockchain:** Base L2, ERC-721, Privy/Dynamic (Auth).
- **AI:** FLUX.1 via Fal.ai/Replicate.
- **Database:** PostgreSQL 16.
- **Infrastructure:** Docker & Docker Compose.

## Critical Implementation Rules
- **Saga Pattern:** All multi-step transactions (Assembly, Discovery) must use the Saga Pattern with idempotency keys.
- **Clean Architecture:** Go backend must follow Clean Architecture (Entities, Use Cases, Repository, Delivery).
- **Dockerized Environment:** All services must be runnable via `docker-compose up`.
- **AI Guardrails:** Use Structured Output (JSON Schema) and RAG to prevent hallucinations.
- **Housekeeping:** Temporary logs (Combat Logs) must be deleted after 7 days unless paid to save.

## Current Progress (Epic 2: Mech Management)
- [x] Backend Mech Module (Entity, Repository, UseCase, Handler)
- [x] Starter Mech Minting Logic (RNG Stats)
- [x] Frontend Mech Card & Hangar UI
- [x] Backend-Frontend Integration for Mech Listing
- [x] Claim Starter Mech Flow

## Directory Structure
- `/backend`: Go source code.
- `/frontend`: Next.js source code.
- `/_bmad-output`: Project documentation (PRD, Architecture).
- `/_bmad`: BMad Method configuration and agents.

## Development Workflow
1. Use `docker-compose up --build` to start the environment.
2. Backend runs on `:8080`, Frontend on `:3000`.
3. Database is accessible on `:5432`.
