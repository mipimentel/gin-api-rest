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
	"github.com/stretchr/testify/assert"
)

func SetupRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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

	r := SetupRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}
