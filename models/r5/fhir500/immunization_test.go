package fhir500_test

import (
	"encoding/json"
	"testing"

	"github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

// immunizationR5JSON exercises the elements that R5 changed relative to R4B:
// manufacturer and reason are CodeableReference, programEligibility is a
// backbone element, reaction.manifestation replaced reaction.detail, and
// protocolApplied.doseNumber is a plain string.
const immunizationR5JSON = `{
  "resourceType": "Immunization",
  "id": "imm-1",
  "status": "completed",
  "vaccineCode": {"coding": [{"system": "http://hl7.org/fhir/sid/cvx", "code": "158"}]},
  "patient": {"reference": "Patient/p1"},
  "occurrenceDateTime": "2026-01-15",
  "manufacturer": {"reference": {"reference": "Organization/org1"}},
  "administeredProduct": {"concept": {"text": "Fluzone"}},
  "informationSource": {"concept": {"text": "self-reported"}},
  "reason": [{"concept": {"text": "seasonal campaign"}}],
  "programEligibility": [{
    "program": {"text": "KEPI"},
    "programStatus": {"text": "not-eligible"}
  }],
  "reaction": [{
    "date": "2026-01-16",
    "manifestation": {"reference": {"reference": "Observation/obs1"}},
    "reported": true
  }],
  "protocolApplied": [{
    "series": "2-dose",
    "doseNumber": "1",
    "seriesDoses": "2"
  }]
}`

func TestUnmarshalImmunization_R5Elements(t *testing.T) {
	imm, err := fhir500.UnmarshalImmunization([]byte(immunizationR5JSON))
	if err != nil {
		t.Fatalf("failed to unmarshal Immunization: %v", err)
	}

	if imm.Status == nil || *imm.Status != fhir500.ImmunizationStatusCompleted {
		t.Errorf("status = %v, want %q", imm.Status, fhir500.ImmunizationStatusCompleted)
	}

	// manufacturer moved from Reference to CodeableReference in R5.
	if imm.Manufacturer == nil || imm.Manufacturer.Reference == nil ||
		imm.Manufacturer.Reference.Reference == nil ||
		*imm.Manufacturer.Reference.Reference != "Organization/org1" {
		t.Errorf("manufacturer did not decode as a CodeableReference: %+v", imm.Manufacturer)
	}

	// reasonCode/reasonReference collapsed into reason (CodeableReference).
	if len(imm.Reason) != 1 || imm.Reason[0].Concept == nil {
		t.Errorf("reason did not decode as a CodeableReference list: %+v", imm.Reason)
	}

	// programEligibility became a backbone element carrying program + programStatus.
	if len(imm.ProgramEligibility) != 1 || imm.ProgramEligibility[0].Program == nil ||
		imm.ProgramEligibility[0].ProgramStatus == nil {
		t.Errorf("programEligibility did not decode as a backbone element: %+v", imm.ProgramEligibility)
	}

	// reaction.detail was replaced by reaction.manifestation (CodeableReference).
	if len(imm.Reaction) != 1 || imm.Reaction[0].Manifestation == nil ||
		imm.Reaction[0].Manifestation.Reference == nil {
		t.Errorf("reaction.manifestation did not decode: %+v", imm.Reaction)
	}

	// doseNumber lost its positiveInt variant and is a bare string in R5.
	if len(imm.ProtocolApplied) != 1 || imm.ProtocolApplied[0].DoseNumber == nil ||
		*imm.ProtocolApplied[0].DoseNumber != "1" {
		t.Errorf("protocolApplied.doseNumber did not decode as a string: %+v", imm.ProtocolApplied)
	}
}

func TestImmunization_MarshalJSON_AddsResourceType(t *testing.T) {
	id := "imm-1"
	status := fhir500.ImmunizationStatusNotDone

	payload, err := json.Marshal(fhir500.Immunization{ID: &id, Status: &status})
	if err != nil {
		t.Fatalf("failed to marshal Immunization: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("failed to decode marshalled Immunization: %v", err)
	}

	if decoded["resourceType"] != "Immunization" {
		t.Errorf("resourceType = %v, want Immunization", decoded["resourceType"])
	}

	// The FHIR code is hyphenated; an underscore would be rejected by the server.
	if decoded["status"] != "not-done" {
		t.Errorf("status = %v, want not-done", decoded["status"])
	}

	// Unset optional elements must stay out of the payload.
	if _, ok := decoded["vaccineCode"]; ok {
		t.Error("vaccineCode should be omitted when unset")
	}
}

func TestImmunizationStatus_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		status fhir500.ImmunizationStatus
		want   bool
	}{
		{name: "completed is a valid R5 immunization status", status: "completed", want: true},
		{name: "entered-in-error is a valid R5 immunization status", status: "entered-in-error", want: true},
		{name: "not-done is a valid R5 immunization status", status: "not-done", want: true},
		{name: "in-progress was dropped from the R5 value set", status: "in-progress", want: false},
		{name: "entered_in_error is not a FHIR code", status: "entered_in_error", want: false},
		{name: "an empty status is invalid", status: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
