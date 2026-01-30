package fhir430

import "encoding/json"

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
// Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient struct {
	ID                   *string                `json:"id,omitempty"`
	Meta                 *Meta                  `json:"meta,omitempty"`
	ImplicitRules        *string                `json:"implicitRules,omitempty"`
	Language             *string                `json:"language,omitempty"`
	Text                 *Narrative             `json:"text,omitempty"`
	Contained            []json.RawMessage      `json:"contained,omitempty"`
	Extension            []Extension            `json:"extension,omitempty"`
	ModifierExtension    []Extension            `json:"modifierExtension,omitempty"`
	Identifier           []Identifier           `json:"identifier,omitempty"`
	Active               *bool                  `json:"active,omitempty"`
	Name                 []HumanName            `json:"name,omitempty"`
	Telecom              []ContactPoint         `json:"telecom,omitempty"`
	Gender               *AdministrativeGender  `json:"gender,omitempty"`
	BirthDate            *string                `json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                  `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *string                `json:"deceasedDateTime,omitempty"`
	Address              []Address              `json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                  `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int                   `json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment           `json:"photo,omitempty"`
	Contact              []PatientContact       `json:"contact,omitempty"`
	Communication        []PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner  []Reference            `json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference             `json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `json:"link,omitempty"`
}

// A contact party (e.g. guardian, partner, friend) for the patient.
// Contact covers all kinds of contact parties: family members, business contacts, guardians, caregivers. Not applicable to register pedigree and family ties beyond use of having contact.
type PatientContact struct {
	ID                *string               `json:"id,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Relationship      []CodeableConcept     `json:"relationship,omitempty"`
	Name              *HumanName            `json:"name,omitempty"`
	Telecom           []ContactPoint        `json:"telecom,omitempty"`
	Address           *Address              `json:"address,omitempty"`
	Gender            *AdministrativeGender `json:"gender,omitempty"`
	Organization      *Reference            `json:"organization,omitempty"`
	Period            *Period               `json:"period,omitempty"`
}

// A language which may be used to communicate with the patient about his or her health.
// If no language is specified, this *implies* that the default local language is spoken.  If you need to convey proficiency for multiple modes, then you need multiple Patient.Communication associations.   For animals, language is not a relevant field, and should be absent from the instance. If the Patient does not speak the default local language, then the Interpreter Required Standard can be used to explicitly declare that an interpreter is required.
type PatientCommunication struct {
	ID                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Language          CodeableConcept `json:"language"`
	Preferred         *bool           `json:"preferred,omitempty"`
}

// Link to another patient resource that concerns the same actual patient.
// There is no assumption that linked patient records have mutual links.
type PatientLink struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Other             Reference   `json:"other"`
	Type              LinkType    `json:"type"`
}

// This function returns resource reference information
func (r Patient) ResourceRef() (string, *string) {
	return "Patient", r.ID
}

// This function returns resource reference information
func (r Patient) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherPatient Patient

// MarshalJSON marshals the given Patient as JSON into a byte slice
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

// UnmarshalPatient unmarshals a Patient.
func UnmarshalPatient(b []byte) (Patient, error) {
	var patient Patient
	if err := json.Unmarshal(b, &patient); err != nil {
		return patient, err
	}
	return patient, nil
}
