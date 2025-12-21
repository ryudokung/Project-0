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
Project-0 is a sustainable Crypto Web Game that bridges traditional gaming mechanics with AI-driven innovation. It features a hybrid economy designed to eliminate "FOMO" cycles and high entry barriers, utilizing AI to generate unique, seasonal NFT assets (Mechs, Tanks, Ships) with genuine aesthetic and functional rarity.

### 1.2 Project Classification
- **Project Type:** Blockchain/Web3 & Web Application
- **Domain:** Gaming & Fintech (Hybrid)
- **Complexity:** High
- **Key Differentiator:** AI-Powered Seasonal Rarity & Direct-to-System Hybrid Economy

### 1.3 Strategic Alignment
The project aligns with the goal of creating a long-term, sustainable ecosystem where players can enter for free, progress through effort (Grinding), and participate in a high-value collector's market driven by AI-generated scarcity.

## 2. Success Criteria

### 2.1 User Success
- **The "Aha!" Moment:** Players receive a high-quality AI-generated Mech or Item for the first time, experiencing the thrill of unique visual ownership and social prestige.
- **Emotional Outcome:** A deep sense of "True Ownership" and "Collection Pride," especially for rare seasonal assets that will never be generated again.
- **Immersion:** Players feel like "Engineers" and "Commanders" through the story-driven assembly process, rather than just clicking a button.

### 2.2 Business Success
- **Aggressive Early Revenue:** Maximizing revenue through a tiered Season Pass model (Standard and Season Pass+).
- **Monetization Efficiency:** High conversion from free-to-play to premium tiers via exclusive benefits like 2x Exploration Rate (with daily caps) and Dynamic Profile Frames.
- **Sustainable Economy:** A "Direct-to-System" model that prevents hyperinflation while maintaining high asset value for premium users.

### 2.3 Technical Success
- **AI Generation Efficiency:** Implementing a "Time-Gated Assembly" system to manage GPU load and prevent spamming.
- **Tiered Quality Control:** Delivering High-Fidelity AI assets for USDT transactions while maintaining theme consistency for in-game currency builds.
- **System Resilience:** Using Saga patterns and distributed tracing to ensure 100% consistency between Web2 (AI/DB) and Web3 (On-chain) states.

## 3. Product Scope

### 3.1 MVP - Minimum Viable Product
- **Multi-Vehicle System:** Mechs, Tanks, and Ships with distinct gameplay roles.
- **Engineering & Assembly System:** A story-driven crafting system where asset generation requires resources, time, and engineers.
- **Tiered AI Assets:** 
    - **Premium (USDT):** High-fidelity, complex AI designs.
    - **Standard (In-game):** Theme-consistent but simpler AI designs.
- **Complex Combat Engine:** Stat-based battle system integrating vehicle attributes and AI-generated item "Options."
- **Monetization Engine:** Season Pass and Season Pass+ with exploration buffs and social prestige items.
- **Discord Integration:** Real-time alerts for star discoveries and achievement broadcasting.

### 3.2 Growth Features (Post-MVP)
- **Dynamic NFT Evolution:** Profile frames and assets that evolve based on seasonal achievements.
- **Advanced Social Systems:** Guilds, territory control, and integrated marketplaces.
- **Automated Star Discovery:** Procedural generation of new exploration zones.

### 3.3 Vision (Future)
- **Cross-Chain Interoperability:** Expanding the asset ecosystem to multiple L2s.
- **Full AI Universe:** Procedural star systems and lore generated dynamically by AI.

## 4. User Journeys

### 4.1 Journey 1: Somchai - The Grinder (The Foundation)
Somchai is a dedicated player who enjoys steady progress. He starts for free, farming basic resources like "Scrap Metal" and "Energy" through daily exploration. After two weeks of consistent effort, he gathers enough materials to initiate his first "Mech Assembly." He experiences a 24-hour waiting period, simulating the engineering process. When complete, he receives a unique AI-generated Mech with a "Digital Gold" camo—a rare find that he proudly showcases in Discord, validating his hard work and encouraging him to aim for higher-tier upgrades.

### 4.2 Journey 2: Alex - The Discord Socialite (The Catalyst)
Alex thrives on community engagement and being "first." He monitors the `#star-discovery` channel for real-time alerts. When a new "Abandoned Planet" is detected, he rallies his guild. Using his **Season Pass+** 2x Exploration Buff, his team reaches the planet first. After defeating the Raid Boss, Alex receives a "Mysterious Blueprint." He chooses the **USDT Premium Build** to ensure the highest quality AI generation. He shares the live "Assembly Status" link in Discord, building hype until a high-fidelity "Light-Wing Mech" is revealed, cementing his status as a community leader.

