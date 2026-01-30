package fhir430

import "encoding/json"

// ValueSet is documented here http://hl7.org/fhir/StructureDefinition/ValueSet
// A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context. Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
type ValueSet struct {
	ID                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []json.RawMessage  `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Url               *string            `json:"url,omitempty"`
	Identifier        []Identifier       `json:"identifier,omitempty"`
	Version           *string            `json:"version,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Title             *string            `json:"title,omitempty"`
	Status            PublicationStatus  `json:"status"`
	Experimental      *bool              `json:"experimental,omitempty"`
	Date              *string            `json:"date,omitempty"`
	Publisher         *string            `json:"publisher,omitempty"`
	Contact           []ContactDetail    `json:"contact,omitempty"`
	Description       *string            `json:"description,omitempty"`
	UseContext        []UsageContext     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept  `json:"jurisdiction,omitempty"`
	Immutable         *bool              `json:"immutable,omitempty"`
	Purpose           *string            `json:"purpose,omitempty"`
	Compose           *ValueSetCompose   `json:"compose,omitempty"`
	Expansion         *ValueSetExpansion `json:"expansion,omitempty"`
}

// A set of criteria that define the contents of the value set by including or excluding codes selected from the specified code system(s) that the value set draws from. This is also known as the Content Logical Definition (CLD).
type ValueSetCompose struct {
	ID                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	LockedDate        *string                  `json:"lockedDate,omitempty"`
	Inactive          *bool                    `json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `json:"include"`
	Exclude           []ValueSetComposeInclude `json:"exclude,omitempty"`
}

// Include one or more codes from a code system or other value set(s).
// All the conditions in an include must be true. If a system is listed, all the codes from the system are listed. If one or more filters are listed, all of the filters must apply. If one or more value sets are listed, the codes must be in all the value sets. E.g. each include is 'include all the codes that meet all these conditions'.
type ValueSetComposeInclude struct {
	ID                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	System            *string                         `json:"system,omitempty"`
	Version           *string                         `json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet          []string                        `json:"valueSet,omitempty"`
}

// Specifies a concept to be included or excluded.
// The list of concepts is considered ordered, though the order might not have any particular significance. Typically, the order of an expansion follows that defined in the compose element.
type ValueSetComposeIncludeConcept struct {
	ID                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Code              string                                     `json:"code"`
	Display           *string                                    `json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

// Additional representations for this concept when used in this value set - other languages, aliases, specialized purposes, used for particular purposes, etc.
// Concepts have both a ```display``` and an array of ```designation```. The display is equivalent to a special designation with an implied ```designation.use``` of "primary code" and a language equal to the [Resource Language](resource.html#language).
type ValueSetComposeIncludeConceptDesignation struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// Select concepts by specify a matching criterion based on the properties (including relationships) defined by the system, or on filters defined by the system. If multiple filters are specified, they SHALL all be true.
// Selecting codes by specifying filters based on properties is only possible where the underlying code system defines appropriate properties. Note that in some cases, the underlying code system defines the logical concepts but not the literal codes for the concepts. In such cases, the literal definitions may be provided by a third party.
type ValueSetComposeIncludeFilter struct {
	ID                *string        `json:"id,omitempty"`
	Extension         []Extension    `json:"extension,omitempty"`
	ModifierExtension []Extension    `json:"modifierExtension,omitempty"`
	Property          string         `json:"property"`
	Op                FilterOperator `json:"op"`
	Value             string         `json:"value"`
}

// A value set can also be "expanded", where the value set is turned into a simple collection of enumerated codes. This element holds the expansion, if it has been performed.
/*
Expansion is performed to produce a collection of codes that are ready to use for data entry or validation. Value set expansions are always considered to be stateless - they are a record of the set of codes in the value set at a point in time under a given set of conditions, and are not subject to ongoing maintenance.

Expansion.parameter is  a simplified list of parameters - a subset of the features of the [Parameters](parameters.html) resource.
*/
type ValueSetExpansion struct {
	ID                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        *string                      `json:"identifier,omitempty"`
	Timestamp         string                       `json:"timestamp"`
	Total             *int                         `json:"total,omitempty"`
	Offset            *int                         `json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Contains          []ValueSetExpansionContains  `json:"contains,omitempty"`
}

// A parameter that controlled the expansion process. These parameters may be used by users of expanded value sets to check whether the expansion is suitable for a particular purpose, or to pick the correct expansion.
// The server decides which parameters to include here, but at a minimum, the list SHOULD include all of the parameters that affect the $expand operation. If the expansion will be persisted all of these parameters SHALL be included. If the codeSystem on the server has a specified version then this version SHALL be provided as a parameter in the expansion (note that not all code systems have a version).
type ValueSetExpansionParameter struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Name              string       `json:"name"`
	ValueString       *string      `json:"valueString,omitempty"`
	ValueBoolean      *bool        `json:"valueBoolean,omitempty"`
	ValueInteger      *int         `json:"valueInteger,omitempty"`
	ValueDecimal      *json.Number `json:"valueDecimal,omitempty"`
	ValueUri          *string      `json:"valueUri,omitempty"`
	ValueCode         *string      `json:"valueCode,omitempty"`
	ValueDateTime     *string      `json:"valueDateTime,omitempty"`
}

// The codes that are contained in the value set expansion.
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
	Contains          []ValueSetExpansionContains                `json:"contains,omitempty"`
}

// This function returns resource reference information
func (r ValueSet) ResourceRef() (string, *string) {
	return "ValueSet", r.ID
}

// This function returns resource reference information
func (r ValueSet) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherValueSet ValueSet

// MarshalJSON marshals the given ValueSet as JSON into a byte slice
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
