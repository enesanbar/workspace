package main

import (
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/enesanbar/workspace/golang/microservices/grpcjson/keyvalue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// GetHandler wraps our RPC Get call
func (c *Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	kv := keyvalue.GetKeyValueRequest{Key: key}

	gresp, err := c.Get(r.Context(), &kv)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Errorf("failed to get: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
