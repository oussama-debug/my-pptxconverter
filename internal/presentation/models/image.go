package models

type Image struct {
	NonVisualProps struct {
		CNvPr struct {
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"cNvPr"`
		CNvPicPr struct {
			PicLocks struct {
				NoChangeAspect string `xml:"noChangeAspect,attr"`
			} `xml:"picLocks"`
		} `xml:"cNvPicPr"`
	} `xml:"nvPicPr"`
	BlipFill struct {
		Blip struct {
			Embed string `xml:"embed,attr"`
		} `xml:"blip"`
	} `xml:"blipFill"`
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
	} `xml:"spPr"`
}
