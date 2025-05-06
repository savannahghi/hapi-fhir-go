package models

import "encoding/json"

// CodeSystem is documented here http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
	ID                     *string                     `json:"id,omitempty"`
	Meta                   *Meta                       `json:"meta,omitempty"`
	ImplicitRules          *string                     `json:"implicitRules,omitempty"`
	Language               *string                     `json:"language,omitempty"`
	Text                   *Narrative                  `json:"text,omitempty"`
	Extension              []Extension                 `json:"extension,omitempty"`
	ModifierExtension      []Extension                 `json:"modifierExtension,omitempty"`
	URL                    *string                     `json:"url,omitempty"`
	Identifier             []Identifier                `json:"identifier,omitempty"`
	Version                *string                     `json:"version,omitempty"`
	VersionAlgorithmString *string                     `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                     `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                     `json:"name,omitempty"`
	Title                  *string                     `json:"title,omitempty"`
	Status                 PublicationStatus           `json:"status,omitempty"`
	Experimental           *bool                       `json:"experimental,omitempty"`
	Date                   *string                     `json:"date,omitempty"`
	Publisher              *string                     `json:"publisher,omitempty"`
	Contact                []ContactDetail             `json:"contact,omitempty"`
	Description            *string                     `json:"description,omitempty"`
	UseContext             []UsageContext              `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept           `json:"jurisdiction,omitempty"`
	Purpose                *string                     `json:"purpose,omitempty"`
	Copyright              *string                     `json:"copyright,omitempty"`
	CopyrightLabel         *string                     `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                     `json:"approvalDate,omitempty"`
	LastReviewDate         *string                     `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                     `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept           `json:"topic,omitempty"`
	Author                 []ContactDetail             `json:"author,omitempty"`
	Editor                 []ContactDetail             `json:"editor,omitempty"`
	Reviewer               []ContactDetail             `json:"reviewer,omitempty"`
	Endorser               []ContactDetail             `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact           `json:"relatedArtifact,omitempty"`
	CaseSensitive          *bool                       `json:"caseSensitive,omitempty"`
	ValueSet               *string                     `json:"valueSet,omitempty"`
	HierarchyMeaning       *CodeSystemHierarchyMeaning `json:"hierarchyMeaning,omitempty"`
	Compositional          *bool                       `json:"compositional,omitempty"`
	VersionNeeded          *bool                       `json:"versionNeeded,omitempty"`
	Content                CodeSystemContentMode       `json:"content,omitempty"`
	Supplements            *string                     `json:"supplements,omitempty"`
	Count                  *int                        `json:"count,omitempty"`
	Filter                 []CodeSystemFilter          `json:"filter,omitempty"`
	Property               []CodeSystemProperty        `json:"property,omitempty"`
	Concept                []CodeSystemConcept         `json:"concept,omitempty"`
}

type CodeSystemFilter struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              string           `json:"code"`
	Description       *string          `json:"description,omitempty"`
	Operator          []FilterOperator `json:"operator"`
	Value             string           `json:"value"`
}

type CodeSystemProperty struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Code              string       `json:"code"`
	URI               *string      `json:"uri,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Type              PropertyType `json:"type"`
}

type CodeSystemConcept struct {
	ID                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              string                         `json:"code"`
	Display           *string                        `json:"display,omitempty"`
	Definition        *string                        `json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `json:"property,omitempty"`
	Concept           []CodeSystemConcept            `json:"concept,omitempty"`
}

type CodeSystemConceptDesignation struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	AdditionalUse     []Coding    `json:"additionalUse,omitempty"`
	Value             string      `json:"value"`
}

type CodeSystemConceptProperty struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	ValueCode         string      `json:"valueCode"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueString       string      `json:"valueString"`
	ValueInteger      int         `json:"valueInteger"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueDecimal      json.Number `json:"valueDecimal"`
}

type OtherCodeSystem CodeSystem

// MarshalJSON marshals the given CodeSystem as JSON into a byte slice.
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

// UnmarshalCodeSystem unmarshals a CodeSystem.
func UnmarshalCodeSystem(b []byte) (CodeSystem, error) {
	var codeSystem CodeSystem
	if err := json.Unmarshal(b, &codeSystem); err != nil {
		return codeSystem, err
	}

	return codeSystem, nil
}
