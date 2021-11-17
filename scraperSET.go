package main

import (
	"encoding/csv"
	"fmt"
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
			// titles = tabla.find('tr')  # titulos
			// valores = tabla.findAll('tr', class_="chico")  # valores
			e.ChildText("td"),
		})
	})

	// 	if ((year == 2010) or (year == 2011) or (year == 2012) or (year == 2013)):
	//     aux = 'https://www.set.gov.py/portal/PARAGUAY-SET/detail?folder-id=repository:collaboration:/sites/PARAGUAY-SET/categories/SET/Informes%20Periodicos/cotizaciones-historicos/ano1&content-id=/repository/collaboration/sites/PARAGUAY-SET/documents/informes-periodicos/cotizaciones/ano1/a-mes-de-mes1'
	//     aux1 = aux.replace("ano1", year)
	//     url = aux1.replace("mes1", month)
	//     print(url)
	// if ((year == 2014) or (year == 2017) or (year == 2019)):
	//     aux = 'https://www.set.gov.py/portal/PARAGUAY-SET/detail?folder-id=repository:collaboration:/sites/PARAGUAY-SET/categories/SET/Informes%20Periodicos/cotizaciones-historicos/ano1/a-mes-de-enero&content-id=/repository/collaboration/sites/PARAGUAY-SET/documents/informes-periodicos/cotizaciones/ano1/a-mes-de-mes1'
	//     aux1 = aux.replace("ano1", year)
	//     url = aux1.replace("mes1", month)
	//     print(url)

	// #no funciona correctamente, se debe ingresar la primera letra en mayuscula
	// if ((year == 2015) or (year == 2016)):
	//     aux = 'https://www.set.gov.py/portal/PARAGUAY-SET/detail?folder-id=repository:collaboration:/sites/PARAGUAY-SET/categories/SET/Informes%20Periodicos/cotizaciones-historicos/ano1/a-mes-de-enero&content-id=/repository/collaboration/sites/PARAGUAY-SET/documents/informes-periodicos/cotizaciones/ano1/A_-_Mes_de_mes1'
	//     aux1 = aux.replace("ano1", year)
	//     url = aux1.replace("mes1", month)
	//     print(url)
	// if ((year == 2018) or (year == 2020)):
	//     aux = 'https://www.set.gov.py/portal/PARAGUAY-SET/detail?folder-id=repository:collaboration:/sites/PARAGUAY-SET/categories/SET/Informes%20Periodicos/cotizaciones-historicos/ano1/a-mes-de-enero&content-id=/repository/collaboration/sites/PARAGUAY-SET/documents/informes-periodicos/cotizaciones/ano1/A%20-%20Mes%20de%20mes1'
	//     aux1 = aux.replace("ano1", year)
	//     url = aux1.replace("mes1", month)
	//     print(url)

	for i := 0; i < 10; i++ {
		fmt.Printf("scraped page: %d\n", i+1)
		c.Visit("https://www.set.gov.py/portal/PARAGUAY-SET/detail?folder-id=repository:collaboration:/sites/PARAGUAY-SET/categories/SET/Informes%20Periodicos/cotizaciones-historicos/2010&content-id=/repository/collaboration/sites/PARAGUAY-SET/documents/informes-periodicos/cotizaciones/2010/a-mes-de-enero")
	}

	log.Printf("scraped completo\n")
	log.Println(c)

}
