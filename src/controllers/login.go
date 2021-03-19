package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioDB, erro := repo.BuscarPorEmail(usuario.Email)
	if erro := seguranca.VerificarSenha(usuarioDB.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("Email ou senha incorretos."))
		return
	}

	usuarioDB.Senha = ""
	token, _ := autenticacao.CriarToken(usuarioDB.ID)
	respostas.JSON(w, http.StatusOK, struct {
		User  modelos.Usuario `json:"user"`
		Token string          `json:"token"`
	}{
		User:  usuarioDB,
		Token: token,
	})
}
