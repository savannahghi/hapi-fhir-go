package search_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/savannahghi/hapi-fhir-go/search"
)

func TestInlineSearchInput_BuildQueryPath(t *testing.T) {
	id := gofakeit.UUID()
	type args struct {
		input search.InlineSearchInput
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case: build simple query path successful",
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   id,
					Params: []search.SearchParams{
						{
							Type:           "include",
							TargetResource: "ServiceRequest",
							SearchField:    "performer",
						},
					},
				},
			},
			want: fmt.Sprintf("ServiceRequest/?_id=%s&_include=ServiceRequest%%3Aperformer", id),
		},
		{
			name: "Happy case: build multiple query successful",
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   id,
					Params: []search.SearchParams{
						{
							Type:           "include",
							TargetResource: "ServiceRequest",
							SearchField:    "performer",
						},
						{
							Type:           "revinclude",
							TargetResource: "Observation",
							SearchField:    "subject",
						},
						{
							Type:           "revinclude",
							TargetResource: "Document",
							SearchField:    "owner",
						},
					},
				},
			},
			want: fmt.Sprintf("ServiceRequest/?_id=%s&_include=ServiceRequest%%3Aperformer&_revinclude=Observation%%3Asubject&_revinclude=Document%%3Aowner", id),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.input.BuildQueryPath()
			if got != tt.want {
				t.Errorf("InlineSearchInput.BuildQueryPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchBuilder_BuildPath(t *testing.T) {
	id := gofakeit.UUID()
	type args struct {
		input    search.SearchBuilder
		id       string
		resource string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy case: build path successful",
			args: args{
				input:    search.SearchBuilder{},
				id:       id,
				resource: "ServiceRequest",
			},
			want: fmt.Sprintf("ServiceRequest/?_id=%s", id),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := search.NewSearchBuilder(tt.args.resource, tt.args.id)
			got := s.BuildPath()
			if got != tt.want {
				t.Errorf("BuildPath() = %v, want %v", got, tt.want)
			}
		})
	}
}