package search_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/savannahghi/hapi-fhir-go/search"
)

func TestInlineSearchInput_Validate(t *testing.T) {
	type args struct {
		input search.InlineSearchInput
	}
	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{
			name:    "Happy case: valid inputs",
			wantErr: false,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   gofakeit.UUID(),
					Params: []search.SearchParams{
						{
							Type:           search.Include,
							TargetResource: "ServiceRequest",
							SearchField:    "performer",
						},
					},
				},
			},
		},
		{
			name:    "Sad case: missing base resource inputs",
			wantErr: true,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "",
					ResourceID:   gofakeit.UUID(),
					Params: []search.SearchParams{
						{
							Type:           search.Include,
							TargetResource: "ServiceRequest",
							SearchField:    "performer",
						},
					},
				},
			},
		},
		{
			name:    "Sad case: missing resource id",
			wantErr: true,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   "",
					Params: []search.SearchParams{
						{
							Type:           search.Include,
							TargetResource: "ServiceRequest",
							SearchField:    "performer",
						},
					},
				},
			},
		},
		{
			name:    "Sad case: missing params",
			wantErr: true,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   gofakeit.UUID(),
					Params:       []search.SearchParams{},
				},
			},
		},
		{
			name:    "Sad case: missing target resource",
			wantErr: true,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   gofakeit.UUID(),
					Params: []search.SearchParams{
						{
							Type:           search.Include,
							TargetResource: "",
							SearchField:    "performer",
						},
					},
				},
			},
		},
		{
			name:    "Sad case: missing search field",
			wantErr: true,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   gofakeit.UUID(),
					Params: []search.SearchParams{
						{
							Type:           search.Include,
							TargetResource: "ServiceRequest",
							SearchField:    "",
						},
					},
				},
			},
		},
		{
			name:    "Sad case: missing search type",
			wantErr: true,
			args: args{
				input: search.InlineSearchInput{
					BaseResource: "ServiceRequest",
					ResourceID:   gofakeit.UUID(),
					Params: []search.SearchParams{
						{
							TargetResource: "ServiceRequest",
							SearchField:    "performer",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("InlineSearchInput.Validate() = %v want = %v", err, tt.wantErr)
			}
		})
	}
}
