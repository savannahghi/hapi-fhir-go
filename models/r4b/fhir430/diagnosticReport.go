package fhir430

import "encoding/json"

// DiagnosticReport is documented here http://hl7.org/fhir/StructureDefinition/DiagnosticReport
// The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport struct {
	ID                 *string                 `json:"id,omitempty"`
	Meta               *Meta                   `json:"meta,omitempty"`
	ImplicitRules      *string                 `json:"implicitRules,omitempty"`
	Language           *string                 `json:"language,omitempty"`
	Text               *Narrative              `json:"text,omitempty"`
	Contained          []json.RawMessage       `json:"contained,omitempty"`
	Extension          []Extension             `json:"extension,omitempty"`
	ModifierExtension  []Extension             `json:"modifierExtension,omitempty"`
	Identifier         []Identifier            `json:"identifier,omitempty"`
	BasedOn            []Reference             `json:"basedOn,omitempty"`
	Status             DiagnosticReportStatus  `json:"status"`
	Category           []CodeableConcept       `json:"category,omitempty"`
	Code               CodeableConcept         `json:"code"`
	Subject            *Reference              `json:"subject,omitempty"`
	Encounter          *Reference              `json:"encounter,omitempty"`
	EffectiveDateTime  *string                 `json:"effectiveDateTime,omitempty"`
	EffectivePeriod    *Period                 `json:"effectivePeriod,omitempty"`
	Issued             *string                 `json:"issued,omitempty"`
	Performer          []Reference             `json:"performer,omitempty"`
	ResultsInterpreter []Reference             `json:"resultsInterpreter,omitempty"`
	Specimen           []Reference             `json:"specimen,omitempty"`
	Result             []Reference             `json:"result,omitempty"`
	ImagingStudy       []Reference             `json:"imagingStudy,omitempty"`
	Media              []DiagnosticReportMedia `json:"media,omitempty"`
	Conclusion         *string                 `json:"conclusion,omitempty"`
	ConclusionCode     []CodeableConcept       `json:"conclusionCode,omitempty"`
	PresentedForm      []Attachment            `json:"presentedForm,omitempty"`
}

// A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
type DiagnosticReportMedia struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Link              Reference   `json:"link"`
}

// This function returns resource reference information
func (r DiagnosticReport) ResourceRef() (string, *string) {
	return "DiagnosticReport", r.ID
}

// This function returns resource reference information
func (r DiagnosticReport) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDiagnosticReport DiagnosticReport

// MarshalJSON marshals the given DiagnosticReport as JSON into a byte slice
func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDiagnosticReport
		ResourceType string `json:"resourceType"`
	}{
		OtherDiagnosticReport: OtherDiagnosticReport(r),
		ResourceType:          "DiagnosticReport",
	})
}

// UnmarshalDiagnosticReport unmarshals a DiagnosticReport.
func UnmarshalDiagnosticReport(b []byte) (DiagnosticReport, error) {
	var diagnosticReport DiagnosticReport
	if err := json.Unmarshal(b, &diagnosticReport); err != nil {
		return diagnosticReport, err
	}
	return diagnosticReport, nil
}
