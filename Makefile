redis:
	docker run -p 6379:6379 --name fhir-rejson redislabs/rejson:latest --appendonly yes

redis-insight:
	docker run -v redisinsight:/db -p 8001:8001 redislabs/redisinsight:latest


build:
	go build -o fhir

server:
	docker run --name -p 4141:4141

proxy:
	caddy run



#docker build -t fhir-server .