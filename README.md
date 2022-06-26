# Example REST Processing in Go

This project demonstrates how to handle each HTTP Verb using go

It includes examples of:

- GET / POST / PUT / PATCH / DELETE / HEAD verb handling
- Query parameters
- URL parameters
- JSON Marshal / Unmarshal processing
- Error handling
- File retrieval
- ETag generation
- Swagger usage
- ... anything else i can think of and have time to include 


## Install Swagger Tooling
```
dir=$(mktemp -d)
git clone https://github.com/go-swagger/go-swagger "$dir"
cd "$dir"
go install ./cmd/swagger
```
## Default Swagger Project Page

http://localhost:8090/docs

## Various Components used

OpenAPI project for golang

[OpenAPI Initiative golang toolkit](https://github.com/go-openapi)

The runtime component for use in codegeneration or as untyped usage.

[openapi runtime interfaces](https://github.com/go-openapi)
```
go get github.com/go-openapi/runtime/middleware
```
Gorilla is a web toolkit for the Go programming language that provides useful, composable packages for writing HTTP-based applications.

[Gorilla Web Toolkit](https://github.com/gorilla)
```
go get -u github.com/gorilla/mux
```
## Generate Swagger YAML
```
swagger generate spec -o ./swagger.yaml --scan-models
```