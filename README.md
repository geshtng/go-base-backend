# go-base-backend

## Description

Example implementation go backend architecture.

## Setup

1.  Clone project

    ```bash
    $ git clone https://github.com/geshtng/go-base-backend
    ```

2.  Init Database

    - Create new database. Example database name: `go_base_backend`.<br>
    - After you run the server, it will automatically create tables and relations on database `go_base_backend`.<br>

3.  Change config
    <br>
    Make a new file named `config.yaml` inside folder `/config`.<br>
    Use `config.yaml.example` to see the example or see the config sample below.<br>

        ```yaml
        app:
            name: go-base-backend

        server:
            host: localhost
            port: 8080

        database:
            db_host: localhost
            db_port: 5432
            db_name: go_base_backend
            db_username: postgres
            db_password: postgres
            db_postgres_ssl_mode: disable

        jwt:
            expired: 60
            issuer: go-base-backend
            secret: sKk6E5gpVD

        ```

## Run the Project

1. Without nodemon
   ```bash
   $ make run
   ```
2. With nodemon
   ```bash
   $ make xrun
   ```

## API List

API list on files `routes/routes.go`

## Example API with Authentication

```http
GET localhost:8080/profiles
```

## Postman Collection

Import files `go-base-backend.postman_collection.json` to your postman
