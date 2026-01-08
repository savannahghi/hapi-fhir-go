package fhir430

import "encoding/json"

// Contract is documented here http://hl7.org/fhir/StructureDefinition/Contract
// Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type Contract struct {
	ID                       *string                      `json:"id,omitempty"`
	Meta                     *Meta                        `json:"meta,omitempty"`
	ImplicitRules            *string                      `json:"implicitRules,omitempty"`
	Language                 *string                      `json:"language,omitempty"`
	Text                     *Narrative                   `json:"text,omitempty"`
	Contained                []json.RawMessage            `json:"contained,omitempty"`
	Extension                []Extension                  `json:"extension,omitempty"`
	ModifierExtension        []Extension                  `json:"modifierExtension,omitempty"`
	Identifier               []Identifier                 `json:"identifier,omitempty"`
	Url                      *string                      `json:"url,omitempty"`
	Version                  *string                      `json:"version,omitempty"`
	Status                   *ContractResourceStatusCodes `json:"status,omitempty"`
	LegalState               *CodeableConcept             `json:"legalState,omitempty"`
	InstantiatesCanonical    *Reference                   `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri          *string                      `json:"instantiatesUri,omitempty"`
	ContentDerivative        *CodeableConcept             `json:"contentDerivative,omitempty"`
	Issued                   *string                      `json:"issued,omitempty"`
	Applies                  *Period                      `json:"applies,omitempty"`
	ExpirationType           *CodeableConcept             `json:"expirationType,omitempty"`
	Subject                  []Reference                  `json:"subject,omitempty"`
	Authority                []Reference                  `json:"authority,omitempty"`
	Domain                   []Reference                  `json:"domain,omitempty"`
	Site                     []Reference                  `json:"site,omitempty"`
	Name                     *string                      `json:"name,omitempty"`
	Title                    *string                      `json:"title,omitempty"`
	Subtitle                 *string                      `json:"subtitle,omitempty"`
	Alias                    []string                     `json:"alias,omitempty"`
	Author                   *Reference                   `json:"author,omitempty"`
	Scope                    *CodeableConcept             `json:"scope,omitempty"`
	TopicCodeableConcept     *CodeableConcept             `json:"topicCodeableConcept,omitempty"`
	TopicReference           *Reference                   `json:"topicReference,omitempty"`
	Type                     *CodeableConcept             `json:"type,omitempty"`
	SubType                  []CodeableConcept            `json:"subType,omitempty"`
	ContentDefinition        *ContractContentDefinition   `json:"contentDefinition,omitempty"`
	Term                     []ContractTerm               `json:"term,omitempty"`
	SupportingInfo           []Reference                  `json:"supportingInfo,omitempty"`
	RelevantHistory          []Reference                  `json:"relevantHistory,omitempty"`
	Signer                   []ContractSigner             `json:"signer,omitempty"`
	Friendly                 []ContractFriendly           `json:"friendly,omitempty"`
	Legal                    []ContractLegal              `json:"legal,omitempty"`
	Rule                     []ContractRule               `json:"rule,omitempty"`
	LegallyBindingAttachment *Attachment                  `json:"legallyBindingAttachment,omitempty"`
	LegallyBindingReference  *Reference                   `json:"legallyBindingReference,omitempty"`
}

// Precusory content developed with a focus and intent of supporting the formation a Contract instance, which may be associated with and transformable into a Contract.
type ContractContentDefinition struct {
	ID                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                        `json:"type"`
	SubType           *CodeableConcept                       `json:"subType,omitempty"`
	Publisher         *Reference                             `json:"publisher,omitempty"`
	PublicationDate   *string                                `json:"publicationDate,omitempty"`
	PublicationStatus ContractResourcePublicationStatusCodes `json:"publicationStatus"`
}

// One or more Contract Provisions, which may be related and conveyed as a group, and may contain nested groups.
type ContractTerm struct {
	ID                   *string                     `json:"id,omitempty"`
	Extension            []Extension                 `json:"extension,omitempty"`
	ModifierExtension    []Extension                 `json:"modifierExtension,omitempty"`
	Identifier           *Identifier                 `json:"identifier,omitempty"`
	Issued               *string                     `json:"issued,omitempty"`
	Applies              *Period                     `json:"applies,omitempty"`
	TopicCodeableConcept *CodeableConcept            `json:"topicCodeableConcept,omitempty"`
	TopicReference       *Reference                  `json:"topicReference,omitempty"`
	Type                 *CodeableConcept            `json:"type,omitempty"`
	SubType              *CodeableConcept            `json:"subType,omitempty"`
	Text                 *string                     `json:"text,omitempty"`
	SecurityLabel        []ContractTermSecurityLabel `json:"securityLabel,omitempty"`
	Offer                ContractTermOffer           `json:"offer"`
	Asset                []ContractTermAsset         `json:"asset,omitempty"`
	Action               []ContractTermAction        `json:"action,omitempty"`
	Group                []ContractTerm              `json:"group,omitempty"`
}

// Security labels that protect the handling of information about the term and its elements, which may be specifically identified..
type ContractTermSecurityLabel struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Number            []int       `json:"number,omitempty"`
	Classification    Coding      `json:"classification"`
	Category          []Coding    `json:"category,omitempty"`
	Control           []Coding    `json:"control,omitempty"`
}

// The matter of concern in the context of this provision of the agrement.
type ContractTermOffer struct {
	ID                  *string                   `json:"id,omitempty"`
	Extension           []Extension               `json:"extension,omitempty"`
	ModifierExtension   []Extension               `json:"modifierExtension,omitempty"`
	Identifier          []Identifier              `json:"identifier,omitempty"`
	Party               []ContractTermOfferParty  `json:"party,omitempty"`
	Topic               *Reference                `json:"topic,omitempty"`
	Type                *CodeableConcept          `json:"type,omitempty"`
	Decision            *CodeableConcept          `json:"decision,omitempty"`
	DecisionMode        []CodeableConcept         `json:"decisionMode,omitempty"`
	Answer              []ContractTermOfferAnswer `json:"answer,omitempty"`
	Text                *string                   `json:"text,omitempty"`
	LinkId              []string                  `json:"linkId,omitempty"`
	SecurityLabelNumber []int                     `json:"securityLabelNumber,omitempty"`
}

// Offer Recipient.
type ContractTermOfferParty struct {
	ID                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Reference         []Reference     `json:"reference"`
	Role              CodeableConcept `json:"role"`
}

// Response to offer text.
type ContractTermOfferAnswer struct {
	ID                *string     `json:"id,omitempty"`
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

// Contract Term Asset List.
type ContractTermAsset struct {
	ID                  *string                       `json:"id,omitempty"`
	Extension           []Extension                   `json:"extension,omitempty"`
	ModifierExtension   []Extension                   `json:"modifierExtension,omitempty"`
	Scope               *CodeableConcept              `json:"scope,omitempty"`
	Type                []CodeableConcept             `json:"type,omitempty"`
	TypeReference       []Reference                   `json:"typeReference,omitempty"`
	Subtype             []CodeableConcept             `json:"subtype,omitempty"`
	Relationship        *Coding                       `json:"relationship,omitempty"`
	Context             []ContractTermAssetContext    `json:"context,omitempty"`
	Condition           *string                       `json:"condition,omitempty"`
	PeriodType          []CodeableConcept             `json:"periodType,omitempty"`
	Period              []Period                      `json:"period,omitempty"`
	UsePeriod           []Period                      `json:"usePeriod,omitempty"`
	Text                *string                       `json:"text,omitempty"`
	LinkId              []string                      `json:"linkId,omitempty"`
	Answer              []ContractTermOfferAnswer     `json:"answer,omitempty"`
	SecurityLabelNumber []int                         `json:"securityLabelNumber,omitempty"`
	ValuedItem          []ContractTermAssetValuedItem `json:"valuedItem,omitempty"`
}

// Circumstance of the asset.
type ContractTermAssetContext struct {
	ID                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Reference         *Reference        `json:"reference,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Text              *string           `json:"text,omitempty"`
}

