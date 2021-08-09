package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

//go:embed share/template.html
var template string

//go:embed share/templatefull.html
var templatefull string

func print(rw http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) < 3 {
		fmt.Fprint(rw, "incorrect argument used path should be /p/magneturi")
		return
	}
	if len(p) >= 3 {
		_, ok := torrents[p[2]]
		if !ok {
			// Captcha
			if !checkCaptcha(r.URL.Query().Get("id"), r.URL.Query().Get("captcha")) {
				rw.Write(message("Are you a robot?", `Fetching things from BitTorrent network is a hard thing to do, to make sure that no DoS attack will happen please complete this captcha:
<form>
    <input type="hidden" name="id" value="`+p[2]+r.RemoteAddr+`"><br />
	<img src="/captcha?id=`+p[2]+r.RemoteAddr+`"><br />
	<input type="text" name="captcha"><br />
	<input type="submit">
</form>`))
				return
			}
			uri := getMagnetURI(p[2])
			T, err := cl.AddMagnet(uri)
			if err != nil {
				log.Println(err)
				fmt.Fprint(rw, err.Error())
				return
			}
			T.SetDisplayName(p[2])
			torrents[p[2]] = T
			rw.Write(message("Fetching article", "The article's metadata is now being downloaded from the BitTorrent network (and a couple of trackers), please be patient, it should load within a few seconds.<meta http-equiv=\"refresh\" content=\"1\">"))
			return
		}
		<-torrents[p[2]].GotInfo()
		b := torrents[p[2]].BytesMissing()
		if b == 0 {
			if len(p) == 3 {
				rw.Header().Set("Location", "/p/"+p[2]+"/index.md")
				rw.WriteHeader(301)
				return
			}
			path := strings.Join(p[3:], "/")
			files := torrents[p[2]].Files()
			for i := range files {
				rw.Header().Add("Content-Security-Policy", "default-src 'self'")
				if files[i].DisplayPath() == path {
					tr := files[i].NewReader()
					bs := make([]byte, files[i].FileInfo().Length)
					tr.SetReadahead(files[i].FileInfo().Length)
					j, err := tr.Read(bs)
					if err != nil && int64(j) != files[i].FileInfo().Length {
						rw.Write(message("Unable to read file.", err.Error()+" "+strconv.Itoa(j)+" "+strconv.Itoa(len(bs))))
					}
					if strings.HasSuffix(files[i].DisplayPath(), ".md") {
						extensions := parser.CommonExtensions | parser.Footnotes
						ps := parser.NewWithExtensions(extensions)
						rw.Write(fullpage(bluemonday.UGCPolicy().SanitizeBytes(markdown.ToHTML(bs, ps, nil))))
					} else {
						rw.Write(bs)
					}
					return
				}
			}
			rw.Write(message("404 - Not Found", "Unable to find requested file."))
			return
			// Actually load content
		} else if b < 256*1024 {
			// Let's download this thing, it's small enough
			rw.Write(message("Fetching article", "The article is now being downloaded from the BitTorrent network, please be patient, it should load within a few seconds.<meta http-equiv=\"refresh\" content=\"5\">"))
			torrents[p[2]].SetDisplayName(p[2])
			go torrents[p[2]].DownloadAll()
			return
		} else {
			// S-Senpai? It's too big!
			rw.Write(message("Unable to fetch article", "Oups! I was unable to fetch this article for you, it was too big! Max allowed size of all files is 256kb (it's a subject to change)"))
			return
		}
	}
}

func getMagnetURI(inputmagnet string) string {
	magnet := "magnet:?xt=urn:btih:" + inputmagnet
	// TEMP: Add tracker
	magnet = magnet + "&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce"
	return magnet
}

func message(title string, content string) []byte {
	return []byte(fmt.Sprintf(template, title, content, "", ""))
}

func fullpage(content []byte) []byte {
	return []byte(fmt.Sprintf(string(templatefull), content))
}
