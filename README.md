# exo-planets
Example Golang server to serve api requests

# Pre-requisites
* Install Golang
* install mysql server with default configuration.
* create a Database name - planetDb

# How to Run program 
There are ways to run this server.
simplest is to run main.go file directly by just cloning the repository.

* Go to root of project.

**Command ro run:**
go run ./cmd/main.go

It will start the golang server on port 8080.

# How to Test
* After starting the server, hit the example Apis
  http://localhost/list
  http://localhost/get/{id}
