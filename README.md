# go-base-backend

## Description

Example implementation go backend architecture.

## Release List
1. Database using PostgreSQL: master branch
2. Database using MySQL: checkout branch [mysql](https://github.com/geshtng/go-base-backend/tree/mysql)
3. Support for gRPC (database using PostgreSQL): checkout branch [grpc-postgresql](https://github.com/geshtng/go-base-backend/tree/grpc-postgresql)
4. Support for gRPC (database using MySQL): checkout branch [grpc-mysql](https://github.com/geshtng/go-base-backend/tree/grpc-mysql)

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

        database:
            db_host: localhost
            db_port: 3306
            db_name: go_base_backend
            db_username: root
            db_password: 
            db_charset: utf8mb4
            db_parse_time: 1
            db_local: Asia%2FJakarta
        ```

## Run the Project

   ```bash
   $ make run
   ```

Compile `.proto` file, use command: `make pb in=<file.proto>`<br>
Example:
```bash
make pb in=internal/handlers/article/grpc/article.proto
```
