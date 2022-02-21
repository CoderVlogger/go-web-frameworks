# API Specs

This document explains API specification for `/entities` JSON API. This API is implemented in all mentioned Go Web Frameworks. Hence, the specification is required to keep them aligned and follow the same API contract.

## Endpoints

### `GET /entities`

List entities.

### Named Params

None.

#### Query Params

- `page` - page number

#### Status Codes

- `404` - if the given page number does not exist

### `GET /entities/1`

Get a single entity by entity's ID.

### Named Params

- `id` - id of a requested entity

#### Query Params

None.

#### Status Codes

- `404` - if an entity with the specified entity id does not exist
- `500` - if an error occurred during the storage operation
