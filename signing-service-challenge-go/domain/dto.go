package domain

//dtos for sending responses and receiving requests to/from users in the api/handler layer

type CreateSignatureDeviceRequest struct {
	Algorithm string `json:"algorithm"`
	Label     string `json:"label"`
}

type CreateSignatureDeviceResponse struct {
	Status string `json:"status"`
}

type SignatureRequest struct {
	DeviceID string `json:"device_id"`
	Data     string `json:"data"`
}

type SignatureResponse struct {
	Signature  string `json:"signature"`
	SignedData string `json:"signed_data"`
}
