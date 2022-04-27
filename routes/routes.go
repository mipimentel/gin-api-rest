package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mipimentel/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
