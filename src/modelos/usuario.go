package modelos

import (
	"errors"
	"strings"
	"time"
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
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O email é obrigatório.")
	}

	return nil
}

// Preparar ...
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
