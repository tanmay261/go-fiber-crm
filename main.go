package main

import (
	"fmt"
	
	"github.com/tanmay261/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/tanmay261/go-fiber-crm/database"
	"github.com/jinzhu/gorm/dialects/sqlite"
)
func setupRoutes(app *fiber.App){
	app.Get("api/v1/leads",lead.GetLeads)
	app.Get("api/v1/leads/:id",lead.GetLead)
	app.Post("api/v1/leads",lead.NewLead)
	app.Delete("api/v1/leads/:id",lead.DeleteLead)
}

func initDatabase(){
var err error
database.DBConn,err=gorm.Open("sqlite3","leads.db")
if err!=nil{
	panic("failed to connect to Database")
}
fmt.Println("Connected to database")
database.DBConn.AutoMigrate(&lead.Lead{})
fmt.Println("Database migrated")
}
func main(){
	initDatabase()
app:=fiber.New()	
setupRoutes(app)
app.Listen(3000)

// Defer means to run the command at the end , i.e.e after every other function returns
defer database.DBConn.Close()
}

