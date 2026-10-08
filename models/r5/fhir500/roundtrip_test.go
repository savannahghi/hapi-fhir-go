package fhir500_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

const researchStudyJSON = `{
  "resourceType": "ResearchStudy",
  "id": "screening-study",
  "url": "https://example.org/fhir/ResearchStudy/screening-study",
  "identifier": [{"system": "https://example.org/study", "value": "SS-001"}],
  "version": "2",
  "name": "ScreeningStudy",
  "title": "Retinal screening for chronic kidney disease",
  "label": [{"type": {"text": "short title"}, "value": "SS"}],
  "status": "active",
  "primaryPurposeType": {"coding": [{"system": "http://hl7.org/fhir/research-study-prim-purp-type", "code": "screening"}]},
  "phase": {"coding": [{"system": "http://hl7.org/fhir/research-study-phase", "code": "n-a"}]},
  "studyDesign": [{"text": "diagnostic accuracy"}],
  "focus": [{"concept": {"text": "retinal imaging"}}],
  "condition": [{"text": "chronic kidney disease"}],
  "period": {"start": "2026-10-01"},
  "site": [{"reference": "Organization/site-1"}],
  "associatedParty": [{
    "name": "Sponsor",
    "role": {"coding": [{"system": "http://hl7.org/fhir/research-study-party-role", "code": "sponsor"}]},
    "party": {"reference": "Organization/sponsor"}
  }],
  "progressStatus": [{
    "state": {"coding": [{"system": "http://hl7.org/fhir/research-study-status", "code": "active"}]},
    "actual": true,
    "period": {"start": "2026-10-01"}
  }],
  "recruitment": {"targetNumber": 400, "actualNumber": 12, "eligibility": {"reference": "Group/eligibility"}},
  "comparisonGroup": [{"linkId": "screened", "name": "Screened", "description": "All screened participants"}],
  "objective": [{
    "name": "accuracy",
    "type": {"coding": [{"system": "http://hl7.org/fhir/research-study-objective-type", "code": "primary"}]},
    "description": "Sensitivity of the retinal score"
  }]
}`

const researchSubjectJSON = `{
  "resourceType": "ResearchSubject",
  "id": "rs-1",
  "identifier": [{"system": "https://example.org/participant", "value": "SS-001-0001"}],
  "status": "active",
  "progress": [{
    "type": {"text": "state"},
    "subjectState": {"coding": [{"system": "http://hl7.org/fhir/research-subject-state", "code": "on-study"}]},
    "startDate": "2026-10-02"
  }],
  "period": {"start": "2026-10-02"},
  "study": {"reference": "ResearchStudy/screening-study"},
  "subject": {"reference": "Patient/p1"},
  "assignedComparisonGroup": "screened",
  "consent": [{"reference": "Consent/c1"}]
}`

const encounterJSON = `{
  "resourceType": "Encounter",
  "id": "visit-1",
  "status": "in-progress",
  "class": [{"coding": [{"system": "http://terminology.hl7.org/CodeSystem/v3-ActCode", "code": "AMB"}]}],
  "subject": {"reference": "Patient/p1"},
  "basedOn": [{"reference": "CarePlan/cp1"}],
  "serviceProvider": {"reference": "Organization/site-1"},
  "participant": [{"type": [{"text": "screener"}], "actor": {"reference": "Practitioner/pr1"}}],
  "actualPeriod": {"start": "2026-10-02T09:00:00+03:00"},
  "plannedStartDate": "2026-10-02T08:30:00+03:00",
  "reason": [{"value": [{"concept": {"text": "baseline visit"}}]}],
  "location": [{"location": {"reference": "Location/l1"}, "status": "active"}]
}`