// Contract Valued Item List.
type ContractTermAssetValuedItem struct {
	ID                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `json:"entityReference,omitempty"`
	Identifier            *Identifier      `json:"identifier,omitempty"`
	EffectiveTime         *string          `json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `json:"quantity,omitempty"`
	UnitPrice             *Money           `json:"unitPrice,omitempty"`
	Factor                *json.Number     `json:"factor,omitempty"`
	Points                *json.Number     `json:"points,omitempty"`
	Net                   *Money           `json:"net,omitempty"`
	Payment               *string          `json:"payment,omitempty"`
	PaymentDate           *string          `json:"paymentDate,omitempty"`
	Responsible           *Reference       `json:"responsible,omitempty"`
	Recipient             *Reference       `json:"recipient,omitempty"`
	LinkId                []string         `json:"linkId,omitempty"`
	SecurityLabelNumber   []int            `json:"securityLabelNumber,omitempty"`
}

// An actor taking a role in an activity for which it can be assigned some degree of responsibility for the activity taking place.
// Several agents may be associated (i.e. has some responsibility for an activity) with an activity and vice-versa.For example, in cases of actions initiated by one user for other users, or in events that involve more than one user, hardware device, software, or system process. However, only one user may be the initiator/requestor for the event.
type ContractTermAction struct {
	ID                  *string                     `json:"id,omitempty"`
	Extension           []Extension                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                 `json:"modifierExtension,omitempty"`
	DoNotPerform        *bool                       `json:"doNotPerform,omitempty"`
	Type                CodeableConcept             `json:"type"`
	Subject             []ContractTermActionSubject `json:"subject,omitempty"`
	Intent              CodeableConcept             `json:"intent"`
	LinkId              []string                    `json:"linkId,omitempty"`
	Status              CodeableConcept             `json:"status"`
	Context             *Reference                  `json:"context,omitempty"`
	ContextLinkId       []string                    `json:"contextLinkId,omitempty"`
	OccurrenceDateTime  *string                     `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                     `json:"occurrenceTiming,omitempty"`
	Requester           []Reference                 `json:"requester,omitempty"`
	RequesterLinkId     []string                    `json:"requesterLinkId,omitempty"`
	PerformerType       []CodeableConcept           `json:"performerType,omitempty"`
	PerformerRole       *CodeableConcept            `json:"performerRole,omitempty"`
	Performer           *Reference                  `json:"performer,omitempty"`
	PerformerLinkId     []string                    `json:"performerLinkId,omitempty"`
	ReasonCode          []CodeableConcept           `json:"reasonCode,omitempty"`
	ReasonReference     []Reference                 `json:"reasonReference,omitempty"`
	Reason              []string                    `json:"reason,omitempty"`
	ReasonLinkId        []string                    `json:"reasonLinkId,omitempty"`
	Note                []Annotation                `json:"note,omitempty"`
	SecurityLabelNumber []int                       `json:"securityLabelNumber,omitempty"`
}

