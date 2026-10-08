package fhir500_test

import (
	"encoding/json"
	"testing"

	"github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

// TestHandWrittenResources_MarshalResourceType pins that the hand-written resources put their
// resourceType on the wire the way the generated ones do, so a model can be sent as it is.
func TestHandWrittenResources_MarshalResourceType(t *testing.T) {
	id := "x"

	tests := []struct {
		name     string
		resource any
		want     string
	}{
		{name: "happy case: a patient", resource: fhir500.Patient{ID: &id}, want: "Patient"},
		{name: "happy case: a patient by pointer", resource: &fhir500.Patient{ID: &id}, want: "Patient"},
		{name: "happy case: an operation outcome", resource: fhir500.OperationOutcome{}, want: "OperationOutcome"},
		{name: "happy case: a bundle", resource: fhir500.Bundle{}, want: "Bundle"},
		{name: "happy case: a generated resource", resource: fhir500.Encounter{}, want: "Encounter"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, err := json.Marshal(tt.resource)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}

			var typed struct {
				ResourceType string `json:"resourceType"`
				ID           string `json:"id"`
			}

			if err := json.Unmarshal(raw, &typed); err != nil {
				t.Fatalf("decode: %v", err)
			}

			if typed.ResourceType != tt.want {
				t.Errorf("resourceType = %q, want %q", typed.ResourceType, tt.want)
			}
		})
	}
}

func TestUnmarshalPatient(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantID  string
		wantErr bool
	}{
		{name: "happy case: a patient reads back", json: `{"resourceType":"Patient","id":"p1","gender":"female"}`, wantID: "p1"},
		{name: "sad case: not json", json: `<html>`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patient, err := fhir500.UnmarshalPatient([]byte(tt.json))

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}

				return
			}

			if err != nil {
				t.Fatalf("unmarshal: %v", err)
			}

			if patient.ID == nil || *patient.ID != tt.wantID {
				t.Errorf("id = %v, want %q", patient.ID, tt.wantID)
			}
		})
	}
}
