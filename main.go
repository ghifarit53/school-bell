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

func playSchoolTime(w http.ResponseWriter, r *http.Request) {
	// very rough, and everything is hardcoded
	// but this is the best I could do for now
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/school-time.mp3").Run()
}

func playRecessTime(w http.ResponseWriter, r *http.Request) {
	exec.Command("pkill", "mpv").Run()
	exec.Command("mpv", "--no-video", "--loop-playlist=no", "./static/audio/recess-time.mp3").Run()
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
	http.HandleFunc("/play-school-time", playSchoolTime)
	http.HandleFunc("/play-recess-time", playRecessTime)
	http.HandleFunc("/play-anno-start", playAnnoStart)
	http.HandleFunc("/play-anno-stop", playAnnoStop)

	// danger siren
	http.HandleFunc("/play-danger-siren", playDangerSiren)
	http.HandleFunc("/play-earthquake-siren", playEarthquakeSiren)
	http.HandleFunc("/stop-audio", stopAudio)

	// TODO: use port from argv
	var address = "localhost:8012"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
