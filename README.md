# Ecommerce App using Echo



![Workflow](https://github.com/mukulmantosh/go-ecommerce-app/actions/workflows/badge.yml/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-60.6%25-yellow)

We intend to develop an e-commerce application with a REST
architecture using the Echo framework. Echo is a Go web framework
known for its exceptional performance, extensibility, and
minimalistic design.

![terminal](./misc/background.png)




### Prerequisites

Before starting up this project, make sure you have the necessary dependencies installed in your machine.

### Software Installation

- [x] [Go](https://go.dev/) - Go is an open source programming language that makes it simple to build secure, scalable systems.

- [x] [Docker](https://www.docker.com/) - Docker helps developers bring their ideas to life by conquering the complexity of app development.

- [x] [PostgreSQL](https://www.postgresql.org/) - The World's Most Advanced Open Source Relational Database

- [x] [golangci-lint](https://golangci-lint.run/) - is a fast Go linters runner. It runs linters in parallel, uses caching, supports yaml config, etc.  


For running Postgres locally using Docker, run the following command: 

```bash
docker run --name ecommerce-local-db -p 5432:5432 -e POSTGRES_PASSWORD=******** -d postgres
```

Execute in Postgres DB Shell

```sql
create database ecommerce;
```
### Database Schema
![db_schema](./misc/ecommerce-db-design.png)


### Environment Variables

Before launching the application, be certain to configure the necessary environment variables.

```
- JWT_SECRET
- DB_HOST
- DB_USERNAME
- DB_PASSWORD
- DB_NAME
- DB_PORT
```
Using **Windows**? Run the following command

![windows-env](./misc/windows-env.png)


### Application Startup

#### Running App

```bash
make run
```
![run-app](./misc/run-app.png)

#### Building App

```bash
make build
```
![build-app](./misc/build-app.png)


#### Executing Tests

```bash
make test
```
![test-app](./misc/gotest.gif)


## REST Examples

You can find the Postman Collection/HTTP Client for GoLand under `postman_httpclient` directory.

![postman-app](./misc/postman.png)

## Docker

You have the option to retrieve the image from [DockerHub](https://hub.docker.com/r/mukulmantosh/go-ecommerce-echo).


### Running Application in Local Kubernetes

Proceed with the instructions to launch your
application within a local Kubernetes cluster, 
such as Docker Desktop or Minikube.

Before proceeding, make sure to update the ConfigMap.

![k8s-terminal](./misc/k8s.png)

