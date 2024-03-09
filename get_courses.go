package kgp_scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Course struct {
	Name    string
	Code    string
	Credits string
	LTP     string
}

func WriteCoursesToJSON() {
	courses := make(map[string]string)

	codes := []string{"AE", "AE1", "AEA", "AEJ", "AEK", "AG", "AG2", "AG4", "AG5", "AG6", "AG8", "AGA", "AGD", "AGJ", "AGK", "AR", "BS4", "BT", "BT1", "BTA", "BTJ", "BTK", "AI", "CH", "CH1", "CHA", "CHJ", "CHK", "CHP", "CY1", "CY2", "CY3", "CY4", "CE", "CE1", "CE3", "CE4", "CE5", "CE6", "CEA", "CED", "CEJ", "CEK", "CS", "CS2", "CSA", "CSJ", "CSK", "PE", "SD1", "SD1", "SD1", "SD1", "SD1", "SD1", "SD1", "SD1", "SD1", "SD1", "SD1", "SD2", "SD2", "SD3", "SD4", "SD5", "SD6", "SD7", "SD8", "SD9", "EE", "EE2", "EE3", "EE6", "EE7", "EE8", "EEA", "EED", "EEJ", "EEK", "IE", "IE2", "IE4", "IE5", "IEA", "IEJ", "IEK", "EC", "EC3", "EC7", "EC8", "EC9", "ECA", "ECD", "ECJ", "ECK", "EX", "EX2", "EX3", "EX4", "EX6", "EXP", "GG", "GG2", "GG3", "GG5", "GG6", "GGP", "HS", "HS2", "HS3", "IM", "IM1", "IMA", "IMJ", "IMK", "QD1", "QDE", "QDM", "MA", "MA2", "MA3", "MA4", "MA5", "ME", "ME1", "ME3", "ME6", "MEA", "MED", "MEJ", "MEK", "MEP", "MF", "MF1", "MFA", "MFJ", "MFK", "MT", "MT1", "MTA", "MTJ", "MTK", "MI", "MI1", "MI3", "MIA", "MIJ", "MIK", "MIP", "NA", "NA1", "NAA", "NAJ", "NAK", "PH", "PH2", "PH3", "PH4", "PP", "EP", "MM4", "MM5", "MM6", "FP"}

	c := colly.NewCollector(
		colly.AllowedDomains(ERP_DOMAIN),
		colly.MaxDepth(1),
	)

	c.OnHTML("tr", func(h *colly.HTMLElement) {
		s := strings.Split(strings.TrimSpace(h.Text), "\n")
		var ss []string
		for i, a := range s {
			s[i] = strings.TrimSpace(a)
			if len(s[i]) != 0 {
				ss = append(ss, s[i])
			}
		}
		if len(ss) == 5 {
			if ss[1] != "Subject No" {
				courses[ss[1]] = ss[2]
			}
		}
	})

	for _, code := range codes {
		c.Visit(fmt.Sprintf("https://erp.iitkgp.ac.in/ERPWebServices/curricula/CurriculaSubjectsList.jsp?stuType=UG&curr_type=NEW&splCode=%s", code))
	}

	c.Visit(fmt.Sprintf("https://erp.iitkgp.ac.in/ERPWebServices/curricula/CurriculaSubjectsList.jsp?stuType=UG&curr_type=NEW&splCode=%s", "AE"))

	b, err := json.MarshalIndent(courses, "", "	")
	if err != nil {
		fmt.Println("error parsing json from courses map:", err)
	}
	err = os.WriteFile("courses.json", b, 0644)
	if err != nil {
		fmt.Println("error writing json file")
	}

}
