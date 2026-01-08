package fhir430

import "encoding/json"

// SearchParameter is documented here http://hl7.org/fhir/StructureDefinition/SearchParameter
// A search parameter that defines a named search item that can be used to search/filter on a resource.
type SearchParameter struct {
	ID                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []json.RawMessage          `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Url               string                     `json:"url"`
	Version           *string                    `json:"version,omitempty"`
	Name              string                     `json:"name"`
	DerivedFrom       *string                    `json:"derivedFrom,omitempty"`
	Status            PublicationStatus          `json:"status"`
	Experimental      *bool                      `json:"experimental,omitempty"`
	Date              *string                    `json:"date,omitempty"`
	Publisher         *string                    `json:"publisher,omitempty"`
	Contact           []ContactDetail            `json:"contact,omitempty"`
	Description       string                     `json:"description"`
	UseContext        []UsageContext             `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept          `json:"jurisdiction,omitempty"`
	Purpose           *string                    `json:"purpose,omitempty"`
	Code              string                     `json:"code"`
	Base              []ResourceType             `json:"base"`
	Type              SearchParamType            `json:"type"`
	Expression        *string                    `json:"expression,omitempty"`
	Xpath             *string                    `json:"xpath,omitempty"`
	XpathUsage        *XPathUsageType            `json:"xpathUsage,omitempty"`
	Target            []ResourceType             `json:"target,omitempty"`
	MultipleOr        *bool                      `json:"multipleOr,omitempty"`
	MultipleAnd       *bool                      `json:"multipleAnd,omitempty"`
	Comparator        []SearchComparator         `json:"comparator,omitempty"`
	Modifier          []SearchModifierCode       `json:"modifier,omitempty"`
	Chain             []string                   `json:"chain,omitempty"`
	Component         []SearchParameterComponent `json:"component,omitempty"`
}

// Used to define the parts of a composite search parameter.
type SearchParameterComponent struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Definition        string      `json:"definition"`
	Expression        string      `json:"expression"`
}

// This function returns resource reference information
func (r SearchParameter) ResourceRef() (string, *string) {
	return "SearchParameter", r.ID
}

// This function returns resource reference information
func (r SearchParameter) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSearchParameter SearchParameter

// MarshalJSON marshals the given SearchParameter as JSON into a byte slice
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}

// UnmarshalSearchParameter unmarshals a SearchParameter.
func UnmarshalSearchParameter(b []byte) (SearchParameter, error) {
	var searchParameter SearchParameter
	if err := json.Unmarshal(b, &searchParameter); err != nil {
		return searchParameter, err
	}
	return searchParameter, nil
}
