package main

import (
	"fmt"
	"net/http"
)

type HLSTube struct {
	paths   map[string]string
	streams map[string]*Stream
}

func (h *HLSTube) handler(w http.ResponseWriter, r *http.Request) {
	// TODO streams instead
	if h.paths[r.URL.Path] == "" {
		// TODO concurrency limting of some sort?

		fmt.Fprint(w, "starting stream")
		// h.paths[r.URL.Path] = r.URL.Path
		// cmd := exec.Command("bash", "-c", fmt.Sprintf("echo %s && sleep 5", r.URL.Path))
		// err := cmd.Start()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(cmd.Process.Pid)
	} else {
		// TODO
		http.FileServer(http.Dir("")).ServeHTTP(w, r)
	}
}
