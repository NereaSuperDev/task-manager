package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/controllers"
)

// SetTaskRoutes configures routes for task entity
func SetTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("/tasks/users/{id}", controllers.GetTasksByUser).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}

// A FEW NOTES TO REMEMBER 
// when we build web applications, you might need some shared functionality to be executed for some or all HTTP handlers.
// middleware is pluggable and self-contained piece of code that wraps a web application.
// here are some example senarios in which you can use middleware:
// logging HTTP requests and responses
// compressing HTTP responses 
// writting common response headers
// creating database session objects
// implementing security and validating authentication credentials
