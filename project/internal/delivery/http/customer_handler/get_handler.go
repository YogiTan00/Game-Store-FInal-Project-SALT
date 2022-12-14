package customer_handler

import (
	"fmt"
	"game-store-final-project/project/internal/delivery/http_response"
	"net/http"

	"github.com/gorilla/mux"
)

func (s_handler *CustomerHandlerInteractor) IndexController(w http.ResponseWriter, r *http.Request) {
	var (
		// ctx  = context.Background()
		vars = mux.Vars(r)
	)

	nik := vars["nik"]

	customer, errCustomerFromUseCase := s_handler.CustomerUseCase.IndexCustomerTrx(nik)
	if errCustomerFromUseCase != nil {
		fmt.Println(errCustomerFromUseCase)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error on Usecase"))
		return
	}

	response, errMap := http_response.MapResponseCustomer(200, "Success", customer)
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
