---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]
inputDocuments:
  - "_bmad-output/analysis/product-brief-Project-0-2025-12-21.md"
  - "_bmad-output/analysis/research/technical-Project-0-research-2025-12-21.md"
documentCounts:
  briefs: 1
  research: 1
  brainstorming: 0
  projectDocs: 0
workflowType: 'prd'
lastStep: 11
project_name: 'Project-0'
user_name: 'Vatcharin'
date: '2025-12-21'
---

# Product Requirements Document - Project-0

**Author:** Vatcharin
**Date:** 2025-12-21

## 1. Executive Summary

### 1.1 Product Vision
Project-0 is a sustainable Crypto Web Game that bridges traditional gaming mechanics with AI-driven innovation. It features a hybrid economy designed to eliminate "FOMO" cycles and high entry barriers by utilizing a **Web2.5 Onboarding strategy** (Social Login first, Wallet later). It utilizes AI to generate unique, seasonal NFT assets (**Vehicles**: Mechs, Tanks, Ships) with genuine aesthetic and functional rarity. The game is an **Intense, Evolving Live-Service Universe** where the creator gradually releases a dark, gripping narrative, new star systems, and context-driven patches to keep the exploration loop fresh and meaningful.

### 1.2 Project Classification
- **Project Type:** Blockchain/Web3 & Web Application
- **Domain:** Gaming & Fintech (Hybrid)
- **Complexity:** High
- **Key Differentiator:** AI-Powered Seasonal Rarity, Direct-to-System Hybrid Economy, and **Unified Vehicle & Pilot System**.

### 1.3 Strategic Alignment
The project aligns with the goal of creating a long-term, sustainable ecosystem where players can enter for free, progress through effort (Grinding), and participate in a high-value collector's market driven by AI-generated scarcity and **continuous world expansion**.

## 2. Success Criteria

### 2.1 User Success
- **Zero-Friction Entry:** Players can start playing within seconds using Google/Email without needing a crypto wallet or gas fees.
- **The "Aha!" Moment:** Players receive a high-quality AI-generated Vehicle or Item for the first time, experiencing the thrill of unique visual ownership and social prestige.
- **Emotional Outcome:** A deep sense of "Collection Pride" and "Social Dominance," especially for rare seasonal assets that trigger envy in other players.
- **Immersion:** Players feel like "Operatives" and "Commanders" caught in an intense, dark narrative where every mission advances a gripping plot.
- **Visual Pride:** Players feel a "Gasp" moment when they see their unique, high-fidelity AI-generated Vehicle for the first time—a design that is uniquely theirs and visually stunning.

### 2.2 Business Success
- **Aggressive Early Revenue:** Maximizing revenue through a tiered Season Pass model (Standard and Season Pass+).
- **Monetization Efficiency:** High conversion from free-to-play to premium tiers via exclusive benefits like 2x Exploration Rate (with daily caps) and Dynamic Profile Frames.
- **Sustainable Economy:** A "Direct-to-System" model that prevents hyperinflation while maintaining high asset value for premium users.

