# The Blog #

## Objective ##
To demonstrate a typical rest service using golang, fiber, gorm, postgres/sqlite, with tests, auth, jwt, and basic relational modeling all wrapped
inside a well architected codebase.

### TL;DR ###
Run `make` to build everything and then run `make docker-up` to docker compose the app and *postgresql*. If you want to run it without docker, you can
run `make run` - see `.env` for environment variables. `make test` will run through all the unit tests and write a coverage report.

### Architecture ###
GoFiber was used to quickly create an "express.js" like application (with room to grow) and setup with JWT middleware for authentication with a sql database. `Gorm` was used as an ORM for SQLite/PostgreSQL to abstract away the database layer. Basic operations are done in the service with the controllers handling validation and input/reponse.

Before we start our service, we prep the data store for use. `Init` is called on the `database` package. This will initialize `gorm` based on `.env` file and environment variables (see `config` package) and call `Migrate` which will auto migrate the models of this service.

A `Service` interface provides the basic blocks for creating a web service - complete with `Start` `Stop` methods.

The `NewWebService` creates a fiber web service that satifies the `Service` interface and returns it as a `Service` instance.
We then register our routes with the `RegisterRoutes` method and start the server.


### API ###
*Main/Misc Endpoints*
```
 /       - Index
 /status - Health Check
```

*Auth Endpoints*
```
 /auth/login      - Login
 /auth/logout     - Logout
 /auth/me         - User Identity
 /auth/register   - Register
 /auth/token      - Token Refresh
```

*Post Endpoints*
```
 /posts                                  - Posts
 /posts?page={}&limit={}                 - Posts with pagination
 /posts/create                           - Create a blog post
 /posts/:id/update                       - Update a blog post
 /posts/:id/delete                       - Delete a blog post
 /posts/:id                              - A specific post
 /posts/:id/comments?page={}&limit={}    - A posts comments list
 /posts/:id/comments/create              - Creates a comment on a post
 /posts/:id/comments/:id/update          - Update a specific comment on a post
 /posts/:id/comments/:id/delete          - Delete a comment from a post
 ```

 ### TODO ###
 Add administration - add / remove / edit users/posts/comments
 Add drafts - YOLO publishing currently but you can revise... with no revision history :(
 Revision History - Keep the old new again or just audit trail your way to "who wrote that?"


 I spent about 4 hours on this...