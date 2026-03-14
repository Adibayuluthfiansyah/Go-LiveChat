# Contributing to RAWVE

RAWVE is an early-stage open source project — maintained by one developer, built for creators who deserve better.
There are no corporate contributors yet. If you're reading this, you might be one of the first.

Contributions of any size are welcome and genuinely appreciated.

---

## Ways to Contribute

- Report a bug
- Suggest a feature
- Improve documentation
- Submit a pull request
- Review someone else's PR
- Ask a question that leads to better docs
- Share the project with someone who'd find it useful

No contribution is too small. A typo fix counts.

---

## Getting Started

**1. Fork the repository**

Go to https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform and click **Fork**.

**2. Clone your fork**

```bash
git clone https://github.com/YOUR_USERNAME/RAWVE-LiveStream-Platform.git
cd RAWVE-LiveStream-Platform
```

**3. Add the upstream remote**

```bash
git remote add upstream https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform.git
```

**4. Create a branch**

```bash
git checkout -b feature/your-feature-name
```

---

## Development Setup

Prerequisites: Go 1.25+, Docker, Docker Compose.

```bash
# Copy environment config
cp .env.example .env

# Start the database
docker-compose up -d

# Run the server
go run cmd/api/main.go
```

Server runs at `http://localhost:8080`.

---

## Code Guidelines

RAWVE uses Clean Architecture. The layer boundaries matter — please respect them.

| Directory | Responsibility |
|---|---|
| `domain/` | Entities and interface contracts only. No dependencies on other layers. |
| `usecase/` | Business logic. Depends on domain interfaces, not concrete implementations. |
| `repository/` | Database and external service calls. Implements domain interfaces. |
| `delivery/` | HTTP handlers, WebSocket layer, middleware. Calls usecases, never repositories directly. |

Beyond architecture:

- Write idiomatic Go — keep functions small, handle errors explicitly
- Avoid adding dependencies unless clearly necessary
- Document complex logic inline, not in the PR description
- Don't break existing behavior without a clear reason

---

## Commit Convention

RAWVE follows [Conventional Commits](https://www.conventionalcommits.org/).

```
feat: add stream analytics endpoint
fix: resolve websocket reconnect issue
docs: update installation guide
refactor: simplify authentication middleware
test: add unit tests for user usecase
chore: update dependencies
```

| Prefix | Use for |
|---|---|
| `feat:` | New feature |
| `fix:` | Bug fix |
| `docs:` | Documentation only |
| `refactor:` | Code refactoring, no behavior change |
| `test:` | Adding or updating tests |
| `chore:` | Build process, tooling, dependencies |

---

## Submitting a Pull Request

Before opening a PR:

- Make sure the project builds: `go build ./...`
- Keep the change focused — one PR, one concern
- Write a clear description of what changed and why
- Reference related issues with `Closes #issue-number` if applicable

```bash
git add .
git commit -m "feat: describe your change"
git push origin feature/your-feature-name
```

Then open a Pull Request on GitHub. There's no formal review SLA — this project is maintained in spare time, but every PR will be read.

---

## Reporting Bugs

Open an issue at:
https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform/issues

Include:
- What you expected to happen
- What actually happened
- Steps to reproduce
- Go version, OS, and any relevant error output

---

## Suggesting Features

Open an issue and describe:
- The problem you're trying to solve
- Your proposed solution
- Any alternatives you considered

Feature ideas that align with RAWVE's core values — lower cut, creator control, open source, transparency — are more likely to be prioritized.

---

## Community Guidelines

Be respectful. Be constructive. Assume good intent.

This project is about building something fairer for creators. That spirit should carry into how we treat each other as contributors too.

---

## Maintainer

**Adibayu Luthfiansyah**
GitHub: https://github.com/adibayuluthfiansyah
Website: https://adibayuluthfiansyah.dev

---

Thank you for helping make RAWVE better.