package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vitalikir156/tasker2/handlers"
)

func RegisterProductRoutes(app *fiber.App) {
	app.Get("/tasks", handlers.GetTasks)          // Получить все продукты
	app.Post("/tasks", handlers.CreateTask)       // Создать новый продукт
	app.Put("/tasks/:id", handlers.UpdateTask)    // Обновить продукт
	app.Delete("/tasks/:id", handlers.DeleteTask) // Удалить продукт
}
