package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (s *BashfulServer) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name, ok := vars["name"]
	if !ok {
		log.WithFields(log.Fields{
			"request": "GET",
			"vars":    vars,
		}).Error("Invalid request, name missing")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd, err := s.store.GetCommand(name)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"request": "GET",
			"name":    name,
		}).Error("Error retrieving command")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(cmd)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"request": "GET",
			"name":    name,
			"cmd":     cmd,
		}).Error("Could not encode the response")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
