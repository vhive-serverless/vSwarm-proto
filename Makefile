
## Protocol buffer

VERSION ?=LOCAL

proto-all: proto-all-python proto-all-go

proto-all-python:
	pipenv install
	pipenv run \
	python3 -m grpc_tools.protoc -I. \
		--python_out=. \
		--grpc_python_out=. \
		./proto/*/*.proto

proto-all-go:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		./proto/*/*.proto


client: test-client/main.go
	sed -i "s|LOCAL|$(VERSION)|" $<;
	go mod tidy
	go build -o $@ $<

# build: bin-dir
# 	if [ -z "$(shell git status --porcelain)" ]; then \
# 		sed -i "s|LOCAL|$$(git rev-parse --short HEAD)|" ./cmd/version.go; \
# 		go build -o $(BIN_DIR)/$(BIN); \
# 		git checkout -- ./cmd/version.go; \
# 	else \
# 		echo Working directory not clean, commit changes; \
# 	fi

