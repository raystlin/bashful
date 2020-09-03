package server

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (s *BashfulServer) list(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string)

	list, err := s.store.ListCommands()
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"request": "LIST",
		}).Error("Error listing commands")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, cmd := range list {
		response[cmd.Name] = cmd.Status
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"request": "LIST",
			"list":    list,
			"map":     response,
		}).Error("Could not encode the response")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
