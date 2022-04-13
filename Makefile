
## Protocol buffer
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
