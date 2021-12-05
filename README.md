# Sau Sheong

## WEB PROGRAMMING GO

`go version 1.17.3`

### Подключение к БД Postgres супер пользователь Linux

`sudo su postgres`

### Создание нового пользователя

`createuser --interactive <username>`

### Вывод списка фалгов ролей пользователя.

`createuser --help`

### Создать базу данных

`createdb <database name>`

### Подключение к Postgres и создание таблицы

`sudo -i -u postgres`

### Подключение к созданному пользователю

`psql -d gwp`

### Создание таблицы при помощи setuo.sql

`psql -U gwp -f setup.sql -d gwp`
или `\i path to file.sql`
