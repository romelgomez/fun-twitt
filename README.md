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




## Setting Up Local Development

1. Run migrations to initialize the database:
   ```bash
   go run github.com/steebchen/prisma-client-go db push


## test

`go test ./... -v` 


