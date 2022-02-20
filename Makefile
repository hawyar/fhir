build:
	go build -o fhir

server:
	make build && ./fhir 

proxy:
	caddy run
