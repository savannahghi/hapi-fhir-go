
package fhir430

import "encoding/json"
// Task is documented here http://hl7.org/fhir/StructureDefinition/Task
// A task to be performed.
type Task struct {
	ID                    *string           `json:"ID,omitempty"`
	Meta                  *Meta             `json:"meta,omitempty"`
	ImplicitRules         *string           `json:"implicitRules,omitempty"`
	Language              *string           `json:"language,omitempty"`
	Text                  *Narrative        `json:"text,omitempty"`
	Contained             []json.RawMessage `json:"contained,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Identifier            []Identifier      `json:"identifier,omitempty"`
	InstantiatesCanonical *string           `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string           `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference       `json:"basedOn,omitempty"`
	GroupIdentifier       *Identifier       `json:"groupIdentifier,omitempty"`
	PartOf                []Reference       `json:"partOf,omitempty"`
	Status                TaskStatus        `json:"status"`
	StatusReason          *CodeableConcept  `json:"statusReason,omitempty"`
	BusinessStatus        *CodeableConcept  `json:"businessStatus,omitempty"`
	Intent                string            `json:"intent"`
	Priority              *RequestPriority  `json:"priority,omitempty"`
	Code                  *CodeableConcept  `json:"code,omitempty"`
	Description           *string           `json:"description,omitempty"`
	Focus                 *Reference        `json:"focus,omitempty"`
	For                   *Reference        `json:"for,omitempty"`
	Encounter             *Reference        `json:"encounter,omitempty"`
	ExecutionPeriod       *Period           `json:"executionPeriod,omitempty"`
	AuthoredOn            *string           `json:"authoredOn,omitempty"`
	LastModified          *string           `json:"lastModified,omitempty"`
	Requester             *Reference        `json:"requester,omitempty"`
	PerformerType         []CodeableConcept `json:"performerType,omitempty"`
	Owner                 *Reference        `json:"owner,omitempty"`
	Location              *Reference        `json:"location,omitempty"`
	ReasonCode            *CodeableConcept  `json:"reasonCode,omitempty"`
	ReasonReference       *Reference        `json:"reasonReference,omitempty"`
	Insurance             []Reference       `json:"insurance,omitempty"`
	Note                  []Annotation      `json:"note,omitempty"`
	RelevantHistory       []Reference       `json:"relevantHistory,omitempty"`
	Restriction           *TaskRestriction  `json:"restriction,omitempty"`
	Input                 []TaskInput       `json:"input,omitempty"`
	Output                []TaskOutput      `json:"output,omitempty"`
}

// If the Task.focus is a request resource and the task is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
type TaskRestriction struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Repetitions       *int        `json:"repetitions,omitempty"`
	Period            *Period     `json:"period,omitempty"`
	Recipient         []Reference `json:"recipient,omitempty"`
}

// Additional information that may be needed in the execution of the task.
type TaskInput struct {
	ID                       *string             `json:"ID,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Type                     CodeableConcept     `json:"type"`
	ValueBase64Binary        string              `json:"valueBase64Binary"`
	ValueBoolean             bool                `json:"valueBoolean"`
	ValueCanonical           string              `json:"valueCanonical"`
	ValueCode                string              `json:"valueCode"`
	ValueDate                string              `json:"valueDate"`
	ValueDateTime            string              `json:"valueDateTime"`
	ValueDecimal             json.Number         `json:"valueDecimal"`
	ValueId                  string              `json:"valueId"`
	ValueInstant             string              `json:"valueInstant"`
	ValueInteger             int                 `json:"valueInteger"`
	ValueMarkdown            string              `json:"valueMarkdown"`
	ValueOid                 string              `json:"valueOid"`
	ValuePositiveInt         int                 `json:"valuePositiveInt"`
	ValueString              string              `json:"valueString"`
	ValueTime                string              `json:"valueTime"`
	ValueUnsignedInt         int                 `json:"valueUnsignedInt"`
	ValueUri                 string              `json:"valueUri"`
	ValueUrl                 string              `json:"valueUrl"`
	ValueUuid                string              `json:"valueUuid"`
	ValueAddress             Address             `json:"valueAddress"`
	ValueAge                 Age                 `json:"valueAge"`
	ValueAnnotation          Annotation          `json:"valueAnnotation"`
	ValueAttachment          Attachment          `json:"valueAttachment"`
	ValueCodeableConcept     CodeableConcept     `json:"valueCodeableConcept"`
	ValueCoding              Coding              `json:"valueCoding"`
	ValueContactPoint        ContactPoint        `json:"valueContactPoint"`
	ValueCount               Count               `json:"valueCount"`
	ValueDistance            Distance            `json:"valueDistance"`
	ValueDuration            Duration            `json:"valueDuration"`
	ValueHumanName           HumanName           `json:"valueHumanName"`
	ValueIdentifier          Identifier          `json:"valueIdentifier"`
	ValueMoney               Money               `json:"valueMoney"`
	ValuePeriod              Period              `json:"valuePeriod"`
	ValueQuantity            Quantity            `json:"valueQuantity"`
	ValueRange               Range               `json:"valueRange"`
	ValueRatio               Ratio               `json:"valueRatio"`
	ValueReference           Reference           `json:"valueReference"`
	ValueSampledData         SampledData         `json:"valueSampledData"`
	ValueSignature           Signature           `json:"valueSignature"`
	ValueTiming              Timing              `json:"valueTiming"`
	ValueContactDetail       ContactDetail       `json:"valueContactDetail"`
	ValueContributor         Contributor         `json:"valueContributor"`
	ValueDataRequirement     DataRequirement     `json:"valueDataRequirement"`
	ValueExpression          Expression          `json:"valueExpression"`
	ValueParameterDefinition ParameterDefinition `json:"valueParameterDefinition"`
	ValueRelatedArtifact     RelatedArtifact     `json:"valueRelatedArtifact"`
	ValueTriggerDefinition   TriggerDefinition   `json:"valueTriggerDefinition"`
	ValueUsageContext        UsageContext        `json:"valueUsageContext"`
	ValueDosage              Dosage              `json:"valueDosage"`
	ValueMeta                Meta                `json:"valueMeta"`
}

