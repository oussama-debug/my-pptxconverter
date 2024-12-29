package models

type Background struct {
	BGPr struct {
		SolidFill struct {
			SrgbClr string `xml:"srgbClr>val,attr"`
		} `xml:"solidFill"`
	} `xml:"bgPr"`
}