const consentJSON = `{
  "resourceType": "Consent",
  "id": "c1",
  "status": "active",
  "category": [{"coding": [{"system": "http://loinc.org", "code": "59284-0"}]}],
  "subject": {"reference": "Patient/p1"},
  "date": "2026-10-02",
  "period": {"start": "2026-10-02"},
  "grantor": [{"reference": "Patient/p1"}],
  "grantee": [{"reference": "Organization/sponsor"}],
  "controller": [{"reference": "Organization/sponsor"}],
  "sourceAttachment": [{"contentType": "image/jpeg", "url": "https://example.org/consent/c1.jpg", "title": "Signed consent form v2"}],
  "regulatoryBasis": [{"text": "ICH GCP"}],
  "policyBasis": {"url": "https://example.org/consent-form/v2"},
  "verification": [{
    "verified": true,
    "verificationType": {"text": "impartial witness"},
    "verifiedBy": {"reference": "Practitioner/w1"},
    "verificationDate": ["2026-10-02T10:00:00+03:00"]
  }],
  "decision": "permit",
  "provision": [{
    "period": {"start": "2026-10-02"},
    "actor": [{"role": {"text": "investigator"}, "reference": {"reference": "Organization/sponsor"}}],
    "action": [{"text": "collect"}],
    "purpose": [{"system": "http://terminology.hl7.org/CodeSystem/v3-ActReason", "code": "HRESCH"}],
    "provision": [{
      "action": [{"text": "access"}],
      "data": [{"meaning": "related", "reference": {"reference": "ResearchStudy/screening-study"}}]
    }]
  }]
}`

const questionnaireJSON = `{
  "resourceType": "Questionnaire",
  "id": "screening",
  "meta": {"profile": ["http://hl7.org/fhir/uv/sdc/StructureDefinition/sdc-questionnaire-extr-smap"]},
  "extension": [{
    "url": "http://hl7.org/fhir/uv/sdc/StructureDefinition/sdc-questionnaire-targetStructureMap",
    "valueCanonical": "https://example.org/fhir/StructureMap/screening"
  }],
  "url": "https://example.org/fhir/Questionnaire/screening",
  "name": "Screening",
  "status": "active",
  "subjectType": ["Patient"],
  "item": [
    {"linkId": "smoker", "text": "Do you smoke?", "type": "boolean", "required": true},
    {"linkId": "sbp", "text": "Systolic blood pressure", "type": "decimal"},
    {
      "linkId": "diabetes",
      "text": "Diabetes",
      "type": "coding",
      "answerConstraint": "optionsOnly",
      "answerOption": [{"valueCoding": {"system": "http://snomed.info/sct", "code": "44054006", "display": "Diabetes mellitus type 2"}}]
    }
  ]
}`

const questionnaireResponseJSON = `{
  "resourceType": "QuestionnaireResponse",
  "id": "qr-1",
  "questionnaire": "https://example.org/fhir/Questionnaire/screening",
  "status": "completed",
  "subject": {"reference": "Patient/p1"},
  "encounter": {"reference": "Encounter/visit-1"},
  "authored": "2026-10-02T10:15:00+03:00",
  "item": [
    {"linkId": "smoker", "answer": [{"valueBoolean": false}]},
    {"linkId": "sbp", "answer": [{"valueDecimal": 128.50}]},
    {"linkId": "diabetes", "answer": [{"valueCoding": {"system": "http://snomed.info/sct", "code": "44054006"}}]}
  ]
}`

const observationJSON = `{
  "resourceType": "Observation",
  "id": "obs-1",
  "status": "final",
  "category": [{"coding": [{"system": "http://terminology.hl7.org/CodeSystem/observation-category", "code": "survey"}]}],
  "code": {"coding": [{"system": "http://loinc.org", "code": "72166-2"}]},
  "subject": {"reference": "Patient/p1"},
  "encounter": {"reference": "Encounter/visit-1"},
  "effectiveDateTime": "2026-10-02T10:15:00+03:00",
  "valueCodeableConcept": {"coding": [{"system": "http://snomed.info/sct", "code": "8517006"}]},
  "derivedFrom": [{"reference": "QuestionnaireResponse/qr-1"}],
  "triggeredBy": [{"observation": {"reference": "Observation/obs-0"}, "type": "reflex"}],
  "component": [
    {"code": {"text": "time of last cigarette"}, "valueTime": "07:30:00"},
    {"code": {"text": "cigarettes per day"}, "valueInteger": 0}
  ]
}`

const organizationJSON = `{
  "resourceType": "Organization",
  "id": "site-1",
  "identifier": [{"system": "https://example.org/facility", "value": "12345"}],
  "active": true,
  "type": [{"coding": [{"system": "http://terminology.hl7.org/CodeSystem/organization-type", "code": "prov"}]}],
  "name": "Example Site Hospital",
  "contact": [{"telecom": [{"system": "phone", "value": "+254700000000"}], "address": {"city": "Nairobi", "country": "KE"}}],
  "partOf": {"reference": "Organization/sponsor"}
}`

