package ip

import "net/http"

func IP(req *http.Request) string {
	addr := req.Header.Get("X-Real-Ip")
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
	}
	if addr == "" {
		addr = req.RemoteAddr
	}
	return addr
}
