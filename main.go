package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct{
	ID			string	`json: "id"`
	Item		string	`json: "item"`
	Completed	bool	`json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: true},
	{ID: "2", Item: "Read Book", Completed: true},
	{ID: "3", Item: "REST API Course Completed", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos);
}

func main(){
	router := gin.Default();
	
	router.GET("/todos", getTodos);

	router.Run("localhost:8080")
}
