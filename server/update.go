package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/raystlin/bashful/storage"
)

func (s *BashfulServer) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name, ok := vars["name"]
	if !ok {
		log.WithFields(log.Fields{
			"request": "Update",
			"vars":    vars,
		}).Error("Invalid request, name missing")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := struct {
		Status string `json:"status"`
	}{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"request": "Update",
			"vars":    vars,
		}).Error("Could not decode the body")

		w.WriteHeader(http.StatusBadRequest)
		return
	} else if data.Status == "" {
		log.WithFields(log.Fields{
			"request": "Update",
			"vars":    vars,
		}).Error("Invalid body")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = storage.SetCommandStatus(s.store, name, data.Status)
	if err != nil {
		log.WithFields(log.Fields{
			"request": "Update",
			"vars":    vars,
		}).Error("Internal error")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}