### 4.3 Journey 3: Sarah - The Collector/Trader (The Value Seeker)
Sarah focuses on the long-term value of AI-generated art. As "Season 1: Iron Age" nears its end, she realizes certain AI designs will never be produced again. She browses the **Web-Based Marketplace** and purchases a "Key" for a rare "Rusty Chrome" Mech from a Grinder. She uses the **Dynamic NFT Evolution** system to apply a seasonal prestige frame, increasing its aesthetic appeal. Sarah holds the asset, knowing its rarity will increase as the player base grows, feeling pride in her curated collection of unique digital masterpieces.

### 4.4 Journey 4: Marcus - The Game Master (Admin/Ops)
Marcus ensures the universe remains balanced. He monitors the **Admin Dashboard** for anomalies. When he detects a bot cluster attempting to spam low-quality item generation, he triggers the **Difficulty Adjustment** mechanism, increasing the resource requirements and assembly time for small items temporarily. He also updates the **AI Prompt Library** for the upcoming season, ensuring fresh visual styles. His goal is to maintain a fair playing field and a stable economy for all participants.

### 4.5 Journey 5: The Adaptive Universe (System Difficulty Adjustment)
The system itself acts as a participant. As more players flock to a specific star system to mine "Iron," the **Resource Difficulty** increases, lowering the drop rate and pushing players to explore uncharted territories. Similarly, as more Legendary Mechs are generated in a season, the **Rarity Difficulty** adjusts, making the next Legendary harder to obtain. This "Bitcoin-style" adjustment ensures that the economy remains self-balancing, preventing hyperinflation and maintaining the prestige of high-tier assets without constant manual intervention.

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

### 9.2 Exploration & Discovery
- **FR5:** ผู้เล่นสามารถส่งยาน/หุ่นยนต์ออกไปสำรวจ (Explore) ในพิกัดต่างๆ ได้
- **FR6:** ระบบสามารถสุ่มเหตุการณ์ (Events) ระหว่างการสำรวจ เช่น การเจอทรัพยากร หรือการเจอศัตรู
- **FR7:** ระบบสามารถคำนวณระยะเวลาการสำรวจตามระยะทางและบัฟของผู้เล่น
- **FR8:** ผู้เล่นที่มี Season Pass+ สามารถรับบัฟความเร็ว/ประสิทธิภาพการสำรวจเพิ่มขึ้น 2 เท่า (ภายใต้ Daily Cap)
- **FR9:** ระบบสามารถสุ่มปล่อยพิกัดดาวดวงใหม่ (Star Discovery) โดยอัตโนมัติตามเงื่อนไขที่กำหนด (เช่น เวลา หรือความคืบหน้าของชุมชน)

### 9.3 Combat System (MVP - Stat-based)
- **FR10:** ระบบสามารถเริ่มการต่อสู้ (Combat) เมื่อผู้เล่นพบศัตรูระหว่างการสำรวจ
- **FR11:** ระบบสามารถคำนวณผลการแพ้/ชนะ โดยใช้ค่าพลัง (Stats) ของยาน/หุ่นยนต์และไอเทมที่สวมใส่
- **FR12:** ระบบสามารถมอบทรัพยากรที่แตกต่างกัน (Resource Differentiation) ตามผลการต่อสู้ (เช่น ชนะได้ Rare Parts, แพ้ได้ Scrap)

### 9.4 Engineering & Assembly (AI Generation)
- **FR13:** ผู้เล่นสามารถเริ่มกระบวนการประกอบร่าง (Assembly) โดยใช้ทรัพยากรที่สะสมมา
- **FR14:** ระบบสามารถสร้างภาพ Mechs/Tanks/Ships ด้วย AI (FLUX.1) ตามเกรดที่เลือก (Standard/Premium)
- **FR15:** ระบบสามารถจัดการคิวการประกอบร่างแบบรอเวลา (Time-gated)
- **FR16:** ผู้เล่นสามารถแชร์สถานะการประกอบร่าง (Assembly Status) ไปยัง Discord ได้

### 9.5 Economy & Marketplace
- **FR17:** ผู้เล่นสามารถซื้อ Season Pass และ Season Pass+ ด้วย USDT
- **FR18:** ผู้เล่นสามารถลงขาย "Key" (NFT) ใน Marketplace แบบ P2P ได้
- **FR19:** ระบบสามารถหักค่าธรรมเนียม (Marketplace Fee) จากการซื้อขายระหว่างผู้เล่น
- **FR20:** ระบบสามารถแสดงข้อมูลความยาก (Difficulty) และอัตราการดรอป (Rarity) บน Transparency Dashboard

