package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFromUrl(url string, path ...string) {
	re := regexp.MustCompile("[0-9]+")
	filename := re.FindStringSubmatch(url)[0]
	p := "download"
	fmt.Println("Downloading", url, "to", filename)

	// TODO: check file existence first with io.IsExist
	if len(path) > 0 {
		p = path[0]

	}
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}

	output, err := os.Create("download/" + filename + ".mp4")
	if err != nil {
		fmt.Println("Error while creating", filename, "-", err)
		return
	}
	defer output.Close()
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
	}

	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)

	response, _ := client.Do(req)

	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func ReadFile(filePath string) []string {
	var myslice []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		link := scanner.Text() + "download/?search_query=&tracking_id=jzhpmbsmrvn"
		link = strings.Replace(link, "/zh-cn", "", 1)
		fmt.Println(link)

		myslice = append(myslice, link)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return myslice
}
func main() {
	filePath := os.Args[1]
	path := os.Args[2]
	myslice := ReadFile(filePath)
	//var wg sync.WaitGroup
	for index, url := range myslice {
		//wg.Add(1)
		defer func() {
			r := recover()
			fmt.Println("recovered:", r)
		}()
		fmt.Printf("downloading %d\n", index)
		downloadFromUrl(url, path)
	}

}
