package routing

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-news/controllers"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"/get-token", "POST", controllers.AuthGetTokenHandler},
	Route{"/news", "GET", controllers.NewsListHandler},
	Route{"/news", "POST", controllers.NewsCreateHandler},
	Route{"/news/{id}", "GET", controllers.NewsGetHandler},
	Route{"/news/{id}", "PUT", controllers.NewsUpdateHandler},
	Route{"/news/{id}/publish", "PUT", controllers.NewsPublishHandler},
	Route{"/news/{id}", "DELETE", controllers.NewsDeleteHandler},
	Route{"/topics", "GET", controllers.TopicListHandler},
	Route{"/topics", "POST", controllers.TopicCreateHandler},
	Route{"/topics/{id}", "GET", controllers.TopicGetHandler},
	Route{"/topics/{id}", "PUT", controllers.TopicUpdateHandler},
	Route{"/topics/{id}", "DELETE", controllers.TopicDeleteHandler},
}

var socketRoutes = Routes{
	Route{"/newssocket", "", controllers.NewsListSocketHandler},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range socketRoutes {
		router.HandleFunc(route.Path, route.Handler)
	}

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	router.
		PathPrefix("/documentation/").
		Handler(http.StripPrefix("/documentation/", http.FileServer(http.Dir("./documentation"))))

	return router
}
