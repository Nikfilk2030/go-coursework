# auth

--- 

## Обязательные этапы
- Установить go
- Установить PostgreSQL
- `go get github.com/gin-gonic/gin`
- `go get -u go.uber.org/zap`
- `go get github.com/golang-jwt/jwt/v4`
- `go get github.com/caarlos0/env`
- `go get github.com/jackc/pgx/v4`
- `go get github.com/juju/zaputil/zapctx`
- `go get github.com/gin-contrib/zap`


## Описание происходящего

>`Dockerfile` - загрузка в Docker и исполнение `main`

>`gitlab-ci.yml` - настройки для CI в gitlab

>`cmd` - исполнение `main`. В целом, этот код представляет собой основной шаблон для приложения
на Gin, включающий запуск и остановку приложения,
обработку сигналов и логирование.
Он позволяет гарантировать корректное завершение приложения при получении
сигналов от операционной системы.

> `tokens` - имплементация библиотеки jwt4 для работы с JSON Web Token (для общения с фронтом).

> `probes` - определяет структуру Probes и связанные с ней функции для обработки HTTP 
> запросов в процессе мониторинга состояния приложения. Он используется для реализации проб
> (probes) в приложении для проверки состояния работы сервиса.
