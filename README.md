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

      database:
          db_host: localhost
          db_port: 5432
          db_name: go_base_backend
          db_username: postgres
          db_password: postgres
          db_postgres_ssl_mode: disable
      ```

## Run the Project

   ```bash
   $ make run
   ```

Compile `.proto` file, use command: `make pb in=<file.proto>`<br>
Example:
```bash
$ make pb in=internal/handlers/article/grpc/article.proto
```
