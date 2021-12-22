package alive

import (
	"encoding/json"
	"net/http"
)

var (
	BuildVersion = "none"
	Timestamp    = "none"
)

type Body struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	body := Body{
		Version:   BuildVersion,
		Timestamp: Timestamp,
	}
	res, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
