**Запустить базу данных в контейнере**  
```
docker-compose -f deployments/docker-compose.yaml up       
```

**Установить Гуся**  
`go get -u github.com/pressly/goose/cmd/goose`

**Создать миграцию**  
`goose -dir migrations postgres "host=localhost port=54321 user=calendar_app password=12345678 dbname=calendar sslmode=disable" create <migration-name> sql`

**Применить миграцию**  
`goose -dir migrations postgres "host=localhost port=54321 user=calendar_app password=12345678 dbname=calendar sslmode=disable" up`

**HELP**  
`goose`
