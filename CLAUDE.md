# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`hapi-fhir-go` is a Go SDK for talking to HAPI FHIR servers. It supports both FHIR R4B and R5: a single `Client` issues the HTTP requests, and callers choose which version's model structs to unmarshal responses into. The library has no required runtime config beyond a base URL.

## Commands

```bash
# Run all tests with race detector and coverage (matches CI)
go test -coverprofile cover.out -v -race ./...
go tool cover -func=cover.out

# Run a single test
go test -run TestName -v ./...

# Lint (CI uses golangci-lint v8; config in .golangci.yaml)
golangci-lint run

# Build
go build ./...
```

Go 1.25 is required (see `go.mod` / CI matrix).

## Architecture

The top-level package `hapifhirgo` (root dir) holds the client and request machinery. Models live under `models/` and are version-segregated.

- **`client.go`** — `Client` struct, `NewClient`, and all `ClientOption` functional options (`WithBasicAuth`, `WithTransport`, `WithHTTPClient`, `WithRetry`, `WithDefaultHeaders`, `WithoutCacheControlHeader`, `WithCREnabledHAPIFHIR`). Also defines `defaultTransport()`, a connection-pool-tuned `http.Transport` that overrides Go's stdlib default (whose `MaxIdleConnsPerHost: 2` starves under concurrency).
- **`operations.go`** — the public FHIR REST surface (`CreateFHIRResource`, `GetFHIRResource`, `SearchFHIRResource`, `PutFHIRResource`, `PatchFHIRResource`, `FHIRPathPatch`, `DeleteFHIRResource`, `PostFHIRBundle`, plus extended ops `GetPatientEverything`, `GetEncounterEverything`, `ValidateResource`, `ExpandValueSet`, `ExtractFHIRResource`). Almost every method funnels through `makeRequest`.
- **`http.go`** — the request pipeline: `makeRequest` → `newRequest` (URL composition, header/auth/body setup) → `c.HTTP.Do` → `readResponse` (status handling, stream-decode into caller's destination). Also defines `APIError` and validation-response handling.
- **`retry.go`** — `RetryPolicy` and the `retryRoundTripper` installed by `WithRetry`.
- **`json.go`** — `structToMap` helper.
- **`models/r4b/fhir430/`** and **`models/r5/fhir500/`** — generated FHIR resource structs and enums (~400 files). Treat as generated data types, not hand-edited logic.

### Key conventions and gotchas

- **Caller-supplied destination.** Read/search methods take a `resource interface{}` / `bundle interface{}` out-param and stream-decode the response body into it. The caller picks the version-specific type (`r4b.Patient` vs `r5.Patient`) — the client itself is version-agnostic.
- **Create auto-validates.** `CreateFHIRResource` calls `ValidateResource` (`$validate`) before POSTing, except for `Questionnaire`/`QuestionnaireResponse`. It also injects `resourceType` and `language: "en"` into the payload.
- **`$validate` returns HTTP 200 even on failure.** `readResponse` special-cases any path containing `$validate` (`isValidateInPath`) and routes it to `handleValidationResponse`, which inspects `OperationOutcome.issue[].severity` — only `error`/`fatal` produce an `APIError`; `warning`/`information`/`success` pass.
- **Errors.** Non-2xx responses become an `APIError` carrying the parsed `OperationOutcome`. Use `APIError.GetOperationOutcome()` for programmatic access.
- **Retry excludes POST by default.** FHIR `create` is not idempotent without `If-None-Exist`, so the default `RetryableMethods` omits POST. Bodies that can't be replayed (`Body` set, `GetBody` nil) are sent once. Add POST to `RetryableMethods` only behind an idempotency-key gateway.
- **`Cache-Control: no-cache`** is set on every request by default to avoid stale search bundles; disable with `WithoutCacheControlHeader()` for cacheable static reads.
- **`PostFHIRBundle`** bypasses `makeRequest` and sets `Prefer: return=representation` so transaction responses include full resources.
- **CR-enabled server.** `useCREnabledServer` params and `WithCREnabledHAPIFHIR` route specific ops (e.g. `$extract`) to a separate Clinical Reasoning Module base URL.

## Releases

Versioning/changelog is automated via release-please (`.github/workflows/release-please.yml`); releases are driven by Conventional Commits (`feat:`, `fix:`, `chore:`). Do not hand-edit `CHANGELOG.md`.
