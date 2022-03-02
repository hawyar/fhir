## FHIR

> FHIR server for testing purposes

This is an implementation of [FHIR Version R4](http://hl7.org/fhir/R4/index.html) - a spec for healthcare data exchange.

## Prerequisites

-   [Go](https://golang.org/doc/install)
-   [Docker](https://www.docker.com/community-edition)
-   [Redis](https://redis.io/topics/quickstart)

## Setup

1. Clone repo

```bash
git clone https://github.com/hawyar/fhir
```

2. Start containers in background

```bash
docker-compose build --no-cache && docker-compose up -d
```

**Server available at:** `http://localhost:8080/v1/`

## Usage

The FHIR server provides a RESTful API to access and exchange healthcare data. The base URL serves as the root of the FHIR server. The `v1` refers to the server's version **NOT** the FHIR spec in use. To access the different resources just append the resource name to the base URL.

**Base URL**: `http://localhost:8080/v1/`

### Currently supported resources:

-   [`Patient`](http://hl7.org/fhir/R4/patient.html): `http://localhost:8080/v1/Patient`
-   [`Procedure`](http://hl7.org/fhir/R4/patient.html): `http://localhost:8080/v1/Procedure`
-   [`CapabilityStatement`](http://hl7.org/fhir/R4/capabilitystatement.html): `http://localhost:8080/v1/metadata`

### Usage

Create a new patient

```bash
curl -X POST -H "Content-Type: application/json" -d '{"resourceType": "Patient", "name": [{"given": ["John"], "family": "Doe"}]}' http://localhost:8080/v1/Patient
```

or a new procedure

```bash
curl -X POST -H "Content-Type: application/json" -d '{"subject":{"reference":"25oYHe8zCfx52wp9S8RKEVjEyTw"}}' http://localhost:8080/v1/Patient
```
