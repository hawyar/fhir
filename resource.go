package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func CreatePatient(r *http.Request) (fhir.Patient, error) {
	var patient fhir.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)

	if err != nil {
		return patient, err
	}

	id := NewID()

	patient.Id = &id

	sys := "https://example.com/"
	identifier := fhir.Identifier{
		System: &sys,
		Value:  &id,
	}

	var identifiers []fhir.Identifier

	patient.Identifier = append(identifiers, identifier)

	return patient, nil
}

func CreateProcedure(r *http.Request) (fhir.Procedure, error) {
	var procedure fhir.Procedure

	err := json.NewDecoder(r.Body).Decode(&procedure)

	if err != nil {
		fmt.Println(err)
		return procedure, err
	}
	id := NewID()
	procedure.Id = &id

	sys := "https://example.com/"

	identifier := fhir.Identifier{
		System: &sys,
		Value:  &id,
	}

	var identifiers []fhir.Identifier

	procedure.Identifier = append(identifiers, identifier)
	
	// just for demo purposes
	code := "http://demo.info/"
	system := "http://demo.info/"
	display := "Procedure"

	singleCoding := fhir.Coding{
		System:  &system,
		Code:    &code,
		Display: &display,
	}

	procedure.Category = &fhir.CodeableConcept{
		Coding: []fhir.Coding{singleCoding},
	}

	bilboa := "Bilboa Medical Hospital"
	locId := "1"

	location := fhir.Location{
		Id:   &locId,
		Name: &bilboa,
	}

	procedure.Location = &fhir.Reference{
		Reference: location.Id,
		Display:   &bilboa,
	}
	return procedure, nil
}

func GetCapabilityStatement(_ *http.Request) (fhir.CapabilityStatement, error) {
	var capabilityStatement fhir.CapabilityStatement

	url := "http://127.0.0.1:8080/v1/metadata"
	title := "Capability Statement for FHIR Server"
	purpose := "Main EHR capability statement, published for contracting and operational support"
	name := "fhir-test-server"
	publisher := "Consensus Networks"
	experimental := true
	version := "1.0.0"

	capabilityStatement = fhir.CapabilityStatement{
		Name:         &name,
		Id:           nil,
		Url:          &url,
		Purpose:      &purpose,
		Title:        &title,
		FhirVersion:  fhir.FHIRVersion4_0_1,
		Experimental: &experimental,
		Publisher:    &publisher,
		Status:       fhir.PublicationStatusDraft,
		Date:         time.Now().Format(time.RFC3339),
		Kind:         fhir.CapabilityStatementKindCapability,
		Software: &fhir.CapabilityStatementSoftware{
			Name:    "FHIR Test Server",
			Version: &version,
		},
		Format: []string{"application/json+fhir", "application/xml+fhir"},
	}

	return capabilityStatement, nil
}
