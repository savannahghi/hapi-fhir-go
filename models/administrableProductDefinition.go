package models

import "encoding/json"

// AdministrableProductDefinition is documented here
// http://hl7.org/fhir/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinition struct {
	ID                    *string                                               `json:"id,omitempty"`
	Meta                  *Meta                                                 `json:"meta,omitempty"`
	ImplicitRules         *string                                               `json:"implicitRules,omitempty"`
	Language              *string                                               `json:"language,omitempty"`
	Text                  *Narrative                                            `json:"text,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                          `json:"identifier,omitempty"`
	Status                PublicationStatus                                     `json:"status"`
	FormOf                []Reference                                           `json:"formOf,omitempty"`
	AdministrableDoseForm *CodeableConcept                                      `json:"administrableDoseForm,omitempty"`
	UnitOfPresentation    *CodeableConcept                                      `json:"unitOfPresentation,omitempty"`
	ProducedFrom          []Reference                                           `json:"producedFrom,omitempty"`
	Ingredient            []CodeableConcept                                     `json:"ingredient,omitempty"`
	Device                *Reference                                            `json:"device,omitempty"`
	Description           *string                                               `json:"description,omitempty"`
	Property              []AdministrableProductDefinitionProperty              `json:"property,omitempty"`
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration"`
}

type AdministrableProductDefinitionProperty struct {
	ID                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
	Status               *CodeableConcept `json:"status,omitempty"`
}

//nolint:lll
type AdministrableProductDefinitionRouteOfAdministration struct {
	ID                        *string                                                            `json:"id,omitempty"`
	Extension                 []Extension                                                        `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `json:"code"`
	FirstDose                 *Quantity                                                          `json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `json:"maxDosePerTreatmentPeriod,omitempty"`
	TargetSpecies             []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

//nolint:lll
type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	ID                *string                                                                            `json:"id,omitempty"`
	Extension         []Extension                                                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `json:"code"`
	WithdrawalPeriod  []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	ID                    *string         `json:"id,omitempty"`
	Extension             []Extension     `json:"extension,omitempty"`
	ModifierExtension     []Extension     `json:"modifierExtension,omitempty"`
	Tissue                CodeableConcept `json:"tissue"`
	Value                 Quantity        `json:"value"`
	SupportingInformation *string         `json:"supportingInformation,omitempty"`
}

type OtherAdministrableProductDefinition AdministrableProductDefinition

// MarshalJSON marshals the given AdministrableProductDefinition as JSON into a byte slice.
func (r AdministrableProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdministrableProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherAdministrableProductDefinition: OtherAdministrableProductDefinition(r),
		ResourceType:                        "AdministrableProductDefinition",
	})
}

// UnmarshalAdministrableProductDefinition unmarshals a AdministrableProductDefinition.
func UnmarshalAdministrableProductDefinition(b []byte) (AdministrableProductDefinition, error) {
	var administrableProductDefinition AdministrableProductDefinition
	if err := json.Unmarshal(b, &administrableProductDefinition); err != nil {
		return administrableProductDefinition, err
	}

	return administrableProductDefinition, nil
}
