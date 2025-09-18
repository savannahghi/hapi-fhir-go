
package fhir430

import "encoding/json"
// CapabilityStatement is documented here http://hl7.org/fhir/StructureDefinition/CapabilityStatement
// A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server for a particular version of FHIR that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement struct {
	ID                  *string                            `json:"ID,omitempty"`
	Meta                *Meta                              `json:"meta,omitempty"`
	ImplicitRules       *string                            `json:"implicitRules,omitempty"`
	Language            *string                            `json:"language,omitempty"`
	Text                *Narrative                         `json:"text,omitempty"`
	Contained           []json.RawMessage                  `json:"contained,omitempty"`
	Extension           []Extension                        `json:"extension,omitempty"`
	ModifierExtension   []Extension                        `json:"modifierExtension,omitempty"`
	Url                 *string                            `json:"url,omitempty"`
	Version             *string                            `json:"version,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	Title               *string                            `json:"title,omitempty"`
	Status              PublicationStatus                  `json:"status"`
	Experimental        *bool                              `json:"experimental,omitempty"`
	Date                string                             `json:"date"`
	Publisher           *string                            `json:"publisher,omitempty"`
	Contact             []ContactDetail                    `json:"contact,omitempty"`
	Description         *string                            `json:"description,omitempty"`
	UseContext          []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose             *string                            `json:"purpose,omitempty"`
	Kind                CapabilityStatementKind            `json:"kind"`
	Instantiates        []string                           `json:"instantiates,omitempty"`
	Imports             []string                           `json:"imports,omitempty"`
	Software            *CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation      *CapabilityStatementImplementation `json:"implementation,omitempty"`
	FhirVersion         FHIRVersion                        `json:"fhirVersion"`
	Format              []string                           `json:"format"`
	PatchFormat         []string                           `json:"patchFormat,omitempty"`
	ImplementationGuide []string                           `json:"implementationGuide,omitempty"`
	Rest                []CapabilityStatementRest          `json:"rest,omitempty"`
	Messaging           []CapabilityStatementMessaging     `json:"messaging,omitempty"`
	Document            []CapabilityStatementDocument      `json:"document,omitempty"`
}

// Software that is covered by this capability statement.  It is used when the capability statement describes the capabilities of a particular software version, independent of an installation.
type CapabilityStatementSoftware struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Version           *string     `json:"version,omitempty"`
	ReleaseDate       *string     `json:"releaseDate,omitempty"`
}

// Identifies a specific implementation instance that is described by the capability statement - i.e. a particular installation, rather than the capabilities of a software program.
type CapabilityStatementImplementation struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	Url               *string     `json:"url,omitempty"`
	Custodian         *Reference  `json:"custodian,omitempty"`
}

// A definition of the restful capabilities of the solution, if any.
// Multiple repetitions allow definition of both client and/or server behaviors or possibly behaviors under different configuration settings (for software or requirements statements).
type CapabilityStatementRest struct {
	ID                *string                                      `json:"ID,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Mode              RestfulCapabilityMode                        `json:"mode"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Security          *CapabilityStatementRestSecurity             `json:"security,omitempty"`
	Resource          []CapabilityStatementRestResource            `json:"resource,omitempty"`
	Interaction       []CapabilityStatementRestInteraction         `json:"interaction,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
	Compartment       []string                                     `json:"compartment,omitempty"`
}

// Information about security implementation from an interface perspective - what a client needs to know.
type CapabilityStatementRestSecurity struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Cors              *bool             `json:"cors,omitempty"`
	Service           []CodeableConcept `json:"service,omitempty"`
	Description       *string           `json:"description,omitempty"`
}

// A specification of the restful capabilities of the solution for a specific resource type.
// Max of one repetition per resource type.
type CapabilityStatementRestResource struct {
	ID                *string                                      `json:"ID,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Type              ResourceType                                 `json:"type"`
	Profile           *string                                      `json:"profile,omitempty"`
	SupportedProfile  []string                                     `json:"supportedProfile,omitempty"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Interaction       []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	Versioning        *ResourceVersionPolicy                       `json:"versioning,omitempty"`
	ReadHistory       *bool                                        `json:"readHistory,omitempty"`
	UpdateCreate      *bool                                        `json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                        `json:"conditionalCreate,omitempty"`
	ConditionalRead   *ConditionalReadStatus                       `json:"conditionalRead,omitempty"`
	ConditionalUpdate *bool                                        `json:"conditionalUpdate,omitempty"`
	ConditionalDelete *ConditionalDeleteStatus                     `json:"conditionalDelete,omitempty"`
	ReferencePolicy   []ReferenceHandlingPolicy                    `json:"referencePolicy,omitempty"`
	SearchInclude     []string                                     `json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                     `json:"searchRevInclude,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
}

