package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esnchez/coding-challenges/signing-service-challenge/domain"
	"github.com/esnchez/coding-challenges/signing-service-challenge/persistence"
	"github.com/google/uuid"
)

// TODO: REST endpoints ...

// handleCreateSignatureDevice handles the creation of a signature device and stores it in the repository
// only POST method allowed, checks and validates request values and returns response/errors and http codes
func (s *Server) handleCreateSignatureDevice(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(rw, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	req := new(CreateSignatureDeviceRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		WriteErrorResponse(rw, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		})
		return
	}

	dev, err := s.svc.CreateSignatureDevice(req.Algorithm, req.Label)
	if err != nil {

		switch err {
		case domain.ErrInvalidAlgorithm:
			WriteErrorResponse(rw, http.StatusBadRequest, []string{
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			})
			return

		case persistence.ErrSaveFailure:
			WriteErrorResponse(rw, http.StatusInternalServerError, []string{
				http.StatusText(http.StatusInternalServerError),
				err.Error(),
			})
			return

		default:
			WriteErrorResponse(rw, http.StatusNotImplemented, []string{
				http.StatusText(http.StatusNotImplemented),
				err.Error(),
			})
			return
		}
	}

	resp := CreateSignatureDeviceResponse{
		Status: fmt.Sprintf("Signature device object created with ID: %s", dev.ID),
	}

	WriteAPIResponse(rw, http.StatusOK, resp)
}

// handleSignTransaction handles the signature of arbitrary data after retrieving successfully a signature device from persistence via its unique id
// only POST method allowed, checks and validates request values and returns response/errors and http codes
func (s *Server) handleSignTransaction(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(rw, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	req := new(SignatureRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		WriteErrorResponse(rw, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		})
		return
	}

	id, err := uuid.Parse(req.DeviceID)
	if err != nil {
		WriteErrorResponse(rw, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
			fmt.Sprintf("Id format is not valid: %s", err.Error()),
		})
	}

	sig, data, err := s.svc.SignTransaction(id, req.Data)
	if err != nil {

		switch err {
		case domain.ErrInvalidDataToBeSigned:
			WriteErrorResponse(rw, http.StatusBadRequest, []string{
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			})
			return

		case persistence.ErrNotFound:
			WriteErrorResponse(rw, http.StatusBadRequest, []string{
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			})
			return

		case domain.ErrSignOperation:
			WriteErrorResponse(rw, http.StatusInternalServerError, []string{
				http.StatusText(http.StatusInternalServerError),
				err.Error(),
			})
			return

		default:
			WriteErrorResponse(rw, http.StatusNotImplemented, []string{
				http.StatusText(http.StatusNotImplemented),
				err.Error(),
			})
			return
		}

	}

	resp := SignatureResponse{
		Signature:  string(sig),
		SignedData: string(data),
	}

	WriteAPIResponse(rw, http.StatusOK, resp)
}

func (s *Server) handleGetAllDevices(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteErrorResponse(rw, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	list, err := s.svc.GetAllDevices()
	if err != nil {
		WriteErrorResponse(rw, http.StatusInternalServerError, []string{
			http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	WriteAPIResponse(rw, http.StatusOK, list)
}
