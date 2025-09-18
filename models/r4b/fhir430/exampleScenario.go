
package fhir430

import "encoding/json"
// ExampleScenario is documented here http://hl7.org/fhir/StructureDefinition/ExampleScenario
// Example of workflow instance.
type ExampleScenario struct {
	ID                *string                   `json:"ID,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []json.RawMessage         `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Url               *string                   `json:"url,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	Version           *string                   `json:"version,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	Status            PublicationStatus         `json:"status"`
	Experimental      *bool                     `json:"experimental,omitempty"`
	Date              *string                   `json:"date,omitempty"`
	Publisher         *string                   `json:"publisher,omitempty"`
	Contact           []ContactDetail           `json:"contact,omitempty"`
	UseContext        []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose           *string                   `json:"purpose,omitempty"`
	Actor             []ExampleScenarioActor    `json:"actor,omitempty"`
	Instance          []ExampleScenarioInstance `json:"instance,omitempty"`
	Process           []ExampleScenarioProcess  `json:"process,omitempty"`
	Workflow          []string                  `json:"workflow,omitempty"`
}

// Actor participating in the resource.
type ExampleScenarioActor struct {
	ID                *string                  `json:"ID,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	ActorId           string                   `json:"actorId"`
	Type              ExampleScenarioActorType `json:"type"`
	Name              *string                  `json:"name,omitempty"`
	Description       *string                  `json:"description,omitempty"`
}

// Each resource and each version that is present in the workflow.
type ExampleScenarioInstance struct {
	ID                *string                                    `json:"ID,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	ResourceId        string                                     `json:"resourceId"`
	ResourceType      ResourceType                               `json:"resourceType"`
	Name              *string                                    `json:"name,omitempty"`
	Description       *string                                    `json:"description,omitempty"`
	Version           []ExampleScenarioInstanceVersion           `json:"version,omitempty"`
	ContainedInstance []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`
}

// A specific version of the resource.
type ExampleScenarioInstanceVersion struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	VersionId         string      `json:"versionId"`
	Description       string      `json:"description"`
}

// Resources contained in the instance (e.g. the observations contained in a bundle).
type ExampleScenarioInstanceContainedInstance struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ResourceId        string      `json:"resourceId"`
	VersionId         *string     `json:"versionId,omitempty"`
}

// Each major process - a group of operations.
type ExampleScenarioProcess struct {
	ID                *string                      `json:"ID,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Title             string                       `json:"title"`
	Description       *string                      `json:"description,omitempty"`
	PreConditions     *string                      `json:"preConditions,omitempty"`
	PostConditions    *string                      `json:"postConditions,omitempty"`
	Step              []ExampleScenarioProcessStep `json:"step,omitempty"`
}

// Each step of the process.
type ExampleScenarioProcessStep struct {
	ID                *string                                 `json:"ID,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Process           []ExampleScenarioProcess                `json:"process,omitempty"`
	Pause             *bool                                   `json:"pause,omitempty"`
	Operation         *ExampleScenarioProcessStepOperation    `json:"operation,omitempty"`
	Alternative       []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`
}

// Each interaction or action.
type ExampleScenarioProcessStepOperation struct {
	ID                *string                                   `json:"ID,omitempty"`
	Extension         []Extension                               `json:"extension,omitempty"`
	ModifierExtension []Extension                               `json:"modifierExtension,omitempty"`
	Number            string                                    `json:"number"`
	Type              *string                                   `json:"type,omitempty"`
	Name              *string                                   `json:"name,omitempty"`
	Initiator         *string                                   `json:"initiator,omitempty"`
	Receiver          *string                                   `json:"receiver,omitempty"`
	Description       *string                                   `json:"description,omitempty"`
	InitiatorActive   *bool                                     `json:"initiatorActive,omitempty"`
	ReceiverActive    *bool                                     `json:"receiverActive,omitempty"`
	Request           *ExampleScenarioInstanceContainedInstance `json:"request,omitempty"`
	Response          *ExampleScenarioInstanceContainedInstance `json:"response,omitempty"`
}

// Indicates an alternative step that can be taken instead of the operations on the base step in exceptional/atypical circumstances.
type ExampleScenarioProcessStepAlternative struct {
	ID                *string                      `json:"ID,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Title             string                       `json:"title"`
	Description       *string                      `json:"description,omitempty"`
	Step              []ExampleScenarioProcessStep `json:"step,omitempty"`
}

// This function returns resource reference information
func (r ExampleScenario) ResourceRef() (string, *string) {
	return "ExampleScenario", r.ID
}

// This function returns resource reference information
func (r ExampleScenario) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherExampleScenario ExampleScenario

// MarshalJSON marshals the given ExampleScenario as JSON into a byte slice
func (r ExampleScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExampleScenario
		ResourceType string `json:"resourceType"`
	}{
		OtherExampleScenario: OtherExampleScenario(r),
		ResourceType:         "ExampleScenario",
	})
}

// UnmarshalExampleScenario unmarshals a ExampleScenario.
func UnmarshalExampleScenario(b []byte) (ExampleScenario, error) {
	var exampleScenario ExampleScenario
	if err := json.Unmarshal(b, &exampleScenario); err != nil {
		return exampleScenario, err
	}
	return exampleScenario, nil
}
