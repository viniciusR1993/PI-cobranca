package routes

import (
	"fatec/controllers"
	"fatec/cookie"
	"fatec/middlewares"
	"fatec/pages"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(rtr *mux.Router) {
	fmt.Println("Inicializando servidor na porta 8000")

	fileServer := http.FileServer(http.Dir("./assets/"))
	rtr.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	//Metodos implantados
	rtr.HandleFunc("/login", controllers.EnterUser).Methods("POST")
	rtr.HandleFunc("/", pages.CarregarTelaLogin).Methods("GET")
	rtr.HandleFunc("/pginicial", pages.CarregarTelaPgInicial).Methods("GET")

	//Teste de Cookie
	rtr.HandleFunc("/testeToken", middlewares.Logger(middlewares.Autenticar(pages.CarregarTelaCadastro))).Methods("GET")
	rtr.HandleFunc("/deleteCookie", cookie.DeleteCookie).Methods("GET")

	//Teste de User
	rtr.HandleFunc("/AllUser", controllers.AllUsers).Methods("GET")
	rtr.HandleFunc("/GetUser", controllers.GetUser).Methods("POST")
	rtr.HandleFunc("/UpdateUser", controllers.UpdateUser).Methods("POST")
	rtr.HandleFunc("/InsertUser", controllers.InsertUser).Methods("POST")
	rtr.HandleFunc("/DeletetUser", controllers.DeletetUser).Methods("POST")

	//Teste de Bank
	rtr.HandleFunc("/AllBank", controllers.AllBank).Methods("GET")
	rtr.HandleFunc("/GetBank", controllers.GetBank).Methods("POST")
	rtr.HandleFunc("/UpdateBank", controllers.UpdateBank).Methods("POST")
	rtr.HandleFunc("/InsertBank", controllers.InsertBank).Methods("POST")
	rtr.HandleFunc("/DeleteBank", controllers.DeleteBank).Methods("POST")

	//Teste de Client
	rtr.HandleFunc("/AllClient", controllers.AllClient).Methods("GET")
	rtr.HandleFunc("/GetClient", controllers.GetClient).Methods("POST")
	rtr.HandleFunc("/UpdateClient", controllers.UpdateClient).Methods("POST")
	rtr.HandleFunc("/InsertClient", controllers.InsertClient).Methods("POST")
	rtr.HandleFunc("/DeleteClient", controllers.DeleteClient).Methods("POST")

	//Teste de Office
	rtr.HandleFunc("/AllOffice", controllers.AllOffice).Methods("GET")
	rtr.HandleFunc("/GetOffice", controllers.GetOffice).Methods("POST")
	rtr.HandleFunc("/UpdateOffice", controllers.UpdateOffice).Methods("POST")
	rtr.HandleFunc("/InsertOffice", controllers.InsertOffice).Methods("POST")
	rtr.HandleFunc("/DeleteOffice", controllers.DeleteOffice).Methods("POST")

	//Teste de Transaction
	rtr.HandleFunc("/AllTransaction", controllers.AllTransaction).Methods("GET")
	rtr.HandleFunc("/GetTransaction", controllers.GetTransaction).Methods("POST")
	rtr.HandleFunc("/UpdateTransaction", controllers.UpdateTransaction).Methods("POST")
	rtr.HandleFunc("/InsertTransaction", controllers.InsertTransaction).Methods("POST")
	rtr.HandleFunc("/DeleteTransaction", controllers.DeleteTransaction).Methods("POST")

	//Teste de File
	rtr.HandleFunc("/AllFile", controllers.AllFile).Methods("GET")
	rtr.HandleFunc("/GetFile", controllers.GetFile).Methods("POST")
	rtr.HandleFunc("/UpdateFile", controllers.UpdateFile).Methods("POST")
	rtr.HandleFunc("/InsertFile", controllers.InsertFile).Methods("POST")
	rtr.HandleFunc("/DeleteFile", controllers.DeleteFile).Methods("POST")
}
