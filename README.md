## fhir

> FHIR server for testing purposes

This is an implementation of [FHIR Version R4](http://hl7.org/fhir/R4/index.html) - a high level specification for healthcare data exchange.


## Prerequisites

 - [Go](https://golang.org/doc/install)
 - [Docker](https://www.docker.com/community-edition)
 - [Redis](https://redis.io/topics/quickstart)
 

## Setup
1. Clone repo
```bash
git clone https://github.com/hawyar/fhir
```

2. Download dependencies

```bash
go mod download
```

3. Start docker containers in background

```bash
docker-compose up -d
```
#### Now the FHIR server is available at [`http://localhost:8080/`](http://127.0.0.1:8080/)




## Usage

Make sure the server is running.

```bash
curl http://localhost:8080/ping
```

### FHIR Server
The FHIR server provides a REST API to access the data. It supports XML and JSON output formats.

### Base Url

The base url is the root of the FHIR server. After that you can access the resources by adding the name as a path. Also notice the upper case resource path names.

**Base Url**: `http://localhost:8080/v1`


For example to access the [`Patient`](http://hl7.org/fhir/R4/patient.html): 

**Patient**: `http://localhost:8080/v1/Patient`

or [`CapabilityStatement`](http://hl7.org/fhir/R4/capabilitystatement.html):

**CapabilityStatement**: `http://localhost:8080/v1/metadata`


### Format

By default the server returns in JSON format. You can change that by adding the `_format` parameter to the url.
Supported formats are:


Patient in xml format: `http://localhost:8080/v1/Patient?_format=xml`

Patient in json format: `http://localhost:8080/v1/Patient?_format=json`


### Other examples:

Get the capabilities statement of the server:

```bash
curl http://127.0.0.1:8080/v1/metadata
```

Create a new patient

```bash
curl -X POST -H "Content-Type: application/json" -d '{"resourceType": "Patient", "name": [{"given": ["John"], "family": "Doe"}]}' http://localhost:8080/v1/Patient
```

### Supported Resources (wip)
- Patient
- CapabilityStatement