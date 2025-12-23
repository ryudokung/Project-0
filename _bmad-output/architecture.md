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
    - **New Sagas:**
        - `ColonyUpgradeSaga`: Resource deduction, timer management, and facility state update.
        - `ExplorationMissionSaga`: Fuel consumption, encounter resolution, and loot/durability updates.
        - `SalvageOperationSaga`: Unit capture processing (Sell/Research/Scrap) and inventory sync.
        - `StoryProgressionSaga`: Narrative milestone validation and fixed reward distribution.
        - `CombatSimulationSaga`: Turn-based simulation, durability deduction, and AI narrative trigger.

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
- **หน้าที่:** เชื่อมต่อกับ FLUX.1 (via Fal.ai/Replicate) และจัดการบริบทของ AI.
- **Logic:** 
    - **Modular NFT Synthesis:** อ่าน Metadata และ Visual Traits จาก NFT แต่ละชิ้น (Railgun, Shield, Pilot Suit) เพื่อนำมาสังเคราะห์เป็นภาพเดียวที่สมบูรณ์.
    - **Model Context Protocol (MCP):** ใช้ MCP เพื่อให้ AI Narrative Engine สามารถเข้าถึงข้อมูล Game State (เช่น HP, O2, Location) จาก Backend ได้โดยตรงแบบ Real-time เพื่อสร้างเนื้อเรื่องที่แม่นยำโดยไม่ต้องส่ง Data ผ่าน Cloud ทั้งหมด.
    - **Context-Aware Prompting:** สร้าง Prompt ที่รวมทั้งไอเทมที่สวมใส่, สถานะความเสียหาย, และบริบทของเหตุการณ์.

### 3.3 Blockchain Service
- **หน้าที่:** สื่อสารกับ Base L2 Smart Contracts.
- **Logic:** 
    - **Virtual-to-On-chain (V2O) Bridge:** จัดการกระบวนการเปลี่ยนสถานะจาก Virtual Asset (Server-side) เป็น On-chain NFT เมื่อผู้เล่นกด Mint หรือต้องการขาย.
    - **ERC-6551 (Token Bound Accounts):** ใช้มาตรฐาน ERC-6551 เพื่อให้ NFT หลัก (เช่น Mech Chassis) มี Wallet ของตัวเองสำหรับเก็บชิ้นส่วนอุปกรณ์ (Weapons, Shields) ทำให้การซื้อขายใน Marketplace ทำได้แบบยกชุด (Bundled Assets).
    - **Modular NFT Management:** จัดการการ Mint และโอน NFT แยกตามชิ้นส่วน.
    - **Metadata Sync:** อัปเดต Metadata ของ NFT แต่ละชิ้นตามสถานะความเสียหาย (Durability) ที่เกิดขึ้นจริง.

### 3.4 Game Engine Service
- **หน้าที่:** ประมวลผล Logic ของเกมที่ไม่ต้องอยู่บน Chain ทั้งหมด.
- **Logic:** 
    - **Hybrid Energy Management:** คำนวณการใช้ Standard Energy (Free) และ Premium Energy (Paid) สำหรับการสำรวจ.
    - **Web2 Backend Fog of War:** ใช้ Server-side Validation ในการจัดการ Fog of War โดย Backend จะส่งข้อมูลเฉพาะสิ่งที่ผู้เล่น "มองเห็น" (ตามระยะ Scan) ไปยัง Client เท่านั้น เพื่อป้องกันการโกง (Map Hack) โดยไม่ต้องใช้ ZKP ในช่วงแรก.
    - **Radar & Risk Assessment Logic:** คำนวณโอกาสรอดชีวิตและระดับความอันตราย (Threat Level) ตาม Scanner ของ Mothership.
    - **AI Event Trigger:** หากเกิดอุบัติเหตุ ระบบจะส่งบริบท (Context) ผ่าน MCP ไปให้ AI Service เพื่อสร้างเหตุการณ์สุ่มที่สมจริง.
    - **Multi-Stage Exploration:** 
        - **Space Travel:** บังคับใช้ `Ship` (Mothership).
        - **Orbital Approach & Atmospheric Entry:** ตรวจสอบเงื่อนไขการลงจอด (Atmospheric Shielding).