// Identifies a restful operation supported by the solution.
// In general, a Resource will only appear in a CapabilityStatement if the server actually has some capabilities - e.g. there is at least one interaction supported. However interactions can be omitted to support summarization (_summary = true).
type CapabilityStatementRestResourceInteraction struct {
	ID                *string                `json:"ID,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Code              TypeRestfulInteraction `json:"code"`
	Documentation     *string                `json:"documentation,omitempty"`
}

// Search parameters for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
// The search parameters should include the control search parameters such as _sort, _count, etc. that also apply to this resource (though many will be listed at [CapabilityStatement.rest.searchParam](capabilitystatement-definitions.html#CapabilityStatement.rest.searchParam)). The behavior of some search parameters may be further described by other code or extension elements, or narrative within the capability statement or linked [SearchParameter](searchparameter.html#) definitions.
type CapabilityStatementRestResourceSearchParam struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Name              string          `json:"name"`
	Definition        *string         `json:"definition,omitempty"`
	Type              SearchParamType `json:"type"`
	Documentation     *string         `json:"documentation,omitempty"`
}

// Definition of an operation or a named query together with its parameters and their meaning and type. Consult the definition of the operation for details about how to invoke the operation, and the parameters.
/*
Operations linked from CapabilityStatement.rest.resource.operation must have OperationDefinition.type = true or OperationDefinition.instance = true.

If an operation that is listed in multiple CapabilityStatement.rest.resource.operation (e.g. for different resource types), then clients should understand that the operation is only supported on the specified resource types, and that may be a subset of those listed in OperationDefinition.resource.
*/
type CapabilityStatementRestResourceOperation struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Definition        string      `json:"definition"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// A specification of restful operations supported by the system.
type CapabilityStatementRestInteraction struct {
	ID                *string                  `json:"ID,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Code              SystemRestfulInteraction `json:"code"`
	Documentation     *string                  `json:"documentation,omitempty"`
}

// A description of the messaging capabilities of the solution.
// Multiple repetitions allow the documentation of multiple endpoints per solution.
type CapabilityStatementMessaging struct {
	ID                *string                                        `json:"ID,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache     *int                                           `json:"reliableCache,omitempty"`
	Documentation     *string                                        `json:"documentation,omitempty"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

// An endpoint (network accessible address) to which messages and/or replies are to be sent.
type CapabilityStatementMessagingEndpoint struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Protocol          Coding      `json:"protocol"`
	Address           string      `json:"address"`
}

// References to message definitions for messages this system can send or receive.
// This is a proposed alternative to the messaging.event structure.
type CapabilityStatementMessagingSupportedMessage struct {
	ID                *string             `json:"ID,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Mode              EventCapabilityMode `json:"mode"`
	Definition        string              `json:"definition"`
}

// A document definition.
type CapabilityStatementDocument struct {
	ID                *string      `json:"ID,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Mode              DocumentMode `json:"mode"`
	Documentation     *string      `json:"documentation,omitempty"`
	Profile           string       `json:"profile"`
}

// This function returns resource reference information
func (r CapabilityStatement) ResourceRef() (string, *string) {
	return "CapabilityStatement", r.ID
}

// This function returns resource reference information
func (r CapabilityStatement) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCapabilityStatement CapabilityStatement

// MarshalJSON marshals the given CapabilityStatement as JSON into a byte slice
func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCapabilityStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherCapabilityStatement: OtherCapabilityStatement(r),
		ResourceType:             "CapabilityStatement",
	})
}

// UnmarshalCapabilityStatement unmarshals a CapabilityStatement.
func UnmarshalCapabilityStatement(b []byte) (CapabilityStatement, error) {
	var capabilityStatement CapabilityStatement
	if err := json.Unmarshal(b, &capabilityStatement); err != nil {
		return capabilityStatement, err
	}
	return capabilityStatement, nil
}
