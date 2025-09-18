
package fhir430

import "encoding/json"
// SubstanceNucleicAcid is documented here http://hl7.org/fhir/StructureDefinition/SubstanceNucleicAcid
// Nucleic acids are defined by three distinct elements: the base, sugar and linkage. Individual substance/moiety IDs will be created for each of these elements. The nucleotide sequence will be always entered in the 5’-3’ direction.
type SubstanceNucleicAcid struct {
	ID                  *string                       `json:"ID,omitempty"`
	Meta                *Meta                         `json:"meta,omitempty"`
	ImplicitRules       *string                       `json:"implicitRules,omitempty"`
	Language            *string                       `json:"language,omitempty"`
	Text                *Narrative                    `json:"text,omitempty"`
	Contained           []json.RawMessage             `json:"contained,omitempty"`
	Extension           []Extension                   `json:"extension,omitempty"`
	ModifierExtension   []Extension                   `json:"modifierExtension,omitempty"`
	SequenceType        *CodeableConcept              `json:"sequenceType,omitempty"`
	NumberOfSubunits    *int                          `json:"numberOfSubunits,omitempty"`
	AreaOfHybridisation *string                       `json:"areaOfHybridisation,omitempty"`
	OligoNucleotideType *CodeableConcept              `json:"oligoNucleotideType,omitempty"`
	Subunit             []SubstanceNucleicAcidSubunit `json:"subunit,omitempty"`
}

// Subunits are listed in order of decreasing length; sequences of the same length will be ordered by molecular weight; subunits that have identical sequences will be repeated multiple times.
type SubstanceNucleicAcidSubunit struct {
	ID                 *string                              `json:"ID,omitempty"`
	Extension          []Extension                          `json:"extension,omitempty"`
	ModifierExtension  []Extension                          `json:"modifierExtension,omitempty"`
	Subunit            *int                                 `json:"subunit,omitempty"`
	Sequence           *string                              `json:"sequence,omitempty"`
	Length             *int                                 `json:"length,omitempty"`
	SequenceAttachment *Attachment                          `json:"sequenceAttachment,omitempty"`
	FivePrime          *CodeableConcept                     `json:"fivePrime,omitempty"`
	ThreePrime         *CodeableConcept                     `json:"threePrime,omitempty"`
	Linkage            []SubstanceNucleicAcidSubunitLinkage `json:"linkage,omitempty"`
	Sugar              []SubstanceNucleicAcidSubunitSugar   `json:"sugar,omitempty"`
}

// The linkages between sugar residues will also be captured.
type SubstanceNucleicAcidSubunitLinkage struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Connectivity      *string     `json:"connectivity,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Name              *string     `json:"name,omitempty"`
	ResidueSite       *string     `json:"residueSite,omitempty"`
}

// 5.3.6.8.1 Sugar ID (Mandatory).
type SubstanceNucleicAcidSubunitSugar struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Name              *string     `json:"name,omitempty"`
	ResidueSite       *string     `json:"residueSite,omitempty"`
}

// This function returns resource reference information
func (r SubstanceNucleicAcid) ResourceRef() (string, *string) {
	return "SubstanceNucleicAcid", r.ID
}

// This function returns resource reference information
func (r SubstanceNucleicAcid) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSubstanceNucleicAcid SubstanceNucleicAcid

// MarshalJSON marshals the given SubstanceNucleicAcid as JSON into a byte slice
func (r SubstanceNucleicAcid) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceNucleicAcid
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceNucleicAcid: OtherSubstanceNucleicAcid(r),
		ResourceType:              "SubstanceNucleicAcid",
	})
}

// UnmarshalSubstanceNucleicAcid unmarshals a SubstanceNucleicAcid.
func UnmarshalSubstanceNucleicAcid(b []byte) (SubstanceNucleicAcid, error) {
	var substanceNucleicAcid SubstanceNucleicAcid
	if err := json.Unmarshal(b, &substanceNucleicAcid); err != nil {
		return substanceNucleicAcid, err
	}
	return substanceNucleicAcid, nil
}