### 3.5 Frontend Design System ([UI/UX Pro Max](https://github.com/nextlevelbuilder/ui-ux-pro-max-skill/))
- **หน้าที่:** จัดการความสวยงามและการใช้งานของผู้ใช้งาน (User Experience).
- **Logic:**
    - **UI Styles:** ใช้ Glassmorphism สำหรับ HUD/Cockpit และ Bento Grid สำหรับ Dashboard ข้อมูล.
    - **Color Systems:** ใช้ Industry-specific palettes (Fintech/SaaS) สำหรับ Marketplace และ Cyberpunk/Neon สำหรับ Gameplay.
    - **Animation:** ใช้ Framer Motion และ React Three Fiber สำหรับการเปลี่ยนผ่านที่ลื่นไหล (Seamless Transitions).
        - **Planetary Surface:** บังคับใช้ `Mech` ในการสำรวจและทำภารกิจ.
        - **Interior/Precision Salvage:** บังคับใช้ `Pilot` (Character) ในการลงจากหุ่นเพื่อเก็บไอเทมหายาก.
    - **Combat Mode:** คำนวณผลการต่อสู้ตาม Stats ของ Mothership, Mech, และ Pilot.
    - **Star Discovery:** ระบบ Trigger สำหรับการขยายจักรวาลตามเงื่อนไขที่กำหนด.

### 3.5 Economy & Revenue Service
- **หน้าที่:** จัดการระบบการเงินและ Monetization.
- **Logic:** ระบบ Season Pass, การขายไอเทม (รวมถึง Pilot Gear), และการตรวจสอบ Revenue Flow (No USDT Out policy).

### 3.6 Auth & User Service
- **หน้าที่:** จัดการ Identity และ Profile.
- **Logic:** 
    - **Customizable Profile:** จัดการข้อมูลการปรับแต่งหน้า Profile และการเลือกโชว์ไอเทม (Showcase).
    - **Social Graph:** ระบบเพื่อน (Friendship) และการติดตาม (Following).
    - **Web3 Auth:** เชื่อมต่อกับ Privy/Dynamic สำหรับ Web3 Auth และเก็บข้อมูล Metadata ของผู้เล่น.

### 3.11 Notification & Bot Service
- **หน้าที่:** จัดการการสื่อสารภายนอกและระบบแจ้งเตือน.
- **Logic:** 
    - **Multi-Channel Dispatcher:** ส่งการแจ้งเตือนผ่าน Web Push, Discord Webhooks, และ X API.
    - **Telegram Bot:** บอทโต้ตอบสำหรับเช็คสถานะการสำรวจและรับข่าวสาร Patch ใหม่.
    - **Social Sharing:** ระบบสร้างภาพ/ลิงก์สำหรับแชร์ความสำเร็จไปยัง Social Media.

### 3.7 Pilot Service
- **หน้าที่:** จัดการข้อมูลตัวละครและอุปกรณ์สวมใส่ (Gear).
- **Logic:** จัดการค่า O2, ความทนทานของชุด, และการอัปเกรดอาวุธ (Swords/Guns).
- **Logic:** เชื่อมต่อกับ Privy/Dynamic สำหรับ Web3 Auth และเก็บข้อมูล Metadata ของผู้เล่น.

### 3.7 Colony & Exploration Service
- **หน้าที่:** จัดการระบบฐานที่มั่นและการสำรวจ.
- **Logic:** 
    - **Colony Management:** การอัปเกรดสิ่งอำนวยความสะดวกและการเคลื่อนย้าย Colony.
    - **Exploration Logic:** การคำนวณการใช้เชื้อเพลิง (Fuel) และการสุ่มพบเจอเหตุการณ์ (Encounters) ในแต่ละ Stage.

### 3.8 Salvage & Research Service
- **หน้าที่:** จัดการระบบการกู้ซากและการวิจัยเทคโนโลยี.
- **Logic:** 
    - **Salvage Processing:** การตัดสินใจจัดการกับยูนิตที่จับได้ (Sell/Research/Scrap).
    - **Resource Refining:** ระบบแปรรูป Scrap Metal ให้เป็นวัสดุซ่อมแซม (Repair Kits) หรือเชื้อเพลิง.
    - **Tech Tree:** การปลดล็อกความสามารถใหม่จากการวิจัยซากศัตรู.

### 3.9 Hangar & Maintenance Service
- **หน้าที่:** จัดการการซ่อมบำรุงหุ่นและยานพาหนะ.
- **Logic:** 
    - **Repair Logic:** การคำนวณทรัพยากรและเวลาที่ใช้ในการซ่อมแซมตามระดับความเสียหาย (Durability).
    - **Refuel & O2 Management:** การเติมเชื้อเพลิงยานแม่และออกซิเจนให้นักบิน.
    - **Visual Restoration:** การอัปเดต Metadata เพื่อลบรอย Wear & Tear เมื่อมีการซ่อมแซมเต็มรูปแบบ.

