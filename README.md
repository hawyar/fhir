## FHIR

> FHIR server for testing purposes

The FHIR server provides a REST API to access and exchange healthcare data with other client/servers. The FHIR version implementation is [R4 (v4.0.1)](http://hl7.org/fhir/R4/index.html).

## Prerequisites

-   [Go](https://golang.org/doc/install)
-   [Docker](https://www.docker.com/community-edition)
-   [Redis](https://redis.io/topics/quickstart)


## Usage

###  Setup

1. Clone repo

```bash
git clone https://github.com/
hawyar/fhir
```

2. Start containers in background

```bash
docker-compose build --no-cache && docker-compose up -d
```

### Server
Base URL: `http://localhost:8080/v1/`

The base URL serves as the root of the FHIR server. The **v1** refers to the server's version **NOT** the FHIR spec in use. The FHIR version implementation is R4 (v4.0.1). To access the different resources just append the resource name to the base URL.

### Resources (wip)

-   [`CapabilityStatement`](http://hl7.org/fhir/R4/capabilitystatement.html): `http://localhost:8080/v1/metadata`
-   [`Patient`](http://hl7.org/fhir/R4/patient.html): `http://localhost:8080/v1/Patient`
-   [`Procedure`](http://hl7.org/fhir/R4/patient.html): `http://localhost:8080/v1/Procedure`

## cURL examples

Create patient

```bash
curl -X POST -H "Content-Type: application/json" -d '{"resourceType": "Patient", "name": [{"given": ["John"], "family": "Doe"}]}' http://localhost:8080/v1/Patient
```

Create procedure

```bash
curl -X POST -H "Content-Type: application/json" -d '{"subject":{"reference":"25oYHe8zCfx52wp9S8RKEVjEyTw"}}' http://localhost:8080/v1/Procedure
```
