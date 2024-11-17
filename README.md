# fun-twitt
fun-twitt is api mock of twitter/X in golang


## models

db provider: `sqlite`

- [X] users  ~ a model
- [ ] tweets ~ a model 
- [ ] follow ~ a model
- [ ] timeline ~ a query 

## User model

```prisma
model User {
  id      String    @id @default(uuid())
  sort_id String    @unique
  name    String    @default("")
  email   String    @unique
  created DateTime  @default(now())
  updated DateTime  @updatedAt
  deleted DateTime?
}
```

Add new user example API response 

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "ac176ba7-5a2a-4bd8-8b17-e54059691136",
        "name": "romel",
        "sort_id": "user_CUHO8G6XT8tC",
        "email": "bmxandcode+004@gmail.com",
        "created": "2024-11-16T00:52:55+00:00",
        "updated": "2024-11-16T00:52:55+00:00"
    },
    "message": ""
}
```

Add tweet response example

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "496b7fa0-6bbd-4783-9c04-b9887ea1913c",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
        "user_id": "ac176ba7-5a2a-4bd8-8b17-e54059691136",
        "sort_id": "tweet_bWpai6lxgyMNZlUVbctjvKwc",
        "created_at": "2024-11-17T04:38:46+00:00"
    },
    "message": ""
}
```



## Setting Up Local Development

1. Run migrations to initialize the database:
   ```bash
   go run github.com/steebchen/prisma-client-go db push


## test

`go test ./... -v` 


