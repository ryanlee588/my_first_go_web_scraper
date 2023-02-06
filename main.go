package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"
	

	"github.com/PuerkitoBio/goquery"
	
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile(data, filename string) {
	file, error := os.Create(filename)
	defer file.Close()
	check(error)

	file.WriteString(data)


}

func main() {
	url := "https://techcrunch.com/"

	response, error1 := http.Get(url)

	defer response.Body.Close()

	check(error1)

	if response.StatusCode > 400 {
		fmt.Println("status Code:", response.StatusCode)
	}

	doc, error2 := goquery.NewDocumentFromReader(response.Body)
	check(error2)

	file, error3 := os.Create("posts.csv")
	check(error3)

	writer := csv.NewWriter(file)
	headers := []string{"title", "url", "excerpt"}
	writer.Write(headers)

	doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection){
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url, _ := h2.Find("a").Attr("href")
		excerpt := strings.TrimSpace(item.Find("div.post-block__content").Text())
		posts := []string{title, url, excerpt}

		writer.Write(posts)
	})

	writer.Flush()

	// fmt.Println(river)
	// writeFile(river, "index.html")

}