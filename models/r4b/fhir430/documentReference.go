
package fhir430

import "encoding/json"
// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
// A reference to a document of any kind for any purpose. Provides metadata about the document so that the document can be discovered and managed. The scope of a document is any seralized object with a mime-type, so includes formal patient centric documents (CDA), cliical notes, scanned paper, and non-patient specific documents like policy text.
type DocumentReference struct {
	ID                *string                      `json:"ID,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Contained         []json.RawMessage            `json:"contained,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier                  `json:"masterIdentifier,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Status            DocumentReferenceStatus      `json:"status"`
	DocStatus         *CompositionStatus           `json:"docStatus,omitempty"`
	Type              *CodeableConcept             `json:"type,omitempty"`
	Category          []CodeableConcept            `json:"category,omitempty"`
	Subject           *Reference                   `json:"subject,omitempty"`
	Date              *string                      `json:"date,omitempty"`
	Author            []Reference                  `json:"author,omitempty"`
	Authenticator     *Reference                   `json:"authenticator,omitempty"`
	Custodian         *Reference                   `json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `json:"content"`
	Context           *DocumentReferenceContext    `json:"context,omitempty"`
}

// Relationships that this document has with other document references that already exist.
// This element is labeled as a modifier because documents that append to other documents are incomplete on their own.
type DocumentReferenceRelatesTo struct {
	ID                *string                  `json:"ID,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `json:"code"`
	Target            Reference                `json:"target"`
}

// The document and format referenced. There may be multiple content element repetitions, each with a different format.
type DocumentReferenceContent struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Attachment        Attachment  `json:"attachment"`
	Format            *Coding     `json:"format,omitempty"`
}

// The clinical context in which the document was prepared.
// These values are primarily added to help with searching for interesting/relevant documents.
type DocumentReferenceContext struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Encounter         []Reference       `json:"encounter,omitempty"`
	Event             []CodeableConcept `json:"event,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	FacilityType      *CodeableConcept  `json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept  `json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference        `json:"sourcePatientInfo,omitempty"`
	Related           []Reference       `json:"related,omitempty"`
}

// This function returns resource reference information
func (r DocumentReference) ResourceRef() (string, *string) {
	return "DocumentReference", r.ID
}

// This function returns resource reference information
func (r DocumentReference) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDocumentReference DocumentReference

// MarshalJSON marshals the given DocumentReference as JSON into a byte slice
func (r DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentReference
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentReference: OtherDocumentReference(r),
		ResourceType:           "DocumentReference",
	})
}

// UnmarshalDocumentReference unmarshals a DocumentReference.
func UnmarshalDocumentReference(b []byte) (DocumentReference, error) {
	var documentReference DocumentReference
	if err := json.Unmarshal(b, &documentReference); err != nil {
		return documentReference, err
	}
	return documentReference, nil
}
