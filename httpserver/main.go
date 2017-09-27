package main

import (
	. "asdf"
	"io/ioutil"
	"net/http"
)

const (
	PORT = "8740"
)

type httpServer struct{}

func (me *httpServer) post(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	Log.Debug("http server recv len:%d, body:%s", len(body), string(body))
	r.Body.Close()

	err := 0

	if 0 == len(body) {
		err = StdErrZeroBody
	}

	if 0 == err {
		HttpError(w, 0, nil)
	} else {
		HttpError(w, err, nil)
	}
}

func (me *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	switch r.Method {
	case "POST":
		me.post(w, r)
	default:
		HttpError(w, StdErrNoSupport, nil)
		return
	}
}

func (me *httpServer) listen() {
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: me,
	}

	Log.Debug("http server listen ...")
	err := server.ListenAndServe()
	Log.Debug("http server listen ...", err.Error())
}

var httpserver = &httpServer{}

func main() {
	httpserver.listen()
}
