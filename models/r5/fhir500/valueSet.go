package fhir500

import "encoding/json"

// ValueSet is documented here http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSet struct {
	ID                     *string            `json:"id,omitempty"`
	Meta                   *Meta              `json:"meta,omitempty"`
	ImplicitRules          *string            `json:"implicitRules,omitempty"`
	Language               *string            `json:"language,omitempty"`
	Text                   *Narrative         `json:"text,omitempty"`
	Extension              []Extension        `json:"extension,omitempty"`
	ModifierExtension      []Extension        `json:"modifierExtension,omitempty"`
	URL                    *string            `json:"url,omitempty"`
	Identifier             []Identifier       `json:"identifier,omitempty"`
	Version                *string            `json:"version,omitempty"`
	VersionAlgorithmString *string            `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding            `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string            `json:"name,omitempty"`
	Title                  *string            `json:"title,omitempty"`
	Status                 PublicationStatus  `json:"status"`
	Experimental           *bool              `json:"experimental,omitempty"`
	Date                   *string            `json:"date,omitempty"`
	Publisher              *string            `json:"publisher,omitempty"`
	Contact                []ContactDetail    `json:"contact,omitempty"`
	Description            *string            `json:"description,omitempty"`
	UseContext             []UsageContext     `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept  `json:"jurisdiction,omitempty"`
	Immutable              *bool              `json:"immutable,omitempty"`
	Purpose                *string            `json:"purpose,omitempty"`
	Copyright              *string            `json:"copyright,omitempty"`
	CopyrightLabel         *string            `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string            `json:"approvalDate,omitempty"`
	LastReviewDate         *string            `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period            `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept  `json:"topic,omitempty"`
	Author                 []ContactDetail    `json:"author,omitempty"`
	Editor                 []ContactDetail    `json:"editor,omitempty"`
	Reviewer               []ContactDetail    `json:"reviewer,omitempty"`
	Endorser               []ContactDetail    `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact  `json:"relatedArtifact,omitempty"`
	Compose                *ValueSetCompose   `json:"compose,omitempty"`
	Expansion              *ValueSetExpansion `json:"expansion,omitempty"`
	Scope                  *ValueSetScope     `json:"scope,omitempty"`
}

type ValueSetCompose struct {
	ID                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *string                  `json:"lockedDate,omitempty"`
	Inactive          *bool                    `json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `json:"include"`
	Exclude           []ValueSetComposeInclude `json:"exclude,omitempty"`
	Property          []string                 `json:"property,omitempty"`
}

type ValueSetComposeInclude struct {
	ID                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	System            *string                         `json:"system,omitempty"`
	Version           *string                         `json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet          []string                        `json:"valueSet,omitempty"`
	Copyright         *string                         `json:"copyright,omitempty"`
}

type ValueSetComposeIncludeConcept struct {
	ID                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Code              string                                     `json:"code"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

type ValueSetComposeIncludeConceptDesignation struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	AdditionalUse     []Coding    `json:"additionalUse,omitempty"`
	Value             string      `json:"value"`
}

type ValueSetComposeIncludeFilter struct {
	ID                *string        `json:"id,omitempty"`
	Extension         []Extension    `json:"extension,omitempty"`
	ModifierExtension []Extension    `json:"modifierExtension,omitempty"`
	Property          string         `json:"property"`
	Op                FilterOperator `json:"op"`
	Value             string         `json:"value"`
}

type ValueSetExpansion struct {
	ID                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	Next              *string                      `json:"next,omitempty"`
	Timestamp         string                       `json:"timestamp"`
	Total             *int                         `json:"total,omitempty"`
	Offset            *int                         `json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Property          []ValueSetExpansionProperty  `json:"property,omitempty"`
	Contains          []ValueSetExpansionContains  `json:"contains,omitempty"`
}

type ValueSetExpansionParameter struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Name              string       `json:"name"`
	ValueString       *string      `json:"valueString,omitempty"`
	ValueBoolean      *bool        `json:"valueBoolean,omitempty"`
	ValueInteger      *int         `json:"valueInteger,omitempty"`
	ValueDecimal      *json.Number `json:"valueDecimal,omitempty"`
	ValueURI          *string      `json:"valueUri,omitempty"`
	ValueCode         *string      `json:"valueCode,omitempty"`
	ValueDateTime     *string      `json:"valueDateTime,omitempty"`
}

type ValueSetExpansionProperty struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	URI               *string     `json:"uri,omitempty"`
}

type ValueSetExpansionContains struct {
	ID                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	System            *string                                    `json:"system,omitempty"`
	Abstract          *bool                                      `json:"abstract,omitempty"`
	Inactive          *bool                                      `json:"inactive,omitempty"`
	Version           *string                                    `json:"version,omitempty"`
	Code              *string                                    `json:"code,omitempty"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
	Property          []ValueSetExpansionContainsProperty        `json:"property,omitempty"`
	Contains          []ValueSetExpansionContains                `json:"contains,omitempty"`
}

type ValueSetExpansionContainsProperty struct {
	ID                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Code              string                                         `json:"code"`
	ValueCode         string                                         `json:"valueCode"`
	ValueCoding       Coding                                         `json:"valueCoding"`
	ValueString       string                                         `json:"valueString"`
	ValueInteger      int                                            `json:"valueInteger"`
	ValueBoolean      bool                                           `json:"valueBoolean"`
	ValueDateTime     string                                         `json:"valueDateTime"`
	ValueDecimal      json.Number                                    `json:"valueDecimal"`
	SubProperty       []ValueSetExpansionContainsPropertySubProperty `json:"subProperty,omitempty"`
}

type ValueSetExpansionContainsPropertySubProperty struct {
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

type ValueSetScope struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	InclusionCriteria *string     `json:"inclusionCriteria,omitempty"`
	ExclusionCriteria *string     `json:"exclusionCriteria,omitempty"`
}

type OtherValueSet ValueSet

// MarshalJSON marshals the given ValueSet as JSON into a byte slice.
func (r ValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherValueSet
		ResourceType string `json:"resourceType"`
	}{
		OtherValueSet: OtherValueSet(r),
		ResourceType:  "ValueSet",
	})
}

// UnmarshalValueSet unmarshals a ValueSet.
func UnmarshalValueSet(b []byte) (ValueSet, error) {
	var valueSet ValueSet
	if err := json.Unmarshal(b, &valueSet); err != nil {
		return valueSet, err
	}

	return valueSet, nil
}
