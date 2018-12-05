package main

import (
	"fmt"
	"io"
	"golang.org/x/net/html"
	"strings"
)



type htmlReader struct{
	s        string
	i        int64 
	prevRune int
};


func (r *htmlReader) Read(b []byte) (n int, err error){
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func newReader(htmlStr string) (io.Reader){
	return &htmlReader{htmlStr,0,-1};
}

func newReaderFunc(htmlStr string) (io.Reader){
	var hr htmlReader
	hr.Read([]byte(htmlStr))
	return &hr;
}


func main(){
	htmlStr := [] byte(`<a href='www.google.com'>aaaa</a>`)
	doc, err := html.Parse(newReader(string(htmlStr)))
	if err != nil {
		fmt.Println("there is an error")
	}
	fmt.Println(doc.Type)
	fmt.Println(doc.Data)

	docR, _ := html.Parse(strings.NewReader(string(htmlStr)))

	fmt.Println(docR.Type)
	fmt.Println(docR.Data)
}