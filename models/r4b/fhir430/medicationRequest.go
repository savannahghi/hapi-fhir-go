
package fhir430

import "encoding/json"
// MedicationRequest is documented here http://hl7.org/fhir/StructureDefinition/MedicationRequest
// An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest struct {
	ID                        *string                           `json:"ID,omitempty"`
	Meta                      *Meta                             `json:"meta,omitempty"`
	ImplicitRules             *string                           `json:"implicitRules,omitempty"`
	Language                  *string                           `json:"language,omitempty"`
	Text                      *Narrative                        `json:"text,omitempty"`
	Contained                 []json.RawMessage                 `json:"contained,omitempty"`
	Extension                 []Extension                       `json:"extension,omitempty"`
	ModifierExtension         []Extension                       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                      `json:"identifier,omitempty"`
	Status                    string                            `json:"status"`
	StatusReason              *CodeableConcept                  `json:"statusReason,omitempty"`
	Intent                    string                            `json:"intent"`
	Category                  []CodeableConcept                 `json:"category,omitempty"`
	Priority                  *RequestPriority                  `json:"priority,omitempty"`
	DoNotPerform              *bool                             `json:"doNotPerform,omitempty"`
	ReportedBoolean           *bool                             `json:"reportedBoolean,omitempty"`
	ReportedReference         *Reference                        `json:"reportedReference,omitempty"`
	MedicationCodeableConcept CodeableConcept                   `json:"medicationCodeableConcept"`
	MedicationReference       Reference                         `json:"medicationReference"`
	Subject                   Reference                         `json:"subject"`
	Encounter                 *Reference                        `json:"encounter,omitempty"`
	SupportingInformation     []Reference                       `json:"supportingInformation,omitempty"`
	AuthoredOn                *string                           `json:"authoredOn,omitempty"`
	Requester                 *Reference                        `json:"requester,omitempty"`
	Performer                 *Reference                        `json:"performer,omitempty"`
	PerformerType             *CodeableConcept                  `json:"performerType,omitempty"`
	Recorder                  *Reference                        `json:"recorder,omitempty"`
	ReasonCode                []CodeableConcept                 `json:"reasonCode,omitempty"`
	ReasonReference           []Reference                       `json:"reasonReference,omitempty"`
	InstantiatesCanonical     []string                          `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri           []string                          `json:"instantiatesUri,omitempty"`
	BasedOn                   []Reference                       `json:"basedOn,omitempty"`
	GroupIdentifier           *Identifier                       `json:"groupIdentifier,omitempty"`
	CourseOfTherapyType       *CodeableConcept                  `json:"courseOfTherapyType,omitempty"`
	Insurance                 []Reference                       `json:"insurance,omitempty"`
	Note                      []Annotation                      `json:"note,omitempty"`
	DosageInstruction         []Dosage                          `json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitution    `json:"substitution,omitempty"`
	PriorPrescription         *Reference                        `json:"priorPrescription,omitempty"`
	DetectedIssue             []Reference                       `json:"detectedIssue,omitempty"`
	EventHistory              []Reference                       `json:"eventHistory,omitempty"`
}

// Indicates the specific details for the dispense or medication supply part of a medication request (also known as a Medication Prescription or Medication Order).  Note that this information is not always sent with the order.  There may be in some settings (e.g. hospitals) institutional or system support for completing the dispense details in the pharmacy department.
type MedicationRequestDispenseRequest struct {
	ID                     *string                                      `json:"ID,omitempty"`
	Extension              []Extension                                  `json:"extension,omitempty"`
	ModifierExtension      []Extension                                  `json:"modifierExtension,omitempty"`
	InitialFill            *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`
	DispenseInterval       *Duration                                    `json:"dispenseInterval,omitempty"`
	ValidityPeriod         *Period                                      `json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed *int                                         `json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               *Quantity                                    `json:"quantity,omitempty"`
	ExpectedSupplyDuration *Duration                                    `json:"expectedSupplyDuration,omitempty"`
	Performer              *Reference                                   `json:"performer,omitempty"`
}

// Indicates the quantity or duration for the first dispense of the medication.
// If populating this element, either the quantity or the duration must be included.
type MedicationRequestDispenseRequestInitialFill struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	Duration          *Duration   `json:"duration,omitempty"`
}

// Indicates whether or not substitution can or should be part of the dispense. In some cases, substitution must happen, in other cases substitution must not happen. This block explains the prescriber's intent. If nothing is specified substitution may be done.
type MedicationRequestSubstitution struct {
	ID                     *string          `json:"ID,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	AllowedBoolean         bool             `json:"allowedBoolean"`
	AllowedCodeableConcept CodeableConcept  `json:"allowedCodeableConcept"`
	Reason                 *CodeableConcept `json:"reason,omitempty"`
}

// This function returns resource reference information
func (r MedicationRequest) ResourceRef() (string, *string) {
	return "MedicationRequest", r.ID
}

// This function returns resource reference information
func (r MedicationRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicationRequest MedicationRequest

// MarshalJSON marshals the given MedicationRequest as JSON into a byte slice
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
}

// UnmarshalMedicationRequest unmarshals a MedicationRequest.
func UnmarshalMedicationRequest(b []byte) (MedicationRequest, error) {
	var medicationRequest MedicationRequest
	if err := json.Unmarshal(b, &medicationRequest); err != nil {
		return medicationRequest, err
	}
	return medicationRequest, nil
}