const taskJSON = `{
  "resourceType": "Task",
  "id": "t1",
  "status": "requested",
  "intent": "order",
  "priority": "routine",
  "code": {"text": "complete the screening form"},
  "focus": {"reference": "Questionnaire/screening"},
  "for": {"reference": "Patient/p1"},
  "requestedPerformer": [{"concept": {"text": "screener"}}],
  "performer": [{"actor": {"reference": "Practitioner/pr1"}}],
  "reason": [{"concept": {"text": "baseline visit"}}],
  "input": [{"type": {"text": "questionnaire"}, "valueCanonical": "https://example.org/fhir/Questionnaire/screening"}],
  "output": [{"type": {"text": "response"}, "valueReference": {"reference": "QuestionnaireResponse/qr-1"}}]
}`

const serviceRequestJSON = `{
  "resourceType": "ServiceRequest",
  "id": "sr1",
  "instantiatesCanonical": ["https://example.org/fhir/PlanDefinition/referral"],
  "status": "active",
  "intent": "order",
  "code": {"concept": {"coding": [{"system": "https://example.org/tests", "code": "TEST", "display": "Retinal photograph"}]}},
  "subject": {"reference": "Patient/p1"},
  "occurrenceTiming": {"repeat": {"frequency": 1, "period": 1, "periodUnit": "d", "when": ["MORN"], "timeOfDay": ["08:00:00"]}},
  "location": [{"reference": {"reference": "Location/l1"}}],
  "reason": [{"concept": {"text": "raised creatinine"}}],
  "patientInstruction": [{"instructionMarkdown": "Fast for 8 hours"}]
}`

// sameAfterRoundTrip decodes in into a T, encodes it again, and fails unless the
// result is the same JSON document. A dropped, renamed or invented element
// shows up as a difference.
func sameAfterRoundTrip[T any](t *testing.T, in string) {
	t.Helper()

	var resource T
	if err := json.Unmarshal([]byte(in), &resource); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	out, err := json.Marshal(resource)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var want, got any
	if err := json.Unmarshal([]byte(in), &want); err != nil {
		t.Fatalf("decode the fixture: %v", err)
	}

	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatalf("decode the output: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("round trip changed the resource\n got: %s\nwant: %s", out, in)
	}
}

func TestR5Resources_RoundTrip(t *testing.T) {
	tests := []struct {
		name      string
		json      string
		roundTrip func(*testing.T, string)
	}{
		{name: "ResearchStudy keeps recruitment, comparison groups and objectives", json: researchStudyJSON, roundTrip: sameAfterRoundTrip[fhir500.ResearchStudy]},
		{name: "ResearchSubject keeps progress and the assigned group", json: researchSubjectJSON, roundTrip: sameAfterRoundTrip[fhir500.ResearchSubject]},
		{name: "Encounter keeps R5 class, actualPeriod and reason", json: encounterJSON, roundTrip: sameAfterRoundTrip[fhir500.Encounter]},
		{name: "Consent keeps R5 decision, controller, verification and nested provisions", json: consentJSON, roundTrip: sameAfterRoundTrip[fhir500.Consent]},
		{name: "Questionnaire keeps the SDC targetStructureMap extension", json: questionnaireJSON, roundTrip: sameAfterRoundTrip[fhir500.Questionnaire]},
		{name: "QuestionnaireResponse keeps a false answer and decimal precision", json: questionnaireResponseJSON, roundTrip: sameAfterRoundTrip[fhir500.QuestionnaireResponse]},
		{name: "Observation keeps a coded value, a time value and triggeredBy", json: observationJSON, roundTrip: sameAfterRoundTrip[fhir500.Observation]},
		{name: "Organization keeps R5 contact and partOf", json: organizationJSON, roundTrip: sameAfterRoundTrip[fhir500.Organization]},
		{name: "Task keeps R5 performer, reason and one value per input", json: taskJSON, roundTrip: sameAfterRoundTrip[fhir500.Task]},
		{name: "ServiceRequest keeps R5 code, location, reason and repeated timing", json: serviceRequestJSON, roundTrip: sameAfterRoundTrip[fhir500.ServiceRequest]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.roundTrip(t, tt.json)
		})
	}
}
