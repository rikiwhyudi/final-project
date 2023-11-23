

## Final Project: WebChat Application with Hexagonal/Clean Architecture

### `docker-compose up --build`
Run and build the app with docker-compose.\
Server started at [http://localhost:8080](http://localhost:8080).

### `install docker`
Install docker here [https://www.docker.com/get-started/](https://www.docker.com/get-started/)

**features:**
- multiple chat channels

**tech-stack used:**
- golang
- echo framework
- websocket
- docker
- html, css, javascript

## `Structure of the Project`
```md
└── chatting-app
    ├── cmd
    │   └── main.go
    ├── internal
    │   ├── adapters
    │   │   ├── handlers
    │   │   │   └── chat_handlers.go
    │   │   └── repository
    │   └── core
    │       ├── models
    │       │   ├── payload.go
    │       │   ├── response.go
    │       │   └── websocket.go
    │       ├── ports
    │       │   └── chat.go
    │       └── services
    │           └── chat_services.go
    ├── pkg
    │   ├── db
    │   └── websockets
    │       └── websockets.go
    ├── routes
    │   ├── chat.go
    │   └── routes.go
    ├── views
    │   └── index.html
    ├── .env
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    └── go.sum
```
