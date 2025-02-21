# TODO Service

Это простое REST API для управления задачами (TODO), написанное на Go с использованием Fiber и PostgreSQL.

---

## Оглавление

1. [Требования](#требования)
2. [Настройка](#настройка)
    - [Конфигурация](#конфигурация)
    - [База данных](#база-данных)
3. [Запуск](#запуск)
    - [Локально](#локально)
4. [API Endpoints](#api-endpoints)
5. [Тестирование](#тестирование)
6. [Лицензия](#лицензия)

---

## Требования

- Go 1.20 или выше
- PostgreSQL 12 или выше

---

## Настройка

### Конфигурация

1. Скопируйте шаблон конфигурации:
   ```bash
   cp config/config.example.yaml config/config.yaml
   ```
   
2. Отредактируйте config/config.yaml
   ```yaml
   server:
      host: "localhost"
      port: 8080
      database:
      user: "your_db_user"
      password: "your_db_password"
      name: "your_db_name"
      host: "localhost"
      port: 5432
      type: "postgres"
   ```
### База данных
1. Убедитесь, что PostgreSQL запущен и доступен.
2. Сервис автоматически создаст таблицу tasks при первом запуске.

### Установка зависимостей
Перед запуском проекта установите все необходимые зависимости:
```bash
go mod tidy
```

---

## Запуск

### Локально

1. Установите переменную окружения для конфигурации:
```bash
# Для Linux/macOS
export CONFIG_PATH=./config/config.yaml

# Для Windows (PowerShell)
$env:CONFIG_PATH=".\config\config.yaml"
```

2. Запустите сервис:

```bash
go run ./cmd/todo/main.go
```

3. Сервис будет доступен по адресу: ```http://localhost:8080```

---

## API Endpoints

### Создать задачу

- Метод: ```POST /tasks```
- Тело запроса:
```json
{
  "title": "Новая задача",
  "description": "Описание задачи",
  "status": "new"
}
```
- Ответ:
```json
{
   "message": "Задача успешно создана",
   "id": 1
}
```
### Получить список задач

- Метод: ```GET /tasks```
- Ответ:
```json
[
   {
      "id": 1,
      "title": "Новая задача",
      "description": "Описание задачи",
      "status": "new",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:00:00Z"
   }
]
```

### Обновить задачу

- Метод: ```PUT /tasks/:id```
- Тело запроса:
```json
{
   "title": "Обновленная задача",
   "description": "Новое описание",
   "status": "in_progress"
}
```
- Ответ:
```json
{
   "message": "Задача успешно обновлена",
   "task": {
      "id": 1,
      "title": "Обновленная задача",
      "description": "Новое описание",
      "status": "in_progress",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:05:00Z"
   }
}
```

### Получить список задач

- Метод: ```DELETE /tasks/:id```
- Ответ:
```json
{
   "message": "Задача успешно удалена",
   "task_id": 1
}
```

---

## Тестирование

Для тестирования API можно использовать ```Postman```. Ниже приведены примеры запросов.

### Примеры запросов

1. Создать задачу
- Метод: ```POST```
- URL: ```http://localhost:8080/tasks```
- Тело запроса (JSON):
```json
{
  "title": "Новая задача",
  "description": "Описание задачи",
  "status": "new"
}
```
- Пример ответа:
```json
{
  "message": "Задача успешно создана",
  "id": 1
}
```

2. Получить список задач
- Метод: ```GET```
- URL: ```http://localhost:8080/tasks```
- Пример ответа:
```json
[
   {
      "id": 1,
      "title": "Новая задача",
      "description": "Описание задачи",
      "status": "new",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:00:00Z"
   }
]
```

3. Обновить задачу
- Метод: ```PUT```
- URL: ```http://localhost:8080/tasks/1``` (где 1 — ID задачи)
- Тело запроса (JSON):
```json
{
   "title": "Обновленная задача",
   "description": "Новое описание",
   "status": "in_progress"
}
```
- Пример ответа:
```json
{
   "message": "Задача успешно обновлена",
   "task": {
      "id": 1,
      "title": "Обновленная задача",
      "description": "Новое описание",
      "status": "in_progress",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:05:00Z"
   }
}
```

4. Удалить задачу
- Метод: ```DELETE```
- URL: ```http://localhost:8080/tasks/1``` (где 1 — ID задачи)
- Пример ответа:
```json
{
   "message": "Задача успешно удалена",
   "task_id": 1
}
```

---

## Скриншоты

![image](https://github.com/user-attachments/assets/1af8e2f7-ec2c-4e9c-913c-92aba173069a)

![image](https://github.com/user-attachments/assets/225232dc-0cef-4bfc-b5c6-23e89560379f)

![image](https://github.com/user-attachments/assets/3adc1047-c2d8-4243-bc40-bdfebcd6868c)

![image](https://github.com/user-attachments/assets/6399d2aa-52aa-4892-aca6-9fa12b99e8cf)

![image](https://github.com/user-attachments/assets/6b2b63cb-c581-4ecf-a55d-f5c86ea89327)







