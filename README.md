# The Blog #

## Objective ##
To demonstrate a typical rest service using golang, fiber, gorm, postgres/sqlite, with tests, auth, jwt, and basic relational modeling all wrapped
inside a well architected codebase.

### Architecture ###

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