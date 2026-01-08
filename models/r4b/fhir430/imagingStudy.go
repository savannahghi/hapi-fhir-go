package fhir430

import "encoding/json"

// ImagingStudy is documented here http://hl7.org/fhir/StructureDefinition/ImagingStudy
// Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudy struct {
	ID                 *string              `json:"id,omitempty"`
	Meta               *Meta                `json:"meta,omitempty"`
	ImplicitRules      *string              `json:"implicitRules,omitempty"`
	Language           *string              `json:"language,omitempty"`
	Text               *Narrative           `json:"text,omitempty"`
	Contained          []json.RawMessage    `json:"contained,omitempty"`
	Extension          []Extension          `json:"extension,omitempty"`
	ModifierExtension  []Extension          `json:"modifierExtension,omitempty"`
	Identifier         []Identifier         `json:"identifier,omitempty"`
	Status             ImagingStudyStatus   `json:"status"`
	Modality           []Coding             `json:"modality,omitempty"`
	Subject            Reference            `json:"subject"`
	Encounter          *Reference           `json:"encounter,omitempty"`
	Started            *string              `json:"started,omitempty"`
	BasedOn            []Reference          `json:"basedOn,omitempty"`
	Referrer           *Reference           `json:"referrer,omitempty"`
	Interpreter        []Reference          `json:"interpreter,omitempty"`
	Endpoint           []Reference          `json:"endpoint,omitempty"`
	NumberOfSeries     *int                 `json:"numberOfSeries,omitempty"`
	NumberOfInstances  *int                 `json:"numberOfInstances,omitempty"`
	ProcedureReference *Reference           `json:"procedureReference,omitempty"`
	ProcedureCode      []CodeableConcept    `json:"procedureCode,omitempty"`
	Location           *Reference           `json:"location,omitempty"`
	ReasonCode         []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference    []Reference          `json:"reasonReference,omitempty"`
	Note               []Annotation         `json:"note,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Series             []ImagingStudySeries `json:"series,omitempty"`
}

// Each study has one or more series of images or other content.
type ImagingStudySeries struct {
	ID                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Uid               string                        `json:"uid"`
	Number            *int                          `json:"number,omitempty"`
	Modality          Coding                        `json:"modality"`
	Description       *string                       `json:"description,omitempty"`
	NumberOfInstances *int                          `json:"numberOfInstances,omitempty"`
	Endpoint          []Reference                   `json:"endpoint,omitempty"`
	BodySite          *Coding                       `json:"bodySite,omitempty"`
	Laterality        *Coding                       `json:"laterality,omitempty"`
	Specimen          []Reference                   `json:"specimen,omitempty"`
	Started           *string                       `json:"started,omitempty"`
	Performer         []ImagingStudySeriesPerformer `json:"performer,omitempty"`
	Instance          []ImagingStudySeriesInstance  `json:"instance,omitempty"`
}

// Indicates who or what performed the series and how they were involved.
// If the person who performed the series is not known, their Organization may be recorded. A patient, or related person, may be the performer, e.g. for patient-captured images.
type ImagingStudySeriesPerformer struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// A single SOP instance within the series, e.g. an image, or presentation state.
type ImagingStudySeriesInstance struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Uid               string      `json:"uid"`
	SopClass          Coding      `json:"sopClass"`
	Number            *int        `json:"number,omitempty"`
	Title             *string     `json:"title,omitempty"`
}

// This function returns resource reference information
func (r ImagingStudy) ResourceRef() (string, *string) {
	return "ImagingStudy", r.ID
}

// This function returns resource reference information
func (r ImagingStudy) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherImagingStudy ImagingStudy

// MarshalJSON marshals the given ImagingStudy as JSON into a byte slice
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingStudy: OtherImagingStudy(r),
		ResourceType:      "ImagingStudy",
	})
}

// UnmarshalImagingStudy unmarshals a ImagingStudy.
func UnmarshalImagingStudy(b []byte) (ImagingStudy, error) {
	var imagingStudy ImagingStudy
	if err := json.Unmarshal(b, &imagingStudy); err != nil {
		return imagingStudy, err
	}
	return imagingStudy, nil
}
