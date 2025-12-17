PROJECT = serve
MAIN_FILE = ./cmd/$(PROJECT)/main.go
INSTALL_PREFIX = ~/.local/bin/

run:
	go run $(MAIN_FILE)

build:
	go build -o ./bin/$(PROJECT) $(MAIN_FILE)

install: build
	mkdir -p ${INSTALL_PREFIX} &&\
	mv -v ./bin/${PROJECT} ${INSTALL_PREFIX}

.PHONY: run install
