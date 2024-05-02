package status

import (
	"encoding/json"
	"github.com/brxyxn/go-phonebook-api/internal/response"
	"net/http"

	"github.com/brxyxn/go-phonebook-api/pkg/logger"
)

type Response struct {
	Status string `json:"status"`
}

func (c *Handlers) GetStatus(w http.ResponseWriter, r *http.Request) {
	status := Response{
		Status: "active",
	}

	res, err := json.Marshal(status)
	if err != nil {
		logger.Error(err)
		return
	}

	w = response.SetContentType(w, r, http.StatusOK)

	_, err = w.Write(res)
	if err != nil {
		logger.Error(err)
		return
	}
}
