package status

import (
	"encoding/json"
	"github.com/brxyxn/go-phonebook-api/internal/response"
	"net/http"

	"github.com/brxyxn/go-phonebook-api/pkg/logger"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (c *Handlers) GetStatus(w http.ResponseWriter, r *http.Request) {
	status := Response{
		Status: "active",
		Error:  "",
	}

	res, err := json.Marshal(status)
	if err != nil {
		logger.Error(err)
		return
	}

	response.Success(w, r, status, "success")

	_, err = w.Write(res)
	if err != nil {
		logger.Error(err)
		return
	}
}
