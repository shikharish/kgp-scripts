package kgp_scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func WriteDepsToJSON() {
	var deps map[string]string = make(map[string]string)

	c := colly.NewCollector(
		colly.AllowedDomains("www.iitkgp.ac.in"),
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "https://www.iitkgp.ac.in/department/") && len(e.Attr("href")) == 38 {
			deps[e.Attr("href")[36:]] = e.Text[1:]
		}
	})

	c.Visit("https://www.iitkgp.ac.in/department/AE")

	b, err := json.MarshalIndent(deps, "", "	")
	if err != nil {
		fmt.Println("error parsing json from deps map:", err)
	}

	err = os.WriteFile("deps.json", b, 0644)
	if err != nil {
		fmt.Println("error writing json file")
	}
}
