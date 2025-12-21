---
stepsCompleted: [1, 2, 3, 4, 5]
inputDocuments:
  - "_bmad-output/prd.md"
  - "_bmad-output/analysis/research/technical-Project-0-research-2025-12-21.md"
  - "_bmad-output/analysis/product-brief-Project-0-2025-12-21.md"
documentCounts:
  prd: 1
  research: 1
  briefs: 1
workflowType: 'architecture'
lastStep: 5
project_name: 'Project-0'
user_name: 'Vatcharin'
date: '2025-12-21'
---

# Architecture Decisions - Project-0

**Author:** Vatcharin
**Date:** 2025-12-21

## 1. Introduction & Context

This document captures the architectural decisions for Project-0, a Crypto Web Game featuring AI-generated seasonal NFTs and a hybrid economy. These decisions are designed to ensure consistency across AI agents and prevent implementation conflicts.

### 1.1 Input Context
- **PRD:** Defines the "Closed-Loop Premium Economy," "Adaptive Universe," and "Engineering & Assembly" systems.
- **Technical Research:** Identifies Base L2, Go Orchestrator, and FLUX.1 AI as core technologies.

## 2. Project Context & Complexity Analysis

### 2.1 Architectural Scope
Project-0 is a high-complexity hybrid system integrating Web3 (Blockchain), AI (Generative Art), and real-time gaming mechanics. The architecture must handle high-stakes financial transactions (USDT) while maintaining an immersive, story-driven user experience.

### 2.2 Key Architectural Challenges
- **Hybrid State Consistency:** Ensuring 100% alignment between off-chain AI generation and on-chain NFT minting.
- **Economic Integrity:** Implementing a transparent, self-balancing difficulty adjustment mechanism (Bitcoin-style).
- **Scalability:** Supporting 10,000 concurrent users with real-time updates and load-adaptive AI processing.
- **Security:** Multi-layered anti-bot protection and secure non-custodial marketplace operations.

### 2.3 Architectural Decision Records (ADR) - Summary

#### ADR 001: Distributed Transaction Management
- **Decision:** Use **Saga Pattern (Orchestration-based)**.
- **Rationale:** Essential for managing long-running, multi-step processes across heterogeneous systems (Go Backend, Fal.ai API, Base L2 Smart Contracts).
- **Key Implementation Details:**
    - Centralized Orchestrator in Go to manage the "Mech Assembly" state machine.
    - Mandatory **Idempotency Keys** for all external service calls to prevent duplicate transactions.
    - Compensating transactions (e.g., credit refunds) for failed steps in the assembly pipeline.

#### ADR 002: Backend Architecture Pattern
- **Decision:** Use **Modular Monolith** with **Clean Architecture**.
- **Rationale:** Provides the best balance between development speed (MVP) and future scalability. Allows for clear separation of concerns while avoiding the operational overhead of Microservices in the early stages.
- **Key Implementation Details:**
    - Domain-driven folder structure (e.g., `internal/ai`, `internal/blockchain`).
    - Shared kernel for common types and utilities.
    - In-process communication via interfaces to allow future extraction into Microservices.

## 3. System Components & Services (Go Backend)

## 3. System Components & Services (Go Backend)

เพื่อให้ระบบมีความยืดหยุ่นและรองรับ Saga Pattern เราจะใช้โครงสร้างแบบ **Modular Monolith** ด้วย **Clean Architecture** ในภาษา Go โดยแบ่ง Service/Module หลักดังนี้:

### 3.1 Orchestrator Service (The Brain)
- **หน้าที่:** ควบคุม State Machine ของการสร้าง Mech (Saga Pattern).
- **Logic:** ประสานงานระหว่าง AI Service และ Blockchain Service เพื่อให้มั่นใจว่าถ้า AI Gen สำเร็จ ต้องมีการ Mint NFT ตามมา หรือถ้า Mint พลาดต้องมีการ Rollback/Refund.

### 3.2 AI Integration Service
- **หน้าที่:** เชื่อมต่อกับ FLUX.1 (via Fal.ai/Replicate).
- **Logic:** 
    - จัดการ Prompt Engineering, Queue การ Generate, และระบบ **HITL (Human-in-the-loop)**.
    - **Combat Visualization:** สร้างภาพการปะทะกัน (e.g., Mech vs Mech) เพื่อเพิ่มความตื่นเต้นในรายงานผลการต่อสู้.

### 3.3 Blockchain Service
- **หน้าที่:** สื่อสารกับ Base L2 Smart Contracts.
- **Logic:** จัดการการ Mint NFT, การโอน USDT, และคำนวณ **Difficulty Adjustment** (Bitcoin-style) สำหรับการค้นพบ Star ใหม่.

