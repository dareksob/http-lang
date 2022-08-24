# http-lang
a collection of http sources written with different languages. All sub folders represent an language wir core and framework version.

## Run an server example
Use compose to run a service selective by use the serice name.

### Examples
- `docker compose up deno-core`
- `docker compose up go-gin`

## Specification
All server should have the same endpoints and behaivors.

All endpoints should be log each request

### Possible endpoints
- `GET` `/` will return a `string` 'hello world'
- `GET` `/get-json` will return a `json` string { msg: 'hello world' } and header Content-type: application/json
- `GET` `/get/:id` will return the id inside json { msg: 'hallo :id' } and header Content-Type: application/json
- `ANY` 404