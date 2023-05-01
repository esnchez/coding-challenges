package api

import (
	"net/http"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type CreateSignatureDeviceResponse struct {
	ID      uuid.UUID `json:"id"`
	Label   string    `json:"label"`
	Status  string    `json:"status"`
	Version string    `json:"version"`
}

type SignatureResponse struct {
	// ID      uuid.UUID `json:"id"`
	//some data after tx generated?
	Status  string `json:"status"`
	Version string `json:"version"`
}

// TODO: REST endpoints ...

// Description
func (s *Server) CreateSignatureDevice(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(rw, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	//from where do we take label and algortihm? url or request body?
	//change hello for that info
	//svc?
	sd := domain.NewSigDevice("Hello")

	resp := CreateSignatureDeviceResponse{
		ID:      sd.GetID(),
		Label:   sd.GetLabel(),
		Status:  "Signature Device object successfully created",
		Version: "v0",
	}

	WriteAPIResponse(rw, http.StatusOK, resp)
}

// Description
func (s *Server) SignTransaction(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(rw, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	//from where do we take device id and data? url or request body?
	//change hello for that info
	//svc?
	// sd := domain.NewSigDevice("Hello")

	resp := SignatureResponse{
		// ID:      sd.GetID(),
		Status:  "Signature completed",
		Version: "v0",
	}

	WriteAPIResponse(rw, http.StatusOK, resp)
}
