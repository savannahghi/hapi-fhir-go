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

## Authenticating as a service

A service calls the store as itself with the client credentials grant. `ClientCredentials` mints
tokens from any OAuth2 token endpoint (Keycloak's is `{base}/realms/{realm}/protocol/openid-connect/token`)
and `CachedTokenProvider` keeps one token for every request until it is about to expire, with
concurrent callers waiting on one mint. Any other minter, such as an identity SDK, fits the same
`TokenMinter` signature.

```go
mint := hapifhirgo.ClientCredentials{
    TokenURL:     "https://keycloak.example/realms/study/protocol/openid-connect/token",
    ClientID:     "study-service",
    ClientSecret: secret,
}.Mint

client, _ := hapifhirgo.NewClient(baseURL,
    hapifhirgo.WithTokenProvider(hapifhirgo.CachedTokenProvider(mint)),
)
```

## Errors

Every answer of 400 or above comes back as an `APIError`, whatever the body is. A FHIR server's
`OperationOutcome` is decoded into `OperationOutcome` and parsed into `Issues`; a gateway's HTML
page leaves both empty and keeps the bytes in `Body`. `Diagnostics()` joins the text of the error
issues, which is what to show or log.

```go
var apiErr hapifhirgo.APIError
if errors.As(err, &apiErr) {
    switch apiErr.StatusCode {
    case http.StatusNotFound:
        // not there
    case http.StatusUnprocessableEntity:
        log.Println(apiErr.Diagnostics())
    }
}
```

## Readiness

`Metadata` reads the capability statement without credentials, so a readiness probe never asks
the identity provider for a token.

```go
var capability struct {
    FHIRVersion string `json:"fhirVersion"`
}
err := client.Metadata(ctx, map[string]any{"_summary": "true"}, &capability)
```

## Writing a model

`PutResource` and `CreateResource` take a model or a map as it is. The resource's `resourceType`
is checked against the type in the path and stamped when absent, and numbers are sent as written.
`PutResource` writes under an id the caller chose, so a write that failed part way can be sent
again without a duplicate. Neither runs `$validate` first; `CreateFHIRResource` still does. A
resource that cannot be sent as asked, because it is not an object or carries another type, is
refused before anything goes out with an error that wraps `ErrBadResource`.

```go
patient := r5.Patient{ID: &id, Name: []*r5.HumanName{{Family: &family}}}

var stored r5.Patient
err := client.PutResource(ctx, "Patient", id, patient, &stored)
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
├── tools/r5import/        # Copies generated R5 models into models/r5/fhir500
└── patient_test.go        # Tests
```

## R5 models

`models/r5/fhir500` covers every resource, datatype and code list in FHIR R5 (5.0.0). Most files are generated by [golang-fhir-models](https://github.com/samply/golang-fhir-models) and keep its licence header and "do not edit" notice. The rest are hand-written models kept for existing callers; their JSON elements match R5.

When using them:

- Required codes such as `Encounter.status` are Go enums that marshal to the FHIR code. An unknown code fails to unmarshal, so an R4 value like `finished` is rejected.
- A choice element such as `value[x]` is one pointer field per type. Set exactly one.
- Decimals are `json.Number`, so `128.50` keeps its precision. `integer64` is a string, as FHIR JSON sends it.
- Generated resources add `resourceType` when marshalled and come with an `UnmarshalX` function, for example `fhir500.UnmarshalEncounter`.

### Regenerating

The import needs a generator build that makes choice variants optional and maps `integer64`. Both fixes are on the `fix/r5-choice-fields-and-integer64` branch of [Salaton/golang-fhir-models](https://github.com/Salaton/golang-fhir-models).

```sh
git clone -b fix/r5-choice-fields-and-integer64 https://github.com/Salaton/golang-fhir-models
(cd golang-fhir-models/fhir-models-gen && go install .)
(cd golang-fhir-models/fhir-models && go generate ./fhir)

# from this repository
go run ./tools/r5import -src ../golang-fhir-models/fhir-models/fhir -dry-run
go run ./tools/r5import -src ../golang-fhir-models/fhir-models/fhir
```

`r5import` only adds files. It copies a generated file when none of the types it declares exists in the package, renames the package to `fhir500`, drops the bson tags and renames `Id` to `ID`. It never overwrites a hand-written model. To replace one with the generated version, delete it and run the import again.

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
