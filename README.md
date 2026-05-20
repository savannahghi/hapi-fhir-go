# hapi-fhir-go

A Go SDK for interacting with HAPI FHIR servers supporting both FHIR R4B and R5.

## Features

- **Multi-version support**: Use FHIR R4B or R5 models
- **Simple client**: One client for all FHIR versions
- **Smaller binaries**: Import only the FHIR version you need
- **Clean separation**: Clear distinction between FHIR versions
- **Production-tuned HTTP transport** out of the box (HTTP/2, large idle pool, no admission cap)
- **Pluggable retry** for idempotent operations (`WithRetry`)
- **Customisable headers** (`WithDefaultHeaders`, `WithoutCacheControlHeader`)

## Tuning for high concurrency

The default transport is sized for a "medium" deployment (4-8 service pods talking to one HAPI host):

| Setting | Default |
|---|---|
| `MaxIdleConns` | 400 |
| `MaxIdleConnsPerHost` | 150 |
| `MaxConnsPerHost` | 0 (unlimited) |
| `IdleConnTimeout` | 90s |
| `TLSHandshakeTimeout` | 5s |
| `ForceAttemptHTTP2` | true |
| `Client.Timeout` | 60s |

If you run hotter (1000+ RPS per pod) or talk to multiple distinct hosts, inject your own transport:

```go
import "net/http"
import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

transport := &http.Transport{
    MaxIdleConns:        800,
    MaxIdleConnsPerHost: 300,
    IdleConnTimeout:     120 * time.Second,
    ForceAttemptHTTP2:   true,
    // ... see net/http docs for the full set
}

client, _ := hapifhirgo.NewClient(
    baseURL,
    hapifhirgo.WithTransport(otelhttp.NewTransport(transport)),
    hapifhirgo.WithRetry(hapifhirgo.RetryPolicy{
        MaxAttempts:    3,
        InitialBackoff: 100 * time.Millisecond,
        MaxBackoff:     2 * time.Second,
    }),
    hapifhirgo.WithoutCacheControlHeader(),
)
```

### POST is not retried by default

FHIR `create` is not idempotent without `If-None-Exist`. The retry policy excludes POST. If you must retry POST (e.g. behind an idempotency-key gateway), add it explicitly:

```go
WithRetry(RetryPolicy{
    MaxAttempts:      3,
    RetryableMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
})
```

## Quick Start

### Using R4B Models

```go
import (
    r4b "github.com/savannahghi/hapi-fhir-go/models/r4b/fhir430"
    "github.com/savannahghi/hapi-fhir-go"
)

client, err := hapifhirgo.NewClient("http://localhost:8080/fhir")
patient := r4b.Patient{...}
err = client.CreateFHIRResource(ctx, "Patient", payload, &patient)
```

### Using R5 Models

```go
import (
    r5 "github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
    "github.com/savannahghi/hapi-fhir-go"
)

client, err := hapifhirgo.NewClient("http://localhost:8080/fhir")
patient := r5.Patient{...}
err = client.CreateFHIRResource(ctx, "Patient", payload, &patient)
```

## Structure

```
hapi-fhir-go/
├── client.go              # HTTP client
├── operations.go          # FHIR operations
├── http.go                # HTTP utilities
├── json.go                # JSON utilities
├── models/
│   ├── r4b/fhir430/       # FHIR R4B models
│   └── r5/fhir500/        # FHIR R5 models
└── patient_test.go        # Tests
```

## Usage Examples

### Creating a Patient

```go
// R4B Patient
import r4b "github.com/savannahghi/hapi-fhir-go/models/r4b/fhir430"

patient := r4b.Patient{
    Name: []*r4b.HumanName{{
        Family: stringPtr("Doe"),
        Given:  []*string{stringPtr("Jane")},
    }},
    Gender: &r4b.PatientGenderEnumMale,
}

client, _ := hapifhirgo.NewClient("http://localhost:8080/fhir")
err := client.CreateFHIRResource(ctx, "Patient", map[string]interface{}{}, &patient)
```

### Searching Resources

```go
var bundle r5.Bundle
err := client.SearchFHIRResource(ctx, "", "Patient", map[string]any{
    "name": "Doe",
}, &bundle)
```

### Getting a Resource

```go
var patient r5.Patient
err := client.GetFHIRResource(ctx, "Patient", "patient-id", &patient)
```
