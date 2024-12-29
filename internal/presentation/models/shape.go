package models

type Shape struct {
	NonVisualProps struct {
		CNvPr struct {
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"cNvPr"`
	} `xml:"nvSpPr"`
	SpPr struct {
		XFrm struct {
			Off struct {
				X int `xml:"x,attr"`
				Y int `xml:"y,attr"`
			} `xml:"off"`
			Ext struct {
				Cx int `xml:"cx,attr"`
				Cy int `xml:"cy,attr"`
			} `xml:"ext"`
		} `xml:"xfrm"`
		PrstGeom struct {
			Prst string `xml:"prst,attr"`
		} `xml:"prstGeom"`
		SolidFill struct {
			SrgbClr string `xml:"srgbClr>val,attr"`
		} `xml:"solidFill"`
		Ln struct {
			W       int    `xml:"w,attr"`
			SrgbClr string `xml:"solidFill>srgbClr>val,attr"`
		} `xml:"ln"`
	} `xml:"spPr"`
	TxBody struct {
		P []Paragraph `xml:"p"`
	} `xml:"txBody"`
}

type ShapeTree struct {
	Shapes []Shape `xml:"sp"`
	Images []Image `xml:"pic"`
}