// Entity of the action.
type ContractTermActionSubject struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Reference         []Reference      `json:"reference"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// Parties with legal standing in the Contract, including the principal parties, the grantor(s) and grantee(s), which are any person or organization bound by the contract, and any ancillary parties, which facilitate the execution of the contract such as a notary or witness.
// Signers who are principal parties to the contract are bound by the Contract.activity related to the Contract.topic, and the Contract.term(s), which either extend or restrict the overall action on the topic by, for example, stipulating specific policies or obligations constraining actions, action reason, or agents with respect to some or all of the topic.For example, specifying how policies or obligations shall constrain actions and action reasons permitted or denied on all or a subset of the Contract.topic (e.g., all or a portion of property being transferred by the contract), agents (e.g., who can resell, assign interests, or alter the property being transferred by the contract), actions, and action reasons; or with respect to Contract.terms, stipulating, extending, or limiting the Contract.period of applicability or valuation of items under consideration.
type ContractSigner struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              Coding      `json:"type"`
	Party             Reference   `json:"party"`
	Signature         []Signature `json:"signature"`
}

// The "patient friendly language" versionof the Contract in whole or in parts. "Patient friendly language" means the representation of the Contract and Contract Provisions in a manner that is readily accessible and understandable by a layperson in accordance with best practices for communication styles that ensure that those agreeing to or signing the Contract understand the roles, actions, obligations, responsibilities, and implication of the agreement.
type ContractFriendly struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// List of Legal expressions or representations of this Contract.
type ContractLegal struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// List of Computable Policy Rule Language Representations of this Contract.
type ContractRule struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// This function returns resource reference information
func (r Contract) ResourceRef() (string, *string) {
	return "Contract", r.ID
}

// This function returns resource reference information
func (r Contract) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherContract Contract

// MarshalJSON marshals the given Contract as JSON into a byte slice
func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContract
		ResourceType string `json:"resourceType"`
	}{
		OtherContract: OtherContract(r),
		ResourceType:  "Contract",
	})
}

// UnmarshalContract unmarshals a Contract.
func UnmarshalContract(b []byte) (Contract, error) {
	var contract Contract
	if err := json.Unmarshal(b, &contract); err != nil {
		return contract, err
	}
	return contract, nil
}
