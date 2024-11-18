package Routes

import (
	"github.com/gorilla/mux"
	service "PencraftB/Services"
	controller "PencraftB/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter();

	router.Handle("/createblog", service.RateLimiter(controller.CreateBlogController)).Methods("POST")
	router.Handle("/addtag",service.RateLimiter(controller.CreateTagController)).Methods("POST")
	
	router.Handle("/getalltags",service.RateLimiter(controller.FetchAllTagController)).Methods("GET")

	return router;
}