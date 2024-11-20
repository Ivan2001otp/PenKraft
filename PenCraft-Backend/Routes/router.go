package Routes

import (
	service "PencraftB/Services"
	controller "PencraftB/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/createblog", service.RateLimiter(controller.CreateBlogController)).Methods("POST")
	router.Handle("/addtag", service.RateLimiter(controller.CreateTagController)).Methods("POST")

	router.Handle("/getalltags", service.RateLimiter(controller.FetchAllTagController)).Methods("GET")
	router.Handle("/getallblogs", service.RateLimiter(controller.FetchAllBlogController)).Methods("GET")

	router.Handle("/updateblog", service.RateLimiter(controller.UpdateBlogController)).Methods("PUT")

	router.Handle("/deleteblog/{blog_id}", service.RateLimiter(controller.HardDeleteBlogbyBlogidController)).Methods("DELETE")

	return router
}
