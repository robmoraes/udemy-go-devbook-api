package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken criar token jwt web
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

// ValidarToken realiza validação do token gerado na autenticação
func ValidarToken(r *http.Request) error {
	ts := extrairToken(r)
	token, erro := jwt.Parse(ts, retornarSecretKey)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

// ExtrairUsuarioID retorna o ID do usuário que está armazenado no token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	ts := extrairToken(r)
	token, erro := jwt.Parse(ts, retornarSecretKey)
	if erro != nil {
		return 0, erro
	}
	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tuid := fmt.Sprintf("%.0f", permissoes["usuarioId"])
		usuarioID, erro := strconv.ParseUint(tuid, 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}

	return 0, errors.New("Token inválido")

}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	ts := strings.Split(token, " ")
	if len(ts) == 2 {
		return ts[1]
	}
	return ""
}

func retornarSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
