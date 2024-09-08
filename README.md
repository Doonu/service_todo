# service_todo

End-point

    POST            /AUTH/SIGN-UP
    POST            /AUTH/SIGN-IN

    GET             /LISTS
    GET             /LISTS/{ID}
    POST            /LISTS
    PUT             /LISTS/{ID}
    DELETE          /LISTS/{ID}
    GET             /LISTS/{ID}/ITEMS 
    POST            /LISTS/{ID}/ITEMS 

Таблицы баз данных
![img.png](tables.png)

Пакеты

- Используется фреймворк **gin** для генерации маршрутов
- viper - библиотека для работы с конфигурацией
- sqlx - библиотека надстройка над sql
- lib/pq - go-драйвер для взаимодействия с базой данных PostgreSQL
- godotenv - библиотека для хранения переменных окружения
- 

----
Чистая архиткектура
----

 - ПОЗВОЛЯЮТ БЕЗБОЛЕЗНЕННО МЕНЯТЬ ФРЕЙМВОРКИ И ИНФРАСТРУКТУРНЫЕ ЭЛЕМЕНТЫ
 - ИМЕЮТ ВОЗМОЖНОСТЬ ИЗОЛИРОВАННО ТЕСТИРОВАТЬ КАЖДЫЙ СЛОЙ СИСТЕМЫ БЕЗ ЗАВИСИМОСТЕЙ
 - НЕ ЗАВИСЯТ ОТ РЕАЛИЗАЦИИ ПОЛЬЗОВАТЕЛЬСКОГО ИНТЕРФЕЙСА
 - БИЗНЕС ЛОГИКА НЕ ЗАВИСИТ ОТ КОНКРЕТНОГО ВЫБОРА БАЗЫ ДАННЫХ ИЛИ ДРУГИХ ТЕХНИЧЕСКИХ РЕШЕНИЙ СИСТЕМЫ

![img_1.png](architectures.png)

Repository - Работа с бд
Service - Работа с бизнес логикой
Handler - Работа с роутером и end поинтами

----
Использование базы данных Postgres + Docker
----
- ***docker pull postgres***

Создание базы данных с именем todo-db и паролем - qwerty, порта для использования 5436
После остановки контейнера - он удаляется 
- ***docker run --name=todo-db -e POSTGRES_PASSWORD="qwerty" -p 5436:5432 -d --rm postgres***

Подключение к базе данных
- docker ps (вытаскиваем id контейнер)
- docker exec -it <id Container> /bin/bash
- psql -U postgres
- \d

-----------------------
Миграции баз данных
-----------------------

**Используется утилита migrate**
- ***migrate create -ext sql -dir ./schema -seq init***

Миграция
- migrate -path ./schema\ -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable" up 

Откат миграции
- migrate -path ./schema\ -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable" down

Удаление миграции
- migrate -path ./schema\ -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable" drop