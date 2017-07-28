package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"
	"log"
)

//const addr = "http://127.0.0.1:4000"


type HelloHttp struct {}
func (h HelloHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============handle req:", r)
	fmt.Fprint(w, "Hello, Golang!")
}

func main() {
	println("http test:", addr)
	var h HelloHttp
	go func() {
		err := http.ListenAndServe("127.0.0.1:4000", h)
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Second * 2)
	postBody()
	postForm()
	get()
	complexRequest()

	fmt.Println("---------Successful----------")
}

func postBody() {
	body := strings.NewReader("post-body")
	resp, err := http.Post(addr, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Fatalln("post error:", err)
		return
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read post resp body error:", err)
		return
	}
	fmt.Println("post resp body: ", string(rbody))
}

func postForm() {
	resp, err := http.PostForm(addr, url.Values{"uid":{"1"}, "token":{"test_token"}})
	if err != nil {
		log.Fatalln("post form error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read post form resp body error:", err)
		return
	}
	fmt.Println("post form resp body:", string(body))
}

func get() {
	resp, err := http.Get(addr)
	if err != nil {
		log.Fatalln("get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read get body error:", err)
		return
	}
	fmt.Println("get body:", string(body))
}

func complexRequest() {
	client := &http.Client{}
	body := strings.NewReader("complexRequestBody")
	req, err := http.NewRequest("POST", addr, body)
	if err != nil {
		log.Fatalln("new request error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=fuxiao")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("client do req error:", err)
		return
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read client do resp body error:", err)
		return
	}
	fmt.Println("client do resp body:", string(rbody))
}











