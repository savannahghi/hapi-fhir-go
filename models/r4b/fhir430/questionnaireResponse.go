
package fhir430

import "encoding/json"
// QuestionnaireResponse is documented here http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
// A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponse struct {
	ID                *string                     `json:"ID,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []json.RawMessage           `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `json:"identifier,omitempty"`
	BasedOn           []Reference                 `json:"basedOn,omitempty"`
	PartOf            []Reference                 `json:"partOf,omitempty"`
	Questionnaire     *string                     `json:"questionnaire,omitempty"`
	Status            QuestionnaireResponseStatus `json:"status"`
	Subject           *Reference                  `json:"subject,omitempty"`
	Encounter         *Reference                  `json:"encounter,omitempty"`
	Authored          *string                     `json:"authored,omitempty"`
	Author            *Reference                  `json:"author,omitempty"`
	Source            *Reference                  `json:"source,omitempty"`
	Item              []QuestionnaireResponseItem `json:"item,omitempty"`
}

// A group or question item from the original questionnaire for which answers are provided.
// Groups cannot have answers and therefore must nest directly within item. When dealing with questions, nesting must occur within each answer because some questions may have multiple answers (and the nesting occurs for each answer).
type QuestionnaireResponseItem struct {
	ID                *string                           `json:"ID,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	LinkId            string                            `json:"linkId"`
	Definition        *string                           `json:"definition,omitempty"`
	Text              *string                           `json:"text,omitempty"`
	Answer            []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`
	Item              []QuestionnaireResponseItem       `json:"item,omitempty"`
}

// The respondent's answer(s) to the question.
// The value is nested because we cannot have a repeating structure that has variable type.
type QuestionnaireResponseItemAnswer struct {
	ID                *string                     `json:"ID,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	ValueBoolean      *bool                       `json:"valueBoolean,omitempty"`
	ValueDecimal      *json.Number                `json:"valueDecimal,omitempty"`
	ValueInteger      *int                        `json:"valueInteger,omitempty"`
	ValueDate         *string                     `json:"valueDate,omitempty"`
	ValueDateTime     *string                     `json:"valueDateTime,omitempty"`
	ValueTime         *string                     `json:"valueTime,omitempty"`
	ValueString       *string                     `json:"valueString,omitempty"`
	ValueUri          *string                     `json:"valueUri,omitempty"`
	ValueAttachment   *Attachment                 `json:"valueAttachment,omitempty"`
	ValueCoding       *Coding                     `json:"valueCoding,omitempty"`
	ValueQuantity     *Quantity                   `json:"valueQuantity,omitempty"`
	ValueReference    *Reference                  `json:"valueReference,omitempty"`
	Item              []QuestionnaireResponseItem `json:"item,omitempty"`
}

// This function returns resource reference information
func (r QuestionnaireResponse) ResourceRef() (string, *string) {
	return "QuestionnaireResponse", r.ID
}

// This function returns resource reference information
func (r QuestionnaireResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherQuestionnaireResponse QuestionnaireResponse

// MarshalJSON marshals the given QuestionnaireResponse as JSON into a byte slice
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaireResponse: OtherQuestionnaireResponse(r),
		ResourceType:               "QuestionnaireResponse",
	})
}

// UnmarshalQuestionnaireResponse unmarshals a QuestionnaireResponse.
func UnmarshalQuestionnaireResponse(b []byte) (QuestionnaireResponse, error) {
	var questionnaireResponse QuestionnaireResponse
	if err := json.Unmarshal(b, &questionnaireResponse); err != nil {
		return questionnaireResponse, err
	}
	return questionnaireResponse, nil
}
