package Routes

import (
	"github.com/gorilla/mux"

//	service "PencraftB/Services"
)

func Router() *mux.Router {
	router := mux.NewRouter();
	return router;
}