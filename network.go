package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func serveTurtleNet(t *Turtle, h *History, reader *bufio.Reader) {
	// fmt.Println("In Server Turtle Net")
	// openBrowser("http://localhost:8080/")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
		cmdInput(t, h, reader, r.URL.Query().Get("query"))
		fileBytes, err := ioutil.ReadFile("./static/turtle.png")
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
	})
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