### 3.4 Game Engine Service
- **หน้าที่:** ประมวลผล Logic ของเกมที่ไม่ต้องอยู่บน Chain ทั้งหมด.
- **Logic:** 
    - **Multi-Stage Exploration:** 
        - **Space Travel:** บังคับใช้ `Ship`.
        - **Orbital Approach:** ต้องใช้ `Ship` + `Mech` ทำงานร่วมกัน.
        - **Planetary Surface:** บังคับใช้ `Mech` ในการสำรวจและทำภารกิจ.
    - **Combat Mode:** คำนวณผลการต่อสู้ตาม Stats และประเภทของ Vehicle (e.g., Mech ปะทะ Mech บนพื้นผิว).
    - **Star Discovery:** ระบบ Trigger สำหรับการขยายจักรวาลตามเงื่อนไขที่กำหนด.

### 3.5 Economy & Revenue Service
- **หน้าที่:** จัดการระบบการเงินและ Monetization.
- **Logic:** ระบบ Season Pass, การขายไอเทม, และการตรวจสอบ Revenue Flow (No USDT Out policy).

### 3.6 Auth & User Service
- **หน้าที่:** จัดการ Identity และ Profile.
- **Logic:** เชื่อมต่อกับ Privy/Dynamic สำหรับ Web3 Auth และเก็บข้อมูล Metadata ของผู้เล่น.

## 4. Microservices Architecture (Alternative/Evolution)

หากต้องการขยายเป็น Microservices เต็มรูปแบบ โครงสร้างจะเปลี่ยนจากการเรียก Function ภายใน (In-process) เป็นการสื่อสารผ่าน Network (RPC/Events) ดังนี้:

### 4.1 Service Communication
- **Synchronous (gRPC):** ใช้สำหรับการอ่านข้อมูลข้าม Service ที่ต้องการความเร็วสูง (เช่น Game Engine ถาม User Stats).
- **Asynchronous (Message Broker - NATS/RabbitMQ):** ใช้สำหรับ Saga Pattern และ Event-driven logic (เช่น เมื่อ AI Gen เสร็จ จะส่ง Event ไปยัง Blockchain Service).

### 4.2 Infrastructure Components
- **API Gateway:** จุดรับ Request เดียวจาก Frontend (Next.js) ทำหน้าที่ Routing, Rate Limiting และ Auth Validation.
- **Event Bus:** หัวใจหลักของ Saga Pattern ในแบบ Microservices เพื่อให้แต่ละ Service ทำงานแบบ Decoupled.
- **Distributed Tracing (Jaeger/OpenTelemetry):** จำเป็นมากเพื่อดูว่า Request หนึ่งวิ่งผ่าน Service ไหนบ้าง (สำคัญมากตอน Debug Saga ที่ซับซ้อน).

### 4.3 Database Strategy
- **Database-per-Service:** แต่ละ Service มี DB ของตัวเอง (เช่น AI Service ใช้ PostgreSQL เก็บ Prompt, Game Service ใช้ Redis เก็บ Real-time State).

## 5. Infrastructure & Deployment (Step 5)

เราจะใช้ **Docker** และ **Docker Compose** ในการจัดการสภาพแวดล้อมการพัฒนาและการ Deploy เพื่อให้มั่นใจว่าระบบทำงานได้เหมือนกันในทุก Environment.

### 5.1 Container Strategy
- **Backend Container:** รัน Go API Server (Port 8080).
- **Frontend Container:** รัน Next.js App (Port 3000).
- **Database Container:** รัน PostgreSQL 16 (Port 5432).

### 5.2 Environment Variables
- `DB_URL`: Connection string สำหรับ PostgreSQL.
- `NEXT_PUBLIC_API_URL`: URL ของ Backend API สำหรับ Frontend.
- `AI_API_KEY`: (Pending) สำหรับเชื่อมต่อ Fal.ai/Replicate.
- `BLOCKCHAIN_RPC_URL`: (Pending) สำหรับเชื่อมต่อ Base L2.

## 6. Data Schema & API Contracts (Step 4)

เพื่อให้ระบบทำงานร่วมกันได้แบบไร้รอยต่อ เราจะใช้ **PostgreSQL** เป็นฐานข้อมูลหลัก (เนื่องจากรองรับ ACID Transactions ที่จำเป็นสำหรับ Saga Pattern) โดยมี Schema เบื้องต้นดังนี้:

### 5.1 Core Entities

#### 1. User Entity
*   `id`: UUID (Primary Key)
*   `wallet_address`: string (Unique, Indexed)
*   `username`: string
*   `credits`: decimal (In-game currency)
*   `last_login`: timestamp

