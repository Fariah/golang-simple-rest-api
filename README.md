# RESTful API with Docker, PostgreSQL, and go-chi

- app/handler: handlers for endpoints
- app/models: struct for db items and etc.
- app/main.go: entry point, start server
- app/routes.go: routes list
- db/: connect to db, models maybe
- docker/: docker files
- migrations/: migration files

## to start

- run `composer up -d`
- go to http://127.0.0.1:8080/