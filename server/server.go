package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/raystlin/bashful/storage"
)

const (
	APIPrefix     = "/api/v1"
	CommandPrefix = APIPrefix + "/cmd"
)

type BashfulServer struct {
	store storage.Store
}

func NewBashfulServer(store storage.Store) *BashfulServer {
	return &BashfulServer{
		store: store,
	}
}

func (bs *BashfulServer) newRouter() *mux.Router {
	router := mux.NewRouter()
	s := router.PathPrefix(CommandPrefix).Subrouter()

	// Commands
	s.HandleFunc("/{name}", bs.get).Methods("GET")
	s.HandleFunc("/{name}", bs.update).Methods("POST")
	s.HandleFunc("/", bs.list).Methods("GET")

	return router
}

func (bs *BashfulServer) ListenAndServe(addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: bs.newRouter(),
	}

	return srv.ListenAndServe()
}

func (bs *BashfulServer) ListenAndServeTLS(addr, certPath, keyPath string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: bs.newRouter(),
	}

	return srv.ListenAndServeTLS(certPath, keyPath)
}
