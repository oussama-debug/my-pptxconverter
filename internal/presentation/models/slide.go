package models

import "encoding/xml"

type Slide struct {
	XMLName    xml.Name   `xml:"sld"`
	Background Background `xml:"cSld>bg"`
	ShapeTree  ShapeTree  `xml:"cSld>spTree"`
}
