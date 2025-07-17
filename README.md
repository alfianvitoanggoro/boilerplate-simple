```
boilerplate-simple
├─ .air.toml
├─ Makefile
├─ README.md
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
│  │  ├─ factory
│  │  │  ├─ auth_factory.go
│  │  │  ├─ factory.go
│  │  │  └─ user_factory.go
│  │  └─ routes
│  │     ├─ auth_routes.go
│  │     ├─ routes.go
│  │     └─ user_routes.go
│  ├─ config
│  │  ├─ app.go
│  │  ├─ config.go
│  │  └─ db.go
│  ├─ domain
│  │  ├─ auth
│  │  │  ├─ dto
│  │  │  │  └─ auth.go
│  │  │  ├─ handler
│  │  │  │  └─ handler.go
│  │  │  └─ usecase
│  │  │     └─ usecase.go
│  │  └─ user
│  │     ├─ dto
│  │     │  └─ user.go
│  │     ├─ handler
│  │     │  └─ handler.go
│  │     ├─ model
│  │     │  └─ model.go
│  │     └─ usecase
│  │        └─ usecase.go
│  ├─ infrastructure
│  │  ├─ db
│  │  │  └─ db.go
│  │  ├─ migration
│  │  │  └─ migrate.go
│  │  └─ repository
│  │     └─ user_repository.go
│  └─ middleware
│     ├─ jwt.go
│     └─ rbac.go
└─ pkg
   ├─ hash
   │  └─ hash.go
   ├─ logger
   │  └─ logger.go
   └─ response
      └─ response.go

```

# Run Application

```
make docker-up
```

# Run Migrate

- Run application before run migrate in the different CLI

```
make migrate
```
