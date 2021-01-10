# Bexs Backend Challenge

The goal of this code is to presents a solution to the problem **Rota de Viagem** described at
[bexs-backend-challenge](https://bitbucket.org/bexstech/bexs-backend-exam/src/master/).

> This document has been built following markdown [lint rules](https://github.com/markdownlint/markdownlint).

## Requirements

To accomplish the recommendation to avoid the usage of frameworks and
external libraries as much as possible this project has been built using only
Golang standard librires.

In order to run the application is necessary to have:

- Golang version 1.13
- Unix based system

I don't have any experience with Golang in production, so I decided to use this project as an excuse
to learn more about this programming language ;)

## How to run the tests

All the tests were written before the implementation following the TDD practices.
The commits represent the coding timeline which makes easy to follow
my train of thought

To run all the tests

```shell
make test
```

## How to run the Application

To run the application it's necessary to compile the code and then execute the binary file.

The easiest way to compile and execute the application using *input.csv* as an 
input file, run the following command:

```shell
make
```

To only compile the application and create the binary run the command:

```shell
make build
```

After that to execute the binary passing the file name as parameter

```shell
./mysolution input.csv
```

## Rest API

This application has 3 different endpoints

- GET /route/best
- POST /route
- GET /route/all

### GET /route/best

This endpoint returns the best route between `source` and `target`

Query string

1. source
1. target

*Request example:*

```shell
curl -XGET 'http://localhost:3000/route/best?source=GRU&target=SCL'
```

*Response:*

```json
{
    "Paths": ["GRU - BRC - SCL"],
    "Cost": 15
}
```

### POST /route

This endpoint adds a new connection from `source` to `target`

Body

The body must be JSON type with the following fields:

1. source: `string`
1. target: `string`
1. cost: `integer`

*Body example:*

```json
{ 
    "source": "BRC",
    "target": "ORX",
    "cost": 15
}
```

*Request example:*

```shell
$ curl -XPOST 'http://localhost:3000/route' \
    --data '{ "source": "BRC", "target": "ORX", "cost": 15 }'

The path has been added with success!
```

### GET /route/all

This endpoint returns all the routes between `source` and `traget`

Query string

1. source
1. target

*Request example:*

```shell
curl -XGET 'http://localhost:3000/route/all?source=GRU&target=ORL'
```

*Response:*

```json
[
    {
        "Paths": [
            "GRU - BRC - SCL - ORL",
            "GRU - BRC - ORL"
        ],
        "Cost": 35
    },
    {
        "Paths": ["GRU - SCL - ORL"],
        "Cost": 40
    },
    {
        "Paths": ["GRU - ORL"],
        "Cost": 56
    }
]
```


## Structure of packages

The structure of the project it's very simple and can be described by the tree below

```
├── Makefile
├── README.md
├── api
│   ├── controller.go
│   └── server.go
├── go.mod
├── input.csv
├── main.go
└── service
    ├── file.go
    ├── file_test.go
    ├── graph.go
    ├── graph_test.go
    ├── route.go
    ├── route_test.go
    └── types.go
```

*OBS: Due to the fact that I never experience golang in a production
environment, I think this is the worst part of this code kkkkkk!*

## Edge cases

After starting to write the solution I found some edge cases that
I choose to resolve, the cases are:

1. When exists more than one path with the minimum cost

    > That is the reason for the response of the endpoints returns the path as `array`.

2. When trying to get the best route between `source` and `target` but they
don't have any way to connect each other

3. When trying to get a route between `source` and `target` but they don't even exist.

## Conclusion

After all that I want to thanks the opportunity to take this test, that helps
me a lot to understand more about Golang and also it was very fun trying to
solve that problem. Since college, I never took a look in a book about
algorithms, and now I catch myself opening the Cormen book again kkkkkkkkkkk =D

> Thanks for all and That's all
