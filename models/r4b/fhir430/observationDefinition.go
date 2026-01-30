package fhir430

import "encoding/json"

// ObservationDefinition is documented here http://hl7.org/fhir/StructureDefinition/ObservationDefinition
// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
type ObservationDefinition struct {
	ID                     *string                                   `json:"id,omitempty"`
	Meta                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules          *string                                   `json:"implicitRules,omitempty"`
	Language               *string                                   `json:"language,omitempty"`
	Text                   *Narrative                                `json:"text,omitempty"`
	Contained              []json.RawMessage                         `json:"contained,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	Category               []CodeableConcept                         `json:"category,omitempty"`
	Code                   CodeableConcept                           `json:"code"`
	Identifier             []Identifier                              `json:"identifier,omitempty"`
	PermittedDataType      []ObservationDataType                     `json:"permittedDataType,omitempty"`
	MultipleResultsAllowed *bool                                     `json:"multipleResultsAllowed,omitempty"`
	Method                 *CodeableConcept                          `json:"method,omitempty"`
	PreferredReportName    *string                                   `json:"preferredReportName,omitempty"`
	QuantitativeDetails    *ObservationDefinitionQuantitativeDetails `json:"quantitativeDetails,omitempty"`
	QualifiedInterval      []ObservationDefinitionQualifiedInterval  `json:"qualifiedInterval,omitempty"`
	ValidCodedValueSet     *Reference                                `json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet    *Reference                                `json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet  *Reference                                `json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet  *Reference                                `json:"criticalCodedValueSet,omitempty"`
}

// Characteristics for quantitative results of this observation.
type ObservationDefinitionQuantitativeDetails struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	CustomaryUnit     *CodeableConcept `json:"customaryUnit,omitempty"`
	Unit              *CodeableConcept `json:"unit,omitempty"`
	ConversionFactor  *json.Number     `json:"conversionFactor,omitempty"`
	DecimalPrecision  *int             `json:"decimalPrecision,omitempty"`
}

// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
type ObservationDefinitionQualifiedInterval struct {
	ID                *string                   `json:"id,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Category          *ObservationRangeCategory `json:"category,omitempty"`
	Range             *Range                    `json:"range,omitempty"`
	Context           *CodeableConcept          `json:"context,omitempty"`
	AppliesTo         []CodeableConcept         `json:"appliesTo,omitempty"`
	Gender            *AdministrativeGender     `json:"gender,omitempty"`
	Age               *Range                    `json:"age,omitempty"`
	GestationalAge    *Range                    `json:"gestationalAge,omitempty"`
	Condition         *string                   `json:"condition,omitempty"`
}

// This function returns resource reference information
func (r ObservationDefinition) ResourceRef() (string, *string) {
	return "ObservationDefinition", r.ID
}

// This function returns resource reference information
func (r ObservationDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherObservationDefinition ObservationDefinition

// MarshalJSON marshals the given ObservationDefinition as JSON into a byte slice
func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherObservationDefinition: OtherObservationDefinition(r),
		ResourceType:               "ObservationDefinition",
	})
}

// UnmarshalObservationDefinition unmarshals a ObservationDefinition.
func UnmarshalObservationDefinition(b []byte) (ObservationDefinition, error) {
	var observationDefinition ObservationDefinition
	if err := json.Unmarshal(b, &observationDefinition); err != nil {
		return observationDefinition, err
	}
	return observationDefinition, nil
}
