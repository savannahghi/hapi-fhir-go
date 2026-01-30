# hapi-fhir-go

A Go SDK for interacting with HAPI FHIR servers supporting both FHIR R4B and R5.

## Features

- **Multi-version support**: Use FHIR R4B or R5 models
- **Simple client**: One client for all FHIR versions
- **Smaller binaries**: Import only the FHIR version you need
- **Clean separation**: Clear distinction between FHIR versions

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
