
package fhir430
// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
// Base StructureDefinition for DataRequirement Type: Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement struct {
	ID                     *string                     `json:"ID,omitempty"`
	Extension              []Extension                 `json:"extension,omitempty"`
	Type                   string                      `json:"type"`
	Profile                []string                    `json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept            `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                  `json:"subjectReference,omitempty"`
	MustSupport            []string                    `json:"mustSupport,omitempty"`
	CodeFilter             []DataRequirementCodeFilter `json:"codeFilter,omitempty"`
	DateFilter             []DataRequirementDateFilter `json:"dateFilter,omitempty"`
	Limit                  *int                        `json:"limit,omitempty"`
	Sort                   []DataRequirementSort       `json:"sort,omitempty"`
}

// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data. Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
type DataRequirementCodeFilter struct {
	ID          *string     `json:"ID,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Path        *string     `json:"path,omitempty"`
	SearchParam *string     `json:"searchParam,omitempty"`
	ValueSet    *string     `json:"valueSet,omitempty"`
	Code        []Coding    `json:"code,omitempty"`
}

// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements. Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
type DataRequirementDateFilter struct {
	ID            *string     `json:"ID,omitempty"`
	Extension     []Extension `json:"extension,omitempty"`
	Path          *string     `json:"path,omitempty"`
	SearchParam   *string     `json:"searchParam,omitempty"`
	ValueDateTime *string     `json:"valueDateTime,omitempty"`
	ValuePeriod   *Period     `json:"valuePeriod,omitempty"`
	ValueDuration *Duration   `json:"valueDuration,omitempty"`
}

// Specifies the order of the results to be returned.
// This element can be used in combination with the sort element to specify quota requirements such as "the most recent 5" or "the highest 5". When multiple sorts are specified, they are applied in the order they appear in the resource.
type DataRequirementSort struct {
	ID        *string       `json:"ID,omitempty"`
	Extension []Extension   `json:"extension,omitempty"`
	Path      string        `json:"path"`
	Direction SortDirection `json:"direction"`
}
