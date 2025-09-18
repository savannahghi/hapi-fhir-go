
package fhir430

import "encoding/json"
// Questionnaire is documented here http://hl7.org/fhir/StructureDefinition/Questionnaire
// A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type Questionnaire struct {
	ID                *string             `json:"ID,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []json.RawMessage   `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Url               *string             `json:"url,omitempty"`
	Identifier        []Identifier        `json:"identifier,omitempty"`
	Version           *string             `json:"version,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Title             *string             `json:"title,omitempty"`
	DerivedFrom       []string            `json:"derivedFrom,omitempty"`
	Status            PublicationStatus   `json:"status"`
	Experimental      *bool               `json:"experimental,omitempty"`
	SubjectType       []ResourceType      `json:"subjectType,omitempty"`
	Date              *string             `json:"date,omitempty"`
	Publisher         *string             `json:"publisher,omitempty"`
	Contact           []ContactDetail     `json:"contact,omitempty"`
	Description       *string             `json:"description,omitempty"`
	UseContext        []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose           *string             `json:"purpose,omitempty"`
	ApprovalDate      *string             `json:"approvalDate,omitempty"`
	LastReviewDate    *string             `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period             `json:"effectivePeriod,omitempty"`
	Code              []Coding            `json:"code,omitempty"`
	Item              []QuestionnaireItem `json:"item,omitempty"`
}

// A particular question, question grouping or display text that is part of the questionnaire.
// The content of the questionnaire is constructed from an ordered, hierarchical collection of items.
type QuestionnaireItem struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	LinkId            string                          `json:"linkId"`
	Definition        *string                         `json:"definition,omitempty"`
	Code              []Coding                        `json:"code,omitempty"`
	Prefix            *string                         `json:"prefix,omitempty"`
	Text              *string                         `json:"text,omitempty"`
	Type              QuestionnaireItemType           `json:"type"`
	EnableWhen        []QuestionnaireItemEnableWhen   `json:"enableWhen,omitempty"`
	EnableBehavior    *EnableWhenBehavior             `json:"enableBehavior,omitempty"`
	Required          *bool                           `json:"required,omitempty"`
	Repeats           *bool                           `json:"repeats,omitempty"`
	ReadOnly          *bool                           `json:"readOnly,omitempty"`
	MaxLength         *int                            `json:"maxLength,omitempty"`
	AnswerValueSet    *string                         `json:"answerValueSet,omitempty"`
	AnswerOption      []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`
	Initial           []QuestionnaireItemInitial      `json:"initial,omitempty"`
	Item              []QuestionnaireItem             `json:"item,omitempty"`
}

// A constraint indicating that this item should only be enabled (displayed/allow answers to be captured) when the specified condition is true.
// If multiple repetitions of this extension are present, the item should be enabled when the condition for *any* of the repetitions is true.  I.e. treat "enableWhen"s as being joined by an "or" clause.  This element is a modifier because if enableWhen is present for an item, "required" is ignored unless one of the enableWhen conditions is met. When an item is disabled, all of its descendants are disabled, regardless of what their own enableWhen logic might evaluate to.
type QuestionnaireItemEnableWhen struct {
	ID                *string                   `json:"ID,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Question          string                    `json:"question"`
	Operator          QuestionnaireItemOperator `json:"operator"`
	AnswerBoolean     bool                      `json:"answerBoolean"`
	AnswerDecimal     json.Number               `json:"answerDecimal"`
	AnswerInteger     int                       `json:"answerInteger"`
	AnswerDate        string                    `json:"answerDate"`
	AnswerDateTime    string                    `json:"answerDateTime"`
	AnswerTime        string                    `json:"answerTime"`
	AnswerString      string                    `json:"answerString"`
	AnswerCoding      Coding                    `json:"answerCoding"`
	AnswerQuantity    Quantity                  `json:"answerQuantity"`
	AnswerReference   Reference                 `json:"answerReference"`
}

// One of the permitted answers for a "choice" or "open-choice" question.
// This element can be used when the value set machinery of answerValueSet is deemed too cumbersome or when there's a need to capture possible answers that are not codes.
type QuestionnaireItemAnswerOption struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         string      `json:"valueDate"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueReference    Reference   `json:"valueReference"`
	InitialSelected   *bool       `json:"initialSelected,omitempty"`
}

// One or more values that should be pre-populated in the answer when initially rendering the questionnaire for user input.
// The user is allowed to change the value and override the default (unless marked as read-only). If the user doesn't change the value, then this initial value will be persisted when the QuestionnaireResponse is initially created.  Note that initial values can influence results.  The data type of initial[x] must agree with the item.type, and only repeating items can have more then one initial value.
type QuestionnaireItemInitial struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDecimal      json.Number `json:"valueDecimal"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         string      `json:"valueDate"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueUri          string      `json:"valueUri"`
	ValueAttachment   Attachment  `json:"valueAttachment"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueQuantity     Quantity    `json:"valueQuantity"`
	ValueReference    Reference   `json:"valueReference"`
}

// This function returns resource reference information
func (r Questionnaire) ResourceRef() (string, *string) {
	return "Questionnaire", r.ID
}

// This function returns resource reference information
func (r Questionnaire) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherQuestionnaire Questionnaire

// MarshalJSON marshals the given Questionnaire as JSON into a byte slice
func (r Questionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaire
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaire: OtherQuestionnaire(r),
		ResourceType:       "Questionnaire",
	})
}

// UnmarshalQuestionnaire unmarshals a Questionnaire.
func UnmarshalQuestionnaire(b []byte) (Questionnaire, error) {
	var questionnaire Questionnaire
	if err := json.Unmarshal(b, &questionnaire); err != nil {
		return questionnaire, err
	}
	return questionnaire, nil
}
