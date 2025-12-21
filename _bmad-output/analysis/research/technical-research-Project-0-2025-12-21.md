# Technical Research: Project-0

**Date:** 2025-12-21
**Project:** Project-0 (Crypto RPG)
**Author:** GitHub Copilot (Gemini 3 Flash)

---

## 1. AI Image Generation Integration

### Comparison of APIs for Seasonal Asset Generation

| Feature | Stable Diffusion (FLUX.1) | Midjourney (v6/v7) | DALL-E 3 (OpenAI) |
| :--- | :--- | :--- | :--- |
| **Speed** | **High** (Sub-5s with optimized infra) | **Moderate** (15-60s) | **Moderate** (10-20s) |
| **Cost** | **Low** ($0.001 - $0.01 via Fal.ai/Replicate) | **Moderate** ($10-$120/mo sub) | **High** ($0.04 - $0.08 per image) |
| **Style Consistency** | **Excellent** (via LoRA/ControlNet) | **High** (Artistic but less "lockable") | **Good** (High prompt adherence) |
| **API Maturity** | High (Open ecosystem) | Low (Unofficial/Discord-centric) | High (Robust REST API) |
| **Best Use Case** | **Seasonal Items/Characters** (Fine-tuned) | Marketing/Concept Art | Rapid Prototyping |

**Findings:**
- **Stable Diffusion (FLUX.1)** is the recommended choice for "Project-0" seasonal assets. Its ability to be fine-tuned via LoRA (Low-Rank Adaptation) ensures that all Mechs or Tanks in a specific season maintain a consistent "visual DNA" while remaining unique.
- **Midjourney** produces the highest "out-of-the-box" artistic quality but lacks a robust, high-volume automation API suitable for a real-time game engine.
- **DALL-E 3** is excellent for prompt adherence but becomes cost-prohibitive at scale and offers less control over specific aesthetic parameters required for game assets.

**Sources:**
- [Black Forest Labs (FLUX.1)](https://blackforestlabs.ai/)
- [Fal.ai Stable Diffusion API](https://fal.ai/models/stable-diffusion-xl)
- [OpenAI DALL-E 3 Pricing](https://openai.com/pricing)

---

## 2. Blockchain & NFT Infrastructure

### L2 Comparison for "Direct-to-System" Economy

| Feature | Polygon (PoS) | Base (Coinbase L2) | Arbitrum (Nova) |
| :--- | :--- | :--- | :--- |
| **Avg. Tx Fee** | $0.01 - $0.05 | **<$0.01** (Post-Dencun) | **<$0.001** |
| **Throughput** | 65,000 TPS (Theoretical) | 2,000+ TPS | **40,000+ TPS** |
| **Ecosystem** | Massive, Gaming-heavy | Rapidly growing, Fiat-linked | DeFi & Gaming focused |
| **On-ramping** | Standard Bridges | **Native Coinbase Integration** | Standard Bridges |

**Findings:**
- **Base** is the strongest candidate for the "Direct-to-System" economy. Its native integration with Coinbase allows for seamless fiat-to-crypto onboarding, which is critical for the "Accessible Entry" requirement of Project-0.
- **Arbitrum Nova** is the technical winner for high-frequency game actions (e.g., resource gathering, small trades) due to its AnyTrust technology, which keeps fees significantly lower than standard L2s.
- **Polygon** remains a safe, highly compatible choice but lacks the specific "low-friction" onboarding advantages of Base or the extreme cost-efficiency of Arbitrum Nova.

**Sources:**
- [L2Fees.info](https://l2fees.info/)
- [Base Documentation](https://docs.base.org/)
- [Arbitrum Nova for Gaming](https://arbitrum.io/nova)

---

## 3. Orchestrator Architecture & Monitoring

### System Design: Microservices vs. Serverless

- **Recommendation: Hybrid Orchestrator Architecture**
    - **Core Game Engine (Microservices):** Use Node.js or Go microservices hosted on **Kubernetes (EKS/GKE)** for stateful operations like combat calculations and star discovery.
    - **Async Tasks (Serverless):** Use **AWS Lambda** or **Google Cloud Functions** for stateless, bursty tasks such as triggering AI image generation, minting NFTs on Base, and sending Discord alerts.

### Monitoring & Debugging Tools
- **Infrastructure:** **Prometheus** for metric collection and **Grafana** for real-time dashboards.
- **Application Performance (APM):** **Datadog** for deep tracing across microservices and serverless functions.
- **Web3 Specific:** **Tenderly** for real-time smart contract monitoring, transaction simulations, and debugging failed "Key" mints.
- **AI Monitoring:** **Arize Phoenix** or **LangSmith** to monitor AI generation latency and prompt performance.

**Sources:**
- [AWS Game Backend Reference Architecture](https://aws.amazon.com/gaming/solutions/game-backend/)
- [Tenderly Web3 Monitoring](https://tenderly.co/)
- [Datadog for Microservices](https://www.datadoghq.com/solutions/microservices/)

---

## 4. Discord Integration

### Real-time Alerts & Marketplace Linking

- **Real-time Alerts (Raid Boss/New Planet):**
    - Use **Discord Webhooks** for simple notifications.
    - Implement a **Discord Bot (Discord.js)** for interactive alerts where players can "Join Raid" directly from a button, which then communicates with the Orchestrator.
- **Web Marketplace Linking:**
    - **Discord OAuth2:** Link Discord accounts to player wallets to verify ownership of "Keys" before allowing access to exclusive Discord channels.
    - **Deep Linking:** Use unique transaction IDs in Discord buttons that redirect to the Web Marketplace with the item pre-selected for purchase or trade.

**Sources:**
- [Discord Developer Portal: OAuth2](https://discord.com/developers/docs/topics/oauth2)
- [Discord.js Documentation](https://discord.js.org/)

---

## 5. Technical Stack Summary

### Programming Languages
- **Backend:** TypeScript (Node.js), Go (for high-performance combat engine).
- **Smart Contracts:** Solidity.
- **Frontend:** React / Next.js.

### Frameworks
- **Web:** Next.js 15 (App Router).
- **Blockchain:** Hardhat / Foundry.
- **Discord:** Discord.js.

### Databases
- **Primary:** PostgreSQL (via Supabase or RDS).
- **Caching/Real-time:** Redis (for star discovery state and raid boss health).
- **Metadata:** MongoDB (for flexible AI-generated asset attributes).

### Tools
- **AI:** FLUX.1 (via Fal.ai), OpenAI API.
- **Monitoring:** Prometheus, Grafana, Datadog, Tenderly.
- **CI/CD:** GitHub Actions.

### Cloud/Infrastructure
- **Provider:** AWS (EKS, Lambda, S3, CloudFront).
- **Blockchain:** Base (L2), Alchemy/QuickNode (RPC Nodes).
- **CDN:** Vercel (Frontend), CloudFront (AI Assets).
