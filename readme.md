## Tiny url

---


Для запуска
```shell script
docker-compose up
```
---
Для запуска тестов
```shell script
go run test ./...
```
---
Для запуска сервиса локально
```shell script
cp config/config.dist.toml config/config.toml
cp .env.dist .env
make run
```

