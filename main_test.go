package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var NOME = "Nome Teste"
var CPF = "12312312312"
var RG = "123123123"

func SetupDasRotasDeTeste() *gin.Engine {
	// exibe mensagens mais simplificadas
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: NOME,
		Cpf:  CPF,
		Rg:   RG,
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func excluiAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListarTodosOsAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer excluiAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscaPorCpf(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer excluiAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.ExibeTodosAlunos)
	url := "/alunos/cpf/" + CPF
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscaPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer excluiAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeTodosAlunos)
	url := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMock []models.Aluno
	// faz parse dos dados de um json codificado e salva em uma variável
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	nome := alunoMock[len(alunoMock)-1].Nome
	cpf := alunoMock[len(alunoMock)-1].Cpf
	rg := alunoMock[len(alunoMock)-1].Rg

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, NOME, nome, "Nome diferente")
	assert.Equal(t, CPF, cpf, "CPF diferente")
	assert.Equal(t, RG, rg, "RG diferente")
}

func TestAtualizarAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer excluiAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PUT("/alunos/:id", controllers.EditaAluno)
	url := "/alunos/" + strconv.Itoa(ID)
	nomeAtualizado := "Nome atualizado"
	cpfAtualizado := "55555555555"
	rgAtualizado := "111111111"
	aluno := models.Aluno{Nome: nomeAtualizado, Cpf: cpfAtualizado, Rg: rgAtualizado}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(valorJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, nomeAtualizado, alunoMockAtualizado.Nome, "Nomes diferentes")
	assert.Equal(t, rgAtualizado, alunoMockAtualizado.Rg, "Rg diferente")
	assert.Equal(t, cpfAtualizado, alunoMockAtualizado.Cpf, "CPF diferente")
}

func TestExcluirAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	url := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