### 9.6 Admin & Quality Control (HITL)
- **FR21:** Admin สามารถตรวจสอบและอนุมัติภาพที่ AI สร้างขึ้น (HITL) ก่อนที่จะทำการ Mint เป็น NFT
- **FR22:** ระบบสามารถปรับระดับความยาก (Difficulty Adjustment) ของทรัพยากรและอัตราการดรอปได้โดยอัตโนมัติ
- **FR23:** ระบบสามารถส่งการแจ้งเตือน Star Discovery ไปยัง Discord Channel ที่กำหนดได้ทันทีเมื่อมีการค้นพบใหม่
- **FR24:** ระบบ Multi-Stage Exploration (Space -> Orbital -> Surface) ที่บังคับใช้ประเภท Vehicle ต่างกัน (Ship/Mech)
- **FR25:** ระบบ Combat Log ที่บันทึกผลการต่อสู้และภาพ Action Shot แบบชั่วคราว (Housekeeping 7 วัน)
- **FR26:** ผู้เล่นสามารถจ่าย USDT เพื่อบันทึก Combat Log แบบถาวร (Permanent Save) ลงใน Profile/Lore

## 10. Non-Functional Requirements

### 10.1 Performance
- **NFR1 (Response Time):** หน้า Dashboard และการตอบสนองทั่วไปต้องโหลดเสร็จภายใน < 2 วินาที
- **NFR2 (AI Generation & Delivery):**
    - **Technical Latency:** ระบบต้องประมวลผลภาพ AI ให้เสร็จสิ้นก่อนเวลาที่กำหนดส่งมอบ (Reveal Time) ในเนื้อเรื่อง
    - **Load-Adaptive Assembly:** ระบบสามารถปรับเปลี่ยนระยะเวลาการรอในเนื้อเรื่อง (Assembly Time) ได้โดยอัตโนมัติตามปริมาณ Load ของระบบ เพื่อรักษาความเสถียรและคุมต้นทุน GPU โดยไม่กระทบต่อประสบการณ์ผู้เล่น (Immersive Load Balancing)
- **NFR3 (Blockchain Sync):** ข้อมูล On-chain ต้องสะท้อนบนหน้าเว็บภายใน < 5 วินาทีหลังจาก Transaction ได้รับการยืนยัน

### 10.2 Security & Anti-Bot
- **NFR4 (Anti-Bot Strategy):** 
    - ใช้ Captcha ในจุดที่มีการทำธุรกรรมสำคัญ
    - ใช้ Behavioral Analysis และ Client-side PoW สำหรับการกระทำทั่วไป (Explore/Combat) เพื่อป้องกัน Bot สายฟรี
    - บังคับ Social Verification (Discord/X) เพื่อป้องกันการปั๊มบัญชี
- **NFR5 (Smart Contract Security):** Smart Contract ต้องผ่านการ Audit และมีระบบ Emergency Stop (Circuit Breaker)
- **NFR6 (Data Protection):** ข้อมูลส่วนบุคคลและ Wallet Address ต้องถูกเข้ารหัสและจัดเก็บตามมาตรฐานความปลอดภัยสูงสุด

### 10.3 Scalability & AI Load Handling
- **NFR7 (Concurrency):** ระบบต้องรองรับผู้เล่นพร้อมกัน (Concurrent Users) ได้อย่างน้อย 10,000 คน
- **NFR8 (AI Reliability/Backtrack):** 
    - ระบบต้องมีกลไก Idempotency เพื่อป้องกันการจ่ายเงินซ้ำซ้อน
    - หากกระบวนการ AI Generation ล้มเหลว ระบบต้องมีกลไก Auto-Retry หรือคืนสิทธิ์ (Credit) ให้ผู้เล่นโดยอัตโนมัติ
    - มีระบบ Audit Log ที่ละเอียดเพื่อตรวจสอบย้อนหลังได้ทุกธุรกรรม

### 10.4 Reliability (ความเสถียร)
- **NFR9 (Uptime):** ระบบต้องมี Uptime ไม่ต่ำกว่า 99.9% (High Availability)
- **NFR10 (Data Consistency):** ใช้ Saga Pattern เพื่อรักษาความถูกต้องของข้อมูลระหว่าง Web2 (Database) และ Web3 (Blockchain) 100%
- **NFR11 (AI Guardrails):** ใช้ Structured Output (JSON Schema) และ RAG เพื่อป้องกัน AI Hallucination และรักษาความถูกต้องของ Lore
- **NFR12 (Housekeeping):** ระบบลบข้อมูลชั่วคราว (Temporary Logs) อัตโนมัติเพื่อรักษาประสิทธิภาพของฐานข้อมูล
