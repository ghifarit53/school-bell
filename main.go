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
	// very rough, and everything is hardcoded
	// but this is the best I could do
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/bell-ringing.mp3").Run()
}

func playAnnoStart(w http.ResponseWriter, r *http.Request) {
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/anno-start.mp3").Run()
}

func playAnnoStop(w http.ResponseWriter, r *http.Request) {
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/anno-stop.mp3").Run()
}

func playDangerSiren(w http.ResponseWriter, r *http.Request) {
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/danger-siren.mp3").Run()
}

func playEarthquakeSiren(w http.ResponseWriter, r *http.Request) {
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/earthquake-siren.mp3").Run()
}

func stopAudio(w http.ResponseWriter, r *http.Request) {
	exec.Command("pkill", "mpv").Run()
}

func main() {
	// routing
	http.HandleFunc("/", handlerIndex)

	// regular bell
	http.HandleFunc("/bell-ringing", playBellRinging)
	http.HandleFunc("/anno-start", playAnnoStart)
	http.HandleFunc("/anno-stop", playAnnoStop)

	// danger siren
	http.HandleFunc("/danger-siren", playDangerSiren)
	http.HandleFunc("/earthquake-siren", playEarthquakeSiren)
	http.HandleFunc("/stop-audio", stopAudio)

	// TODO: use port from argv
	var address = "localhost:1337"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
