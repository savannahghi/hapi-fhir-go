package hapifhirgo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/savannahghi/hapi-fhir-go/models"
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

func fakePatient() *models.FHIRPatient {
	id := gofakeit.UUID()
	system := scalarutils.URI("test")
	version := "0.0.1"
	userSelected := false
	active := true
	phoneSystem := models.ContactPointSystemEnumPhone
	use := models.ContactPointUseEnumHome
	rank := int64(1)
	phone := gofakeit.Phone()
	date, err := scalarutils.NewDate(12, 12, 2000)
	if err != nil {
		log.Fatal(err)
	}
	male := models.PatientGenderEnumMale
	maleContact := models.PatientContactGenderEnumMale
	address := gofakeit.Address()
	addrUse := models.AddressUseEnumHome
	postalAddrType := models.AddressTypeEnumPostal
	name := gofakeit.Name()
	nameUse := models.HumanNameUseEnumOfficial

	creation := scalarutils.DateTime("2020-09-24T18:02:38.661033Z")
	typeCode := gofakeit.UUID()

	patient := models.FHIRPatient{
		ID: &id,
		Identifier: []*models.FHIRIdentifier{
			{
				ID:  &id,
				Use: models.IdentifierUseEnumOfficial,
				Type: models.FHIRCodeableConcept{
					Text: "MR",
					Coding: []*models.FHIRCoding{
						{
							System:       &system,
							Version:      &version,
							Code:         (*scalarutils.Code)(&typeCode),
							Display:      id,
							UserSelected: &userSelected,
						},
					},
				},
				System:   &system,
				Value:    id,
				Assigner: &models.FHIRReference{},
			},
		},
		Active: &active,
		Name: []*models.FHIRHumanName{
			{
				Given:  []*string{&name},
				Family: &name,
				Use:    nameUse,
				Text:   name,
			},
		},
		Telecom: []*models.FHIRContactPoint{
			{
				System: &phoneSystem,
				Use:    &use,
				Rank:   &rank,
				Value:  &phone,
			},
		},
		Gender:    &male,
		BirthDate: date,
		Address: []*models.FHIRAddress{
			{
				Use:  &addrUse,
				Type: &postalAddrType,
				Line: []*string{&address.Address},
				Text: address.Address,
			},
		},
		Photo: []*models.FHIRAttachment{
			{
				ID:       &id,
				Creation: &creation,
			},
		},
		Contact: []*models.FHIRPatientContact{
			{
				ID:           new(string),
				Relationship: []*models.FHIRCodeableConcept{},
				Name: &models.FHIRHumanName{
					Given:  []*string{&name},
					Family: &name,
					Use:    nameUse,
					Text:   name,
				},
				Telecom: []*models.FHIRContactPoint{
					{
						System: &phoneSystem,
						Use:    &use,
						Rank:   &rank,
						Value:  &phone,
					},
				},
				Address: &models.FHIRAddress{

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

	gender := models.PatientGenderEnum("invalid")

	type args struct {
		ctx   context.Context
		input *models.FHIRPatient
	}
	tests := []struct {
		name    string
		args    args
		want    *models.PatientPayload
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
		{
			name: "Sad Case - Create an invalid patient",
			args: args{
				ctx: context.Background(),
				input: &models.FHIRPatient{
					Gender: &gender,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL: baseURL,
				HTTP:    &http.Client{Timeout: 10 * time.Second},
			}
			_, err := c.CreateFHIRPatient(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateFHIRPatient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
