package scraper

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/br3w0r/gamelist-scraper/entity"
	"bitbucket.org/br3w0r/gamelist-scraper/format"
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

func Scrape(c chan entity.ScraperResp, pagesToScrape int) {
	rootURL := "https://www.metacritic.com"

	for i := 0; i < pagesToScrape; i++ {
		doc := OpenURL(rootURL + fmt.Sprintf("/browse/games/score/metascore/all/all/filtered?page=%d", i))

		var err error = nil
		doc.Find("table.clamp-list>tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
			if _, exists := s.Attr("class"); exists || err != nil {
				return
			}

			data := entity.GameProperties{}
			var exists bool

			data.ImageUrl, exists = s.Find("td.clamp-image-wrap>a>img").Attr("src")
			if !exists {
				err = errors.New("image doesn't have a src attribute")
				c <- entity.ScraperResp{Game: entity.GameProperties{}, Err: err}
				return
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
				err = errors.New("failed to format date to year: " + err.Error())
				c <- entity.ScraperResp{Game: entity.GameProperties{}, Err: err}
				return
			}

			detailRef, exists := titleRef.Attr("href")
			if !exists {
				err = errors.New("failed to find link for deatiled description")
				c <- entity.ScraperResp{Game: entity.GameProperties{}, Err: err}
				return
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

			c <- entity.ScraperResp{Game: data, Err: nil}
		})
	}
	close(c)
}
