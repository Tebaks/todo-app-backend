# Todo App Backend
## Run
```
go run .
```
## Run tests
```
go test ./...
```

## Development
-   First write test with mock database for find todos for repository.
-   Then create repository structure and Find function to make test pass.
-   After test passed for database pass to service layer and write test with mock repository.
-   Write service logic to make test pass.
-   After test passed for service pass to handler layer and write test with mock service.
- And finaly when test are passed for handler I create router 
-   I repeat this steps for AddTodo Functionality

## CI/CD
For cicd I use github actions and deploy to heroku.
-   First I create heroku app for backend with Postgress addon
### Github Actions
-   Install golang
-   Run tests
-   Deploy to production heroku
## Consumer Driven Contract
Because of error I get when I try to create pacts I didn't manage to create consumer driven contract tests for backend.

## Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/Tebaks/todo-app-backend)