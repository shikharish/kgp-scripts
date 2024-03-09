package kgp_scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func WriteProgramsToJSON() {
	programs := make(map[string]string)

	c := colly.NewCollector(
		colly.AllowedDomains(ERP_DOMAIN),
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		url := h.Attr("href")
		i := strings.Index(url, "splCode")
		var code string
		if i != -1 {
			code = url[i+8 : i+11]
			if code[2:] == "&" {
				code = code[:2]
			}
		}
		programs[code] = h.Text
	})

	c.Visit(UG_NEW_CURRICULA_URL)

	b, err := json.MarshalIndent(programs, "", "	")
	if err != nil {
		fmt.Println("error parsing json from programs map:", err)
	}

	err = os.WriteFile("programs.json", b, 0644)
	if err != nil {
		fmt.Println("error writing json file")
	}
}
