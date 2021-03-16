package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = "./static/index.html"
    }
    http.ServeFile(w, r, p)
}

func playBellRinging(w http.ResponseWriter, r *http.Request) {
	exec.Command("mpv", "./static/audio/bell-ringing.mp3").Run()
}

func playDangerSiren(w http.ResponseWriter, r *http.Request) {
	exec.Command("mpv", "./static/audio/danger-siren.mp3").Run()
}

func playEarthquakeSiren(w http.ResponseWriter, r *http.Request) {
	exec.Command("mpv", "./static/audio/earthquake-siren.mp3").Run()
}

func main() {
	// routing
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/bell-ringing", playBellRinging)
	http.HandleFunc("/danger-siren", playDangerSiren)
	http.HandleFunc("/earthquake-siren", playEarthquakeSiren)

	var address = "localhost:1337"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