// Outputs produced by the Task.
type TaskOutput struct {
	ID                       *string             `json:"ID,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Type                     CodeableConcept     `json:"type"`
	ValueBase64Binary        string              `json:"valueBase64Binary"`
	ValueBoolean             bool                `json:"valueBoolean"`
	ValueCanonical           string              `json:"valueCanonical"`
	ValueCode                string              `json:"valueCode"`
	ValueDate                string              `json:"valueDate"`
	ValueDateTime            string              `json:"valueDateTime"`
	ValueDecimal             json.Number         `json:"valueDecimal"`
	ValueId                  string              `json:"valueId"`
	ValueInstant             string              `json:"valueInstant"`
	ValueInteger             int                 `json:"valueInteger"`
	ValueMarkdown            string              `json:"valueMarkdown"`
	ValueOid                 string              `json:"valueOid"`
	ValuePositiveInt         int                 `json:"valuePositiveInt"`
	ValueString              string              `json:"valueString"`
	ValueTime                string              `json:"valueTime"`
	ValueUnsignedInt         int                 `json:"valueUnsignedInt"`
	ValueUri                 string              `json:"valueUri"`
	ValueUrl                 string              `json:"valueUrl"`
	ValueUuid                string              `json:"valueUuid"`
	ValueAddress             Address             `json:"valueAddress"`
	ValueAge                 Age                 `json:"valueAge"`
	ValueAnnotation          Annotation          `json:"valueAnnotation"`
	ValueAttachment          Attachment          `json:"valueAttachment"`
	ValueCodeableConcept     CodeableConcept     `json:"valueCodeableConcept"`
	ValueCoding              Coding              `json:"valueCoding"`
	ValueContactPoint        ContactPoint        `json:"valueContactPoint"`
	ValueCount               Count               `json:"valueCount"`
	ValueDistance            Distance            `json:"valueDistance"`
	ValueDuration            Duration            `json:"valueDuration"`
	ValueHumanName           HumanName           `json:"valueHumanName"`
	ValueIdentifier          Identifier          `json:"valueIdentifier"`
	ValueMoney               Money               `json:"valueMoney"`
	ValuePeriod              Period              `json:"valuePeriod"`
	ValueQuantity            Quantity            `json:"valueQuantity"`
	ValueRange               Range               `json:"valueRange"`
	ValueRatio               Ratio               `json:"valueRatio"`
	ValueReference           Reference           `json:"valueReference"`
	ValueSampledData         SampledData         `json:"valueSampledData"`
	ValueSignature           Signature           `json:"valueSignature"`
	ValueTiming              Timing              `json:"valueTiming"`
	ValueContactDetail       ContactDetail       `json:"valueContactDetail"`
	ValueContributor         Contributor         `json:"valueContributor"`
	ValueDataRequirement     DataRequirement     `json:"valueDataRequirement"`
	ValueExpression          Expression          `json:"valueExpression"`
	ValueParameterDefinition ParameterDefinition `json:"valueParameterDefinition"`
	ValueRelatedArtifact     RelatedArtifact     `json:"valueRelatedArtifact"`
	ValueTriggerDefinition   TriggerDefinition   `json:"valueTriggerDefinition"`
	ValueUsageContext        UsageContext        `json:"valueUsageContext"`
	ValueDosage              Dosage              `json:"valueDosage"`
	ValueMeta                Meta                `json:"valueMeta"`
}

// This function returns resource reference information
func (r Task) ResourceRef() (string, *string) {
	return "Task", r.ID
}

// This function returns resource reference information
func (r Task) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherTask Task

// MarshalJSON marshals the given Task as JSON into a byte slice
func (r Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTask
		ResourceType string `json:"resourceType"`
	}{
		OtherTask:    OtherTask(r),
		ResourceType: "Task",
	})
}

// UnmarshalTask unmarshals a Task.
func UnmarshalTask(b []byte) (Task, error) {
	var task Task
	if err := json.Unmarshal(b, &task); err != nil {
		return task, err
	}
	return task, nil
}
