package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mipimentel/gin-api-rest/controllers"
	"github.com/mipimentel/gin-api-rest/database"
	"github.com/mipimentel/gin-api-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678910", RG: "12345679"}
	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCode(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET(":nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/kierkegaard", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, fmt.Sprintf("Status code recebido: %d -- valor esperado: %d", resposta.Code, http.StatusOK))

	mockDaResposta := `{"API diz:":"Olá kierkegaard!"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListaTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}
