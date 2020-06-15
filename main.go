// A siple webserver to show cougar angels.
// Displays a simple page with background and foreground images selected
// at random from 2 distinct lists.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var (
	host    = flag.String("host", "127.0.0.1", "Host address to listen on.")
	port    = flag.Int("port", 9010, "Port upon which to listen for requests.")
	docroot = flag.String("docroot", "./pics", "DocumentRoot.")
	bgroot  = flag.String("bgroot", "./backgrounds", "BackGround root")

	// Single template makes the milkshake.
	page = template.Must(template.New("page").Parse(pageTemplate))

	image = []string{}

	background = []string{}

	// Set a random seed object up.
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

type handler struct{}

func newHandler() (*handler, error) {
	return &handler{}, nil
}

func selectRandom(s string) string {
	switch {
	case s == "/bk/img.jpg":
		return filepath.Join(*docroot, image[r.Intn(len(image))])
	case s == "/bk/background.jpg":
		return filepath.Join(*bgroot, background[r.Intn(len(background))])
	default:
		return "#fail"
	}
}
func writeFile(w http.ResponseWriter, r *http.Request) {
	fp := r.URL.Path
	if fp == "#fail" {
		log.Println("not a valid path")
		return
	}
	log.Println(fp)
	if _, err := os.Stat(fp); err != nil {
		fmt.Fprintf(w, "<!-- invalid image: %v -->\n", r.URL.Path)
	}
	// Set a cache-control header prior to sending the file.
	w.Header().Add("Cache-Control", "no-cache")

	http.ServeFile(w, r, fp)
}

func index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Background string
		Main       string
	}{
		Background: fmt.Sprintf("/bk/background.jpg?%d", time.Now().Unix()),
		Main:       fmt.Sprintf("/bk/img.jpg?%d", time.Now().Unix()),
	}

	// Set a cache-control header prior to sending the file.
	w.Header().Add("Cache-Control", "no-cache")

	var b bytes.Buffer
	err := page.Execute(&b, data)
	if err != nil {
		fmt.Fprintf(w, "failed to parse template: %v", err)
	}
	fmt.Fprintf(w, b.String())
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/bk"):
		writeFile(w, r)
	case strings.HasPrefix(r.URL.Path, "/"):
		index(w, r)
	}
}

func main() {
	flag.Parse()

	files, err := ioutil.ReadDir(*docroot)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		image = append(image, f.Name())
	}

	bgs, err := ioutil.ReadDir(*bgroot)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range bgs {
		background = append(background, f.Name())
	}

	worker, err := newHandler()
	if err != nil {
		fmt.Printf("failed to create server: %v\n", err)
		return
	}

	h := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", *host, *port),
		Handler: worker,
	}
	log.Fatal(h.ListenAndServe()) // this should be fcgi but I cant be arsed
}
