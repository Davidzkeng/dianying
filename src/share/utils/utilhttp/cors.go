package utilhttp

import "net/http"

var (
	cors = map[string]bool{"*":true}
)

func Cors(w http.ResponseWriter, r *http.Request) http.ResponseWriter  {
	//跨域处理
	if origin := r.Header.Get("Origin");cors[origin]{
		w.Header().Set("Access-Control-Allow-Origin",origin)
	}else if len(origin) > 0 && cors["*"] {
		w.Header().Set("Access-Control-Allow-Origin",origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	return w
}