package Routes

import (
	service "PencraftB/Services"
	controller "PencraftB/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/api/v1/blog", service.RateLimiter(controller.CreateBlogController)).Methods("POST")
	router.Handle("/api/v1/tag", service.RateLimiter(controller.CreateTagController)).Methods("POST")

	router.Handle("/api/v1/tags", service.RateLimiter(controller.FetchAllTagController)).Methods("GET")
	router.Handle("/api/v1/blogs", service.RateLimiter(controller.FetchAllBlogController)).Methods("GET")
	router.Handle("/api/v1/blog/{blog_id}",service.RateLimiter(controller.FetchBlogbyBlogIdController)).Methods("GET")

	router.Handle("/api/v1/blog", service.RateLimiter(controller.UpdateBlogController)).Methods("PUT")

	router.Handle("/api/v1/blog/{blog_id}", service.RateLimiter(controller.HardDeleteBlogbyBlogidController)).Methods("DELETE")
	router.Handle("/api/v1/deleteall",service.RateLimiter(controller.DeleteAllDataController)).Methods("DELETE");


	return router
}
