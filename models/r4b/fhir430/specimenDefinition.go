
package fhir430

import "encoding/json"
// SpecimenDefinition is documented here http://hl7.org/fhir/StructureDefinition/SpecimenDefinition
// A kind of specimen with associated set of requirements.
type SpecimenDefinition struct {
	ID                 *string                        `json:"ID,omitempty"`
	Meta               *Meta                          `json:"meta,omitempty"`
	ImplicitRules      *string                        `json:"implicitRules,omitempty"`
	Language           *string                        `json:"language,omitempty"`
	Text               *Narrative                     `json:"text,omitempty"`
	Contained          []json.RawMessage              `json:"contained,omitempty"`
	Extension          []Extension                    `json:"extension,omitempty"`
	ModifierExtension  []Extension                    `json:"modifierExtension,omitempty"`
	Identifier         *Identifier                    `json:"identifier,omitempty"`
	TypeCollected      *CodeableConcept               `json:"typeCollected,omitempty"`
	PatientPreparation []CodeableConcept              `json:"patientPreparation,omitempty"`
	TimeAspect         *string                        `json:"timeAspect,omitempty"`
	Collection         []CodeableConcept              `json:"collection,omitempty"`
	TypeTested         []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

// Specimen conditioned in a container as expected by the testing laboratory.
type SpecimenDefinitionTypeTested struct {
	ID                 *string                                `json:"ID,omitempty"`
	Extension          []Extension                            `json:"extension,omitempty"`
	ModifierExtension  []Extension                            `json:"modifierExtension,omitempty"`
	IsDerived          *bool                                  `json:"isDerived,omitempty"`
	Type               *CodeableConcept                       `json:"type,omitempty"`
	Preference         SpecimenContainedPreference            `json:"preference"`
	Container          *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`
	Requirement        *string                                `json:"requirement,omitempty"`
	RetentionTime      *Duration                              `json:"retentionTime,omitempty"`
	RejectionCriterion []CodeableConcept                      `json:"rejectionCriterion,omitempty"`
	Handling           []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`
}

// The specimen's container.
type SpecimenDefinitionTypeTestedContainer struct {
	ID                    *string                                         `json:"ID,omitempty"`
	Extension             []Extension                                     `json:"extension,omitempty"`
	ModifierExtension     []Extension                                     `json:"modifierExtension,omitempty"`
	Material              *CodeableConcept                                `json:"material,omitempty"`
	Type                  *CodeableConcept                                `json:"type,omitempty"`
	Cap                   *CodeableConcept                                `json:"cap,omitempty"`
	Description           *string                                         `json:"description,omitempty"`
	Capacity              *Quantity                                       `json:"capacity,omitempty"`
	MinimumVolumeQuantity *Quantity                                       `json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string                                         `json:"minimumVolumeString,omitempty"`
	Additive              []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`
	Preparation           *string                                         `json:"preparation,omitempty"`
}

// Substance introduced in the kind of container to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	ID                      *string         `json:"ID,omitempty"`
	Extension               []Extension     `json:"extension,omitempty"`
	ModifierExtension       []Extension     `json:"modifierExtension,omitempty"`
	AdditiveCodeableConcept CodeableConcept `json:"additiveCodeableConcept"`
	AdditiveReference       Reference       `json:"additiveReference"`
}

// Set of instructions for preservation/transport of the specimen at a defined temperature interval, prior the testing process.
type SpecimenDefinitionTypeTestedHandling struct {
	ID                   *string          `json:"ID,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	TemperatureQualifier *CodeableConcept `json:"temperatureQualifier,omitempty"`
	TemperatureRange     *Range           `json:"temperatureRange,omitempty"`
	MaxDuration          *Duration        `json:"maxDuration,omitempty"`
	Instruction          *string          `json:"instruction,omitempty"`
}

// This function returns resource reference information
func (r SpecimenDefinition) ResourceRef() (string, *string) {
	return "SpecimenDefinition", r.ID
}

// This function returns resource reference information
func (r SpecimenDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSpecimenDefinition SpecimenDefinition

// MarshalJSON marshals the given SpecimenDefinition as JSON into a byte slice
func (r SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimenDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimenDefinition: OtherSpecimenDefinition(r),
		ResourceType:            "SpecimenDefinition",
	})
}

// UnmarshalSpecimenDefinition unmarshals a SpecimenDefinition.
func UnmarshalSpecimenDefinition(b []byte) (SpecimenDefinition, error) {
	var specimenDefinition SpecimenDefinition
	if err := json.Unmarshal(b, &specimenDefinition); err != nil {
		return specimenDefinition, err
	}
	return specimenDefinition, nil
}
