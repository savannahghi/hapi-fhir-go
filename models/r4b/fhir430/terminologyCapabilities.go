
package fhir430

import "encoding/json"
// TerminologyCapabilities is documented here http://hl7.org/fhir/StructureDefinition/TerminologyCapabilities
// A TerminologyCapabilities resource documents a set of capabilities (behaviors) of a FHIR Terminology Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type TerminologyCapabilities struct {
	ID                *string                                `json:"ID,omitempty"`
	Meta              *Meta                                  `json:"meta,omitempty"`
	ImplicitRules     *string                                `json:"implicitRules,omitempty"`
	Language          *string                                `json:"language,omitempty"`
	Text              *Narrative                             `json:"text,omitempty"`
	Contained         []json.RawMessage                      `json:"contained,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Url               *string                                `json:"url,omitempty"`
	Version           *string                                `json:"version,omitempty"`
	Name              *string                                `json:"name,omitempty"`
	Title             *string                                `json:"title,omitempty"`
	Status            PublicationStatus                      `json:"status"`
	Experimental      *bool                                  `json:"experimental,omitempty"`
	Date              string                                 `json:"date"`
	Publisher         *string                                `json:"publisher,omitempty"`
	Contact           []ContactDetail                        `json:"contact,omitempty"`
	Description       *string                                `json:"description,omitempty"`
	UseContext        []UsageContext                         `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                      `json:"jurisdiction,omitempty"`
	Purpose           *string                                `json:"purpose,omitempty"`
	Kind              CapabilityStatementKind                `json:"kind"`
	Software          *TerminologyCapabilitiesSoftware       `json:"software,omitempty"`
	Implementation    *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`
	LockedDate        *bool                                  `json:"lockedDate,omitempty"`
	CodeSystem        []TerminologyCapabilitiesCodeSystem    `json:"codeSystem,omitempty"`
	Expansion         *TerminologyCapabilitiesExpansion      `json:"expansion,omitempty"`
	CodeSearch        *CodeSearchSupport                     `json:"codeSearch,omitempty"`
	ValidateCode      *TerminologyCapabilitiesValidateCode   `json:"validateCode,omitempty"`
	Translation       *TerminologyCapabilitiesTranslation    `json:"translation,omitempty"`
	Closure           *TerminologyCapabilitiesClosure        `json:"closure,omitempty"`
}

// Software that is covered by this terminology capability statement.  It is used when the statement describes the capabilities of a particular software version, independent of an installation.
type TerminologyCapabilitiesSoftware struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
}

// Identifies a specific implementation instance that is described by the terminology capability statement - i.e. a particular installation, rather than the capabilities of a software program.
type TerminologyCapabilitiesImplementation struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
}

// Identifies a code system that is supported by the server. If there is a no code system URL, then this declares the general assumptions a client can make about support for any CodeSystem resource.
// The code system - identified by its system URL - may also be declared explicitly as a Code System Resource at /CodeSystem, but it might not be.
type TerminologyCapabilitiesCodeSystem struct {
	ID                *string                                    `json:"ID,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Uri               *string                                    `json:"uri,omitempty"`
	Version           []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`
	Subsumption       *bool                                      `json:"subsumption,omitempty"`
}

// For the code system, a list of versions that are supported by the server.
// Language translations might not be available for all codes.
type TerminologyCapabilitiesCodeSystemVersion struct {
	ID                *string                                          `json:"ID,omitempty"`
	Extension         []Extension                                      `json:"extension,omitempty"`
	ModifierExtension []Extension                                      `json:"modifierExtension,omitempty"`
	Code              *string                                          `json:"code,omitempty"`
	IsDefault         *bool                                            `json:"isDefault,omitempty"`
	Compositional     *bool                                            `json:"compositional,omitempty"`
	Language          []string                                         `json:"language,omitempty"`
	Filter            []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty"`
	Property          []string                                         `json:"property,omitempty"`
}

// Filter Properties supported.
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Op                []string    `json:"op"`
}

// Information about the [ValueSet/$expand](valueset-operation-expand.html) operation.
type TerminologyCapabilitiesExpansion struct {
	ID                *string                                     `json:"ID,omitempty"`
	Extension         []Extension                                 `json:"extension,omitempty"`
	ModifierExtension []Extension                                 `json:"modifierExtension,omitempty"`
	Hierarchical      *bool                                       `json:"hierarchical,omitempty"`
	Paging            *bool                                       `json:"paging,omitempty"`
	Incomplete        *bool                                       `json:"incomplete,omitempty"`
	Parameter         []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty"`
	TextFilter        *string                                     `json:"textFilter,omitempty"`
}

// Supported expansion parameter.
type TerminologyCapabilitiesExpansionParameter struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// Information about the [ValueSet/$validate-code](valueset-operation-validate-code.html) operation.
type TerminologyCapabilitiesValidateCode struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Translations      bool        `json:"translations"`
}

// Information about the [ConceptMap/$translate](conceptmap-operation-translate.html) operation.
type TerminologyCapabilitiesTranslation struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NeedsMap          bool        `json:"needsMap"`
}

// Whether the $closure operation is supported.
type TerminologyCapabilitiesClosure struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Translation       *bool       `json:"translation,omitempty"`
}

// This function returns resource reference information
func (r TerminologyCapabilities) ResourceRef() (string, *string) {
	return "TerminologyCapabilities", r.ID
}

// This function returns resource reference information
func (r TerminologyCapabilities) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherTerminologyCapabilities TerminologyCapabilities

// MarshalJSON marshals the given TerminologyCapabilities as JSON into a byte slice
func (r TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTerminologyCapabilities
		ResourceType string `json:"resourceType"`
	}{
		OtherTerminologyCapabilities: OtherTerminologyCapabilities(r),
		ResourceType:                 "TerminologyCapabilities",
	})
}

// UnmarshalTerminologyCapabilities unmarshals a TerminologyCapabilities.
func UnmarshalTerminologyCapabilities(b []byte) (TerminologyCapabilities, error) {
	var terminologyCapabilities TerminologyCapabilities
	if err := json.Unmarshal(b, &terminologyCapabilities); err != nil {
		return terminologyCapabilities, err
	}
	return terminologyCapabilities, nil
}
