# Combat & Gameplay Design: Project-0

## 1. Combat Philosophy
Combat in Project-0 is designed to be **Strategic, High-Stakes, and Narrative-Driven**. The scale is **1:1 (One Pilot, One Vehicle)**, but the focus is on the **Asset Value** and the **Intense Story**. You are a high-stakes operative in a dangerous universe where your Vehicle is a powerful tool for dominance and collection. Every mission is a chapter in a dark, evolving plot.

## 2. Combat Stats & Attributes
Every unit (Bastion, Vehicle, Pilot) has core stats, but their roles differ:

| Unit Type | Primary Role | Key Stats |
| :--- | :--- | :--- |
| **Bastion** | Transport & Support | Fuel Capacity, Armor, **Scanner Range**. |
| **Vehicle** | Heavy Combat & Extraction | Firepower, Durability, **Radar Sensitivity**. |
| **Pilot (Character)** | Precision Salvage & EVA | Agility, O2 Capacity, Melee/Sidearm Skill. |

### 2.1 Combat Power (CP)
Combat Power is the primary metric for evaluating the strength of a Vehicle and its Pilot.
**Formula:**
$$CP = (Total ATK \times 3) + (Total DEF \times 2) + (Total HP / 5)$$

### 2.2 Elemental Types (The "Adaptive" Layer)
Vehicles and Pilot Gear are assigned a "Core Type" based on their AI generation:
- **Kinetic:** Strong against Energy, weak against Explosive.
- **Energy:** Strong against Explosive, weak against Kinetic.
- **Explosive:** Strong against Kinetic, weak against Energy.

## 3. Combat Flow: The "Immersive Cockpit" Model
To maintain the "Premium Economy" feel while providing a seamless 3D experience, combat follows an **Asynchronous Auto-Resolve with Real-time 3D Visualization**:

1. **Radar Scan & Risk Assessment:** Before launching, the player uses the Bastion's **Scanner** to probe the target location.
    - **Information:** The scanner provides a **Threat Level (Low/Med/High/Extreme)** and a **Success Probability (%)**.
    - **Fog of War (Web2 Backend):** The scanner cannot see everything. Hidden elite enemies or environmental hazards are managed by the **Web2 Backend**, which only sends visible data to the client to prevent map-hacking.
    - **The Choice:** Players must decide if their single Vehicle can handle the "Unknown" based on the fuzzy radar data.
2. **Bastion Transit:** Player travels to a star system. Bastion stats determine fuel cost and encounter risk (Space Pirates/Debris).
3. **Atmospheric Entry & Random Events:** If the target is a planet, the system rolls for "Accidents" based on the Risk Assessment.
    - **AI-Generated Accidents:** If a roll fails, the AI Service generates a unique event based on the context via **Model Context Protocol (MCP)**.
    - **Consequences:** Massive durability loss, Vehicle destruction, or Pilot death.
4. **Vehicle Deployment:** Once in orbit or on the surface, the Vehicle is launched. The camera zooms into the **3D Cockpit HUD** rendered with **WebGPU + React Three Fiber**.
5. **Pilot EVA (The Final Stage):** For specific high-value salvage (e.g., inside a derelict ship or narrow cave), the Pilot must exit the Vehicle.
    - **Risk:** High vulnerability.
    - **Reward:** Rare "Fixed Blueprints" or "Ancient Tech" only accessible by hand.
    - **Gear:** O2 levels and Pilot weapons (Swords/Guns) determine survival.
6. **Encounter & Radar:** The 3D Radar displays enemy signatures as they approach. You might find a lone scavenger or an elite squad that outclasses you.
7. **Simulation & Visualization:** The backend runs the simulation (1 vs. 1 or 1 vs. Small Group). The frontend visualizes the "Action Highlights" using **WebGPU** for high-fidelity effects:
    - **Camera Shakes:** Reflecting hits taken.
    - **HUD Glitches:** Indicating system damage.
    - **AI Narrative:** Real-time text logs appearing on the cockpit's side-monitors.
8. **Outcome:**
    - **Victory:** Loot is gathered, enemies are marked for Salvage. The current node on the timeline is marked as **Resolved**, enabling the "Proceed" button.
    - **Defeat:** Squad retreats, units take significant Durability damage. If the Pilot dies during EVA, it's a catastrophic loss (Loss of all loot from that run + Critical Vehicle damage).
9. **Return & Maintenance (The Bastion Phase):**
    - **Debrief:** AI summarizes the mission and visualizes the damage.
    - **Refine:** Convert salvaged scrap into usable materials.
    - **Repair:** Use materials to fix the Vehicle and Pilot gear at the **Visual Equipment Map**.
    - **Rest:** Pilot recovers fatigue, O2 is refilled.

## 4. Visual Wear & Tear & Modular NFT Assembly
Visuals are dynamically generated using AI (FLUX.1) to reflect a **Modular NFT Ecosystem**. Instead of a single static NFT, your loadout is a combination of individual NFT components.

### 4.1 Modular NFT Composition
Every major component is its own NFT with unique stats and visual traits:
- **Bastion Parts:** Hull, Engines, Heat Shields (Individual NFTs).
- **Vehicle Parts:** Chassis, Modules (HEAD, CORE, ARM_L, ARM_R, LEGS) (Individual NFTs).
- **Pilot Gear:** Suit, Helmet, Melee Weapon (Lightsaber), Ranged Weapon (Railgun) (Individual NFTs).

