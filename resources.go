package main

type meta struct {
	versionId string `json:"versionId" xml:"versionId"`
	// instant time format is precise and differernt from dateTime
	// YYYY-MM-DDThh:mm:ss.sss+zz:zz (e.g. 2015-02-07T13:28:17.239+02:00 or 2017-01-01T00:00:00Z)
	lastUpdated string `json:"lastUpdated" xml:"lastUpdated"`
	// source may identify another FHIR server, document, message, database, etc. In the provenance resource, this corresponds to Provenance.entity.what[x]
	source   string `json:"source" xml:"source"`
	profile  string `json:"profile" xml:"profile"`
	security string `json:"security" xml:"security"`
	tag      string `json:"tag" xml:"tag"`
}

type baseResource struct {
	resourceType string `json:"name" xml:"name"`
	id           string `json:"id" xml:"id"`
	meta         meta   `json:"meta" xml:"meta"`
	// for tag exmaples see: https://tools.ietf.org/search/bcp47#appendix-A
	language      string `json:"language" xml:"language"`
	implicitRules string `json:"implicitRules" xml:"implicitRules"`
}

type uri string

type software struct {
	name        string `json:"name" xml:"name"`
	version     string `json:"version" xml:"version"`
	releaseDate string `json:"releaseDate" xml:"releaseDate"`
}

type implementation struct {
	description string `json:"description" xml:"description"`
	url         string `json:"url" xml:"url"`
	custodian   string `json:"custodian" xml:"custodian"`
}

type capabilityStatement struct {
	meta                meta           `json:"meta" xml:"meta"`
	url                 uri            `json:"url" xml:"url"`
	version             string         `json:"version" xml:"version"`
	name                string         `json:"name" xml:"name"`
	title               string         `json:"title" xml:"title"`
	status              string         `json:"status" xml:"status"`
	experimental        bool           `json:"experimental" xml:"experimental"`
	publisher           string         `json:"publisher" xml:"publisher"`
	contact             []string       `json:"contact" xml:"contact"`
	description         string         `json:"description" xml:"description"`
	useContext          []string       `json:"useContext" xml:"useContext"`
	jurisdication       []string       `json:"jurisdication" xml:"jurisdication"`
	purpose             string         `json:"purpose" xml:"purpose"`
	copyright           string         `json:"copyright" xml:"copyright"`
	kind                string         `json:"kind" xml:"kind"`
	instantiates        uri            `json:"instantiates" xml:"instantiates"`
	imports             []string       `json:"imports" xml:"imports"`
	software            software       `json:"software" xml:"software"`
	implementation      implementation `json:"implementation" xml:"implementation"`
	fhirVersion         string         `json:"fhirVersion" xml:"fhirVersion"`
	format              []string       `json:"format" xml:"format"`
	patchFormat         []string       `json:"patchFormat" xml:"patchFormat"`
	implementationGuide []string       `json:"implementationGuide" xml:"implementationGuide"`
}
