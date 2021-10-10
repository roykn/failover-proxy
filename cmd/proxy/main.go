package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var body []byte

func main() {
	url, _ := url.Parse("http://localhost:8080")
	proxy := httputil.NewSingleHostReverseProxy(url)

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			p.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, e error) {
				//log.Println(e)
				rw.Write(body)
			}

			p.ModifyResponse = func(r *http.Response) error {
				body, _ := ioutil.ReadAll(r.Body)
				r.Body = ioutil.NopCloser(bytes.NewReader(body))
				log.Printf("%v\n", string(body))
				return nil
			}
			p.ServeHTTP(w, r)
		}
	}
	http.HandleFunc("/", handler(proxy))

	log.Println("Starting proxy on :8081")
	err := http.ListenAndServe(":8081", nil)
	log.Fatal(err)
}
