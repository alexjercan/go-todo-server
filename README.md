# Todo Application

## Getting Started

### Docker Compose

#### Tested using

- Windows 10
- Docker version 20.10.14, build a224086
- Docker Compose version v2.4.1

#### Requirements

- Docker [download](https://docs.docker.com/get-docker/)
- Docker Compose [download](https://docs.docker.com/compose/install/)

#### Usage

1. Download the repository

    ``` console
    git clone https://github.com/alexjercan/go-todo-server.git
    cd go-todo-server
    ```

2. Run docker compose

    ``` console
    docker-compose up
    ```

### Rest API

#### Routes

1. `GET /api/items/`
    - Get all items in the database.
    - Input: none
    - Output: `[{"ID":number,"Description":string,"Done":boolean},...]`
2. `POST /api/items/`
    - Create a new item.
    - Input: `{"Description":string}`
    - Output: `{"ID":number,"Description":string,"Done":boolean}`
3. `PUT /api/items/`
    - Update an item.
    - Input: `{"ID":number,"Description":string,"Done":boolean}`
    - Output: `{"ID":number,"Description":string,"Done":boolean}`
4. `DELETE /api/items/`
    - Delete an item.
    - Input: `{"ID":number}`
    - Output: `{"ID":number,"Description":string,"Done":boolean}`
