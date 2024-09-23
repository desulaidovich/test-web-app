package domain

type (
	Request struct {
		GUID string `json:"guid"`
	}
	Response struct {
		Error   string             `json:"error,omitempty"`
		Message string             `json:"message,omitempty"`
		Token   *map[string]string `json:"token,omitempty"`
	}
)
