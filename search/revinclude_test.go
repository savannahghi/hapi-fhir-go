package search_test

import(
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/savannahghi/hapi-fhir-go/search"
	"github.com/stretchr/testify/assert"
)

func TestSearchBuilder_RevInclude(t *testing.T) {
	baseResource := "ServiceRequest"
	resourceID := gofakeit.UUID()
	targetResource := "ServiceRequest"
	searchFiled := "performer"
	type args struct {
		baseResource   string
		resourceID     string
		targetResource string
		searchFiled    string
	}
	tests := []struct {
		name string
		args args
		want *search.SearchBuilder
	}{
		{
			name: "Happy case",
			args: args{
				baseResource:   baseResource,
				resourceID:     resourceID,
				targetResource: targetResource,
				searchFiled:    searchFiled,
			},
			want: &search.SearchBuilder{
				Input: search.InlineSearchInput{
					BaseResource: baseResource,
					ResourceID:   resourceID,
					Params: []search.SearchParams{
						{
							Type:           search.RevInclude,
							TargetResource: targetResource,
							SearchField:    searchFiled,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := search.NewSearchBuilder(tt.args.baseResource, tt.args.resourceID)
			got := s.RevInclude(tt.args.targetResource, tt.args.searchFiled)
			assert.Equal(t, tt.want.Input, got.Input)
		})
	}
}
