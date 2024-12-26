package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  uint32 `json:"nome"`
	Email uint32 `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(bodyRequest, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario em struct"))
		return
	}

	db, erro := banco.Connectar()
	if erro != nil {
		log.Fatal("Erro ao conectar no banco de dados")
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}
