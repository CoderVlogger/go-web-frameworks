# API Specs

This document explains API specification for `/entities` JSON API. This API is implemented in all mentioned Go Web Frameworks. Hence, the specification is required to keep them aligned and follow the same API contract.

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

- `400` - invalid invalid input data (payload)
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

- `400` - invalid invalid input data (payload)
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
