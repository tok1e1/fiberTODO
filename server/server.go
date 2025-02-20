package server

import (
	"context"
	"fiberTODO/cmd/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Создание новой задачи
func CreateTasks(c *fiber.Ctx, db database.DB) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	if task.Title == "" || task.Status == "" {
		return c.Status(400).SendString("Поле title и status обязательны")
	}

	now := time.Now()

	var taskID int
	rows, err := db.Query(context.TODO(),
		"INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, task.Status, now, now,
	)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&taskID); err != nil {
			return c.Status(500).SendString("Ошибка при извлечении ID: " + err.Error())
		}
	} else {
		return c.Status(500).SendString("Ошибка: ID не получен после вставки задачи")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Задача успешно создана",
		"id":      taskID,
	})
}

// Получение списка всех задач
func GetTasks(c *fiber.Ctx, db database.DB) error {
	var tasks []Task

	rows, err := db.Query(context.TODO(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return c.Status(500).SendString("Ошибка при обработке данных: " + err.Error())
		}
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		return c.Status(200).JSON([]Task{})
	}

	return c.Status(200).JSON(tasks)
}

// Обновление задачи
func UpdateTask(c *fiber.Ctx, db database.DB) error {
	taskID := c.Params("id")
	if taskID == "" {
		return c.Status(400).SendString("ID задачи обязателен")
	}

	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if task.Title == "" || task.Status == "" {
		return c.Status(400).SendString("Поле title и status обязательны")
	}

	now := time.Now()

	err := db.ExecQuery(
		context.TODO(),
		"UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, task.Status, now, taskID,
	)

	if err != nil {
		return c.Status(500).SendString("Ошибка при обновлении задачи: " + err.Error())
	}

	var updatedTask Task
	rows, err := db.Query(context.TODO(),
		"SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1", taskID,
	)
	if err != nil {
		return c.Status(500).SendString("Ошибка при извлечении обновленных данных: " + err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.Description, &updatedTask.Status, &updatedTask.CreatedAt, &updatedTask.UpdatedAt)
		if err != nil {
			return c.Status(500).SendString("Ошибка при сканировании результата: " + err.Error())
		}
	} else {
		return c.Status(404).SendString("Задача не найдена")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Задача успешно обновлена",
		"task":    updatedTask,
	})
}

// Удалить задачу
func DeleteTask(c *fiber.Ctx, db database.DB) error {
	taskID := c.Params("id")
	if taskID == "" {
		return c.Status(400).SendString("ID задачи обязателен")
	}

	var exists bool
	rows, err := db.Query(context.TODO(),
		"SELECT EXISTS(SELECT 1 FROM tasks WHERE id = $1)", taskID,
	)
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса: " + err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return c.Status(500).SendString("Ошибка при проверке существования задачи: " + err.Error())
		}
	}

	if !exists {
		return c.Status(404).SendString("Задача не найдена")
	}

	err = db.ExecQuery(context.TODO(), "DELETE FROM tasks WHERE id = $1", taskID)
	if err != nil {
		return c.Status(500).SendString("Ошибка при удалении задачи: " + err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Задача успешно удалена",
		"task_id": taskID,
	})
}
