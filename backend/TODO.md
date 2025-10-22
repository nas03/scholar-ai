Backend roadmap and TODO

Scope: Golang backend under `backend/` (Gin + GORM). This plan organizes work into milestones with priorities and clear acceptance criteria.

Legend:
- Priority: P0 (urgent), P1 (important), P2 (nice-to-have)
- Status: [ ] todo, [~] in-progress, [x] done

Key entities (current + planned):
- User, Course, Semester, Schedule/Timetable, Reminder, Lecture Notes, Materials, Quiz

Milestone M0 — Foundation and conventions
- [ ] P0 Standardized project configuration
	- Acceptance: `internal/config/config.go` loads from env and .env; required vars validated with clear errors; sample `.env.example` exists.
- [ ] P0 Structured logging
	- Acceptance: Central logger (e.g., zap/logrus) wired in `global/global.go`; request logging middleware; log levels via env; JSON output in prod.
- [ ] P0 Error and response contract
	- Acceptance: Use `pkg/response` consistently; add error codes, request-id correlation, and uniform error payload; update controllers to use helpers.
- [ ] P0 CORS and security headers
	- Acceptance: `internal/middleware/cors.go` uses allowed origins from config; add secure headers (no sniff, frameguard where relevant).
- [ ] P1 Rate limiting middleware
	- Acceptance: Global/token-bucket per IP and per user on auth endpoints.

Milestone M1 — Auth and user management
- [ ] P0 User registration and login endpoints
	- Acceptance: POST /api/v1/auth/register, /auth/login, /auth/logout; passwords hashed (bcrypt/argon2id); JWT access + refresh tokens; refresh rotates.
- [ ] P0 Email verification
	- Acceptance: Email verification token generation + confirm endpoint; resend with cooldown; mock provider in dev; interface for real provider.
- [ ] P1 Password reset flow
	- Acceptance: request reset, token, reset endpoint; token invalidation; minimum password policy.
- [ ] P1 Phone number update and verification (optional)
	- Acceptance: Store E.164 numbers; OTP verify; pluggable SMS provider; rate limiting.
- [ ] P1 2FA (TOTP)
	- Acceptance: Enable/disable TOTP, QR provisioning, recovery codes, step-up on sensitive actions.
- [ ] P2 SSO (OIDC: Google/Microsoft)
	- Acceptance: OAuth2 login, account linking, new-user onboarding with provider claims.
- [ ] P0 Login attempt tracking and lockout
	- Acceptance: Track failed attempts; temporary lockout with exponential backoff; audit logs.

Milestone M2 — Core domain CRUD
- Courses
	- [ ] P0 CRUD: name, description, length/credits; routes in `internal/router/user.route.go` sibling course routes; service + repo methods; unit tests.
- Semesters
	- [ ] P1 CRUD: name, start/end dates; validation that courses map to semesters.
- Schedule/Timetable
	- [ ] P1 CRUD: blocks with day-of-week, start/end, location; conflict detection on create/update.

Milestone M3 — Productivity features
- Reminders
	- [ ] P1 CRUD; schedule engine (cron/worker) to dispatch reminders; pluggable channels (email, push placeholder).
- Lecture Notes
	- [ ] P1 CRUD; support rich text (store as markdown/json); versioning metadata; basic search by title/tags.
- Materials
	- [ ] P1 Upload & list; storage provider interface (local dev, S3-compatible in prod); signed URLs for download; size/type limits.
- Quiz
	- [ ] P2 CRUD for quizzes and questions; assignment to course/notes; simple scoring endpoint.

Milestone M4 — Data and persistence
- [ ] P0 Database migrations
	- Acceptance: Adopt `golang-migrate` or `goose`; convert `sql/init.sql` into versioned migrations; Makefile targets up/down.
- [ ] P0 GORM models finalized
	- Acceptance: Models in `internal/models/*.go` with constraints, indexes; AutoMigrate only used in dev.
- [ ] P1 Redis integration
	- Acceptance: Connection pool from config; health-check; used for rate limiting, blacklisted tokens, and caching hot reads.

Milestone M5 — Observability and reliability
- [ ] P0 Health, readiness, liveness
	- Acceptance: `/healthz` (process), `/readyz` (DB+Redis), `/livez`; hook into container probes.
- [ ] P1 Metrics (Prometheus)
	- Acceptance: Basic process, HTTP latency, DB latency, cache hit rate; `/metrics` endpoint.
- [ ] P2 Tracing (OpenTelemetry)
	- Acceptance: Traces for HTTP handlers and DB calls; exporter configurable.

Milestone M6 — Quality: tests and linting
- [ ] P0 Unit tests for services/repositories
	- Acceptance: >60% coverage for core packages; table-driven tests; sqlite-in-memory for GORM repos.
- [ ] P0 Integration tests for auth flow
	- Acceptance: Register → verify email → login → refresh → logout happy path passes in CI.
- [ ] P1 Static analysis
	- Acceptance: `golangci-lint` configured; CI job fails on lint errors; `go fmt` enforced.

Milestone M7 — Docs and DX
- [ ] P1 OpenAPI/Swagger docs
	- Acceptance: `swaggo/swag` integrated; `/swagger/index.html` available in dev; CI step to validate spec builds.
- [ ] P1 Developer onboarding
	- Acceptance: README updated with run, test, migrate, and env setup; Makefile targets for common actions.

Milestone M8 — Packaging and deployment
- [ ] P1 Dockerization
	- Acceptance: Multi-stage Dockerfile; minimal image; non-root user; healthcheck.
- [ ] P1 Docker Compose for dev
	- Acceptance: Compose with app + DB + Redis; hot-reload in dev.
- [ ] P2 CI/CD pipeline
	- Acceptance: GitHub Actions: build, test, lint, docker build; artifact or container publish.

Backlog and ideas
- [ ] P2 Full-text search for notes/materials.
- [ ] P2 Webhook callbacks for reminders.
- [ ] P2 Background worker separation (e.g., separate process/queue).

References (current code map)
- Config: `internal/config/config.go`
- Models: `internal/models/*.go`
- Repositories: `internal/repositories/*.go`
- Services: `internal/services/*.go`
- Controllers: `internal/controllers/*.go`
- Router: `internal/router/*.go`
- Init (GORM/router): `internal/initialize/*.go`
- Responses: `pkg/response/*`
- SQL bootstrap: `sql/init.sql`

Notes
- Assume Gin for routing and middleware.
- JWT signing key and secrets must never be committed; ensure local `.env` and CI secrets are used.
- Prefer small, composable PRs per milestone bullet with tests.
