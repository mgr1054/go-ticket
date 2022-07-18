# Go Ticket

A simple Gin REST-API with following features:


- Backend Architecture with Gin-Gonic and PosgretSQL to manage events, tickets & users 
- Data Modeles build with GORM
- Documentation with Swagger
- Authentication and Authorization with JWT
- Password Encryption with Bycrypt
- Containerization with Docker
- CI with Github Actions

## Architecture

```
--- pkg
 |------- controller
 |------- db
 |------- docs
 |------- middleware
 |------- models
 |------- utils
```

- controller
  - implement logic and handle input from router
- db
  - setup database connection
- middleware
  - middleware for authorization
- models
  - database models
- utils
  - JWT generation and database helpers

## Getting started

You need [Docker Desktop](https://www.docker.com/products/docker-desktop/) to start with this tutorial.  
Alternatively you can run the service locally without docker by running `go run .` inside the root folder.  
If you chose to do that a PostgreSQL Database is needed with the following setup: 

| Name     | Value         |
| -------- | ------------- |
| host     | database      |
| user     | admin         |
| password | p             |
| dbname   | postgres      |
| port     | 5432          |

1. Checkout the repository to your local IDE. 

```sh
$ git clone https://github.com/mgr1054/go-ticket.git
```
2. Pull the required images from DockerHub. 

```sh
$ docker compose pull
```

3. Start the Docker container. 

```sh
$ docker compose up
```

4. Now the service is up and running. 

Service accepts requests under `http://localhost:8080/api`  
API Documentation is reachable under `http://localhost:8080/swagger/index.html`  
A pre-defined Postman workspace is available at `go-ticket/extras/Go-Ticket.postman_collection.json`
Project Presentation is also stored at `go-ticket/extras/Präsentation_MaxGreß.pdf`
