package models

type Run struct {
	RPr RunProperties `xml:"rPr"`
	T   string        `xml:"t"`
}

type RunProperties struct {
	Sz      int    `xml:"sz,attr"`
	B       *bool  `xml:"b"`
	I       *bool  `xml:"i"`
	U       *bool  `xml:"u"`
	SrgbClr string `xml:"solidFill>srgbClr>val,attr"`
}
