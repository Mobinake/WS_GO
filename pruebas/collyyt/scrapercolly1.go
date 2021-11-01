package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	//create file
	fName := "data.csv"
	file, err := os.Create(fName)
	//check errors during the creating process
	if err != nil {
		log.Fatalf("No se puede leer el archivo. error :%q", err)
		return
	}
	defer file.Close()

	//escribir en el csv
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {

		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	for i := 0; i < 312; i++ {
		// fmt.Printf("scraped page: %d\n", i)

		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("scraped completo\n")
	log.Println(c)

}
