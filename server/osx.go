package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const uriElCapitan = "https://support.apple.com/en-au/HT205204"
const uriYosemite = "https://support.apple.com/en-au/HT205218"

func ProcessOSX() {
	processAppleTabularData(uriElCapitan, "OS X El Capitan", "osx-elcapitan")
	processAppleTabularData(uriYosemite, "OS X Yosemite", "osx-yosemite")
	//TODO Mavericks
}

func processAppleTabularData(uri string, version string, filenamePrefix string) {
	html, err := goquery.NewDocument(uri)

	if err != nil {
		log.Printf("Unable to process %s\n", version)
	}

	var types = [3]string{"trusted", "alwaysask", "blocked"}

	for _, t := range types {

		var certs []Cert

		counter := 0
		html.Find(fmt.Sprintf("#%s tr", t)).Each(func(_ int, tr *goquery.Selection) {
			counter++

			if counter == 1 {
				return
			}

			tds := tr.Find("td")

			certs = append(certs, Cert{
				CertificateName: strings.TrimSpace(tds.Eq(0).Text()),
				IssuedBy:        strings.TrimSpace(tds.Eq(1).Text()),
				Type:            strings.TrimSpace(tds.Eq(2).Text()),
				KeySize:         strings.TrimSpace(tds.Eq(3).Text()),
				SigAlg:          strings.TrimSpace(tds.Eq(4).Text()),
				SerialNumber:    strings.TrimSpace(tds.Eq(5).Text()),
				Expires:         strings.TrimSpace(tds.Eq(6).Text()),
				EVPolicy:        strings.TrimSpace(tds.Eq(7).Text()),
			})
		})

		j, err := json.Marshal(certs)

		if err != nil {
			log.Printf("Unable to process %s - %s\n", version, t)
			return
		}

		WriteFile(fmt.Sprintf("%s-%s.json", filenamePrefix, t), j)

		if err != nil {
			log.Println("Unable to process %s - %s\n", version, t)
			return
		}
	}
}
