
```
boilerplate-simple
├─ .air.toml
├─ .env
├─ .env.example
├─ Makefile
├─ bin
│  ├─ boilerplate-simple
│  └─ build-errors.log
├─ cmd
│  └─ server
│     └─ main.go
├─ deploy
│  └─ docker
│     ├─ Dockerfile
│     └─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ app
│  │  ├─ app.go
│  │  └─ factory
│  │     ├─ auth_factory.go
│  │     ├─ factory.go
│  │     └─ user_factory.go
│  ├─ config
│  │  ├─ app.go
│  │  ├─ config.go
│  │  └─ db.go
│  ├─ delivery
│  │  ├─ auth_handler.go
│  │  └─ user_handler.go
│  ├─ domain
│  │  └─ user.go
│  ├─ dto
│  │  ├─ auth.go
│  │  └─ user.go
│  ├─ infrastructure
│  │  ├─ db
│  │  │  └─ db.go
│  │  ├─ migration
│  │  │  └─ migrate.go
│  │  └─ repository
│  │     └─ user_repository.go
│  ├─ middleware
│  │  ├─ jwt.go
│  │  └─ rbac.go
│  ├─ routes
│  │  ├─ auth_routes.go
│  │  ├─ routes.go
│  │  └─ user_routes.go
│  └─ usecase
│     ├─ auth_usecase.go
│     └─ user_usecase.go
└─ pkg
   ├─ hash
   │  └─ hash.go
   ├─ logger
   │  └─ logger.go
   └─ response
      └─ response.go

```