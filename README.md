# Getting Started

Clone this repo

Ensure that you have docker installed

In the project directory, you can run:

## `go mod tidy`  

This will install all the required packages

### `make up`

Will spin up docker postgress container and pgadmin

### `make migrate`

to create migrations

### `make server`

Runs the app in the development mode.\
Open [http://localhost:3001/api](http://localhost:3001/api) to view it in your browser.

The application has the following end points

### `POST /api` creates a user, it expects the following in the body

```JSON
{
    "name": <string>
}
```

### `GET /api` this lists all the users

```JSON
[
    {
        "id": 1,
        "name": "Better Tester"
    },
    {
        "id": 2,
        "name": "John Doe"
    }
]
```

### `GET /api/user_id` this lists  a single user with the given id or returns an empty object if not found

```JSON
    {
        "id": 2,
        "name": "John Doe"
    }

```

### `PUT /api/user_id` updates  user with the given id

```JSON
{
    "name": <string>
}
```

### `DELETE /api/user_id` deletes  user with the given id
