package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stdout, r.RequestURI)

		b, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintln(os.Stdout, string(b))

		w.Write([]byte("success"))
	}

	router.Path("/job/notifications").HandlerFunc(f)

	http.ListenAndServe(":60080", router)
}
