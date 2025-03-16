# tasker2
## TODO-list demo

### Зависимости
-GO 1.24.1
-docker (проверено на версии 26.1.3)
-docker-compose (проверено на версии 1.25.0)

### Для запуска на liunx машине
```
 git clone git@github.com:vitalikir156/tasker2.git
 cd tasker2/
 docker-compose up -d
```
По умолчанию используется порт 3000 для HTTP

### Настройки docker-compose.yml
 HTTPPORT - указывает на каком порту работать HTTP серверу (обязательно скорректировать и ports)
 DBSTRING - строка подключения к БД (для понимания синтаксиса: postgres://username:password@localhost:5432/database_name)
