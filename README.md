## fhir

> FHIR server for testing purposes 💅

FHIR Server following the on the [FHIR Version R4](http://hl7.org/fhir/R4/index.html) for testing.

## Usage

Clone

```bash
git clone https://github.com/hawyar/fhir
```

Sync dependencies

```bash
go mod tidy
```

Start the server. Runs on port `4141`

```bash
make server
```

Now start the reverse proxy (using Caddy) on `8080`

```bash
make proxy
```

Before creating any resource check if the server is running

```bash
curl http://127.0.0.1:8080/ping
```

### FHIR Resources

Create a new patient

```bash
curl -X POST -H "Content-Type: application/json" -d '{"resourceType": "Patient", "name": [{"given": ["John"], "family": "Doe"}]}' http://localhost:4141/v1/Patient
```
