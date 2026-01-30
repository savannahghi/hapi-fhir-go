
package fhir430

import (
	"encoding/json"
	"fmt"
)
// BiologicallyDerivedProductCategory is documented here http://hl7.org/fhir/ValueSet/product-category
type BiologicallyDerivedProductCategory int

const (
	BiologicallyDerivedProductCategoryOrgan BiologicallyDerivedProductCategory = iota
	BiologicallyDerivedProductCategoryTissue
	BiologicallyDerivedProductCategoryFluid
	BiologicallyDerivedProductCategoryCells
	BiologicallyDerivedProductCategoryBiologicalAgent
)

func (code BiologicallyDerivedProductCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BiologicallyDerivedProductCategory) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal BiologicallyDerivedProductCategory code `%s`", s)
	}
	switch s {
	case "organ":
		*code = BiologicallyDerivedProductCategoryOrgan
	case "tissue":
		*code = BiologicallyDerivedProductCategoryTissue
	case "fluid":
		*code = BiologicallyDerivedProductCategoryFluid
	case "cells":
		*code = BiologicallyDerivedProductCategoryCells
	case "biologicalAgent":
		*code = BiologicallyDerivedProductCategoryBiologicalAgent
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductCategory code `%s`", s)
	}
	return nil
}
func (code BiologicallyDerivedProductCategory) String() string {
	return code.Code()
}
func (code BiologicallyDerivedProductCategory) Code() string {
	switch code {
	case BiologicallyDerivedProductCategoryOrgan:
		return "organ"
	case BiologicallyDerivedProductCategoryTissue:
		return "tissue"
	case BiologicallyDerivedProductCategoryFluid:
		return "fluid"
	case BiologicallyDerivedProductCategoryCells:
		return "cells"
	case BiologicallyDerivedProductCategoryBiologicalAgent:
		return "biologicalAgent"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductCategory) Display() string {
	switch code {
	case BiologicallyDerivedProductCategoryOrgan:
		return "Organ"
	case BiologicallyDerivedProductCategoryTissue:
		return "Tissue"
	case BiologicallyDerivedProductCategoryFluid:
		return "Fluid"
	case BiologicallyDerivedProductCategoryCells:
		return "Cells"
	case BiologicallyDerivedProductCategoryBiologicalAgent:
		return "BiologicalAgent"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductCategory) Definition() string {
	switch code {
	case BiologicallyDerivedProductCategoryOrgan:
		return "A collection of tissues joined in a structural unit to serve a common function."
	case BiologicallyDerivedProductCategoryTissue:
		return "An ensemble of similar cells and their extracellular matrix from the same origin that together carry out a specific function."
	case BiologicallyDerivedProductCategoryFluid:
		return "Body fluid."
	case BiologicallyDerivedProductCategoryCells:
		return "Collection of cells."
	case BiologicallyDerivedProductCategoryBiologicalAgent:
		return "Biological agent of unspecified type."
	}
	return "<unknown>"
}
