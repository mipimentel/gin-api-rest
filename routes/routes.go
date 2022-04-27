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
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
