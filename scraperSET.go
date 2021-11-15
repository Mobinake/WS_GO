package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	//crear archivo donde se guarda losdatos temporalmente
	fName := "set.csv"
	file, err := os.Create(fName)
	//comprobar errores durante la creacion del archivo
	if err != nil {
		log.Fatalf("No se puede leer el archivo. error :%q", err)
		return
	}
	defer file.Close()

	//escribir en el csv
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//dominios permitidos al scraper
	c := colly.NewCollector(
		colly.AllowedDomains("set.gov.py"),
	)
	//localizamos donde se encuentra la informacion en la tabla
	c.OnHTML(".chico", func(e *colly.HTMLElement) {

		writer.Write([]string{
			e.ChildText("table"),
			// titles = tabla.find('tr')  # titulos
			// valores = tabla.findAll('tr', class_="chico")  # valores
			//e.ChildText("table"),
		})
	})

	//for i := 0; i < 312; i++ {
	// fmt.Printf("scraped page: %d\n", i)

	c.Visit("https://www.set.gov.py/portal/PARAGUAY-SET/detail?folder-id=repository:collaboration:/sites/PARAGUAY-SET/categories/SET/Informes%20Periodicos/cotizaciones-historicos/2010&content-id=/repository/collaboration/sites/PARAGUAY-SET/documents/informes-periodicos/cotizaciones/2010/a-mes-de-enero")
	//}

	//log.Printf("scraped completo\n")
	log.Println(c)

}
