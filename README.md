# Library Management System
Library books management system using GRPC and Golang with PostgreSQL database and Redis cache.

## Features
- Authorization and authentication user
- Management of books
- Management of authors
- Management of categories
- Management of borrow books
- Recommendation system based on user borrowing history

## Usecase Diagram
```mermaid
       +------------------+
       |    User          |
       +------------------+
                 |
                 v
       +------------------+
       |  API Gateway     |
       +------------------+
           |    |    |    |
           v    v    v    v
+---------+ +------+ +------+ +-------+
| Book    | |Author| |Category| |User  |
| Service | |Service| |Service | |Service|
+---------+ +------+ +------+ +-------+
            |    |    |    |
            v    v    v    v
        +------------------+
        |    Database      |
        +------------------+
  ```

## Sequence Diagram
```mermaid
sequenceDiagram
    User ->> API Gateway: Request
    API Gateway ->> Book Service: Request
    Book Service ->> Database: Request
    Database -->> Book Service: Response
    Book Service -->> API Gateway: Response
    API Gateway -->> User: Response
```



