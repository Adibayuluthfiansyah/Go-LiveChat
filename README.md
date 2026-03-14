<div align="center">

<img src="https://readme-typing-svg.demolab.com?font=Bebas+Neue&size=80&pause=1000&color=FF2D4E&center=true&vCenter=true&width=600&height=120&lines=RAWVE" alt="RAWVE" />

### **Lower Cut. Full Control. Open Source.**
*The streaming platform that's actually on your side.*

<br/>

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-FF2D4E?style=for-the-badge)](LICENSE)
[![Architecture](https://img.shields.io/badge/Architecture-Clean-FF6225?style=for-the-badge&logo=buffer&logoColor=white)](#architecture)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-00C853?style=for-the-badge&logo=github&logoColor=white)](CONTRIBUTING.md)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://postgresql.org)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://docker.com)

<br/>

> RAWVE is an **open-source live streaming backend** built for creators who want
> real ownership — of their community, their rules, and their revenue.
> We take a smaller cut so we can keep the lights on. Nothing more, nothing hidden.

<br/>

[**Get Started**](#-getting-started) · [**Why RAWVE**](#-why-rawve) · [**Architecture**](#-architecture) · [**Monetization**](#-how-rawve-makes-money) · [**Roadmap**](#-roadmap) · [**Contributing**](#-contributing)

</div>

---

## ✨ Why RAWVE?

Most streaming platforms treat creators as a product — taking half their revenue, owning their audience data, and making all the rules.

**RAWVE is different in three ways that actually matter.**

### 1. Lower cut than every major platform

We take **15%** — not because we're greedy, but because running global infrastructure costs real money. We'd rather be honest about that than pretend it's free while quietly extracting more.

| Platform | Revenue Cut | What's left for you |
|---|---|---|
| TikTok Live | ~70%+ | ~30% |
| Twitch | ~50% | ~50% |
| YouTube Live | ~45% | ~55% |
| **RAWVE** | **15%** | **85%** |

### 2. Full creator control

On RAWVE, you own your community. That means:

- **Your moderation rules** — set your own content policies, not ours
- **Your audience data** — we don't hold your subscribers hostage
- **Your monetization model** — charge what you want, how you want
- **Your stream key** — take it anywhere, point it at any ingest server

### 3. Fully open source (MIT)

Every line of code that powers RAWVE is public. You can audit it, fork it, self-host it, or contribute to it. No black boxes, no vendor lock-in, no surprises.

---

## 💰 How RAWVE Makes Money

We believe transparency builds trust. Here's exactly how we stay sustainable.

**RAWVE takes 15% of creator monetization transactions** — donations, tips, paid subscriptions between creators and their fans. That's it.

```
Example — a viewer sends a $10 tip:
  Creator receives   $8.50  (85%)
  RAWVE platform fee $1.50  (15%)
```

**Why 15% and not lower?**

Running a global live streaming infrastructure is expensive — CDN bandwidth, media servers, databases, engineering. At our projected scale of 1,000–10,000 active streamers, 15% is the floor that keeps us operational without outside funding.

We publish our infrastructure cost estimates publicly. When we become more efficient at scale, we will lower the cut. That's a commitment, not a marketing line.

**What we will never do:**
- Take a hidden cut on top of the stated 15%
- Sell creator or viewer data to advertisers
- Lock you into our platform or hold your audience data hostage
- Raise the cut without a 90-day public notice period

---

## 🚀 Roadmap

| Status | Feature | Description |
|---|---|---|
| ✅ | **Secure Authentication** | Webhook integration & JWT verification via Clerk |
| ✅ | **Clean Architecture** | Strict Domain / Usecase / Repository / Delivery separation |
| 🔧 | **Real-time Live Chat** | Lightning-fast messaging with Gorilla WebSocket |
| 🔧 | **Stream Management** | Creator live stage creation and control |
| 📋 | **Monetization Engine** | Tip/donation flow with transparent 15% platform fee |
| 📋 | **Creator Analytics** | Real-time viewers, revenue, and engagement dashboard |
| 📋 | **Content Moderation Tools** | Creator-controlled moderation with optional automation |
| 📋 | **Self-host Mode** | Full self-hosting with zero platform fee (open source) |
| 📋 | **Enterprise License** | Private instances for organizations and media companies |

> `✅ Done` · `🔧 In Progress` · `📋 Planned`

> **Note on self-hosting:** Because RAWVE is MIT-licensed, you can always self-host with 0% cut. The 15% only applies to creators using RAWVE's hosted infrastructure. We think that's fair.

---

## 🛠️ Tech Stack

```
Language        →  Golang (Go 1.25+)
HTTP Framework  →  Gin Web Framework
WebSocket       →  Gorilla WebSocket
ORM & Database  →  GORM + PostgreSQL
Authentication  →  Clerk SDK (JWT Verification)
Infrastructure  →  Docker & Docker Compose
CDN             →  Cloudflare / Bunny.net (recommended)
```

---

## 📁 Architecture

RAWVE is built with **Clean Architecture** — strict separation of concerns, zero dependency leakage between layers.

```
rawve/
├── cmd/
│   └── api/
│       └── main.go              # Entry point & Dependency Injection
│
├── internal/
│   ├── domain/                  # ♥  Heart of the app
│   │   ├── user.go              #    Entities & Interface contracts
│   │   └── stream.go
│   │
│   ├── usecase/                 # 🧠 Brain of the app
│   │   ├── user_usecase.go      #    Core business logic
│   │   └── stream_usecase.go
│   │
│   ├── repository/              # 💪 Muscle of the app
│   │   ├── user_repository.go   #    Database & external service calls
│   │   └── stream_repository.go
│   │
│   └── delivery/                # 🚪 Receptionist of the app
│       ├── http/                #    HTTP handlers & routing
│       ├── websocket/           #    WebSocket handlers
│       └── middleware/          #    Auth, logging, rate limiting
│
├── docker-compose.yml
├── .env.example
└── README.md
```

### Data Flow

```
HTTP Request
     │
     ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Delivery  │────▶│   Usecase   │────▶│ Repository  │────▶│  Database   │
│  (Gin/WS)   │     │ (Business)  │     │  (Data)     │     │ (Postgres)  │
└─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
     ▲                    │                    │
     │                    ▼                    ▼
     │             ┌─────────────┐     ┌─────────────┐
     └─────────────│   Domain    │     │    Clerk    │
      Response     │  (Entities) │     │   (Auth)    │
                   └─────────────┘     └─────────────┘
```

---

## ⚡ Getting Started

### Prerequisites

- [Go 1.25+](https://golang.org/doc/install)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- A [Clerk](https://clerk.com) account (free tier works)

### 1. Clone the Repository

```bash
git clone https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform.git
cd RAWVE-LiveStream-Platform
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Open `.env` and fill in your values:

```env
# Server
PORT=8080

# Database
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=rawve
DB_PORT=5435

# Clerk Authentication
CLERK_SECRET_KEY=sk_test_your_clerk_secret_key_here
```

> 🔑 Get your `CLERK_SECRET_KEY` from [Clerk Dashboard](https://dashboard.clerk.com) → API Keys

### 3. Start the Database

```bash
docker-compose up -d
```

### 4. Run the Application

```bash
go run cmd/api/main.go
```

Server will be running at `http://localhost:8080` 🚀

---

## 🔐 Authentication Flow

RAWVE uses **Clerk** for authentication with a webhook-based user sync pattern:

```
User signs in via Clerk
        │
        ▼
Clerk issues JWT token
        │
        ▼
Frontend sends request with Bearer token
        │
        ▼
RAWVE middleware verifies JWT (Clerk public key)
        │
        ▼
Extract claims → SyncUserFromAuth()
        │
        ▼
CreateOrUpdate user in PostgreSQL
        │
        ▼
Request continues with authenticated context
```

---

## 🤝 Contributing

We welcome contributions from everyone — bug fixes, new features, or documentation improvements.

```bash
# 1. Fork this repository

# 2. Create your feature branch
git checkout -b feature/your-awesome-feature

# 3. Commit your changes
git commit -m "feat: add your awesome feature"

# 4. Push to the branch
git push origin feature/your-awesome-feature

# 5. Open a Pull Request
```

### Commit Convention

We follow [Conventional Commits](https://www.conventionalcommits.org/):

| Prefix | Use for |
|---|---|
| `feat:` | New feature |
| `fix:` | Bug fix |
| `docs:` | Documentation only |
| `refactor:` | Code refactoring |
| `test:` | Adding tests |
| `chore:` | Build process or tooling |

---

## 👨‍💻 Creator

<div align="center">

<br/>

<img src="https://github.com/adibayuluthfiansyah.png" width="100" height="100" style="border-radius:50%" alt="Adibayu Luthfiansyah"/>

### **Adibayu Luthfiansyah**

*Founder & Lead Developer of RAWVE*

[![Website](https://img.shields.io/badge/Website-adibayuluthfiansyah.dev-FF2D4E?style=for-the-badge&logo=safari&logoColor=white)](https://adibayuluthfiansyah.dev)
[![GitHub](https://img.shields.io/badge/GitHub-adibayuluthfiansyah-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/adibayuluthfiansyah)

<br/>

> *"I built RAWVE because creators deserve a platform that's honest with them —
> about what it takes, what it costs, and what they get in return."*

<br/>

</div>

---

## 📄 License

RAWVE is open-source software licensed under the [MIT License](LICENSE).

Self-hosting is always free and always will be. The 15% platform fee only applies to creators using RAWVE's hosted infrastructure.

---

<div align="center">

Built with ❤️ for creators who deserve better.

**[rawve.live](https://rawve.live)** · [@rawve](https://github.com/rawve)

<br/>

*"Lower Cut. Full Control. Open Source."*

</div>