package fhir500_test

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

func ptr[T any](v T) *T { return &v }

// keysOf marshals v and returns its top-level JSON keys, sorted.
func keysOf(t *testing.T, v any) []string {
	t.Helper()

	payload, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode: %v", err)
	}

	keys := make([]string, 0, len(decoded))
	for k := range decoded {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func TestExtension_MarshalsFHIRKeys(t *testing.T) {
	tests := []struct {
		name      string
		extension fhir500.Extension
		wantKey   string
		wantValue string
	}{
		{
			name:      "a false boolean is kept",
			extension: fhir500.Extension{URL: "https://example.org/flag", ValueBoolean: ptr(false)},
			wantKey:   "valueBoolean",
			wantValue: "false",
		},
		{
			name:      "a canonical is sent as valueCanonical",
			extension: fhir500.Extension{URL: "https://example.org/map", ValueCanonical: "https://example.org/StructureMap/x"},
			wantKey:   "valueCanonical",
			wantValue: `"https://example.org/StructureMap/x"`,
		},
		{
			name:      "a url is sent as valueUrl",
			extension: fhir500.Extension{URL: "https://example.org/link", ValueURL: "https://example.org/a"},
			wantKey:   "valueUrl",
			wantValue: `"https://example.org/a"`,
		},
		{
			name:      "a uri is sent as valueUri, not valueURI",
			extension: fhir500.Extension{URL: "https://example.org/uri", ValueURI: "urn:example:1"},
			wantKey:   "valueUri",
			wantValue: `"urn:example:1"`,
		},
		{
			name:      "a uuid is sent as valueUuid, not valueUUID",
			extension: fhir500.Extension{URL: "https://example.org/uuid", ValueUUID: "urn:uuid:0d2c7c4e-1b8f-4a59-9d0b-2f4b8f3f8a11"},
			wantKey:   "valueUuid",
			wantValue: `"urn:uuid:0d2c7c4e-1b8f-4a59-9d0b-2f4b8f3f8a11"`,
		},
		{
			name:      "an id is sent as valueId, not valueID",
			extension: fhir500.Extension{URL: "https://example.org/id", ValueID: "abc"},
			wantKey:   "valueId",
			wantValue: `"abc"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := json.Marshal(tt.extension)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}

			var decoded map[string]json.RawMessage
			if err := json.Unmarshal(payload, &decoded); err != nil {
				t.Fatalf("decode: %v", err)
			}

			if got := string(decoded[tt.wantKey]); got != tt.wantValue {
				t.Errorf("%s = %q, want %q (payload %s)", tt.wantKey, got, tt.wantValue, payload)
			}

			if len(decoded) != 2 {
				t.Errorf("want only url and %s, got %s", tt.wantKey, payload)
			}
		})
	}
}

func TestParametersParameter_MarshalsFHIRKeys(t *testing.T) {
	tests := []struct {
		name      string
		parameter fhir500.ParametersParameter
		want      []string
	}{
		{name: "a uri is sent as valueUri", parameter: fhir500.ParametersParameter{Name: "a", ValueUrI: ptr("urn:x")}, want: []string{"name", "valueUri"}},
		{name: "a url is sent as valueUrl", parameter: fhir500.ParametersParameter{Name: "a", ValueUrL: ptr("https://x")}, want: []string{"name", "valueUrl"}},
		{name: "a uuid is sent as valueUuid", parameter: fhir500.ParametersParameter{Name: "a", ValueUUID: ptr("urn:uuid:1")}, want: []string{"name", "valueUuid"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keysOf(t, tt.parameter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keys = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChoiceElements_EmitOnlyTheChosenValue(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  []string
	}{
		{
			name:  "an answer of false emits only valueBoolean",
			value: fhir500.QuestionnaireResponseItemAnswer{ValueBoolean: ptr(false)},
			want:  []string{"valueBoolean"},
		},
		{
			name:  "a coded answer emits only valueCoding",
			value: fhir500.QuestionnaireResponseItemAnswer{ValueCoding: &fhir500.Coding{Code: ptr("44054006")}},
			want:  []string{"valueCoding"},
		},
		{
			name:  "a task input emits its type and one value",
			value: fhir500.TaskInput{Type: fhir500.CodeableConcept{Text: "form"}, ValueCanonical: ptr("https://example.org/Questionnaire/x")},
			want:  []string{"type", "valueCanonical"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keysOf(t, tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keys = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusCodes_RejectR4Values(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		unmarshal func([]byte) error
		wantErr   bool
	}{
		{
			name:      "in-progress is an R5 encounter status",
			payload:   `{"resourceType": "Encounter", "status": "in-progress"}`,
			unmarshal: func(b []byte) error { _, err := fhir500.UnmarshalEncounter(b); return err },
		},
		{
			name:      "finished was dropped from the R5 encounter statuses",
			payload:   `{"resourceType": "Encounter", "status": "finished"}`,
			unmarshal: func(b []byte) error { _, err := fhir500.UnmarshalEncounter(b); return err },
			wantErr:   true,
		},
		{
			name:      "in_progress is not a FHIR code",
			payload:   `{"resourceType": "Encounter", "status": "in_progress"}`,
			unmarshal: func(b []byte) error { _, err := fhir500.UnmarshalEncounter(b); return err },
			wantErr:   true,
		},
		{
			name:      "not-done is an R5 consent status",
			payload:   `{"resourceType": "Consent", "status": "not-done"}`,
			unmarshal: func(b []byte) error { _, err := fhir500.UnmarshalConsent(b); return err },
		},
		{
			name:      "rejected was dropped from the R5 consent statuses",
			payload:   `{"resourceType": "Consent", "status": "rejected"}`,
			unmarshal: func(b []byte) error { _, err := fhir500.UnmarshalConsent(b); return err },
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.unmarshal([]byte(tt.payload))
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmptyStructElements_AreOmitted(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  []string
	}{
		{
			name:  "an identifier without a type does not send an empty type",
			value: fhir500.Identifier{System: ptr("https://example.org/id"), Value: "1"},
			want:  []string{"system", "value"},
		},
		{
			name:  "a meta with only a profile does not send a zero lastUpdated",
			value: fhir500.Meta{Profile: []string{"https://example.org/StructureDefinition/x"}},
			want:  []string{"profile"},
		},
		{
			name:  "a range with only a low bound does not send an empty high",
			value: fhir500.Range{Low: fhir500.Quantity{Value: 90, Unit: "mmHg"}},
			want:  []string{"low"},
		},
		{
			name:  "a quantity without a system or code sends neither",
			value: fhir500.Quantity{Value: 5, Unit: "cigarettes"},
			want:  []string{"unit", "value"},
		},
		{
			name:  "a quantity of zero keeps its value",
			value: fhir500.Quantity{},
			want:  []string{"value"},
		},
		{
			name: "a usage context sends only the value it holds",
			value: fhir500.UsageContext{
				Code:                 fhir500.Coding{Code: ptr("focus")},
				ValueCodeableConcept: fhir500.CodeableConcept{Text: "screening"},
			},
			want: []string{"code", "valueCodeableConcept"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keysOf(t, tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keys = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetServiceRequestIDFromTask(t *testing.T) {
	tests := []struct {
		name    string
		task    *fhir500.Task
		want    string
		wantErr bool
	}{
		{
			name: "the referral service request is found among basedOn",
			task: &fhir500.Task{BasedOn: []fhir500.Reference{
				{ID: ptr("cp1"), Type: ptr("CarePlan")},
				{ID: ptr("sr1"), Type: ptr(fhir500.ReferralServiceRequestType.String())},
			}},
			want: "ServiceRequest/sr1",
		},
		{name: "a task based on nothing has no referral", task: &fhir500.Task{}, want: ""},
		{name: "a nil task is an error", task: nil, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.task.GetServiceRequestIDFromTask()
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestServiceRequest_HelpersReadTheR5Code(t *testing.T) {
	coded := &fhir500.ServiceRequest{Code: &fhir500.CodeableReference{Concept: &fhir500.CodeableConcept{
		Coding: []*fhir500.Coding{
			{Code: ptr("TEST"), Display: "Retinal photograph"},
			{Code: ptr("159623"), Display: "Raised creatinine"},
		},
	}}}

	tests := []struct {
		name           string
		request        *fhir500.ServiceRequest
		wantTest       string
		wantReason     string
		wantRequested  []string
		requestedCIELs string
	}{
		{
			name:           "the code concept is read",
			request:        coded,
			wantTest:       "Retinal photograph",
			wantReason:     "Raised creatinine",
			wantRequested:  []string{"Retinal photograph"},
			requestedCIELs: "TEST",
		},
		{
			name:           "a request without a code falls back instead of panicking",
			request:        &fhir500.ServiceRequest{},
			wantTest:       "",
			wantReason:     "test",
			requestedCIELs: "TEST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.request.GetPatientReferralTest(); got != tt.wantTest {
				t.Errorf("GetPatientReferralTest() = %q, want %q", got, tt.wantTest)
			}

			if got := tt.request.GetPatientReferralReason(); got != tt.wantReason {
				t.Errorf("GetPatientReferralReason() = %q, want %q", got, tt.wantReason)
			}

			if got := tt.request.GetRequestedServices(tt.requestedCIELs); !reflect.DeepEqual(got, tt.wantRequested) {
				t.Errorf("GetRequestedServices() = %v, want %v", got, tt.wantRequested)
			}
		})
	}
}