**AI Assembly Logic:** When a player equips these NFTs, the AI Service "reads" the metadata and visual signatures of each individual NFT and **assembles** them into a single, cohesive "Master Image" and 3D model. The final visual is a unique synthesis of all equipped parts.

### 4.2 Context-Aware Visuals (The "Event" Layer)
The AI doesn't just combine parts; it places them in the **Context** of your journey:
- **Exploration Snapshots:** If you are on a volcanic planet, the AI generates your Pilot (with equipped Nomad Cloak and Lightsaber NFT) standing amidst lava flows.
- **Battle Highlights:** If your Railgun NFT was used to deliver the finishing blow, the AI generates a high-action shot of that specific Railgun firing.
- **Visual Wear & Tear:** Damage is applied to the *specific* anatomical slot that took the hit (e.g., your ARM_L Module shows scorch marks, while your CORE remains pristine). **Repairing** the component will trigger an AI update to restore its visual state.

## 5. The Ultimate Goal: The Evolving Universe
The goal of Project-0 is not just to find "better gear," but to become a legendary figure in an **ever-expanding universe**.
- **Collection Pride:** Completing seasonal sets of AI-generated NFTs that will never be minted again.
- **Creator-Driven Lore:** The universe evolves through "Context Patches" released by the creator. New star systems, mysterious factions, and world-changing events are introduced periodically.
- **Hall of Fame:** Players who discover rare stars or defeat legendary bosses are immortalized in the game's lore and metadata.
- **Continuous Combat:** A never-ending cycle of exploration and battle, where each patch brings new challenges that require smarter strategies and better-maintained gear.

## 7. Progression & Growth: Measuring Strength

เพื่อให้ผู้เล่นรู้สึกถึงความ "เก่งขึ้น" และมีเป้าหมายในระยะยาว ระบบจะวัดความก้าวหน้าผ่าน 4 แกนหลัก:

### 7.1 Pilot Rank & Mastery (XP)
- **Experience Points (XP):** ได้รับจากการทำภารกิจสำเร็จ, การ Salvage ซากหุ่น, และการค้นพบดวงดาวใหม่ๆ
- **Pilot Rank:** การเลื่อนระดับจะปลดล็อก "Advanced Licenses" เพื่อให้สามารถขับ Vehicle หรือ Bastion ในระดับที่สูงขึ้นได้
- **Skill Points:** ทุกครั้งที่ Rank อัป จะได้รับแต้มเพื่ออัปเกรดความสามารถเฉพาะตัวของนักบิน (เช่น Agility, O2 Efficiency, Melee/Sidearm Mastery)

### 7.2 Modular Gear & "Options" (The Loot Loop)
- **AI-Generated Stats:** ชิ้นส่วน NFT แต่ละชิ้น (Module, Chassis, Weapon) จะมีค่าพลังสุ่ม (Options) ที่สร้างโดย AI (เช่น +5% Kinetic Resistance, +10% Scanner Range)
- **Rarity Tiers:** แบ่งระดับความหายากเป็น Standard, Rare, Epic, Legendary และ Seasonal
- **Visual DNA:** ไอเทมระดับสูงจะมี Visual Traits ที่ซับซ้อนและสวยงามกว่า ซึ่งแสดงถึงความเก่งกาจที่มองเห็นได้ชัดเจนในสังคม

### 7.3 Vehicle Sync Rate (Synergy)
- **Familiarity:** การใช้ Vehicle หรือชิ้นส่วนเดิมซ้ำๆ จะเพิ่มค่า "Sync Rate"
- **Bonuses:** ค่า Sync Rate ที่สูงจะมอบ Buff พิเศษ เช่น +Evasion, +Accuracy หรือลดการใช้ Energy ซึ่งเป็นการตอบแทนผู้เล่นที่ดูแลและซ่อมแซมอุปกรณ์คู่ใจแทนการเปลี่ยนใหม่ตลอดเวลา

### 7.4 Bastion Tech Tree (Base Progression)
- **Facility Upgrades:** ผู้เล่นต้องลงทุนทรัพยากรเพื่ออัปเกรด Preparation Area (ความเร็วในการซ่อม), Lab (คุณภาพการคราฟต์) และ Scanner (ระยะการค้นหา)
- **Deep Space Access:** การอัปเกรดเครื่องยนต์และเกราะกันความร้อนของ Bastion เป็นวิธีเดียวที่จะเข้าถึงระบบดาวระดับสูงที่มี Loot ระดับตำนาน

## 8. Strategic Synergy & Upgrades
- **Energy Management:** 
    - **Standard Energy:** ใช้สำหรับการสำรวจทั่วไป (รีเฟรชฟรีทุกวัน).
    - **Premium Energy (Overclock):** ใช้สำหรับการสำรวจต่อเนื่องเมื่อ Standard Energy หมด หรือใช้เพื่อเพิ่ม Success Probability (%) ในภารกิจเสี่ยงสูง.
- **Bastion Upgrades:** Essential for reaching Deep-Space and surviving Atmospheric Entry.
- **Mech Upgrades:** Focus on combat efficiency and specialized extraction tools.
- **Pilot Gear:** Swords, Guns, and O2 Tanks are critical for the final stage of exploration.
- **Colony Position:** Being closer to a "Combat Zone" reduces fuel costs but increases the risk of "Colony Raids" (Defensive Combat).
- **Story Mode:** Boss encounters provide "Fixed Blueprints" for high-tier Mothership/Mech parts and Pilot gear.

```