# go-base-backend

## Release List
1. Database using PostgreSQL: branch [master](https://github.com/geshtng/go-base-backend/tree/master)
2. Database using MySQL: checkout branch [mysql](https://github.com/geshtng/go-base-backend/tree/mysql)
3. Support for gRPC (database using PostgreSQL): checkout branch [grpc-postgresql](https://github.com/geshtng/go-base-backend/tree/grpc-postgresql)
4. Support for gRPC (database using MySQL): checkout branch [grpc-mysql](https://github.com/geshtng/go-base-backend/tree/grpc-mysql)

## Description
Example implementation go backend (clean) architecture. It is very easy to configure.

This project has 4 domain layers:

- Model
  <br>
  This layer will save models that were used in the other domains. Can be accessed from any other layer and other domains.
- Handler
  <br>
  This layer will do the job as the presenter of the application.
- Service
  <br>
  This layer will do the job as a controller and handle the business logic.
- Repository
  <br>
  This layer is the one that stores the database handler. Any operation on database like querying, inserting, updating, and deleting, will be done on this layer.

## Setup
1.  Clone project.
    ```bash
    $ git clone https://github.com/geshtng/go-base-backend
    ```
2.  Init Database.
    - Create a new database. Example database name: `go_base_backend`.<br>
    - After you run the server, it will automatically create tables and relations in the database `go_base_backend`.<br>
3.  Change config.

    Make a new file named `config.yaml` inside the folder `/config`.<br>
    Use `config.yaml.example` to see the example or see the config sample below.<br>
    ```yaml
    app:
        name: go-base-backend

    server:
        host: localhost
        port: 8080

    database:
        db_host: localhost
        db_port: 3306
        db_name: go_base_backend
        db_username: root
        db_password: 
        db_charset: utf8mb4
        db_parse_time: 1
        db_local: Asia%2FJakarta

    jwt:
        expired: 60
        issuer: go-base-backend
        secret: sKk6E5gpVD

    ```

## Run the Project
   ```bash
   $ make run
   ```

## API List
You can find API list on file `routes/routes.go`

## Example API with Authentication
I have set up an example of an API that uses authentication:
```http
GET localhost:8080/profiles
```

## Postman Collection
Import files `go-base-backend.postman_collection.json` to your postman

## Framework and Library
  - [Gin](https://github.com/gin-gonic/gin)
  - [Gorm](https://github.com/go-gorm/gorm)
  - [Copier](https://github.com/jinzhu/copier)
  - [Golang-jwt](https://github.com/golang-jwt/jwt)
  - [Viper](https://github.com/spf13/viper)
  - Other libraries listed in `go.mod`