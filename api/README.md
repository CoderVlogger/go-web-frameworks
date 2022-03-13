# API Specs

This document explains API specifications for the **/entities** JSON API. Its goal is to define the standards and follow them in each used web framework.

## Endpoints

### `GET /entities`

List entities.

#### Named Params

- None

#### Query Params

- `page` - page number

#### Payload

- None

#### Status Codes

- `400` - invalid input data (bad request)
- `404` - entity does not exist

### `GET /entities/<id>`

Get a single entity by entity's ID.

#### Named Params

- `id` - id of a requested entity

#### Query Params

- None

#### Payload

- None

#### Status Codes

- `404` - entity does not exist
- `500` - unclassified internal error

### `POST /entities`

Add a new entity.

#### Named Params

- None

#### Query Params

- None

#### Payload

- Entity object

#### Status Codes

- `400` - invalid input data (bad request)
- `404` - entity does not exist
- `500` - unclassified internal error

### `PUT /entities`

Update an existing entity.

#### Named Params

- None

#### Query Params

- None

#### Payload

- Entity object

#### Status Codes

- `400` - invalid input data (bad request)
- `404` - entity does not exist
- `500` - unclassified internal error

### `DELETE /entities/<id>`

Delete an existing entity.

#### Named Params

- `id` - id of a requested entity

#### Query Params

- None

#### Payload

- Entity object

#### Status Codes

- `404` - entity does not exist
- `500` - unclassified internal error
