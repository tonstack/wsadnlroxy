#!make
include .env
export $(shell sed 's/=.*//' .env)

MAIN 		:= ./cmd/main.go

DOCKERD 	:= docker
COMPOSE 	:= docker-compose
DOCKER_UP 	:= up --build --force-recreate -d

.SILENT:

# -------- gorun --------
run:
	go run $(MAIN)

# -------- nginx --------
nginx-pass:
	bash ./data/pass-nginx.sh

nginx-reload:
	$(COMPOSE) exec nginx nginx -s reload

nginx-reload-no-ssl:
	$(COMPOSE) exec nginx-no-ssl nginx -s reload

# -------- setup --------
setup-no-ssl:
	make pass-nginx
	$(COMPOSE) $(DOCKER_UP) app
	$(COMPOSE) $(DOCKER_UP) nginx-no-ssl

setup-wtih-ssl:
	make pass-nginx
	$(COMPOSE) $(DOCKER_UP) app
	bash init-ssl.sh -d ${DOMAIN}

# -------- docker -------
docker-rm-all:
	-$(DOCKERD) kill	wstcproxy nginx-no-ssl nginx certbot
	-$(DOCKERD) rm 		wstcproxy nginx-no-ssl nginx certbot

docker-stop-all:
	-$(DOCKERD) stop	wstcproxy nginx-no-ssl nginx certbot

docker-run-no-ssl:
	-$(COMPOSE) up -d app nginx-no-ssl

docker-run-with-ssl:
	-$(COMPOSE) up -d app nginx certbot
