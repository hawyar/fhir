redis:
	docker run --name fhir-redis -d redis redis-server --appendonly yes

redis-insight:
	docker run -v redisinsight:/db -p 8001:8001 redislabs/redisinsight:latest

build:
	go build -o fhir

server:
	make build && ./fhir 

proxy:
	caddy run
