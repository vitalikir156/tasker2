package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/vitalikir156/tasker2/db"
	"github.com/vitalikir156/tasker2/models"
)

func GetTasks(c fiber.Ctx) error {
	query := "SELECT id, title, description, status, created_at, updated_at from tasks"
	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных")
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created, &task.Updated)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}
		tasks = append(tasks, task)
	}
	return c.JSON(tasks)
}

func CreateTask(c fiber.Ctx) error {
	task := new(models.Task)
	if err := c.Bind().Body(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	if task.Status != "new" && task.Status != "in_progress" && task.Status != "done" {
		return c.Status(406).SendString("Недопустимое значение поля status")
	}
	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)"
	_, err := db.DB.Exec(context.Background(), query,
		task.Title, task.Description, task.Status)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}
	return c.Status(201).SendString("Задача добавлена")
}

func UpdateTask(c fiber.Ctx) error {
	task := new(models.Task)
	if err := c.Bind().Body(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	id := c.Params("id")
	if task.Status != "new" && task.Status != "in_progress" && task.Status != "done" {
		return c.Status(406).SendString("Недопустимое значение поля status")
	}
	query := "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at =now() where id=$4"
	out, err := db.DB.Exec(context.Background(), query,
		task.Title, task.Description, task.Status, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления задачи")
	}
	if out.RowsAffected() < 1 {
		return c.Status(404).SendString("Нет задач с заданным ID")
	}
	return c.Status(201).SendString("Задача обновлена")
}

func DeleteTask(c fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM tasks where id=$1"
	out, err := db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления задачи")
	}
	if out.RowsAffected() < 1 {
		return c.Status(404).SendString("Нет задач с заданным ID")
	}
	return c.Status(201).SendString("Задача удалена")
}
