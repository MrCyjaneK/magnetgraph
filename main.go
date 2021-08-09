package main

import (
	_ "embed"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/steambap/captcha"
	"github.com/zserge/lorca"
)

//go:embed share/style.css
var style []byte
var captchas = make(map[string]string)
var timec = make(map[string]int)

func main() {
	tinit()
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/style.css" {
			rw.Header().Add("Content-Type", "text/css")
			rw.WriteHeader(200)
			rw.Write(style)
		} else {
			rw.Write(message("Welcome to Magnetgraph!", `
<script>
	function openit() {
		window.location.href = "/p/"+document.getElementById("hash").value
	}
</script
<p>Magnetgraph is a simple program that let you create posts that will never get removed,
because it is impossible by design.</p>

<input type="text" id="hash" name="hash" placeholder="Article Hash">
<a onclick=openit>Open</a>

<hr />
<p>Learn more at <a href="https://github.com/mrcyjanek/magnetgraph" target="_blank">GitHub</a>
or <a href="/p/66911b72b19b4d529714ee07d9b72291bf7b04bd/index.md">open an article</a></p>

`))
		}
	})
	http.HandleFunc("/captcha/", func(rw http.ResponseWriter, r *http.Request) {
		data, _ := captcha.New(480, 240, captcha.SetOption(func(o *captcha.Options) {
			o.TextLength = 5
			o.Noise = 1 + rand.ExpFloat64()*10
		}))
		captchas[r.URL.Query().Get("id")] = data.Text
		data.WriteImage(rw)
	})
	http.HandleFunc("/p/", print)
	go func() {
		for {
			cleanCaptchas()
		}
	}()
	go func() {
		if !capt {
			ui, _ := lorca.New("", "", 480, 320)
			ui.Eval(`window.location.href = "http://127.0.0.1:43132"`)
			<-ui.Done()
		}
	}()
	err := http.ListenAndServe(":43132", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanCaptchas() {
	log.Println("Cleaning up captchas")
	for i := range captchas {
		if timec[i] < 5 {
			timec[i]++
		} else {
			delete(captchas, i)
			delete(timec, i)
		}

	}
	time.Sleep(1 * time.Minute)

}
