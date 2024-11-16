# Prisma Client


## Commons

```bash
go run github.com/steebchen/prisma-client-go validate
go run github.com/steebchen/prisma-client-go generate
go run github.com/steebchen/prisma-client-go migrate dev
```

## ref  

0. **Validate the schema**:

```bash
go run github.com/steebchen/prisma-client-go validate
```

1. **Generate the Go Client**:

```bash
go run github.com/steebchen/prisma-client-go generate
```

2. **Sync Database Schema (Development)**:
   
```bash
go run github.com/steebchen/prisma-client-go db push
```

3. **Pull Existing Database Schema**:
   
```bash
go run github.com/steebchen/prisma-client-go db pull
```

4. **Create a Migration (Development)**:

```bash
go run github.com/steebchen/prisma-client-go migrate dev
```

5. **Apply Migrations to Production**:
   
```bash
go run github.com/steebchen/prisma-client-go migrate deploy
```
