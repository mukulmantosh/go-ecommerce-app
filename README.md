# Ecommerce App using Echo

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
Using Windows?

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
![test-app](./misc/test-app.png)

