## Выполните следующие шаги, чтобы настроить и запустить приложение:

1. Клонируйте репозиторий на свой локальный компьютер:
```
git clone https://github.com/iKayrat/testovoe-zadanie.git
```

2. Создайте файл **.env** в корневом каталоге проекта со следующими переменными среды:
```
DBSOURCE=postgresql://user:password@localhost:5432/dbname?sslmode=disable
SECRET_KEY=secret_key
```
(примечание: переменные среды уже имеются))

3. Запустите контейнеры Docker с помощью docker-сompose:
```
docker-compose up -d
```
или

```
make dbstart
```
Данная команда запустит контейнер c Postgres, создаст базу данных и поднимет миграцию

**Команда для запуска**
```
go run cmd/main.go (номер заказов)
```
eg.
```
go run cmd/main.go 10,11,14,15
```

Дерево приложения:
```
├── api
├── app.env
├── assets
├── build
│   ├── ci
│   └── package
├── cmd
│   └── main.go
├── configs
├── deployments
├── docs
├── examples
├── githooks
├── go.mod
├── go.sum
├── init
├── internal
│   ├── app
│   │   ├── config
│   │   │   └── config.go
│   │   ├── controller
│   │   │   ├── controller.go
│   │   │   └── service.go
│   │   └── db
│   │       ├── migration
│   │       │   ├── 000001_all tables.down.sql
│   │       │   ├── 000001_products.up.sql
│   │       │   ├── 000002_orders.up.sql
│   │       │   ├── 000003_users.up.sql
│   │       │   ├── 000004_shelves.up.sql
│   │       │   ├── 000005_product_shelves.up.sql
│   │       │   └── 000006_Initial_Data.up.sql
│   │       ├── query
│   │       │   ├── order.sql
│   │       │   ├── product_shelves.sql
│   │       │   ├── product.sql
│   │       │   └── shelf.sql
│   │       └── sqlc
│   │           ├── db.go
│   │           ├── models.go
│   │           ├── order.sql.go
│   │           ├── product_shelves.sql.go
│   │           ├── product.sql.go
│   │           ├── querier.go
│   │           ├── shelf.sql.go
│   │           └── store.go
│   └── pkg
│       └── _your_private_lib_
├── LICENSE.md
├── Makefile
├── pkg
├── README.md
├── scripts
├── sqlc.yaml
├── test
├── third_party
├── tools
├── web
│   ├── app
│   ├── static
│   └── template
└── website
```
