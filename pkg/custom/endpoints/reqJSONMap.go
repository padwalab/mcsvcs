package endpoints

type TestStatusRequest struct{}

type TestStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
