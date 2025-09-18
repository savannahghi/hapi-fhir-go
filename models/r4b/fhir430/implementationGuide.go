
package fhir430

import "encoding/json"
// ImplementationGuide is documented here http://hl7.org/fhir/StructureDefinition/ImplementationGuide
// A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide struct {
	ID                *string                        `json:"ID,omitempty"`
	Meta              *Meta                          `json:"meta,omitempty"`
	ImplicitRules     *string                        `json:"implicitRules,omitempty"`
	Language          *string                        `json:"language,omitempty"`
	Text              *Narrative                     `json:"text,omitempty"`
	Contained         []json.RawMessage              `json:"contained,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Url               string                         `json:"url"`
	Version           *string                        `json:"version,omitempty"`
	Name              string                         `json:"name"`
	Title             *string                        `json:"title,omitempty"`
	Status            PublicationStatus              `json:"status"`
	Experimental      *bool                          `json:"experimental,omitempty"`
	Date              *string                        `json:"date,omitempty"`
	Publisher         *string                        `json:"publisher,omitempty"`
	Contact           []ContactDetail                `json:"contact,omitempty"`
	Description       *string                        `json:"description,omitempty"`
	UseContext        []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `json:"jurisdiction,omitempty"`
	PackageId         string                         `json:"packageId"`
	License           *SPDXLicense                   `json:"license,omitempty"`
	FhirVersion       []FHIRVersion                  `json:"fhirVersion"`
	DependsOn         []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`
	Global            []ImplementationGuideGlobal    `json:"global,omitempty"`
	Definition        *ImplementationGuideDefinition `json:"definition,omitempty"`
	Manifest          *ImplementationGuideManifest   `json:"manifest,omitempty"`
}

// Another implementation guide that this implementation depends on. Typically, an implementation guide uses value sets, profiles etc.defined in other implementation guides.
type ImplementationGuideDependsOn struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Uri               string      `json:"uri"`
	PackageId         *string     `json:"packageId,omitempty"`
	Version           *string     `json:"version,omitempty"`
}

// A set of profiles that all resources covered by this implementation guide must conform to.
// See [Default Profiles](implementationguide.html#default) for a discussion of which resources are 'covered' by an implementation guide.
type ImplementationGuideGlobal struct {
	ID                *string      `json:"ID,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Type              ResourceType `json:"type"`
	Profile           string       `json:"profile"`
}

// The information needed by an IG publisher tool to publish the whole implementation guide.
// Principally, this consists of information abuot source resource and file locations, and build parameters and templates.
type ImplementationGuideDefinition struct {
	ID                *string                                  `json:"ID,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Grouping          []ImplementationGuideDefinitionGrouping  `json:"grouping,omitempty"`
	Resource          []ImplementationGuideDefinitionResource  `json:"resource"`
	Page              *ImplementationGuideDefinitionPage       `json:"page,omitempty"`
	Parameter         []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`
	Template          []ImplementationGuideDefinitionTemplate  `json:"template,omitempty"`
}

// A logical group of resources. Logical groups can be used when building pages.
// Groupings are arbitrary sub-divisions of content. Typically, they are used to help build Table of Contents automatically.
type ImplementationGuideDefinitionGrouping struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Description       *string     `json:"description,omitempty"`
}

// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
type ImplementationGuideDefinitionResource struct {
	ID                *string       `json:"ID,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Reference         Reference     `json:"reference"`
	FhirVersion       []FHIRVersion `json:"fhirVersion,omitempty"`
	Name              *string       `json:"name,omitempty"`
	Description       *string       `json:"description,omitempty"`
	ExampleBoolean    *bool         `json:"exampleBoolean,omitempty"`
	ExampleCanonical  *string       `json:"exampleCanonical,omitempty"`
	GroupingId        *string       `json:"groupingId,omitempty"`
}

// A page / section in the implementation guide. The root page is the implementation guide home page.
// Pages automatically become sections if they have sub-pages. By convention, the home page is called index.html.
type ImplementationGuideDefinitionPage struct {
	ID                *string                             `json:"ID,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	NameUrl           string                              `json:"nameUrl"`
	NameReference     Reference                           `json:"nameReference"`
	Title             string                              `json:"title"`
	Generation        GuidePageGeneration                 `json:"generation"`
	Page              []ImplementationGuideDefinitionPage `json:"page,omitempty"`
}

// Defines how IG is built by tools.
type ImplementationGuideDefinitionParameter struct {
	ID                *string            `json:"ID,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Code              GuideParameterCode `json:"code"`
	Value             string             `json:"value"`
}

// A template for building resources.
type ImplementationGuideDefinitionTemplate struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Source            string      `json:"source"`
	Scope             *string     `json:"scope,omitempty"`
}

// Information about an assembled implementation guide, created by the publication tooling.
type ImplementationGuideManifest struct {
	ID                *string                               `json:"ID,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Rendering         *string                               `json:"rendering,omitempty"`
	Resource          []ImplementationGuideManifestResource `json:"resource"`
	Page              []ImplementationGuideManifestPage     `json:"page,omitempty"`
	Image             []string                              `json:"image,omitempty"`
	Other             []string                              `json:"other,omitempty"`
}

// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
type ImplementationGuideManifestResource struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	ExampleBoolean    *bool       `json:"exampleBoolean,omitempty"`
	ExampleCanonical  *string     `json:"exampleCanonical,omitempty"`
	RelativePath      *string     `json:"relativePath,omitempty"`
}

// Information about a page within the IG.
type ImplementationGuideManifestPage struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Title             *string     `json:"title,omitempty"`
	Anchor            []string    `json:"anchor,omitempty"`
}

// This function returns resource reference information
func (r ImplementationGuide) ResourceRef() (string, *string) {
	return "ImplementationGuide", r.ID
}

// This function returns resource reference information
func (r ImplementationGuide) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherImplementationGuide ImplementationGuide

// MarshalJSON marshals the given ImplementationGuide as JSON into a byte slice
func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImplementationGuide
		ResourceType string `json:"resourceType"`
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
}

// UnmarshalImplementationGuide unmarshals a ImplementationGuide.
func UnmarshalImplementationGuide(b []byte) (ImplementationGuide, error) {
	var implementationGuide ImplementationGuide
	if err := json.Unmarshal(b, &implementationGuide); err != nil {
		return implementationGuide, err
	}
	return implementationGuide, nil
}
