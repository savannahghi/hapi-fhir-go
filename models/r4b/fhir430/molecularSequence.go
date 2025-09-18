
package fhir430

import "encoding/json"
// MolecularSequence is documented here http://hl7.org/fhir/StructureDefinition/MolecularSequence
// Raw data describing a biological sequence.
type MolecularSequence struct {
	ID                *string                             `json:"ID,omitempty"`
	Meta              *Meta                               `json:"meta,omitempty"`
	ImplicitRules     *string                             `json:"implicitRules,omitempty"`
	Language          *string                             `json:"language,omitempty"`
	Text              *Narrative                          `json:"text,omitempty"`
	Contained         []json.RawMessage                   `json:"contained,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                        `json:"identifier,omitempty"`
	Type              *string                             `json:"type,omitempty"`
	CoordinateSystem  int                                 `json:"coordinateSystem"`
	Patient           *Reference                          `json:"patient,omitempty"`
	Specimen          *Reference                          `json:"specimen,omitempty"`
	Device            *Reference                          `json:"device,omitempty"`
	Performer         *Reference                          `json:"performer,omitempty"`
	Quantity          *Quantity                           `json:"quantity,omitempty"`
	ReferenceSeq      *MolecularSequenceReferenceSeq      `json:"referenceSeq,omitempty"`
	Variant           []MolecularSequenceVariant          `json:"variant,omitempty"`
	ObservedSeq       *string                             `json:"observedSeq,omitempty"`
	Quality           []MolecularSequenceQuality          `json:"quality,omitempty"`
	ReadCoverage      *int                                `json:"readCoverage,omitempty"`
	Repository        []MolecularSequenceRepository       `json:"repository,omitempty"`
	Pointer           []Reference                         `json:"pointer,omitempty"`
	StructureVariant  []MolecularSequenceStructureVariant `json:"structureVariant,omitempty"`
}

// A sequence that is used as a reference to describe variants that are present in a sequence analyzed.
type MolecularSequenceReferenceSeq struct {
	ID                  *string          `json:"ID,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	Chromosome          *CodeableConcept `json:"chromosome,omitempty"`
	GenomeBuild         *string          `json:"genomeBuild,omitempty"`
	Orientation         *string          `json:"orientation,omitempty"`
	ReferenceSeqId      *CodeableConcept `json:"referenceSeqId,omitempty"`
	ReferenceSeqPointer *Reference       `json:"referenceSeqPointer,omitempty"`
	ReferenceSeqString  *string          `json:"referenceSeqString,omitempty"`
	Strand              *string          `json:"strand,omitempty"`
	WindowStart         *int             `json:"windowStart,omitempty"`
	WindowEnd           *int             `json:"windowEnd,omitempty"`
}

// The definition of variant here originates from Sequence ontology ([variant_of](http://www.sequenceontology.org/browser/current_svn/term/variant_of)). This element can represent amino acid or nucleic sequence change(including insertion,deletion,SNP,etc.)  It can represent some complex mutation or segment variation with the assist of CIGAR string.
type MolecularSequenceVariant struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
	ObservedAllele    *string     `json:"observedAllele,omitempty"`
	ReferenceAllele   *string     `json:"referenceAllele,omitempty"`
	Cigar             *string     `json:"cigar,omitempty"`
	VariantPointer    *Reference  `json:"variantPointer,omitempty"`
}

// An experimental feature attribute that defines the quality of the feature in a quantitative way, such as a phred quality score ([SO:0001686](http://www.sequenceontology.org/browser/current_svn/term/SO:0001686)).
type MolecularSequenceQuality struct {
	ID                *string                      `json:"ID,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Type              string                       `json:"type"`
	StandardSequence  *CodeableConcept             `json:"standardSequence,omitempty"`
	Start             *int                         `json:"start,omitempty"`
	End               *int                         `json:"end,omitempty"`
	Score             *Quantity                    `json:"score,omitempty"`
	Method            *CodeableConcept             `json:"method,omitempty"`
	TruthTP           *json.Number                 `json:"truthTP,omitempty"`
	QueryTP           *json.Number                 `json:"queryTP,omitempty"`
	TruthFN           *json.Number                 `json:"truthFN,omitempty"`
	QueryFP           *json.Number                 `json:"queryFP,omitempty"`
	GtFP              *json.Number                 `json:"gtFP,omitempty"`
	Precision         *json.Number                 `json:"precision,omitempty"`
	Recall            *json.Number                 `json:"recall,omitempty"`
	FScore            *json.Number                 `json:"fScore,omitempty"`
	Roc               *MolecularSequenceQualityRoc `json:"roc,omitempty"`
}

// Receiver Operator Characteristic (ROC) Curve  to give sensitivity/specificity tradeoff.
type MolecularSequenceQualityRoc struct {
	ID                *string       `json:"ID,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Score             []int         `json:"score,omitempty"`
	NumTP             []int         `json:"numTP,omitempty"`
	NumFP             []int         `json:"numFP,omitempty"`
	NumFN             []int         `json:"numFN,omitempty"`
	Precision         []json.Number `json:"precision,omitempty"`
	Sensitivity       []json.Number `json:"sensitivity,omitempty"`
	FMeasure          []json.Number `json:"fMeasure,omitempty"`
}

// Configurations of the external repository. The repository shall store target's observedSeq or records related with target's observedSeq.
type MolecularSequenceRepository struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Url               *string     `json:"url,omitempty"`
	Name              *string     `json:"name,omitempty"`
	DatasetId         *string     `json:"datasetId,omitempty"`
	VariantsetId      *string     `json:"variantsetId,omitempty"`
	ReadsetId         *string     `json:"readsetId,omitempty"`
}

// Information about chromosome structure variation.
type MolecularSequenceStructureVariant struct {
	ID                *string                                 `json:"ID,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	VariantType       *CodeableConcept                        `json:"variantType,omitempty"`
	Exact             *bool                                   `json:"exact,omitempty"`
	Length            *int                                    `json:"length,omitempty"`
	Outer             *MolecularSequenceStructureVariantOuter `json:"outer,omitempty"`
	Inner             *MolecularSequenceStructureVariantInner `json:"inner,omitempty"`
}

// Structural variant outer.
type MolecularSequenceStructureVariantOuter struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
}

// Structural variant inner.
type MolecularSequenceStructureVariantInner struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Start             *int        `json:"start,omitempty"`
	End               *int        `json:"end,omitempty"`
}

// This function returns resource reference information
func (r MolecularSequence) ResourceRef() (string, *string) {
	return "MolecularSequence", r.ID
}

// This function returns resource reference information
func (r MolecularSequence) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMolecularSequence MolecularSequence

// MarshalJSON marshals the given MolecularSequence as JSON into a byte slice
func (r MolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMolecularSequence
		ResourceType string `json:"resourceType"`
	}{
		OtherMolecularSequence: OtherMolecularSequence(r),
		ResourceType:           "MolecularSequence",
	})
}

// UnmarshalMolecularSequence unmarshals a MolecularSequence.
func UnmarshalMolecularSequence(b []byte) (MolecularSequence, error) {
	var molecularSequence MolecularSequence
	if err := json.Unmarshal(b, &molecularSequence); err != nil {
		return molecularSequence, err
	}
	return molecularSequence, nil
}
