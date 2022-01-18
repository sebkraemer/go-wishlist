# go-wishlist

(Yet to become) Wishlist web app with Golang, part of my [#100DaysOfCode](https://github.com/sebkraemer/100-days-of-code/blob/master/log.md#day-64)

# Documentation

## Misc. notes for setup and running

`brew install --cash mongodb-compass` for conveniently inspecting the DB

Start mongodb service:
`docker-compose -f ./docker/docker-compose.yml up`

## API

This is a draft for the API.


### REST resource

```
{
    "id": 1,
    "description": "Writing Without Bullshit book",
}
```

Internally, the resource will have an owner and date information from when the wish was first created.

```
{
    "id": 1,
    "owner": "seb",
    "description": "Writing Without Bullshit book",
}
```

### REST routes

```
{
  "GET /api/wishes": {
    "desc": "returns all wishes",
    "response": "200 application/json",
    "data": [{}, {}, {}]
  },

  "GET /api/wishes/:id": {
    "desc": "returns a wish respresented by its id",
    "response": "200 application/json",
    "data": {}
  },

  "POST /api/wishes": {
    "desc": "create and returns a new wish using the posted object",
    "response": "201 application/json",
    "data": {}
  },

  "PUT /api/wishes/:id": {
    "desc": "updates and returns a wish with the posted update object",
    "response": "200 application/json",
    "data": {}
  },

  "DELETE /api/wishes/:id": {
    "desc": "deletes and returns the matching wish",
    "response": "200 application/json",
    "data": {}
  }
}
```

# Ideas, TODOs

-[] make db connection configurable
-[] implement CRUD functionality
-[] expose REST interface
-[] middleware
  -[] logger
  -[] auth
  -[] web interface
-[] authentication
  -[] mongodb password protection
  -[] authentication by application
  -[] user accounts, create and manage
  -[] authenticate users in web interface, JWT?
-[] deployment, where and how
-[] insert layer to support other nosql dbs, e.g. dynamodb
  (see also link section for aws example)
- wishes
  - add priority
  - add wishlist sharing

# Links and resources

- mongo and setting it up in docker
  - https://www.mongodb.com/compatibility/docker
  - https://docs.docker.com/compose/compose-file/compose-file-v3
  - MongoDB university, free courses and possibly vouchers for certification
    - https://university.mongodb.com/courses/catalog
    - kudos @MohsenKamranii for the hint. https://twitter.com/mohsenkamranii/status/1482500357389701122?s=21
  - https://docs.mongodb.com/drivers/go/current/
  - https://docs.mongodb.com/manual/crud/
- https://dev.to/joojodontoh/build-user-authentication-in-golang-with-jwt-and-mongodb-2igd
- https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.8.0/mongo#example-Connect-AWS
