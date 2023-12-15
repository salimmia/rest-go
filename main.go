package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
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

func addTodos(context *gin.Context) {
	var newTodos todo

	err := context.BindJSON(&newTodos)

	if(err != nil){
		return
	}

	todos = append(todos, newTodos)

	context.IndentedJSON(http.StatusCreated, newTodos)
}

func getTodoById(id string) (*todo, error){
	for i, t := range todos{
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context){
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main(){
	router := gin.Default()
	
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	
	router.PATCH("/todos/:id", toggleTodoStatus)

	router.POST("/todos", addTodos)

	router.Run("localhost:8080")
}