### 2.3 Technical Success
- **AI Generation Efficiency:** Implementing a "Time-Gated Assembly" system to manage GPU load and prevent spamming.
- **Hybrid 2D/3D Immersion:** Delivering High-Fidelity AI assets (2D) for visual storytelling while using **WebGPU + React Three Fiber** for a global 3D space background and an immersive 3D Cockpit and Bastion experience.
- **Single-Page Game Loop (Unified Controller):** Moving away from traditional web routing (`/bastion`, `/explore`) to a unified, state-driven game controller using **XState v5**. This ensures seamless transitions, persistent audio/state, and a true "Game Client" feel.
- **Design Intelligence:** Utilizing [UI/UX Pro Max](https://github.com/nextlevelbuilder/ui-ux-pro-max-skill/) standards to ensure professional-grade aesthetics, including Glassmorphism, Bento Grids, and industry-specific color palettes for a premium feel.
- **Modular NFT Ownership:** Utilizing **ERC-6551 (Token Bound Accounts)** to allow Vehicles to "own" their equipment, simplifying marketplace trading and inventory management.
- **Anatomical Equipment System:** Modules are mapped to specific anatomical slots (HEAD, CORE, ARM_L, ARM_R, LEGS) on a silhouette in the Bastion UI.
- **Hybrid State Management (V2O):** Implementing a "Manifested-to-On-chain" model where gameplay and asset creation are off-chain by default (Server-side), with on-chain minting required only for trading or permanent storage.
- **System Resilience:** Using Saga patterns and distributed tracing to ensure 100% consistency between Web2 (AI/DB) and Web3 (On-chain) states.
- **AI Contextual Intelligence:** Implementing **Model Context Protocol (MCP)** to provide the AI Narrative Engine with real-time game state for dynamic event generation.

## 3. Product Scope

### 3.1 MVP - Minimum Viable Product
- **Multi-Auth Identity System (Web2.5 Hybrid):**
    - **Guest Login:** Play immediately with a single click. Identity tied to `guest_id` (LocalStorage).
    - **Traditional Login:** Signup/Login with Username, Email, and Password.
    - **Social Login:** Google/Email via Privy (Primary Web2.5 path).
    - **Account Binding:** Seamlessly upgrade Guest -> Traditional/Social without losing progress.
    - **Late-Binding Wallet:** Link an external wallet (MetaMask, etc.) only when ready to mint NFTs.
- **Player Identity:** Start as a **Void Scavenger** with a **Resonance Suit** and a **Starter Pack**. Players must complete a **Character Creation** process (Name, Gender, Face, Hair) before receiving their first ship. This identity is tied to a **Character** instance, allowing for multiple characters per user.
- **Mothership Engineering Philosophy:** Motherships (The Bastion) are defined by two independent, non-linear systems:
    1. **Teleport System (Dimensional Tech):** For instant travel across dangerous sectors. High energy cost and instability. Some ships cannot install this due to structural mass.
    2. **Atmospheric Entry System (Structural Tech):** For safe landing on planets. Manages heat and gravity. Independent of Teleport capabilities.
    - **Strategic Choice:** Players choose between Speed (Teleport) vs. Safety (Entry) vs. Cost. Not a linear upgrade.
- **Unified Vehicle & Pilot System:** Motherships (The Bastion), Vehicles (Mechs, Tanks, Ships), and Pilot Gear with distinct gameplay roles.
- **Deep Gameplay Expansion:**
    - **Emergency Retrieval Protocol:** A fail-safe system for resource depletion (0 Fuel/O2) that auto-warps the pilot back to the Bastion with penalties (Stress, Critical Fatigue, and Reward loss).
    - **Neural Overdrive (Active Skills):** Tactical skills like **Overclock** (+30% ECP) and **Emergency Repair** (Restore HP) powered by **Neural Energy (NE)**.
    - **Damage Matrix:** Elemental damage types (**Kinetic, Energy, Void**) with specific strengths and weaknesses.
- **Engineering & Assembly System:** A story-driven crafting system (Synthesize) where asset generation requires resources, time, and engineers. Unlocked after the first mission.
- **Tiered AI Assets:** 
    - **Tiers (T1 - T5):** Power levels.
    - **Rarity Classes:** Standard (White), Refined (Green), Prototype (Blue), Relic (Purple), Singularity (Gold).
- **V2O Minting Rules:** Manifested assets must meet specific criteria (Epic+ rarity, 10+ Expeditions, >80% Durability) to be eligible for on-chain minting.
- **Void Signals (Gacha):** A "Hoyoverse-style" pull system for high-tier assets with a pity mechanism.
- **Seasonal Temporal Wormholes:** Limited-time exploration zones requiring specific gear compatibility, offering exclusive "Lost Tech" and "Rare DNA Fragments."
- **Complex Combat Engine:** Stat-based battle system integrating vehicle attributes (Bastion, Vehicle, Pilot) and AI-generated item "Options."
    - **Combat Power (CP):** A weighted sum of stats: `(ATK*2) + (DEF*2) + (HP/10)`.
    - **Effective CP (ECP):** `(Vehicle_CP + Exosuit_CP) * Suitability_Mod * Resonance_Sync * (1 - Fatigue_Penalty) * Synergy_Mod`.
    - **Real Combat Integration:** Transitioning from placeholder combat to a functional system where encounters are linked to real backend data.
    - **Backend Enemy Seeding:** A system to generate and persist specific NPC enemies (e.g., Striker, Guardian, Scout) within combat encounters, allowing for persistent enemy stats and unique loot tables.
- **Monetization Engine:** Season Pass, Void Signals, and Minting fees for "Void-Touched" (Premium) items.
- **Discord Integration:** Real-time alerts for star discoveries and achievement broadcasting.

### 3.2 Growth Features (Post-MVP)
- **Marketplace:** Peer-to-peer trading of Minted assets (NFTs).
- **Dynamic NFT Evolution:** Profile frames and assets that evolve based on seasonal achievements.
- **Advanced Social Systems:** Guilds, territory control, and integrated marketplaces.
- **Automated Star Discovery:** Procedural generation of new exploration zones.

### 3.3 Vision (Future)
- **Cross-Chain Interoperability:** Expanding the asset ecosystem to multiple L2s.
- **Full AI Universe:** Procedural star systems and lore generated dynamically by AI.

## 4. User Journeys

### 4.1 Journey 1: Somchai - The Grinder (The Foundation)
Somchai is a dedicated player who enjoys steady progress. He starts for free, farming basic resources like "Scrap Metal" and "Energy" through daily exploration with his **Starter Pack**. After two weeks of consistent effort, he gathers enough materials to initiate his first "Mech Assembly." He experiences a 24-hour waiting period, simulating the engineering process. When complete, he receives a unique AI-generated Mech with a "Digital Gold" camo—a rare find that he proudly showcases in Discord, validating his hard work and encouraging him to aim for higher-tier upgrades.

### 4.2 Journey 2: Alex - The Discord Socialite (The Catalyst)
Alex thrives on community engagement and being "first." He monitors the `#star-discovery` channel for real-time alerts. When a new "Abandoned Planet" is detected, he rallies his guild. Using his **Season Pass+** 2x Exploration Buff, his team reaches the planet first. After defeating the Raid Boss, Alex receives a "Mysterious Blueprint." He chooses the **USDT Premium Build** to ensure the highest quality AI generation. He shares the live "Assembly Status" link in Discord, building hype until a high-fidelity "Light-Wing Mech" is revealed, cementing his status as a community leader.

### 4.3 Journey 3: Sarah - The Collector/Trader (The Value Seeker)
Sarah focuses on the long-term value of AI-generated art. As "Season 1: Iron Age" nears its end, she realizes certain AI designs will never be produced again. She browses the **Web-Based Marketplace** and purchases a "Key" for a rare "Rusty Chrome" Mech from a Grinder. She uses the **Dynamic NFT Evolution** system to apply a seasonal prestige frame, increasing its aesthetic appeal. Sarah holds the asset, knowing its rarity will increase as the player base grows, feeling pride in her curated collection of unique digital masterpieces.

### 4.4 Journey 4: Marcus - The Architect (Seasonal Planning)
Marcus มุ่งเน้นไปที่การวางโครงสร้างจักรวาลในระยะยาว เขาไม่ได้ปรับเปลี่ยนเกมแบบ Real-time แต่จะทำงานเป็นรอบ **Season (เช่น ทุก 3 เดือน)**:
1. **Pre-Season:** ออกแบบ "Visual DNA" และ Meta ของ Season ถัดไป (เช่น ปรับสมดุลอาวุธประเภท Kinetic ให้แรงขึ้น)
2. **Patch Deployment:** ปล่อย Patch ใหญ่เพื่อเปิดโซนใหม่, เพิ่มรูหนอน (Wormholes), และอัปเดตเนื้อเรื่องผ่าน AI Narrative Engine
3. **Monitoring:** ติดตาม Feedback และข้อมูลเศรษฐกิจเพื่อนำไปวางแผนใน Patch หรือ Season ถัดไป

### 4.5 Journey 5: The Adaptive Economy (Dynamic Scarcity)
ระบบใช้ **Dynamic Adjustment** ในเชิงเศรษฐศาสตร์ (Bitcoin-style) เพื่อรักษาคุณค่าของไอเทม:
- **Resource Scarcity:** หากใน Season นี้มีการขุดแร่ "Void Crystal" มากเกินไป ระบบจะค่อยๆ ปรับความยากในการหาเพิ่มขึ้นโดยอัตโนมัติ เพื่อป้องกันเงินเฟ้อ
- **Rarity Balancing:** อัตราการดรอปของไอเทมระดับ Singularity จะถูกปรับตามจำนวนที่มีอยู่ในระบบ เพื่อให้ของแรร์ยังคงความแรร์ตลอดทั้ง Season

### 4.6 Journey Requirements Summary
- **Core Gameplay:** Resource farming, stat-based combat, and buff-augmented exploration.
- **Engineering System:** Time-gated assembly, tiered AI quality (Premium/Standard), and crafting status sharing.
- **Economy & Marketplace:** Web-based trading, seasonal scarcity logic, and Dynamic NFT frames.
- **Adaptive Logic:** Bitcoin-style difficulty adjustment for resources, rarity, and assembly time.
- **Management Tools:** Admin dashboard, bot detection, and real-time economy balancing controls.

## 5. Domain-Specific Requirements

### 5.1 Gaming & Fintech Hybrid Overview
Project-0 operates as a "Closed-Loop Premium Economy." USDT flows into the system via premium services and player-to-player trading but never flows out from the platform to users (No Payout/Withdrawal). This model ensures long-term sustainability and simplifies regulatory compliance.

### 5.2 Revenue Streams & Marketplace
- **Direct Revenue (Primary):** USDT income from Season Pass, Season Pass+, and Premium AI Assembly services.
- **Marketplace Fees (Secondary):** A percentage-based transaction fee on all NFT (Key) trades between players. The marketplace is non-custodial, with USDT moving directly between user wallets.
- **No-Withdrawal Policy:** The platform does not handle user payouts, reducing AML risks and the need for complex KYC for standard gameplay.

### 5.3 Economic Drivers: Pride, Scarcity & Legacy
- **Pride (Social Prestige):** Visual Aura effects for premium assets and Dynamic Profile Frames that reflect seasonal achievements.
- **Scarcity (Asset Value):** Seasonal asset retirement (no re-production) and Bitcoin-style Difficulty Adjustment for resources and rarity.
- **Legacy (On-chain History):** 
    - **Seasonal Hall of Fame:** Permanent on-chain record of top collectors and discoverers.
    - **On-chain Lore:** Metadata that records an asset's battle history and previous famous owners, turning NFTs into "Digital Antiques."

### 5.4 Compliance & Security
- **Light KYC:** Implementation of Discord/Email-linked accounts for basic bot prevention.
- **Smart Contract Integrity:** Use of OpenZeppelin standards for ERC-721 (NFTs) and Payment Gateways.
- **Transparency Dashboard:** A public "Difficulty & Rarity" dashboard to ensure players trust the self-balancing economic logic.

## 6. Innovation & Novel Patterns

### 6.1 Detected Innovation Areas
- **AI-Powered Seasonal Rarity:** Generative AI creating unique, time-limited assets that become "Digital Antiques."
- **Bitcoin-style Difficulty Adjustment:** Self-balancing economic and technical control mechanism.
- **Closed-Loop Premium Economy:** A sustainable USDT-in, No-USDT-out model focused on player-to-player value transfer.

### 6.2 Market Context & Competitive Landscape
- **First-Mover Advantage:** Pioneering the use of adaptive difficulty and HITL AI curation in Web3 gaming.
- **Copycat Risk Mitigation:** Focus on "Community Legacy" and "On-chain Lore" to create non-replicable value.

### 6.3 Validation & Quality Control Approach
- **Human-in-the-loop (HITL) AI Curation:** A "Curator Dashboard" where the project owner (Head Engineer) approves AI-generated assets from a batch of options before minting. This ensures 100% theme consistency and visual quality.
- **Difficulty Trust:** Validated through a public Transparency Dashboard showing real-time system difficulty metrics.

### 6.4 Risk Mitigation
- **Economy Beta Test:** Simulating difficulty adjustments under high-load scenarios before full launch.
- **AI Fallback:** Pre-generated high-quality asset sets to maintain experience if real-time generation fails or quality drops.

## 7. Project-Type Specific Requirements

### 7.1 Blockchain & Web3 (Base L2)
- **Chain Selection:** Base L2 (Ethereum Layer 2) for low fees and high throughput.
- **Wallet & Onboarding:**
    - **Hybrid Approach:** Support for MetaMask (Power Users) and Social Login (X, Google) for mainstream players.
    - **Account Abstraction (AA):** Implementation of ERC-4337 or similar to enable gasless transactions or simplified wallet management for social login users.
- **Smart Contract Strategy:**
    - **Gas Optimization:** High-priority optimization of Solidity code (e.g., using `uint256` efficiently, minimizing storage writes, using `calldata`).
    - **Standards:** ERC-721 for NFTs (Mechs/Items) with metadata extensions for lore and battle history.
- **Security:** Smart contracts will undergo a security audit prior to mainnet deployment to ensure asset safety and economic integrity.

### 7.2 Web Application (Next.js/React)
- **Frontend Stack:** Next.js (React) for a performant, SEO-friendly, and scalable user interface.
- **Mobile-First Design:** 100% support for Mobile Browsers with a fully responsive layout, ensuring a seamless experience across smartphones, tablets, and desktops.
- **Real-time Synchronization:**
    - **Live Updates:** Use of WebSockets (or similar technology) to provide real-time updates for "Assembly Status," "Star Discoveries," and "Marketplace Activity" directly on the web dashboard.
    - **Discord Sync:** Bi-directional sync between game state and Discord notifications.

### 7.3 Technical Architecture Considerations
- **Orchestration:** Go-based backend orchestrator managing the Saga pattern between Web2 (AI/DB) and Web3 (On-chain) states.
- **AI Pipeline:** Integration with Fal.ai/Replicate for FLUX.1 generation, managed via a secure HITL (Human-in-the-loop) admin interface.
- **Data Consistency:** Distributed tracing and robust error handling to ensure that a successful AI generation always results in a valid on-chain mint or state update.

## 8. Project Scoping & Phased Development

### 8.1 MVP Strategy & Philosophy
**MVP Approach:** Revenue-Driven & Experience MVP. เน้นการสร้างรายได้จากความภูมิใจในการครอบครอง (Pride) และระบบเศรษฐกิจที่หมุนเวียนได้จริงตั้งแต่วันแรก โดยมี Game Loop ที่สมบูรณ์ (Explore -> Combat -> Craft).

### 8.2 MVP Feature Set (Phase 1)
**Core User Journeys Supported:**
- Somchai (Grinder): สามารถฟาร์มของผ่านการ Explore และ Combat พื้นฐานได้
- Alex (Socialite): แย่งชิงการค้นพบดาวใหม่และโชว์ผลงานการประกอบร่าง
- Sarah (Collector): เริ่มสะสม Seasonal Mechs รุ่นแรก

**Must-Have Capabilities:**
- **Core Loop:** ระบบสำรวจ (Exploration) -> ระบบต่อสู้พื้นฐาน (Stat-based Combat) -> การดรอปทรัพยากรที่แตกต่างกันตามผลการต่อสู้
- **Engineering System:** ระบบประกอบร่าง (Assembly) แบบรอเวลา (Time-gated) พร้อมเกรด AI (Standard/Premium)
- **Monetization:** Season Pass & Season Pass+ (2x Buff, Social Prestige)
- **Marketplace:** ระบบ P2P Trading สำหรับ "Key" (NFT) บนหน้าเว็บ
- **Onboarding:** Social Login (X, Google) พร้อม Embedded Wallet เพื่อความรวดเร็ว

### 8.3 Post-MVP Features
**Phase 2 (Growth):**
- **Visual Combat:** ระบบต่อสู้ที่มีกราฟิกและการเคลื่อนไหวที่สวยงามขึ้น
- **Account Abstraction (AA):** ระบบ Gasless Transactions เต็มรูปแบบ
- **Social Features:** ระบบกิลด์, First Discovery Tag (Metadata), และ Visual Aura สำหรับสายเปย์
- **Dynamic Evolution:** Profile Frame ที่เปลี่ยนไปตามความสำเร็จใน Season

**Phase 3 (Expansion):**
- **Cross-chain Expansion:** รองรับการเล่นและเทรดข้าม Chain
- **Automated Universe:** ระบบสร้างดาวเคราะห์และเนื้อเรื่องอัตโนมัติด้วย AI
- **Territory Control:** การยึดครองพื้นที่และทรัพยากรในระดับกิลด์

### 8.4 Risk Mitigation Strategy
- **Technical Risks:** ใช้ Saga Pattern จัดการความสอดคล้องของข้อมูล AI/Blockchain และใช้ Embedded Wallet ลดความซับซ้อนในช่วงแรก
- **Market Risks:** ใช้ระบบ Difficulty Adjustment (Bitcoin-style) ควบคุมอัตราการดรอปทรัพยากรจาก Combat ไม่ให้เฟ้อ

## 9. Functional Requirements

### 9.1 User Onboarding & Identity
- **FR1:** ผู้เล่นสามารถเข้าสู่ระบบผ่าน Social Accounts (X, Google) ได้
- **FR2:** ระบบสามารถสร้าง Embedded Wallet ให้กับผู้เล่นใหม่โดยอัตโนมัติหลังการเข้าสู่ระบบ
- **FR3:** ผู้เล่นสามารถเชื่อมต่อ MetaMask Wallet เพื่อใช้งานในระดับ Power User ได้
- **FR4:** ระบบสามารถผูกบัญชี Discord เข้ากับบัญชีผู้เล่นเพื่อรับการแจ้งเตือนได้
- **FR41:** หลังจาก Login ครั้งแรก ผู้เล่นต้องสร้างตัวละคร (Character Creation) โดยกำหนดชื่อตัวละคร (แยกจาก Username), เพศ, และรูปลักษณ์ (หน้าตา, ทรงผม)
- **FR42:** ระบบตัวละครต้องรองรับการมีหลายตัวละครต่อหนึ่งบัญชี (Character Instances) เพื่อรองรับการสลับตัวละครหลักจากระบบกาชาในอนาคต
- **FR43:** หน้า Hangar ต้องแสดง **Pilot ID Badge** ที่ระบุรูป Avatar, ชื่อตัวละคร, และ Rank/Level ของตัวละครที่ใช้งานอยู่

### 9.2 Exploration & Discovery
- **FR5:** ระบบ Multi-Stage Exploration (Mothership -> Mech -> Pilot EVA) ที่บังคับใช้ประเภท Vehicle และอุปกรณ์ต่างกันตามระยะทางและสภาพแวดล้อม
- **FR6:** ระบบสุ่มเหตุการณ์ (Events) และอุบัติเหตุ (Accidents) ระหว่างการสำรวจ โดยใช้โครงสร้าง **Expedition & Encounters**:
    - **Expedition:** แทนเซสชันการสำรวจในหนึ่งพื้นที่ (เช่น การลงจอดบนดาว)
    - **Encounter:** แทนเหตุการณ์ย่อยใน Timeline (เช่น การพบซากหุ่น, การถูกโจมตี, การขุดแร่)
    - **AI Narrative:** ใช้ AI สร้างเนื้อเรื่องตามบริบท (Context-aware AI Narrative) ผ่าน **Model Context Protocol (MCP)**
- **FR7:** ระบบสามารถคำนวณระยะเวลาการสำรวจและโอกาสสำเร็จ (Success Probability %) ตาม Stat ของอุปกรณ์ที่สวมใส่
- **FR8:** ผู้เล่นที่มี Season Pass+ สามารถรับบัฟความเร็ว/ประสิทธิภาพการสำรวจเพิ่มขึ้น 2 เท่า (ภายใต้ Daily Cap)
- **FR9:** ระบบ Radar Scan & Fog of War:
    - **Hierarchical Map:** แบ่งระดับพื้นที่เป็น Sector -> Sub-Sector -> Planet Location
    - **Server-side Validation:** ผู้เล่นสามารถสแกนหาพิกัดดาวเพื่อดูระดับความอันตราย (Threat Level) เบื้องต้น โดยใช้ **Web2 Backend Fog of War** (Go + PostgreSQL) เพื่อความรวดเร็วและป้องกันการโกง
- **FR27:** ระบบ Pre-flight Briefing: แสดงความเสี่ยงและคำเตือนก่อน Launch (เช่น ขาดเกราะกันความร้อน หรือ O2 ไม่พอ) โดยผู้เล่นสามารถเลือก "เสี่ยง" (Proceed at Own Risk) ได้
- **FR33:** ระบบ Pilot Death & Penalty: หากนักบินเสียชีวิตระหว่าง EVA จะสูญเสียไอเทมที่เก็บได้ในรอบนั้นทั้งหมด และหุ่น Mech จะได้รับความเสียหายหนัก (Critical Damage)
- **FR34:** ระบบ Evolving Universe: ระบบรองรับการอัปเดต "Context Patches" จากผู้สร้าง (Creator) เพื่อเพิ่มดาวเคราะห์ใหม่, ศัตรูใหม่ และเนื้อเรื่องใหม่ตามฤดูกาล
- **FR40:** ระบบ Real-time Resource Consumption: ทุกการก้าวเดิน (Next Encounter) ในการสำรวจจะมีการหักทรัพยากรจริงจาก Pilot Stats (ค่าเริ่มต้น: O2 -15.0, Fuel -5.0) และบันทึกลงฐานข้อมูลทันที

### 9.3 Combat System (MVP - Immersive 1:1)
- **FR10:** ระบบการต่อสู้แบบ 1:1 (One Pilot, One Mech) เน้นความสมจริงและผลกระทบส่วนบุคคล (Personal Stakes)
- **FR11:** ระบบสามารถคำนวณผลการแพ้/ชนะ โดยใช้ค่าพลัง (Stats) ของ Mothership, Mech และ Pilot Gear
- **FR12:** ระบบสามารถมอบทรัพยากรและชิ้นส่วน NFT (Modular Salvage) ตามผลการต่อสู้
- **FR25:** ระบบ Combat Log & 3D Visualization: แสดงผลการต่อสู้ผ่านหน้าจอ Cockpit HUD (Three.js) พร้อมเอฟเฟกต์สั่นสะเทือนและ Glitch ตามความเสียหายจริง
- **FR35:** ระบบ Hangar & Maintenance: ผู้เล่นต้องทำการซ่อมแซม (Repair) และเติมทรัพยากร (Refuel/O2) ที่ยานแม่หลังจบภารกิจ โดยใช้ทรัพยากรที่ได้จากการ Salvage (Scrap Metal) มาแปรรูป (Refine) เป็นวัสดุซ่อมแซม

### 9.4 Engineering & Assembly (Modular AI Generation)
- **FR13:** ระบบ Modular NFT Assembly: ผู้เล่นสามารถประกอบร่างหุ่นและยานจากชิ้นส่วน NFT แยกชิ้น (Chassis, Weapons, Shields) โดยใช้มาตรฐาน **ERC-6551 (Token Bound Accounts)** เพื่อให้ NFT หลักเป็นเจ้าของชิ้นส่วนอุปกรณ์
- **FR14:** ระบบ AI Visual Synthesis: AI (FLUX.1) สามารถ "อ่าน" ชิ้นส่วน NFT ที่สวมใส่และสร้างภาพ Master Image ที่รวมทุกชิ้นส่วนเข้าด้วยกันตามบริบทของภารกิจ
- **FR15:** ระบบจัดการคิวการประกอบร่างและอัปเกรดชิ้นส่วนแบบรอเวลา (Time-gated)
- **FR16:** ระบบ Visual Wear & Tear: AI สร้างภาพที่สะท้อนความเสียหายจริง (รอยไหม้, รอยกระสุน) ลงใน NFT Metadata
- **FR30:** ระบบ Pilot Gear & Progression: ผู้เล่นสามารถอัปเกรดชุดนักบิน (Nomad Style), อาวุธ (Swords/Guns) และถัง O2 เพื่อใช้ในการสำรวจช่วง EVA

### 9.5 Economy & Marketplace
- **FR17:** ผู้เล่นสามารถซื้อ Season Pass และ Season Pass+ ด้วย USDT
- **FR18:** ผู้เล่นสามารถลงขายชิ้นส่วน NFT (Modular Parts) ใน Marketplace แบบ P2P ได้ (ต้องทำการ Mint ขึ้น Chain ก่อน)
- **FR19:** ระบบสามารถหักค่าธรรมเนียม (Marketplace Fee) จากการซื้อขายระหว่างผู้เล่น
- **FR20:** ระบบสามารถแสดงข้อมูลความยาก (Difficulty) และอัตราการดรอป (Rarity) บน Transparency Dashboard
- **FR31:** ระบบ Hybrid Energy: 
    - **Standard Energy (Off-chain):** รีเฟรชฟรีทุกวัน ใช้สำหรับการสำรวจและต่อสู้ทั่วไป
    - **Premium Energy / Overclock (On-chain):** ผู้เล่นสามารถจ่าย USDT/Token เพื่อซื้อพลังงานเพิ่มสำหรับการสำรวจที่เกินขีดจำกัดรายวัน (Exploration Boost)
- **FR32:** ระบบ Manifested-to-On-chain (V2O): สินค้าและหุ่นที่สร้างขึ้นใหม่จะเป็น "Manifested Assets" บน Server โดยอัตโนมัติ และผู้เล่นสามารถเลือก "Mint to Chain" เมื่อต้องการขายหรือโอนย้ายออกภายนอกเท่านั้น

### 9.7 Social & Engagement
- **FR37:** ระบบ Customizable Pilot Profile: ผู้เล่นสามารถปรับแต่งหน้า Profile, เลือกโชว์หุ่น Mech (Showcase), และแสดงรายการ NFT/Achievements ที่ภาคภูมิใจได้
- **FR38:** ระบบ Social Interaction: รองรับการเพิ่มเพื่อน (Friend System), การดู Profile ผู้เล่นอื่น, และการแชร์ความสำเร็จไปยัง Social Media (X, Discord)
- **FR39:** ระบบ Multi-Channel Notifications: แจ้งเตือนข่าวสาร, Patch ใหม่, และการค้นพบดาวผ่าน Web Push, Discord Webhook, และ Telegram Bot

### 9.6 Admin & Quality Control (HITL)
- **FR21:** Admin สามารถตรวจสอบและอนุมัติภาพที่ AI สังเคราะห์ขึ้น (HITL) ก่อนที่จะทำการ Mint หรืออัปเดต Metadata
- **FR22:** ระบบสามารถปรับระดับความยาก (Difficulty Adjustment) ของทรัพยากรและอัตราการดรอปได้โดยอัตโนมัติ (Economic Balancing)
- **FR23:** ระบบสามารถส่งการแจ้งเตือน Star Discovery ไปยัง Discord Channel ที่กำหนดได้ทันทีเมื่อมีการค้นพบใหม่
- **FR26:** ผู้เล่นสามารถจ่าย USDT เพื่อบันทึก Combat Log และภาพ Snapshot การสำรวจแบบถาวร (Permanent Save) ลงใน Profile/Lore
- **FR36:** ระบบ Creator Studio (The Architect Engine):
    - **Style Guide Management:** ผู้สร้างสามารถกำหนด "Visual DNA" ผ่าน AI Prompts และ LoRA สำหรับแต่ละ Season
    - **Context Patching & Staging:** ระบบจัดการการปล่อย Patch ใหญ่ (Seasonal Updates) พร้อมพื้นที่ทดสอบ (Sandbox/Staging)
    - **Universe Analytics:** ระบบมอนิเตอร์ข้อมูลเศรษฐกิจและพฤติกรรมผู้เล่นเพื่อใช้ในการออกแบบ Meta ใน Season ถัดไป
    - **Safety & Automation:** ระบบ One-Click Rollback สำหรับกู้คืนสถานะจักรวาล และระบบ Procedural Auto-Pilot สำหรับสร้างเหตุการณ์ย่อยอัตโนมัติ
    - **Seasonal Control:** เครื่องมือสำหรับกำหนดวันเปิด-ปิดรูหนอน และการปล่อยเนื้อเรื่องตาม Timeline ที่วางไว้ใน Patch

## 10. Progression & Growth (Measuring Strength)

เพื่อให้ผู้เล่นรู้สึกถึงความ "เก่งขึ้น" และมีเป้าหมายในระยะยาว Project-0 จึงมีระบบวัดความก้าวหน้าหลายระดับ:

### 10.1 Pilot Rank & Mastery (XP)
- **Experience Points (XP):** ได้รับจากการทำภารกิจสำเร็จ, การ Salvage ซากหุ่น, และการค้นพบดวงดาวใหม่ๆ
- **Pilot Rank:** การเลื่อนระดับจะปลดล็อก "Advanced Licenses" เพื่อให้สามารถขับ Mech, Aircraft หรือ Mothership ในระดับที่สูงขึ้นได้
- **Skill Points:** ทุกครั้งที่ Rank อัป จะได้รับแต้มเพื่ออัปเกรดความสามารถเฉพาะตัวของนักบิน (เช่น Agility, O2 Efficiency, Melee/Sidearm Mastery)

### 10.2 Modular Gear & "Options" (The Loot Loop)
- **AI-Generated Stats:** ชิ้นส่วน NFT แต่ละชิ้น (Arm, Chassis, Weapon) จะมีค่าพลังสุ่ม (Options) ที่สร้างโดย AI (เช่น +5% Kinetic Resistance, +10% Scanner Range)
- **Rarity Tiers:** แบ่งระดับความหายากเป็น Standard, Rare, Epic, Legendary และ Seasonal
- **Visual DNA:** ไอเทมระดับสูงจะมี Visual Traits ที่ซับซ้อนและสวยงามกว่า ซึ่งแสดงถึงความเก่งกาจที่มองเห็นได้ชัดเจนในสังคม

### 10.3 Mech Sync Rate (Synergy)
- **Familiarity:** การใช้ Mech หรือชิ้นส่วนเดิมซ้ำๆ จะเพิ่มค่า "Sync Rate"
- **Bonuses:** ค่า Sync Rate ที่สูงจะมอบ Buff พิเศษ เช่น +Evasion, +Accuracy หรือลดการใช้ Energy ซึ่งเป็นการตอบแทนผู้เล่นที่ดูแลและซ่อมแซมอุปกรณ์คู่ใจแทนการเปลี่ยนใหม่ตลอดเวลา

### 10.4 Mothership Engineering Matrix (Tech Web)
แทนที่ระบบ Tech Tree แบบเดิมด้วย **Engineering Matrix** ที่แบ่งเป็น 2 สายหลัก (Dual-Core Paths) ซึ่งผู้เล่นต้องเลือกทิศทางการพัฒนา:

#### 10.4.1 Teleport (Dimensional/Speed) Path
เน้นความรวดเร็วและการข้ามระยะทาง แต่มีความเสี่ยงสูงหากโครงสร้างไม่แข็งแรง:
1.  **Signal Booster I**: เพิ่มระยะการมองเห็น Encounter ล่วงหน้า +1 จุด
2.  **Dimensional Anchor**: ลดการใช้เชื้อเพลิง (Fuel) ระหว่างการ Teleport 10%
3.  **Void Stabilizer**: ลดโอกาสเกิด "Accident" ระหว่างการกระโดดข้ามมิติ 5%
4.  **Frequency Tuner**: เพิ่มโอกาสตรวจพบสัญญาณระดับ "Refined" ขึ้น 5%
5.  **Wormhole Navigator**: ปลดล็อกความสามารถ "Short-range Jump" (ข้ามได้ 1 Encounter ทันที)

#### 10.4.2 Atmospheric Entry (Structural/Safety) Path
เน้นความทนทานและการลงจอดที่ปลอดภัยในสภาพแวดล้อมที่โหดร้าย:
1.  **Heat Shielding I**: ลดการใช้ O2 ระหว่างการเข้าสู่ชั้นบรรยากาศ 15%
2.  **Reinforced Hull**: เพิ่มค่า HP เริ่มต้นให้ Mech หลังการลงจอด 10%
3.  **Shock Absorbers**: ลดโอกาสเกิด "Critical Damage" จากการลงจอดกระแทก 10%
4.  **Cargo Stabilizer**: เพิ่มโอกาส 5% ที่จะรักษาไอเทมทั้งหมดไว้ได้หากเกิดอุบัติเหตุ
5.  **Descent Thrusters**: ปลดล็อก "Precision Landing" (สามารถเลือกจุดลงจอดที่ปลอดภัยกว่าได้)

### 10.5 Starter Gear & Onboarding Loot
ผู้เล่นใหม่จะได้รับชุดอุปกรณ์เริ่มต้น (Starter Gear) ที่มี **Visual DNA** เฉพาะตัวเพื่อสร้างความผูกพัน:
- **Pilot Suit**: "Nomad-01" (Standard, Visual DNA: Scavenged Fabric)
- **Sidearm**: "Rusty Bolt" (Standard, Visual DNA: Industrial Scrap)
- **O2 Tank**: "Old Lung" (Standard, Visual DNA: Dented Steel)
- **Early Game Buff**: เพิ่มอัตราการดรอปไอเทมระดับ **Refined** ขึ้นเล็กน้อยในช่วง 3 ภารกิจแรกเพื่อช่วยในการตั้งตัว

## 11. Void Signals (Gacha System)

### 11.1 Gacha Mechanics (Hoyoverse-style Pity)
- **Relic Pity**: การันตีไอเทมระดับ **Relic** หรือสูงกว่าทุกๆ 10 Pulls
- **Singularity Pity**: การันตีไอเทมระดับ **Singularity** ทุกๆ 80 Pulls
- **Daily Signal**: ผู้เล่นได้รับสิทธิ์สุ่มฟรี 1 ครั้งทุก 24 ชั่วโมง (ไม่สะสม)

### 11.2 Technical Implementation & Logging
- **Enhanced Logging**: ทุกการสุ่มจะถูกบันทึกข้อมูลอย่างละเอียดเพื่อการตรวจสอบ (QA) และป้องกันการโกง:
    - `user_id`, `pull_type` (Daily/Paid), `seed`, `pity_counter_before/after`, `result_item_id`, `timestamp`
- **API Integration**: ยกเลิกการใช้ Mock Data ใน Frontend และเชื่อมต่อกับ Backend API จริงผ่าน `/api/v1/gacha/pull`
- **UI/UX Glitch Effect**: เพิ่มเอฟเฟกต์ Glitch และ Static Noise ในหน้า Reveal เมื่อสุ่มได้ไอเทมระดับ **Relic** หรือ **Singularity** เพื่อสร้างความตื่นเต้น

## 12. Functional Requirements (Updated)
... (FR เดิม)
- **FR41:** ระบบ Daily Signal มอบสิทธิ์สุ่มฟรี 1 ครั้งต่อวัน โดยตรวจสอบจาก `last_free_pull_at` ในฐานข้อมูล
- **FR42:** ระบบ Engineering Matrix บันทึกสถานะการปลดล็อก Node ในตาราง `mothership_upgrades`
- **FR43:** ระบบ Gacha Logging บันทึก Seed และ Pity State ลงในตาราง `gacha_history` ทุกครั้งที่มีการ Pull
- **Deep Space Access:** การอัปเกรดเครื่องยนต์และเกราะกันความร้อนของ Mothership เป็นวิธีเดียวที่จะเข้าถึงระบบดาวระดับสูงที่มี Loot ระดับตำนาน

### 10.5 Seasonal Achievements & Hall of Fame
- **Legacy Status:** ผู้เล่นที่ทำ Milestone สำคัญได้ (เช่น "First to Discover Star-X") จะถูกบันทึกชื่อไว้ใน Metadata ของจักรวาลถาวร
- **Seasonal Prestige:** กรอบโปรไฟล์และฉายาพิเศษที่แสดงถึงความสำเร็จในแต่ละ Season

## 11. Visual Allure & The "Cool" Factor (The Hook)

เพื่อให้เกมมีแรงดึงดูดมหาศาล (Mass Appeal) แม้ผู้เล่นจะยังไม่รู้ระบบการเล่น Project-0 จะเน้นความ "เท่" และ "ความภูมิใจในการครอบครอง" เป็นอันดับหนึ่ง:

### 11.1 Visual DNA & Seasonal Masterpieces
- **Creator-Defined Aesthetics:** ในแต่ละ Season ผู้สร้างจะกำหนด "Visual DNA" ที่เป็นเอกลักษณ์ (เช่น Season 1: Cyber-Samurai, Season 2: Bio-Mechanical Horror) ทำให้ไอเทมในแต่ละช่วงเวลามีคุณค่าทางศิลปะที่หาจากที่อื่นไม่ได้
- **Legendary Visual Traits:** ไอเทมระดับตำนาน (Legendary) จะมีเอฟเฟกต์พิเศษที่มองเห็นได้ชัดเจน เช่น เกราะโปร่งแสง (Translucent Armor), ปีกพลังงาน (Energy Wings), หรือพื้นผิวที่มีอนิเมชั่น (Animated Textures)
- **One-of-a-Kind Generation:** AI จะถูกโปรแกรมให้สร้าง "Masterpiece" ที่มีโอกาสเกิดน้อยมาก (เช่น 0.01%) ซึ่งจะมีดีไซน์ที่หลุดกรอบจากไอเทมทั่วไปอย่างสิ้นเชิง

### 11.2 The 3D Hangar Showcase
- **High-Fidelity 3D Preview:** ผู้เล่นสามารถดูหุ่นและยานของตัวเองใน Hangar แบบ 3D ที่รันด้วย WebGPU พร้อมแสงเงาและเงาสะท้อนระดับ AAA
- **Photo Mode & Social Sharing:** ระบบกล้องที่ปรับแต่งได้เพื่อถ่ายภาพ "Mech Snapshot" ที่สวยงามที่สุด พร้อมปุ่มแชร์ลง X/Discord/Telegram ได้ทันที
- **Public Profiles:** ผู้เล่นคนอื่นสามารถเข้ามา "เยี่ยมชม" Hangar ของเราเพื่อดูคอลเลกชันหุ่นเท่ๆ ได้ สร้างแรงบันดาลใจให้อยากได้ของแบบเดียวกัน

### 11.3 Pilot Identity & Customization
- **Pilot Gear Aesthetics:** ไม่ใช่แค่หุ่นที่เท่ แต่นักบิน (Pilot) ก็มีชุดเกราะและอาวุธ (Swords/Guns) ที่ดีไซน์มาอย่างประณีต
- **Appearance Customization:** ระบบเริ่มต้นรองรับการเลือกหน้าตาและทรงผมแบบ Pixel Art หรือ Icon เพื่อสร้างเอกลักษณ์เฉพาะตัว
- **Dynamic Posing:** ในหน้าโปรไฟล์ นักบินจะยืนโพสต์ท่าคู่กับยาน (Ship) หรือ Mech คู่ใจในสภาพแวดล้อมที่ AI เจนให้ตามดาวที่ผู้เล่นไปสำรวจล่าสุด

## 12. Non-Functional Requirements

### 10.1 Performance & Experience
- **NFR1 (Response Time):** หน้า Dashboard และการตอบสนองทั่วไปต้องโหลดเสร็จภายใน < 2 วินาที
- **NFR13 (Seamless Multi-Vehicle Transition):** การเปลี่ยนผ่านมุมมองระหว่างยานพาหนะ (Mothership -> Mech -> Pilot EVA) ต้องเป็นแบบไร้รอยต่อ (Seamless) โดยไม่มี Loading Screen:
    - **Mothership to Mech:** กล้องพุ่ง (Zoom-in) จากห้องควบคุมยานแม่เข้าสู่ Cockpit ของ Mech
    - **Mech to Pilot:** กล้องแพนออกจาก Cockpit เข้าสู่มุมมองบุคคลที่สาม (Third-person) ของนักบินขณะออกจากหุ่น
    - **Atmospheric Entry:** การเปลี่ยนผ่านจากอวกาศเข้าสู่ชั้นบรรยากาศต้องรันด้วย Real-time 3D Rendering (**WebGPU + React Three Fiber**) ตลอดกระบวนการ
- **NFR2 (AI Generation & Delivery):**
    - **Technical Latency:** ระบบต้องประมวลผลภาพ AI ให้เสร็จสิ้นก่อนเวลาที่กำหนดส่งมอบ (Reveal Time)
    - **Load-Adaptive Assembly:** ระบบสามารถปรับเปลี่ยนระยะเวลาการรอได้ตามปริมาณ Load ของ GPU
- **NFR3 (Blockchain Sync):** ข้อมูล On-chain ของชิ้นส่วน NFT ต้องสะท้อนบนหน้าเว็บภายใน < 5 วินาที
- **NFR15 (Fog of War Security):** ใช้ Web2 Backend Validation สำหรับ Fog of War เพื่อป้องกันการเปิดแมพจากฝั่ง Client โดยไม่ต้องใช้ ZKP ในช่วง MVP เพื่อลดความซับซ้อนและ Latency
- **NFR16 (AI Context Efficiency):** ใช้ **MCP (Model Context Protocol)** เพื่อให้ AI เข้าถึงข้อมูล Game State ในเครื่อง Server ได้โดยตรง ลดการส่งข้อมูลผ่าน Cloud และเพิ่มความเร็วในการสร้าง Narrative Event

### 10.2 Security & Anti-Bot
- **NFR4 (Anti-Bot Strategy):** 
    - ใช้ Captcha ในจุดที่มีการทำธุรกรรมสำคัญ
    - ใช้ Behavioral Analysis สำหรับการกระทำทั่วไป (Explore/Combat)
    - บังคับ Social Verification (Discord/X)
- **NFR5 (Smart Contract Security):** Smart Contract สำหรับ Modular NFTs ต้องผ่านการ Audit และมีระบบ Circuit Breaker
- **NFR6 (Data Protection):** ข้อมูลส่วนบุคคลและ Wallet Address ต้องถูกเข้ารหัสตามมาตรฐานสูงสุด

### 10.3 Scalability & AI Load Handling
- **NFR7 (Concurrency):** ระบบต้องรองรับผู้เล่นพร้อมกันได้อย่างน้อย 10,000 คน
- **NFR8 (AI Reliability/Backtrack):** 
    - ระบบต้องมีกลไก Idempotency สำหรับการสังเคราะห์ภาพ AI จากหลายชิ้นส่วน NFT
    - หากกระบวนการ AI Synthesis ล้มเหลว ต้องมีระบบ Auto-Retry หรือคืนสิทธิ์ให้ผู้เล่น
- **NFR14 (Visual Consistency):** ภาพที่ AI สังเคราะห์ต้องมีความถูกต้องตาม Metadata ของชิ้นส่วน NFT ที่สวมใส่จริง 100%

### 10.4 Reliability (ความเสถียร)
- **NFR9 (Uptime):** ระบบต้องมี Uptime ไม่ต่ำกว่า 99.9%
- **NFR10 (Data Consistency):** ใช้ Saga Pattern เพื่อรักษาความถูกต้องของข้อมูลระหว่าง Web2 และ Web3 100% (โดยเฉพาะการประกอบชิ้นส่วน NFT)
- **NFR11 (AI Guardrails):** ใช้ Structured Output และ RAG เพื่อป้องกัน AI Hallucination ในการสร้างเนื้อเรื่องอุบัติเหตุ
- **NFR12 (Housekeeping):** ระบบลบข้อมูลชั่วคราว (Temporary Logs) อัตโนมัติเพื่อรักษาประสิทธิภาพของฐานข้อมูล
