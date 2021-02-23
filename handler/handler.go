package handler

// Message ...
type Message struct {
    OK               bool        `json:"ok"`
    ErrorCode        int         `json:"error_code,omitempty"`
    ErrorDescription interface{} `json:"description,omitempty"`
    Response         interface{} `json:"response,omitempty"`
}
