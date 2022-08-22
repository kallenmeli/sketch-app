# Sketch Canvas

## How to run?

```bash
docker-compose up
```

The variables that could be used as an example are stored in `.env` file in the root path.

## Running tests

```bash
go test ./...
```

## How to use?

All the endpoints are mapped to in the root path.

**[API] Get a draw by ID**
```bash 
curl http://localhost:8080/your-guid
```

**[API] Write a draw**

```bash
curl --location --request POST 'localhost:8080/' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "x": 30,
        "y": 0,
        "outline": "@",
        "width": 5,
        "height": 3,
        "fill": "."
    },
     {
        "x": 10,
        "y": 0,
        "outline": "@",
        "width": 1,
        "height": 1,
        "fill": "Y"
    }
]'
```

**[VIEW] See a draw:**

Access the following webpage passing your valid draw id.
[http://localhost:8080?id=your-draw-id](http://localhost:8080?id=your-draw-id)


## Techs & Libraries

- Golang 1.19;
- Postgres;
- [Julien Schmidt httprouter](julienschmidt/httprouter)