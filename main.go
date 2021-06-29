package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/br3w0r/gamelist-sraper/entity"
	"bitbucket.org/br3w0r/gamelist-sraper/format"
	"github.com/PuerkitoBio/goquery"
)

func OpenURL(url string) *goquery.Document {
	log.Printf("Opening URL: " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making GET request: " + err.Error())
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading document: " + err.Error())
	}

	return doc
}

func main() {
	rootURL := "https://www.metacritic.com"
	doc := OpenURL(rootURL + "/browse/games/score/metascore/all/all/filtered")

	doc.Find("table.clamp-list>tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		if _, exists := s.Attr("class"); exists {
			return
		}

		data := &entity.GameProperties{}
		var exists bool
		var err error

		data.ImageURL, exists = s.Find("td.clamp-image-wrap>a>img").Attr("src")
		if !exists {
			log.Fatal("Image doesn't have a src attribute")
		}

		summary := s.Find("td.clamp-summary-wrap")

		titleRef := summary.Find("a.title")
		data.Name = titleRef.Text()

		// Format paltfrom
		uPlatform := summary.Find("div.clamp-details>div.platform>span.data").Text()
		data.AddPlatform(format.SinglePlatform(uPlatform))

		// Format year_released
		uDate := summary.Find("div.clamp-details>span").Text()
		data.YearReleased, err = format.YearReleased(uDate)
		if err != nil {
			log.Fatal("Failed to format date to year: " + err.Error())
		}

		detailRef, exists := titleRef.Attr("href")
		if !exists {
			log.Fatal("Failed to find link for deatiled description")
		}

		// Enter deatailed description
		detail := OpenURL(rootURL + detailRef)

		// Add all other platforms
		detail.Find("li.product_platforms>span.data>a").Each(func(i int, s *goquery.Selection) {
			data.AddPlatform(s.Text())
		})

		// Add genres
		detail.Find("li.product_genre>span.data").Each(func(i int, s *goquery.Selection) {
			data.AddGenre(s.Text())
		})

		res, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal("Failed to serialize data body")
		}

		fmt.Println(string(res))
	})
}
