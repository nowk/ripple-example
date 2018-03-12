
default: install-dumb-init

# server runs main.go which contains an http server
server:
	@docker-compose run --rm --service-ports server
.PHONY: server

# install-dumb-init installs dumb-init for our docker entrypoints
install-dumb-init:
	@mkdir -p .bin
	@wget -O .bin/dumb-init \
		https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64
	@sudo chmod +x .bin/dumb-init
.PHONY: install-dumb-init

