# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
SRC=src/*
BIN=keychain

all: build
build:
		$(GOBUILD) -o $(BIN) $(SRC)
clean: 
		$(GOCLEAN)
		rm -f $(BIN)

deps:
		$(GOGET) github.com/go-redis/redis
