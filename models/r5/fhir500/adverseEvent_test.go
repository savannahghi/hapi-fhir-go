package fhir500_test

import (
	"encoding/json"
	"testing"

	"github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

// adverseEventR5JSON exercises the elements R5 reworked relative to R4B:
// event became code, date became occurrence[x], resultingCondition became
// resultingEffect, contributor became participant, suspectEntity.instance
// became a choice, and causality gained assessmentMethod/entityRelatedness.
const adverseEventR5JSON = `{
  "resourceType": "AdverseEvent",
  "id": "ae-1",
  "status": "completed",
  "actuality": "actual",
  "code": {"coding": [{"system": "http://snomed.info/sct", "code": "281647001"}]},
  "subject": {"reference": "Patient/p1"},
  "occurrencePeriod": {"start": "2026-02-01", "end": "2026-02-03"},
  "detected": "2026-02-02",
  "recordedDate": "2026-02-04",
  "resultingEffect": [{"reference": "Condition/c1"}],
  "seriousness": {"text": "serious"},
  "outcome": [{"text": "recovered"}],
  "expectedInResearchStudy": false,
  "participant": [{
    "function": {"text": "informant"},
    "actor": {"reference": "Practitioner/pr1"}
  }],
  "suspectEntity": [{
    "instanceReference": {"reference": "Medication/med1"},
    "causality": {
      "assessmentMethod": {"text": "WHO-UMC"},
      "entityRelatedness": {"text": "probable"},
      "author": {"reference": "Practitioner/pr1"}
    }
  }],
  "contributingFactor": [{"itemCodeableConcept": {"text": "renal impairment"}}],
  "supportingInfo": [{"itemReference": {"reference": "Observation/obs1"}}]
}`

func TestUnmarshalAdverseEvent_R5Elements(t *testing.T) {
	event, err := fhir500.UnmarshalAdverseEvent([]byte(adverseEventR5JSON))
	if err != nil {
		t.Fatalf("failed to unmarshal AdverseEvent: %v", err)
	}

	// status is new in R5; R4B AdverseEvent had none.
	if event.Status == nil || *event.Status != fhir500.AdverseEventStatusCompleted {
		t.Errorf("status = %v, want %q", event.Status, fhir500.AdverseEventStatusCompleted)
	}

	if event.Actuality == nil || *event.Actuality != fhir500.AdverseEventActualityActual {
		t.Errorf("actuality = %v, want %q", event.Actuality, fhir500.AdverseEventActualityActual)
	}

	// R4B event was renamed to code.
	if event.Code == nil || len(event.Code.Coding) == 0 {
		t.Errorf("code did not decode: %+v", event.Code)
	}

	// R4B date became the occurrence[x] choice.
	if event.OccurrencePeriod == nil || event.OccurrencePeriod.Start == "" {
		t.Errorf("occurrencePeriod did not decode: %+v", event.OccurrencePeriod)
	}

	// R4B resultingCondition was renamed to resultingEffect.
	if len(event.ResultingEffect) != 1 {
		t.Errorf("resultingEffect did not decode: %+v", event.ResultingEffect)
	}

	// R4B contributor became the participant backbone element.
	if len(event.Participant) != 1 || event.Participant[0].Actor == nil {
		t.Errorf("participant did not decode: %+v", event.Participant)
	}

	// suspectEntity.instance is a choice in R5.
	if len(event.SuspectEntity) != 1 || event.SuspectEntity[0].InstanceReference == nil {
		t.Fatalf("suspectEntity.instanceReference did not decode: %+v", event.SuspectEntity)
	}

	causality := event.SuspectEntity[0].Causality
	if causality == nil || causality.AssessmentMethod == nil || causality.EntityRelatedness == nil {
		t.Errorf("causality did not decode assessmentMethod/entityRelatedness: %+v", causality)
	}

	if len(event.ContributingFactor) != 1 || event.ContributingFactor[0].ItemCodeableConcept == nil {
		t.Errorf("contributingFactor did not decode: %+v", event.ContributingFactor)
	}

	if len(event.SupportingInfo) != 1 || event.SupportingInfo[0].ItemReference == nil {
		t.Errorf("supportingInfo did not decode: %+v", event.SupportingInfo)
	}
}

func TestAdverseEvent_MarshalJSON_AddsResourceType(t *testing.T) {
	id := "ae-1"
	status := fhir500.AdverseEventStatusInProgress
	actuality := fhir500.AdverseEventActualityPotential

	payload, err := json.Marshal(fhir500.AdverseEvent{
		ID:        &id,
		Status:    &status,
		Actuality: &actuality,
	})
	if err != nil {
		t.Fatalf("failed to marshal AdverseEvent: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("failed to decode marshalled AdverseEvent: %v", err)
	}

	if decoded["resourceType"] != "AdverseEvent" {
		t.Errorf("resourceType = %v, want AdverseEvent", decoded["resourceType"])
	}

	if decoded["status"] != "in-progress" {
		t.Errorf("status = %v, want in-progress", decoded["status"])
	}

	// R4B-only elements must not reappear on the R5 model.
	for _, removed := range []string{"event", "date", "resultingCondition", "severity", "contributor"} {
		if _, ok := decoded[removed]; ok {
			t.Errorf("%q is an R4B element and should not be emitted by the R5 model", removed)
		}
	}
}

func TestAdverseEventStatusAndActuality_IsValid(t *testing.T) {
	statusTests := []struct {
		name   string
		status fhir500.AdverseEventStatus
		want   bool
	}{
		{name: "in-progress is a valid adverse event status", status: "in-progress", want: true},
		{name: "completed is a valid adverse event status", status: "completed", want: true},
		{name: "entered-in-error is a valid adverse event status", status: "entered-in-error", want: true},
		{name: "unknown is a valid adverse event status", status: "unknown", want: true},
		{name: "in_progress is not a FHIR code", status: "in_progress", want: false},
		{name: "an empty status is invalid", status: "", want: false},
	}

	for _, tt := range statusTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}

	actualityTests := []struct {
		name      string
		actuality fhir500.AdverseEventActuality
		want      bool
	}{
		{name: "actual is a valid actuality", actuality: "actual", want: true},
		{name: "potential is a valid actuality", actuality: "potential", want: true},
		{name: "avoided is not in the actuality value set", actuality: "avoided", want: false},
		{name: "an empty actuality is invalid", actuality: "", want: false},
	}

	for _, tt := range actualityTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.actuality.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
