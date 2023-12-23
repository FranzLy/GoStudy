package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
	// 1
	/*_, err := fmt.Fprintln(w, "<h1>Hello Golang!</h1>")
	if err != nil {
		fmt.Printf("return to web err:%v\n", err.Error())
		return
	}*/

	// 2
	b, err := ioutil.ReadFile("E:\\LearningLog\\Coding\\GoStudy\\GoStudy\\gowebstudy\\lesson01\\hello.html")
	fmt.Println(string(b))
	_, err = fmt.Fprintln(w, string(b))
	if err != nil {
		fmt.Printf("read file then return to web err:%v\n", err.Error())
		return
	}

}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http listen server failed. err: %v", err.Error())
		return
	}
}
