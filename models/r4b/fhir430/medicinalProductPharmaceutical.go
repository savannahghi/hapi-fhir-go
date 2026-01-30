package fhir430

import "encoding/json"

// MedicinalProductPharmaceutical is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductPharmaceutical
// A pharmaceutical product described in terms of its composition and dose form.
type MedicinalProductPharmaceutical struct {
	ID                    *string                                               `json:"id,omitempty"`
	Meta                  *Meta                                                 `json:"meta,omitempty"`
	ImplicitRules         *string                                               `json:"implicitRules,omitempty"`
	Language              *string                                               `json:"language,omitempty"`
	Text                  *Narrative                                            `json:"text,omitempty"`
	Contained             []json.RawMessage                                     `json:"contained,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                          `json:"identifier,omitempty"`
	AdministrableDoseForm CodeableConcept                                       `json:"administrableDoseForm"`
	UnitOfPresentation    *CodeableConcept                                      `json:"unitOfPresentation,omitempty"`
	Ingredient            []Reference                                           `json:"ingredient,omitempty"`
	Device                []Reference                                           `json:"device,omitempty"`
	Characteristics       []MedicinalProductPharmaceuticalCharacteristics       `json:"characteristics,omitempty"`
	RouteOfAdministration []MedicinalProductPharmaceuticalRouteOfAdministration `json:"routeOfAdministration"`
}

// Characteristics e.g. a products onset of action.
type MedicinalProductPharmaceuticalCharacteristics struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `json:"code"`
	Status            *CodeableConcept `json:"status,omitempty"`
}

// The path by which the pharmaceutical product is taken into or makes contact with the body.
type MedicinalProductPharmaceuticalRouteOfAdministration struct {
	ID                        *string                                                            `json:"id,omitempty"`
	Extension                 []Extension                                                        `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `json:"code"`
	FirstDose                 *Quantity                                                          `json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `json:"maxDosePerTreatmentPeriod,omitempty"`
	MaxTreatmentPeriod        *Duration                                                          `json:"maxTreatmentPeriod,omitempty"`
	TargetSpecies             []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

// A species for which this route applies.
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies struct {
	ID                *string                                                                            `json:"id,omitempty"`
	Extension         []Extension                                                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `json:"code"`
	WithdrawalPeriod  []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// A species specific time during which consumption of animal product is not appropriate.
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	ID                    *string         `json:"id,omitempty"`
	Extension             []Extension     `json:"extension,omitempty"`
	ModifierExtension     []Extension     `json:"modifierExtension,omitempty"`
	Tissue                CodeableConcept `json:"tissue"`
	Value                 Quantity        `json:"value"`
	SupportingInformation *string         `json:"supportingInformation,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductPharmaceutical) ResourceRef() (string, *string) {
	return "MedicinalProductPharmaceutical", r.ID
}

// This function returns resource reference information
func (r MedicinalProductPharmaceutical) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductPharmaceutical MedicinalProductPharmaceutical

// MarshalJSON marshals the given MedicinalProductPharmaceutical as JSON into a byte slice
func (r MedicinalProductPharmaceutical) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductPharmaceutical
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductPharmaceutical: OtherMedicinalProductPharmaceutical(r),
		ResourceType:                        "MedicinalProductPharmaceutical",
	})
}

// UnmarshalMedicinalProductPharmaceutical unmarshals a MedicinalProductPharmaceutical.
func UnmarshalMedicinalProductPharmaceutical(b []byte) (MedicinalProductPharmaceutical, error) {
	var medicinalProductPharmaceutical MedicinalProductPharmaceutical
	if err := json.Unmarshal(b, &medicinalProductPharmaceutical); err != nil {
		return medicinalProductPharmaceutical, err
	}
	return medicinalProductPharmaceutical, nil
}
