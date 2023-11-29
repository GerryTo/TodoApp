package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct{
	Id int `json:"id" gorm:"column:n_id;primaryKey"`
	Title string `json:"title" gorm:"column:v_title"`
	Body string `json:"body" gorm:"column:v_body"`
	Done int `json:"done" gorm:"column:n_done"`
}

func (Todo) TableName() string {
	return "tbl_todo"
}

func main() {
	
	db,err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_todoapps?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Fatal(err)
	}
	
	defer db.Close()

	app:= fiber.New()
	

	app.Get("/healthcheck",func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})


	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil{
			return err
		}

		if err := db	.Create(todo).Error; err != nil {
			log.Println("Error creating todo:", err)
			return c.Status(500).SendString("Internal Server Error")
		}

		return c.JSON(todo)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		
		if err != nil{
			return c.Status(401).SendString("Invalid data/Data not found")
		}
		
		var todo Todo

		if err := db.First(&todo,"n_id = ?", id ).Error; err != nil {
			return c.Status(404).SendString("Todo not Found")
		}

		db.Model(&todo).Where("n_id = ?", id).Update("n_done", 1)

		return c.JSON(todo)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		var todos []Todo
		db.Where("n_done != 1").Find(&todos)

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}