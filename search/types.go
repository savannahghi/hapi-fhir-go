package search

import (
	"fmt"
)

type SearchModeEnum string

const (
	Include    SearchModeEnum = "include"
	RevInclude SearchModeEnum = "revinclude"
)

type SearchParams struct {
	Type           SearchModeEnum
	TargetResource string
	SearchField    string
}

type InlineSearchInput struct {
	BaseResource string
	ResourceID   string
	Params       []SearchParams
}

type SearchBuilder struct {
	Input InlineSearchInput
}

// Validate performs structural validation on the InlineSearchInput and ensures
// that all required fields needed to build a valid FHIR search query are present.
//
// It returns an error describing the first validation failure encountered.
// If all fields are valid, Validate returns nil.
func (i InlineSearchInput) Validate() error {
	if i.BaseResource == "" {
		return fmt.Errorf("base resource is required")
	}
	if i.ResourceID == "" {
		return fmt.Errorf("resource id is required")
	}
	
	if len(i.Params) <= 0 {
		return fmt.Errorf("missing search parameters")
	}
	
	for _, item := range i.Params {
		if item.SearchField == "" {
			return fmt.Errorf("search field is required")
		}
		if item.TargetResource == "" {
			return fmt.Errorf("target resource is required")
		}
		if item.Type == "" {
			return fmt.Errorf("missing search type")
		}
	}

	return nil
}
