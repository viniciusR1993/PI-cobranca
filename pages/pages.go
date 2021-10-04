package pages

import (
	"fatec/utils"
	"net/http"
)

func CarregarTelaCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)
}

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)
}

func CarregarTelaPgInicial(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "pag_inicial.html", nil)
}
