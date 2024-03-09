package kgp_scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func WriteDepsToJSON() {
	deps := make(map[string]string)

	c := colly.NewCollector(
		colly.AllowedDomains(KGP_DOMAIN),
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		if strings.Contains(h.Attr("href"), DEPARTMENT_URL) && len(h.Attr("href")) == 38 {
			deps[h.Attr("href")[36:]] = h.Text[1:]
		}
	})

	c.Visit(DEPARTMENT_URL + "AE")

	b, err := json.MarshalIndent(deps, "", "	")
	if err != nil {
		fmt.Println("error parsing json from deps map:", err)
	}

	err = os.WriteFile("deps.json", b, 0644)
	if err != nil {
		fmt.Println("error writing json file")
	}
}
