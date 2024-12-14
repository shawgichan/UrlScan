dev-up:
	docker stop urlscan-powerdns-container || true
	docker rm urlscan-powerdns-container || true
	docker build -t urlscan-powerdns-dev -f docker/powerdns/Dockerfile .
	docker run -d   --name urlscan-powerdns-container   -v $$(pwd):/project   -p 53:53/udp   -p 53:53/tcp   --privileged   urlscan-powerdns-dev

docker-up:
	docker exec -it urlscan-powerdns-container /bin/bash

.PHONY: dev-up docker-up