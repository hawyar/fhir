package main

import (
	"encoding/json"
	"net/http"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func CreatePatient(r *http.Request) (fhir.Patient, error) {
	var patient fhir.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)

	if err != nil {
		return patient, err
	}

	id := NewID()
	sys := "https://example.com/"
	identifier := fhir.Identifier{
		System: &sys,
		Value:  &id,
	}

	var identifiers []fhir.Identifier

	patient.Identifier = append(identifiers, identifier)

	return patient, nil
}

func GetCapabilityStatement(r *http.Request) (fhir.CapabilityStatement, error) {
	var capabilityStatement fhir.CapabilityStatement

	url := "http://127.0.0.1:8080/v1/metadata"
	title := "Capaability Statement"
	name := "v1"

	capabilityStatement = fhir.CapabilityStatement{
		Name:        &name,
		Id:          nil,
		Url:         &url,
		Title:       &title,
		FhirVersion: fhir.FHIRVersion4_0_1,
	}

	err := json.NewDecoder(r.Body).Decode(&capabilityStatement)

	if err != nil {
		return capabilityStatement, err
	}

	return capabilityStatement, nil
}
