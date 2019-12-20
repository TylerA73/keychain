# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
SRC=src/*
SETTINGS=settings
BIN=keychain

all: build
build:
		$(GOBUILD) -o $(BIN) $(SRC)
		mkdir -p $(SETTINGS)
clean: 
		$(GOCLEAN)
		rm -f $(BIN)

deps:
		$(GOGET) github.com/go-redis/redis
