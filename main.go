package main

import(
	"fmt"
	"log"
	"net/http"

	"time"
)

var Port = ":5100"

func main(){

	http.HandleFunc("/",ServeFiles)
	fmt.Println("Serving @ : ", "http://127.0.0.1" + Port) // aka http://localhost
	log.Fatal(http.ListenAndServe(Port, nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request){

	switch r.Method{

	case "GET":

		path := r.URL.Path

		fmt.Println(path)

		if path == "/"{

			path = "./static/index.html"
		}
		else{

			path = "." + path
		}

		http.ServeFile(w, r, path)

	case "POST":

		var curTime = time.Now().Format(time.RFC3339)

		r.ParseMultipartForm(0)

		message := r.FormValue("message")

		fmt.Println("----------------------------------")
		fmt.Println("Message from Client: ", message + "  |  " + curTime)
		// respond to client's request
		fmt.Fprintf(w, "Server: %s \n", message + "  |  " + curTime)
	
	default:
		fmt.Fprintf(w,"Request type other than GET or POST not supported")

	}

}


/*

OLD TEST CODE


func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	//w.Write([]byte("<h1 style='color:orange;'>BUENOS</h1>"))
	//w.Write([]byte("<img src='https://pkg.go.dev/static/shared/logo/go-white.svg' alt='SuS'>"))
}


func main() {
	//http.HandleFunc("/hello", Hello)
	//http.Handle("/", http.FileServer(http.Dir("/static"))) // index.html file

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    http.ServeFile(w, r, "static/index.html")
	    w.Write([]byte("<h1 style='color:orange;'>BUENOS</h1>"))
	    dt := time.Now()
		fmt.Printf("\nServer request made: %s /\n", r.Method, dt.String())


	})

	log.Fatal(http.ListenAndServe(":5100", nil))
}

*/