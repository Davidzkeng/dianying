package main

import (
	"dianying/src/domain/apid"
	"dianying/src/share/config"
	"dianying/src/share/utils/utilhttp"
	"encoding/json"
	"github.com/micro/go-micro/config/cmd"
	"io/ioutil"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/",handleRPC)
	http.ListenAndServe(":8082",mux)
}
func handleRPC(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/"{
		w.Write([]byte("It work!"))
		return
	}
	w = utilhttp.Cors(w,r)
	if r.Method == "OPTIONS"{
		return
	}
	handleJONRPC(w,r)
	return
}
func handleJONRPC(w http.ResponseWriter, r *http.Request) {
	service,method := apid.PathToReceiver(config.Namespace,r.URL.Path)

	br,_ := ioutil.ReadAll(r.Body)

	request := json.RawMessage(br)

	var response json.RawMessage
	req := (*cmd.DefaultOptions().Client).NewRequest(service,method,&request)
}


