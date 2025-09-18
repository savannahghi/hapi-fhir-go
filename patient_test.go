package hapifhirgo_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	hapifhirgo "github.com/savannahghi/hapi-fhir-go"
	r5 "github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
	"github.com/savannahghi/scalarutils"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupHAPIFHIRTestContainer(t *testing.T) (string, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "hapiproject/hapi:latest",
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/fhir/metadata").WithStartupTimeout(2 * time.Minute),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %s", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %s", err)
	}

	port, err := container.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatalf("Failed to get container port: %s", err)
	}

	baseURL := fmt.Sprintf("http://%s:%s/fhir", host, port.Port())

	cleanup := func() {
		container.Terminate(ctx)
	}

	return baseURL, cleanup
}

func fakePatient() *r5.Patient {
	id := gofakeit.UUID()
	system := "test"
	version := "0.0.1"
	userSelected := false
	active := true
	phoneSystem := r5.ContactPointSystemEnumPhone
	use := r5.ContactPointUseEnumHome
	rank := int64(1)
	phone := gofakeit.Phone()
	date, err := scalarutils.NewDate(12, 12, 2000)
	if err != nil {
		log.Fatal(err)
	}
	male := r5.PatientGenderEnumMale
	maleContact := r5.PatientContactGenderEnumMale
	address := gofakeit.Address()
	addrUse := r5.AddressUseEnumHome
	postalAddrType := r5.AddressTypeEnumPostal
	name := gofakeit.Name()
	nameUse := r5.HumanNameUseEnumOfficial

	creation := scalarutils.DateTime("2020-09-24T18:02:38.661033Z")
	typeCode := gofakeit.UUID()

	patient := r5.Patient{
		ID: &id,
		Identifier: []*r5.Identifier{
			{
				ID:  &id,
				Use: r5.IdentifierUseEnumOfficial,
				Type: r5.CodeableConcept{
					Text: "MR",
					Coding: []*r5.Coding{
						{
							System:       &system,
							Version:      &version,
							Code:         &typeCode,
							Display:      id,
							UserSelected: &userSelected,
						},
					},
				},
				System:   &system,
				Value:    id,
				Assigner: &r5.Reference{},
			},
		},
		Active: &active,
		Name: []*r5.HumanName{
			{
				Given:  []*string{&name},
				Family: &name,
				Use:    nameUse,
				Text:   name,
			},
		},
		Telecom: []*r5.ContactPoint{
			{
				System: &phoneSystem,
				Use:    &use,
				Rank:   &rank,
				Value:  &phone,
			},
		},
		Gender:    &male,
		BirthDate: date,
		Address: []*r5.Address{
			{
				Use:  &addrUse,
				Type: &postalAddrType,
				Line: []*string{&address.Address},
				Text: address.Address,
			},
		},
		Photo: []*r5.Attachment{
			{
				ID:       &id,
				Creation: &creation,
			},
		},
		Contact: []*r5.PatientContact{
			{
				ID:           new(string),
				Relationship: []*r5.CodeableConcept{},
				Name: &r5.HumanName{
					Given:  []*string{&name},
					Family: &name,
					Use:    nameUse,
					Text:   name,
				},
				Telecom: []*r5.ContactPoint{
					{
						System: &phoneSystem,
						Use:    &use,
						Rank:   &rank,
						Value:  &phone,
					},
				},
				Address: &r5.Address{

					Use:  &addrUse,
					Type: &postalAddrType,
					Line: []*string{&address.Address},
					Text: address.Address,
				},
				Gender: &maleContact,
			},
		},
	}

	return &patient
}

func TestClient_CreateFHIRPatient(t *testing.T) {
	baseURL, cleanup := setupHAPIFHIRTestContainer(t)
	defer cleanup()

	type args struct {
		ctx   context.Context
		input *r5.Patient
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Create a valid patient",
			args: args{
				ctx:   context.Background(),
				input: fakePatient(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := hapifhirgo.NewClient(baseURL)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			err = client.CreateFHIRResource(tt.args.ctx, "Patient", map[string]interface{}{}, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateFHIRResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