#### 2. Mech (NFT) Entity
*   `id`: UUID (Primary Key)
*   `token_id`: uint256 (On-chain ID, Nullable until minted)
*   `owner_id`: UUID (FK to User)
*   `vehicle_type`: enum (MECH, TANK, SHIP) - แยกประเภทตาม PRD
*   `class`: enum (STRIKER, GUARDIAN, SCOUT, ARTILLERY) - สำหรับกลยุทธ์ใน Combat Mode
*   `image_url`: string (AI Generated URL)
*   `stats`: JSONB (HP, Attack, Defense, Speed, Energy) - ใช้ JSONB เพื่อความยืดหยุ่น
*   `rarity`: enum (COMMON, RARE, LEGENDARY)
*   `season`: string (e.g., "Season 1: Iron Age")
*   `status`: enum (PENDING, MINTED, BURNED)

#### 3. Saga Transaction (Orchestration Log)
*   `id`: UUID (Primary Key)
*   `user_id`: UUID (FK to User)
*   `type`: enum (MECH_ASSEMBLY, STAR_DISCOVERY, EXPLORATION_MISSION)
*   `current_step`: string (e.g., "SPACE_TRAVEL", "PLANETARY_SURFACE", "COMBAT_RESOLVING")
*   `status`: enum (STARTED, COMPLETED, FAILED, COMPENSATING)
*   `payload`: JSONB (เก็บ Context เช่น `ship_id`, `mech_id`, `target_star_id`)
*   `idempotency_key`: string (Unique)

#### 4. Star (Universe) Entity
*   `id`: UUID (Primary Key)
*   `name`: string
*   `difficulty_level`: int (Bitcoin-style adjustment)
*   `discovered_by`: UUID (FK to User)
*   `is_active`: boolean

#### 5. Combat Log & Battle Record
*   `id`: UUID (Primary Key)
*   `user_id`: UUID (FK to User)
*   `mission_id`: UUID (FK to Saga Transaction)
*   `battle_data`: JSONB (Stats, Rounds, Outcome)
*   `image_url`: string (AI Generated Action Shot)
*   `is_permanent`: boolean (Default: false)
*   `expires_at`: timestamp (Housekeeping: e.g., 7 days)
*   `saved_at`: timestamp (Null if not paid to save)

### 5.2 API Endpoints (v1)

เราจะใช้ RESTful API สำหรับการสื่อสารระหว่าง Frontend และ Backend:

#### 1. Assembly (Engineering)
- `POST /api/v1/assembly/start`: เริ่มกระบวนการสร้าง Mech/Ship (ใช้ Credits/USDT)
- `GET /api/v1/assembly/:id/status`: เช็คสถานะการสร้างและคิว AI

#### 2. Exploration & Combat
- `POST /api/v1/exploration/start`: ส่งทีมออกสำรวจ (ต้องระบุ `ship_id` และ `mech_id`)
- `GET /api/v1/exploration/:id/events`: ดึงเหตุการณ์ที่เกิดขึ้นระหว่างสำรวจ (Real-time updates)
- `POST /api/v1/combat/:id/save`: จ่าย USDT เพื่อบันทึก Battle Log ถาวร

#### 3. Universe & Marketplace
- `GET /api/v1/universe/stars`: ดึงข้อมูลดวงดาวและระดับความยากปัจจุบัน
- `GET /api/v1/marketplace/listings`: ดึงรายการขาย NFT (Keys)

### 5.3 Saga Flow Definitions

#### Flow A: Mech Assembly (Premium)
1.  **User:** กดปุ่ม "Assemble Premium Mech" -> จ่าย USDT.
2.  **Orchestrator:** สร้าง Saga Record (Status: `STARTED`) -> ล็อคทรัพยากร.
3.  **AI Service:** รับ Job -> Gen ภาพด้วย FLUX.1 -> ส่งเข้า HITL Queue.
4.  **Admin:** ตรวจสอบภาพ -> กด Approve.
5.  **Blockchain Service:** Mint NFT บน Base L2 -> อัปเดต Metadata.
6.  **Orchestrator:** อัปเดต Saga (Status: `COMPLETED`) -> แจ้งเตือนผู้เล่น.

#### Flow B: Exploration & Combat
1.  **User:** เลือกพิกัดและยาน -> กด "Launch Mission".
2.  **Game Engine:** คำนวณระยะทางและสุ่มเหตุการณ์ (Events).
3.  **Game Engine:** หากเจอศัตรู -> คำนวณผลการต่อสู้ (Stat-based).
4.  **AI Service:** Gen ภาพ "Action Shot" ของการปะทะ.
5.  **Orchestrator:** บันทึก Combat Log (Status: `TEMP`) -> ตั้งเวลา Housekeeping (7 วัน).
