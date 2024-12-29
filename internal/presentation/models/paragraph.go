package models

type Paragraph struct {
	PPr ParagraphProperties `xml:"pPr"`
	R   []Run               `xml:"r"`
}

type ParagraphProperties struct {
	Align       string `xml:"algn,attr"`
	LineSpacing int    `xml:"lnSpc>spcPct>val,attr"`
}
