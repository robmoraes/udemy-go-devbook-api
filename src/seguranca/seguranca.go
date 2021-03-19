package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//VerificarSenha compara uma senh com hash e uma senha e retorna se elas são iguais
func VerificarSenha(senhaHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}
