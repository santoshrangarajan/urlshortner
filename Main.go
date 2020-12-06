package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

var (
	urlMap          map[int]string
	shortnerBaseURL string
)

//urlMap := make(map[int32]string)

////// initialize reader at start
///////// get/new

func initializeURLShortner() {
	urlMap = make(map[int]string)
	shortnerBaseURL = "http://myshortner.com/"
}

func readFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL to be shortened: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	return text
}

func generateHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	n := h.Sum32()
	urlMap[int(n)] = s
	nstr := strconv.Itoa(int(n))
	sEnc := b64.StdEncoding.EncodeToString([]byte(nstr))
	//println("nstr", nstr)
	return sEnc
}

func generateShortURL(URL string) string {
	sEnc := generateHash(URL)
	return shortnerBaseURL + sEnc
}

func generateLongURL(URL string) string {

	sEnc := strings.ReplaceAll(URL, shortnerBaseURL, "")

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	nstr := string(sDec)
	n, err := strconv.Atoi(nstr)

	if err != nil {
		println("decode", "cannot decode url")
	}

	url := urlMap[n]

	return url
}

func main() {

	initializeURLShortner()

	URL := readFromTerminal()
	println("original url:", URL)

	u := generateShortURL(URL)

	println("short url:", u)

	d := generateLongURL(u)
	println("long url:", d)

}
