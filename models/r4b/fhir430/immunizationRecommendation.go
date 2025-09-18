
package fhir430

import "encoding/json"
// ImmunizationRecommendation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
// A patient's point-in-time set of recommendations (i.e. forecasting) according to a published schedule with optional supporting justification.
type ImmunizationRecommendation struct {
	ID                *string                                    `json:"ID,omitempty"`
	Meta              *Meta                                      `json:"meta,omitempty"`
	ImplicitRules     *string                                    `json:"implicitRules,omitempty"`
	Language          *string                                    `json:"language,omitempty"`
	Text              *Narrative                                 `json:"text,omitempty"`
	Contained         []json.RawMessage                          `json:"contained,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `json:"identifier,omitempty"`
	Patient           Reference                                  `json:"patient"`
	Date              string                                     `json:"date"`
	Authority         *Reference                                 `json:"authority,omitempty"`
	Recommendation    []ImmunizationRecommendationRecommendation `json:"recommendation"`
}

// Vaccine administration recommendations.
type ImmunizationRecommendationRecommendation struct {
	ID                           *string                                                 `json:"ID,omitempty"`
	Extension                    []Extension                                             `json:"extension,omitempty"`
	ModifierExtension            []Extension                                             `json:"modifierExtension,omitempty"`
	VaccineCode                  []CodeableConcept                                       `json:"vaccineCode,omitempty"`
	TargetDisease                *CodeableConcept                                        `json:"targetDisease,omitempty"`
	ContraindicatedVaccineCode   []CodeableConcept                                       `json:"contraindicatedVaccineCode,omitempty"`
	ForecastStatus               CodeableConcept                                         `json:"forecastStatus"`
	ForecastReason               []CodeableConcept                                       `json:"forecastReason,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`
	Description                  *string                                                 `json:"description,omitempty"`
	Series                       *string                                                 `json:"series,omitempty"`
	DoseNumberPositiveInt        *int                                                    `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString             *string                                                 `json:"doseNumberString,omitempty"`
	SeriesDosesPositiveInt       *int                                                    `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString            *string                                                 `json:"seriesDosesString,omitempty"`
	SupportingImmunization       []Reference                                             `json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                             `json:"supportingPatientInformation,omitempty"`
}

// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
type ImmunizationRecommendationRecommendationDateCriterion struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Value             string          `json:"value"`
}

// This function returns resource reference information
func (r ImmunizationRecommendation) ResourceRef() (string, *string) {
	return "ImmunizationRecommendation", r.ID
}

// This function returns resource reference information
func (r ImmunizationRecommendation) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherImmunizationRecommendation ImmunizationRecommendation

// MarshalJSON marshals the given ImmunizationRecommendation as JSON into a byte slice
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationRecommendation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationRecommendation: OtherImmunizationRecommendation(r),
		ResourceType:                    "ImmunizationRecommendation",
	})
}

// UnmarshalImmunizationRecommendation unmarshals a ImmunizationRecommendation.
func UnmarshalImmunizationRecommendation(b []byte) (ImmunizationRecommendation, error) {
	var immunizationRecommendation ImmunizationRecommendation
	if err := json.Unmarshal(b, &immunizationRecommendation); err != nil {
		return immunizationRecommendation, err
	}
	return immunizationRecommendation, nil
}
