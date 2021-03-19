package modelos

import (
	"api/src/respostas"
	"api/src/seguranca"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um modelo do banco de dados
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório.")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório.")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório.")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido.")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória.")
	}

	return nil
}

func (usuario *Usuario) validar2(w http.ResponseWriter, etapa string) bool {
	erros := []string{}
	if usuario.Nome == "" {
		erros = append(erros, "O nome é obrigatório")
	}
	if usuario.Nick == "" {
		erros = append(erros, "O nick é obrigatório")
	}
	if usuario.Email == "" {
		erros = append(erros, "O e-mail é obrigatório")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); usuario.Email != "" && erro != nil {
		erros = append(erros, "O e-mail inserido é inválido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		erros = append(erros, "A senha é obrigatória")
	}

	if len(erros) > 0 {
		respostas.JSON(w, http.StatusUnprocessableEntity, struct {
			Erros []string `json:"erros"`
		}{
			Erros: erros,
		})
		return false
	}

	return true
}

// Preparar prepara os dados do usuário validando e formatando.
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// Preparar2 prepara os dados do usuário validando e formatando.
func (usuario *Usuario) Preparar2(w http.ResponseWriter, etapa string) bool {
	valido := usuario.validar2(w, etapa)
	if !valido {
		return false
	}

	if erro := usuario.formatar(etapa); erro != nil {
		respostas.Erro(w, 400, erro)
		return false
	}

	return true
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaHash)
	}

	return nil
}
