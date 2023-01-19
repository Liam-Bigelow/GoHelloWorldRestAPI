package main

import (
	"net/http",
	"errors",
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID			string	`json:"id"`
	Item		string	`json:"item"`
	Completed	bool	`json:"completed"`
}

var todos = []todo {
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}


func getTodos( context *gin.Context ){
	context.IndentedJSON( http.StatusOK, todos )
}

func getTodoById( id string ) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	// if here no todo was found
	return nil, errors.New("no todo found");
}

func getTodo( context *gin.Context ) {
	id := context.Param( "id" )
	todo, err := getTodoById( id )

	if err != nul {
		context.IndentedJSON( http.StatusNotFound, gin.H{ "message": "Todo not found" })
		return
	}

	// if here todo is found
	context.IndentedJSON( http.StatusOK, todo )
}

func addTodos( context *gin.Context ){
	var newTodo todo

	if err := context.BindJSON( &newTodo ); err != nil {
		return
	}

	todos = append( todos, newTodo )

	context.IndentedJSON( http.StatusCreated, newTodo )
}

func main() {
	router := gin.Default()

	router.GET( "/todos", getTodos )
	router.GET( "/todos/:d", getTodo )
	router.POST( "/todos", addTodos )

	router.Run("localhost:8000")
}


