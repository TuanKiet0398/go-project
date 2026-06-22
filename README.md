# Golang Learning Plan

A roadmap for learning Go, tracked through the projects in this repo.

## 1. Go Basics — `go-simple/`
- [x] Basic variables, slices, functions
- [x] Plain `net/http` (no framework)
- [ ] Structs & methods
- [ ] Interfaces
- [ ] Pointers
- [ ] Error handling (`error`, `errors.Is/As`, custom errors)
- [ ] Goroutines & channels (basic concurrency)
- [ ] Unit testing with the `testing` package

## 2. Web App — `go-webapp/go-stripe/`
- [x] `application` + `config` structs (dependency injection pattern)
- [x] Custom `http.Server` with timeouts — `cmd/web/main.go`
- [ ] Routing with `chi` — define routes in `cmd/web/routes.go`
- [ ] Middleware (logging, panic recovery, CORS)
- [ ] HTML templates + template caching (`templateCache`)
- [ ] Serving static files (CSS/JS) from `web/`
- [ ] Database connection (Postgres) — `internal` package
  - [ ] Models & repository pattern
  - [ ] Migrations
- [ ] Stripe integration
  - [ ] Create Payment Intent
  - [ ] Webhook handling for payment events
  - [ ] Checkout page (card input form)
- [ ] Sessions & authentication (admin login)
- [ ] Form input validation
- [ ] Environment-based config (dev/staging/production)

## 3. Best Practices & Extensions
- [ ] Structured logging
- [ ] Graceful server shutdown
- [ ] Handler tests (httptest)
- [ ] Dockerize the app
- [ ] Basic CI (build + test)

## Notes
- Prioritize finishing `go-webapp/go-stripe` as a complete end-to-end project before moving on to new topics.
- Check off `[x]` items as they're completed to track progress.
