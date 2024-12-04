# Library Management System
Library books management system using GRPC and Golang with PostgreSQL database and Redis cache.

## Features
- Authorization and authentication user
- Management of books
- Management of authors
- Management of categories
- Management of borrow books
- Recommendation system based on user borrowing history

## Entity Relationship Diagram
```mermaid
erDiagram
    USER ||--o{ BORROW_BOOK : "borrow"
    BORROW_BOOK }o--|| BOOK : "belongs to"
    BOOK ||--o{ AUTHOR : "written by"
    BOOK ||--o{ CATEGORY : "categorized by"

    USER {
        string username
        string password
        string email
    }
    BORROW_BOOK {
        string id
        string user_id
        string book_id
        date borrow_date
        date return_date
    }
    BOOK {
        string id
        string title
        string description
        string author_id
        string category_id
        int stock
    }
    AUTHOR {
        string id
        string name
    }
    CATEGORY {
        string id
        string name
    }
```

![ERD](erd.png)


## Usecase Diagram
```mermaid
graph TD
    User["User"] --> APIGateway["API Gateway"]
    APIGateway --> BookService["Book Service"]
    APIGateway --> AuthorService["Author Service"]
    APIGateway --> CategoryService["Category Service"]
    APIGateway --> UserService["User Service"]
```

## Sequence Diagram Book Service
```mermaid
sequenceDiagram
    User ->> API Gateway: Request
    API Gateway ->> Book Service: Request
    Book Service ->> Database: Request
    Database -->> Book Service: Response
    Book Service -->> API Gateway: Response
    API Gateway -->> User: Response
```

## Sequence Diagram Author Service
```mermaid
sequenceDiagram
    User ->> API Gateway: Request
    API Gateway ->> Author Service: Request
    Author Service ->> Database: Request
    Database -->> Author Service: Response
    Author Service -->> API Gateway: Response
    API Gateway -->> User: Response
```

## Sequence Diagram Category Service
```mermaid
sequenceDiagram
    User ->> API Gateway: Request
    API Gateway ->> Category Service: Request
    Category Service ->> Database: Request
    Database -->> Category Service: Response
    Category Service -->> API Gateway: Response
    API Gateway -->> User: Response
```

## Sequence Diagram User Service
```mermaid
sequenceDiagram
    User ->> API Gateway: Request
    API Gateway ->> User Service: Request
    User Service ->> Database: Request
    Database -->> User Service: Response
    User Service -->> API Gateway: Response
    API Gateway -->> User: Response
```

Link Postman:
```
https://bold-trinity-904712.postman.co/workspace/5f93435f-a044-4476-8e9d-d6270c4700fc/overview
```

## How to run
1. Clone this repository
2. Run `docker-compose up -d` to start the PostgreSQL and Redis


