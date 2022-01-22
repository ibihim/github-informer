BIN := bin/igit
MAIN := cmd/github/main.go
INSTALL_TO := $(GOBIN)

clean:
	@rm -f $(BIN)

build: clean
	@go build -o $(BIN) $(MAIN)

update: # Works as expected :D
	@rm go.mod
	@rm go.sum
	@go mod init
	@go mod tidy

install: build
	cp $(BIN) $(INSTALL_TO)

run: build
	@$(BIN)