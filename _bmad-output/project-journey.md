# Project-0: The Evolution Journey (Tech & Design)

สรุปการเดินทางของ Project-0 จากจุดเริ่มต้นสู่ระบบ Decoupled Systems และ Deep Durability

## 1. The Technical Pivot: From Standard Web to "Game Engine" Architecture
เราได้ทำการ Refactor โครงสร้าง Frontend ใหม่ทั้งหมดเพื่อให้รองรับความซับซ้อนของเกมในระยะยาว

- **Old Way:** React Components จัดการ State กันเองผ่าน Props/Context ทำให้ Code ผูกติดกัน (Tightly Coupled) และขยายยาก
- **New Way (Unity-like Systems):** 
    - **EventBus:** ใช้ระบบ Global Event ในการสื่อสารระหว่างส่วนต่างๆ (เช่น Combat จบ -> ส่ง Event -> Hangar อัปเดต)
    - **Singleton Systems:** แยก Logic ออกจาก UI มาไว้ใน `CombatSystem`, `HangarSystem`, `GachaSystem`
    - **XState (Game Machine):** ใช้ State Machine ควบคุม Flow ใหญ่ของเกม (Loading -> Menu -> Hangar -> Combat)
    - **Result:** โค้ดสะอาดขึ้นมาก UI มีหน้าที่แค่ "แสดงผล" และ "ส่งคำสั่ง" ส่วน Logic อยู่ใน System ทั้งหมด

## 2. The Bastion: Beyond a Simple Menu
เปลี่ยนจาก "Mothership" ที่เป็นแค่เมนูเลือกด่าน กลายเป็น **"The Bastion"** (ฐานทัพเคลื่อนที่)

- **Bridge View:** หน้าจอหลักที่จะเป็นศูนย์กลางความดื่มด่ำ (Immersion) เห็นวิวอวกาศและสถานะของยานแบบ Real-time
- **Modular Hardpoints:** ยานมี Slot สำหรับติดตั้งอุปกรณ์ (Warp Drive, Shields, Turrets) ซึ่งส่งผลต่อการเล่นจริง
- **The Maintenance Hub:** เป็นที่ที่ผู้เล่นต้องกลับมาซ่อมแซมและปรับแต่งอุปกรณ์

## 3. The "Stage Change" NFT Model
เราแก้โจทย์เรื่อง "NFT ในเกมควรเป็นอย่างไร" ด้วยโมเดลที่สมดุลที่สุด

- **Virtual First:** ไอเทมทุกชิ้นเริ่มจากการเป็นข้อมูลใน Database (Virtual) เพื่อความรวดเร็วในการเล่น
- **Stage Change (Minting):** เมื่อผู้เล่นเจอไอเทมที่ถูกใจหรือมีค่า สามารถ "Mint" เพื่อเปลี่ยน Stage เป็น NFT ได้
- **Digital Twin:** NFT ไม่ใช่แค่รูปภาพ แต่เป็น "ฝาแฝด" ของไอเทมในเกม มันยังคงมีค่า Durability, Stats และต้องซ่อมแซมเหมือนเดิม แต่ได้สิทธิ์ในการเทรดและโชว์ความเป็นเจ้าของบน Chain

## 4. Deep Durability System (DDS)
ระบบความทนทานที่ลึกกว่าแค่ "พังแล้วใช้ไม่ได้"

- **Thresholds:** แบ่งสถานะไอเทมเป็น 5 ระดับ (Pristine, Worn, Damaged, Critical, Broken)
- **Impact:** 
    - **Visual:** ยิ่งพัง ยิ่งมี Glitch หรือควันออกในหน้า UI
    - **Gameplay:** ประสิทธิภาพลดลง มีโอกาสเกิด Malfunction (เช่น ปืนขัดลำกล้อง)
- **Economy:** สร้าง Loop การใช้ทรัพยากร (Scrap/Energy) เพื่อซ่อมแซม ทำให้ระบบเศรษฐกิจในเกมหมุนเวียน

## 5. Strategic Encounter & AI Visual Consistency
เราได้ออกแบบระบบการสำรวจที่ลึกขึ้น โดยเชื่อมโยงข้อมูลจาก Database เข้ากับ AI Generation

- **Strategic Choices:** การสำรวจใน Void ไม่ใช่แค่การสุ่มเจอเหตุการณ์ แต่เป็นการตัดสินใจ (Stealth, Assault, Bypass) ที่อิงจาก Stat ของ Pilot, Vehicle และ Equipment จริงๆ
- **Dynamic Visual AI:** ใช้เทคโนโลยี IP-Adapter และ ControlNet เพื่อให้ AI สร้างภาพเหตุการณ์ที่ "หน้าตาเหมือนหุ่นของผู้เล่น" และ "สอดคล้องกับสภาพความเสียหาย (DDS)" ในขณะนั้น
- **Result:** ทุกการสำรวจจะมีภาพประกอบที่เป็นเอกลักษณ์เฉพาะตัวของผู้เล่นคนนั้นจริงๆ

## 6. Current Tech Stack
- **Frontend:** Next.js 15, Framer Motion (Animations), Three.js (3D Elements), XState, TailwindCSS, **Decoupled Systems (EventBus + Singletons)**
- **Backend:** Go (Clean Architecture), PostgreSQL, JWT Auth, Privy (Web3 Auth), **DDS Items System**
- **AI Integration:** FLUX.1 (Image Gen), IP-Adapter (Consistency), ControlNet (Structure)

---
*สถานะปัจจุบัน: เสร็จสิ้นการออกแบบโครงสร้างระบบ Items (DDS) และ Gameplay Technical Spec พร้อมเริ่ม Implement ระบบสำรวจ (Exploration Engine)*
