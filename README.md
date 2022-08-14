# Welcome to Catalog!

Hi! I'm the service which manage all products.

## Features

 - Show all products ready to play
 - Configure a product
 - Close an existing product

## Tech stack

I'm written in Go, version 1.17, and i use MongoDB to store all data.

### Docs
 - Golang ([docs](https://golang.org/)) 
 - MongoDB ([docs](https://www.mongodb.com/)) 

    #### Proto
   	Recompile proto files to generate server and client files.
	- *proto_name*.pdb.go
	- *proto_name*_grpc.pdb.go

	##### Instructions
	- change directory to root
	- run *protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/catalog.proto*
