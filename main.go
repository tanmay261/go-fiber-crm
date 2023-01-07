package main

import(
	"fmt"
	"log"
	"github.com/gofiber/fiber"
	"github.com/tanmay261/go-fiber-crm/database"

)
func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase(){

}
func main(){
app:=fiber.New()	
setupRoutes(app)
app.Listen(3000)
}

func GetLead(){

}

func GetLeads(){

}

func DeleteLead(){}

func NewLead(){}