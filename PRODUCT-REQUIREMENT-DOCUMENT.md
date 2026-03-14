# Product Requirements Document (PRD): RAWVE

## 1. Meta Information

| Field | Detail |
|---|---|
| **Product Name** | RAWVE Live Streaming Platform |
| **Slogan** | "Lower Cut. Full Control. Open Source." |
| **Core Philosophy** | Full transparency — on costs, rules, and code. Creators know exactly what they get and what the platform takes. |
| **Platform** | Web Application (Backend-Heavy, API-First) |
| **License** | MIT — self-hosting is free, forever |

---

## 2. Problem Statement

Content creators today are trapped in platform ecosystems that:

- **Take too much** — Twitch 50%, TikTok 70%+, YouTube 45%. This isn't infrastructure cost, it's extraction.
- **Hide the numbers** — creators don't know the real cost being taken or why the percentage is that high.
- **Lock in communities** — viewer data, subscriber lists, and community history are owned by the platform, not the creator.
- **Stay closed** — there's no way to audit, fork, or migrate without losing everything that's been built.

RAWVE is built on one principle: **creators who build communities deserve to know and control everything** — including how much the platform takes and why.

---

## 3. Product Goals & Vision

### Primary Goals

| Goal | Description |
|---|---|
| **Lower Cut** | 15% platform fee — far below competitors. This number is transparent, justified, and will decrease as we scale efficiently. |
| **Creator Control** | Creators manage their own moderation, monetization, and community data — not the platform. |
| **Open Source** | All code is public under the MIT License. Self-hosting = 0% platform fee, forever. |
| **Extreme Performance** | Real-time chat with very low latency, resilient against sudden traffic spikes. |

### Long-Term Vision

RAWVE aims to become the streaming infrastructure that independent creators and organizations can genuinely trust — not because they have no other option, but because RAWVE is more honest and more capable than the alternatives.

---

## 4. Monetization Model

### Principle

> RAWVE does not hide its costs. The platform fee is taken transparently from every monetization transaction — not from creator data or hidden advertising.

### Fee Structure

```
Every monetization transaction (donations, tips, paid subscriptions):

  Creator receives   85%
  RAWVE platform fee 15%

Example: a viewer sends a $10 tip
  Creator gets  $8.50
  RAWVE gets    $1.50
```

### Transparency Commitments

The platform commits to:

- Publishing infrastructure cost estimates openly
- Never raising the platform fee without a minimum 90-day public notice
- Lowering the fee as operational efficiency improves at scale
- Never selling creator or viewer data to third parties
- Never adding hidden fees on top of the stated 15%

### Self-Hosting = 0% Fee

Because RAWVE is MIT-licensed, anyone can self-host with a 0% platform fee. The 15% only applies to creators using RAWVE's hosted infrastructure.

---

## 5. User Personas

### Persona 1 — The Creator (Streamer)

**Who:** An independent creator tired of high cuts and one-sided rules from major platforms.

**Needs:**
- A simple dashboard to manage stream titles and metadata
- Real-time live chat monitoring
- Clear revenue analytics — how much came in, how much was deducted, and why
- Full control over community moderation rules

**Core pain point:** "I work hard to build my community, but the platform owns the data and takes half my earnings."

---

### Persona 2 — The Viewer

**Who:** A viewer who wants to support their favorite creator directly.

**Needs:**
- Instant access to live chat without lag
- Fast login (Google / Email)
- Confidence that their donation actually reaches the creator in full

**Core pain point:** "I send a $10 donation — how much does the creator actually receive?"

---

### Persona 3 — The Contributor (Developer)

**Who:** A developer who believes in open source and wants to contribute to fairer streaming infrastructure.

**Needs:**
- Clean, documented, easy-to-understand code
- An architecture that isn't confusing to extend
- A community that's open to pull requests and technical discussion

---

## 6. Core Features (MVP)

### A. Identity & Access Management (via Clerk)

- Registration and authentication via Clerk webhooks
- Session validation via JWT Bearer Token on all protected endpoints
- User sync: `CreateOrUpdate` in PostgreSQL on every successful authentication

### B. Stream Control

- **Start / End Stream:** API to toggle stream status (`is_live: boolean`)
- **Stream Metadata:** Configure stream title and description
- **Stream Key Management:** Creators can regenerate their stream key at any time

### C. Real-Time Engagement (Live Chat)

- **WebSocket Connection:** Full-duplex communication via Gorilla WebSocket
- **Chat History:** Last 50 messages loaded automatically when a viewer joins
- **Broadcasting:** Messages distributed to all viewers in the same room simultaneously using goroutine + channel pattern

### D. Monetization Engine (Post-MVP, planned)

- Tip / donation flow with transparent 15% platform fee
- Automatic receipts for both creators and viewers
- Revenue dashboard with clear fee breakdown

---

## 7. Technical Architecture & Stack

### Stack

| Layer | Technology |
|---|---|
| Language | Golang (Go 1.25+) |
| HTTP Framework | Gin Web Framework |
| Real-time | Gorilla WebSocket |
| Database | PostgreSQL via GORM |
| Auth | Clerk SDK (JWT Verification) |
| Infrastructure | Docker & Docker Compose |
| CDN | Cloudflare |

### Architecture Pattern: Clean Architecture

```
rawve/
├── cmd/api/main.go          # Entry point & Dependency Injection
├── internal/
│   ├── domain/              # Entities & interface contracts
│   ├── usecase/             # Core business logic
│   ├── repository/          # Database & external service calls
│   └── delivery/
│       ├── http/            # HTTP handlers & routing
│       ├── websocket/       # WebSocket handlers
│       └── middleware/      # Auth, logging, rate limiting
```

### Data Flow

```
Request → Delivery → Usecase → Repository → Database
                  ↕
              Domain (Entities)     Clerk (Auth)
```

### Frontend Compatibility

Designed API-first — easily consumed by modern frameworks (Next.js, React, Vue) and mobile clients.

---

## 8. Non-Functional Requirements

| Requirement | Target |
|---|---|
| Chat latency | < 100ms p99 |
| WebSocket connections | 10,000+ concurrent per instance |
| API response time | < 200ms p95 |
| Uptime | 99.9% monthly |
| Platform fee transparency | Displayed on every transaction receipt |

---

## 9. Future Roadmap (Post-MVP)

| Priority | Feature | Description |
|---|---|---|
| High | **Redis Pub/Sub** | Migrate from in-memory WebSocket hub to Redis for horizontal scaling across multiple Go server instances |
| High | **Creator Analytics** | Real-time revenue dashboard — viewers, incoming donations, fee breakdown |
| Medium | **Content Moderation Tools** | Creator-controlled moderation rules, keyword filters, timeout/ban system |
| Medium | **Self-host Documentation** | Complete guide to deploying RAWVE independently with 0% platform fee |
| Low | **Web3 / Crypto Monetization** | Smart contract integration (Solidity) for direct tipping to creator wallets (Ethereum/Polygon) — optional, does not replace the main system |
| Low | **Enterprise License** | Private instances for organizations, media companies, or universities |

---

## 10. Success Metrics

| Metric | Target (6 months post-launch) |
|---|---|
| Active streamers | 1,000+ |
| Creator retention (M1 → M3) | > 60% |
| Net Promoter Score (NPS) | > 40 |
| Platform fee disputes | < 1% of total transactions |
| Open source contributors | > 50 contributors |

---

*This is a living document — updated as the product evolves and community feedback comes in.*