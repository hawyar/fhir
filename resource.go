package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func NewCapabilityStatement() (fhir.CapabilityStatement, error) {
	stmt := fhir.CapabilityStatement{}
	return stmt, nil
}

func CreatePatient(r *http.Request) (fhir.Patient, error) {
	var patient fhir.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)

	if err != nil {
		return patient, err
	}

	if patient.Id == nil {
		id := NewID()
		patient.Id = &id
	}

	sys := "https://example.com/"
	identifier := fhir.Identifier{
		System: &sys,
		Value:  patient.Id,
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

	if procedure.Subject.Reference == nil {
		ref := NewID()
		procedure.Subject.Reference = &ref
	}

	bilboa := "Bilboa Medical Hospital"
	locId := "12"

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

// NewMeta instantiates a new Meta object carried by each resource. See https://www.hl7.org/fhir/resource.html#metadata
func NewMeta() (fhir.Meta, error) {
	meta := fhir.Meta{}
	return meta, nil
}

func NewObservation(r *http.Request) (fhir.Observation, error) {
	var observation fhir.Observation

	err := json.NewDecoder(r.Body).Decode(&observation)

	if err != nil {
		return observation, err
	}

	// id := NewID()

	// observation.Id = &id

	// if observation.Subject.Reference == nil {
	// 	ref := NewID()
	// 	observation.Subject.Reference = &ref
	// }

	return observation, nil
}

func NewLocation(r *http.Request) (fhir.Location, error) {
	var location fhir.Location

	err := json.NewDecoder(r.Body).Decode(&location)

	if err != nil {
		return location, err
	}

	id := NewID()

	location.Id = &id

	return location, nil
}

// NewBasic created a special type of resource. It doesn't correspond to a specific pre-defined HL7 concept.
// Instead, it's a placeholder for any resource-like concept that isn't already defined in the HL7 specification.
func NewBasic(r *http.Request) (fhir.Basic, error) {
	var basic fhir.Basic

	err := json.NewDecoder(r.Body).Decode(&basic)

	if err != nil {
		return basic, err
	}

	id := NewID()
	createdAt := time.Now().Format(time.RFC3339)

	basic.Id = &id
	basic.Created = &createdAt

	return basic, nil
}

// NewOrganization creates a new Organization resource.
// The Organization resource is used for collections of people that have come together to achieve an objective.
func NewOrganization(r *http.Request) (fhir.Organization, error) {
	var organization fhir.Organization

	err := json.NewDecoder(r.Body).Decode(&organization)

	if err != nil {
		return organization, err
	}

	orgName := "FHIR Hospital"
	id := NewID()
	active := true

	organization.Name = &orgName
	organization.Id = &id
	organization.Active = &active
	organization.Alias = append(organization.Alias, orgName, "Demo", "FHIR", "Hospital")

	return organization, nil
}