### 3.10 Story & Narrative Service
- **หน้าที่:** ควบคุมเนื้อเรื่องและภารกิจหลัก.
- **Logic:** 
    - **Creator Studio Engine:** ระบบจัดการ "Style Guides" และ "Comprehensive Templates" (Character, Vehicle, Environment, UI, VFX).
    - **Staging & Sandbox Logic:** ระบบจำลอง Game State สำหรับการทดสอบ Patch ก่อน Deploy จริง.
    - **Universe Analytics Engine:** ระบบรวบรวมและวิเคราะห์ข้อมูลพฤติกรรมผู้เล่น (Heatmaps, Death Rates, Salvage Trends).
    - **Rollback & Automation:** ระบบจัดการ Versioning ของจักรวาล และ Procedural Logic สำหรับเหตุการณ์ย่อยอัตโนมัติ.
    - **Context Patch Integration:** ระบบรองรับการโหลดข้อมูล Lore และ Event ใหม่ ๆ จาก Creator เพื่อเปลี่ยน Game State ของทั้งจักรวาล.
    - **Story Progression:** การตรวจสอบเงื่อนไขการผ่านด่านและมอบรางวัลแบบ Fixed.
    - **AI Combat Logs:** การส่งข้อมูลการต่อสู้ให้ AI เพื่อสร้าง Narrative Log ที่เป็นเอกลักษณ์.

## 4. Frontend Architecture (Next.js & 3D Engine)

เพื่อให้ได้ประสบการณ์แบบ Seamless และ High-Fidelity เราจะใช้ Stack ดังนี้:

### 4.1 Core Stack
- **Framework:** Next.js 15+ (App Router)
- **3D Engine:** **React Three Fiber (R3F)** + Three.js
- **Renderer:** **WebGPU** (Fallback to WebGL 2) เพื่อประสิทธิภาพสูงสุดในการเรนเดอร์ Cockpit และ Atmospheric Entry.
- **State Management:** Zustand (สำหรับ Game State ที่รวดเร็ว) และ React Query (สำหรับ Server State).

### 4.2 Seamless Transition Logic
- **Single Canvas Architecture:** ใช้ Canvas เดียวกันทั้งแอปเพื่อหลีกเลี่ยงการ Re-mount 3D Scene เมื่อเปลี่ยนหน้า.
- **Camera Interpolation:** ใช้การคำนวณพิกัดกล้องเพื่อทำ Smooth Zoom/Pan ระหว่าง Mothership -> Mech -> Pilot EVA.
- **Asset Preloading:** ใช้ระบบ Background Loading สำหรับโมเดล 3D และพื้นผิวดาวเคราะห์ขณะที่ผู้เล่นกำลังดูหน้าจอ Radar.

## 5. Microservices Architecture (Alternative/Evolution)

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
*   `status`: enum (VIRTUAL, MINTING, ON_CHAIN, BURNED) - รองรับระบบ Hybrid State

#### 6. Energy & Resource Entity
*   `user_id`: UUID (FK to User)
*   `standard_energy`: int (Off-chain, daily refresh)
*   `premium_energy`: int (On-chain/Paid)
*   `last_refresh`: timestamp

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
2.  **Game Engine:** ตรวจสอบ Energy (Standard/Premium) -> หักแต้มพลังงาน.
3.  **Game Engine:** คำนวณระยะทางและสุ่มเหตุการณ์ (Events).
4.  **Game Engine:** หากเจอศัตรู -> คำนวณผลการต่อสู้ (Stat-based).
5.  **AI Service:** Gen ภาพ "Action Shot" ของการปะทะ.
6.  **Orchestrator:** บันทึก Combat Log (Status: `TEMP`) -> ตั้งเวลา Housekeeping (7 วัน).

#### Flow C: Virtual-to-On-chain (V2O) Promotion
1.  **User:** เลือก Virtual Mech -> กด "Mint to Chain" (จ่าย Gas/USDT).
2.  **Orchestrator:** เปลี่ยนสถานะ Mech เป็น `MINTING` -> ล็อคการใช้งานชั่วคราว.
3.  **Blockchain Service:** ส่ง Transaction ไปยัง Base L2 (Mint ERC-721 + Setup ERC-6551).
4.  **Blockchain Service:** รอการ Confirm จาก Chain.
5.  **Orchestrator:** อัปเดตสถานะเป็น `ON_CHAIN` -> ปลดล็อคให้ใช้งานหรือลงขายใน Marketplace ได้.
