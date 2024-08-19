# Тестовое задание Avito Backend Bootcamp

## Обзор сервиса

Этот backend-сервис позволяет пользователям публиковать объявления о продаже или аренде квартир на платформе Авито. Сервис предоставляет функционал для регистрации пользователей, авторизации, создания квартир и их модерации, обеспечивая соответствие контента стандартам Авито.

## Функционал

- **Авторизация пользователей:**
    - Упрощенная генерация токенов для разных ролей пользователей (клиент, модератор) через `/dummyLogin`.
    - Полноценная система регистрации и входа с использованием токенов для аутентификации через `/register` и `/login`.

- **Управление домами:**
    - Модераторы могут создавать новые дома через `/house/create`.
    - Пользователи могут получать список квартир в доме через `/house/{id}`.

- **Управление квартирами:**
    - Любой авторизованный пользователь может создать квартиру через `/flat/create`.
    - Модераторы могут изменять статус модерации квартиры через `/flat/update`.

- **Модерация квартир:**
    - Квартиры могут иметь четыре статуса модерации: `created`, `approved`, `declined`, `on moderation`.
    - Модераторы могут обновлять статус квартиры и брать её на модерацию, чтобы предотвратить конфликт с другими модераторами.


## API Endpoints

### Авторизация

#### `/dummyLogin`

- **Метод:** `GET`
- Пример запроса ```curl -X 'GET' \
  'https://editor.swagger.io/dummyLogin?user_type=moderator' \
  -H 'accept: application/json```
- **Описание:** Получение токена для аутентификации на основе выбранного типа пользователя (клиент или модератор).
- **Параметры:**
    - `user_type` (обязательный): Тип пользователя (`client` или `moderator`).
- **Ответ:**
    - `200 OK`: Токен для дальнейшей авторизации.
    - `400 BadRequest` ```
      {
      "message": "Incorrect user_type",
      "request_id": "request_id",
      "code": 400
      }```
    - `404`
    - `500` ```{
      "message": "что-то пошло не так",
      "request_id": "request_id",
      "code": 12345
      }```

#### `/register`

- **Метод:** `POST`
- **Описание:** Регистрация нового пользователя.
- **Тело запроса:**
    - `email` (обязательный): Email пользователя.
    - `password` (обязательный): Пароль пользователя.
    - `user_type` (обязательный): Тип пользователя (`client` или `moderator`).
- **Ответ:**
    - `200 OK`: Идентификатор зарегистрированного пользователя.
    - `400 BadRequest` ```
      {
      "message": "Incorrect user_type",
      "request_id": "request_id",
      "code": 400
      }```
    - `404`
    - `409` ```"error": "User with this email already exists```
    - `500` ```{
      "message": "что-то пошло не так",
      "request_id": "request_id",
      "code": 500
      }```

#### `/login`

- **Метод:** `POST`
- **Описание:** Авторизация пользователя по email и паролю.
- **Тело запроса:**
    - `id` (обязательный): Идентификатор пользователя.
    - `password` (обязательный): Пароль пользователя.
- **Ответ:**
    - `200 OK`: Токен для дальнейшей авторизации.
    - `400 BadRequest` ```"error": "Invalid data type",```
    - `404` ```"error": "User not found"```
    - `500` ```{
    "message": "что-то пошло не так",
    "request_id": "request_id",
    "code": 500
    }```

### Управление домами

#### `/house/create`

- **Метод:** `POST`
- **Описание:** Создание нового дома (только для модераторов).
- **Тело запроса:**
    - `address` (обязательный): Адрес дома.
    - `year` (обязательный): Год постройки.
    - `developer` (необязательный): Застройщик.
- **Ответ:**
    - `200 OK`: Информация о созданном доме.
    - `400 BadRequest` ```
        {
        "message": "Invalid data type",
        "request_id": "request_id",
        "code": 400
        }```
    - `401` ```{
      "message": "Insufficient access rights",
      "request_id": "request_id",
      "code": 401
      }```
    - `500` ```{
      "message": "что-то пошло не так",
      "request_id": "request_id",
      "code": 500
      }```

#### `/house/{id}`

- **Метод:** `GET`
- **Описание:** Получение списка квартир в доме по его идентификатору.
- **Параметры:**
    - `id` (обязательный): Идентификатор дома.
- **Ответ:**
    - `200 OK`: Список квартир. Клиент видит только квартиры со статусом `approved`, модератор — все квартиры.
    - `400 BadRequest` ```
      {
      "message": "Invalid id data type",
      "request_id": "request_id",
      "code": 400
      }```
    - `401` ```{
        "message": "Unauthorized access",
        "request_id": "request_id",
        "code": 401
        }```
    - `500` ```{
      "message": "что-то пошло не так",
      "request_id": "request_id",
      "code": 500
      }```
### Управление квартирами

#### `/flat/create`

- **Метод:** `POST`
- **Описание:** Создание новой квартиры (доступно всем авторизованным пользователям).
- **Тело запроса:**
    - `id` (обязательный): Номер квартиры
    - `house_id` (обязательный): Идентификатор дома.
    - `price` (обязательный): Цена квартиры.
    - `rooms` (обязательный): Количество комнат.
- **Ответ:**
    - `200 OK`: Информация о созданной квартире. Квартира получает статус `created`.
    - `400 BadRequest` ```
      {
      "message": "Invalid  data type",
      "request_id": "request_id",
      "code": 400
      }```
    - `500` ```{
        "message": "Flat with this id was created earlier",
        "request_id": "request_id",
        "code": 401
        }```
    - `500` ```{
      "message": "что-то пошло не так",
      "request_id": "request_id",
      "code": 500
      }```

#### `/flat/update`

- **Метод:** `POST`
- **Описание:** Обновление статуса модерации квартиры (только для модераторов).
- **Тело запроса:**
    - `id` (обязательный): Идентификатор квартиры.
    - `house_id` (обязательный): Идентификатор дома
    - `status` (обязательный): Новый статус модерации (`created`, `approved`, `declined`, `on moderation`).
- **Ответ:**
    - `200 OK`: Информация о обновленной квартире.
    - `400 BadRequest` ```
      {
      "message": "Invalid  data type",
      "request_id": "request_id",
      "code": 400
      }```
    - `401` ```{
        "message": "Insufficient access rights ",
        "request_id": "request_id",
        "code": 401
        }```
    - `401` ```{
      "message": "Flat is already in moderation or has been processed",
      "request_id": "request_id",
      "code": 401
      }```
    - `500` ```{
      "message": "Flat with this id not found",
      "request_id": "request_id",
      "code": 401
      }```
    - `500` ```{
      "message": "что-то пошло не так",
      "request_id": "request_id",
      "code": 500
      }```

## Запуск проекта
1. git clone https://github.com/Chigvero/HomeService
2. `cd HomeService`
3. `docker-compose up --build`
4. Сервис будет доступен по адресу `http://localhost:8080`

## Запуск тестов
1. Зайдите в корневую директорию проекта
2. `go test -v ./...`

## Заключение

Этот сервис обеспечивает полный функционал для работы с объявлениями о продаже и аренде квартир на Авито, включая авторизацию пользователей, управление домами и квартирами, а также модерацию объявлений.

## PS
1. 