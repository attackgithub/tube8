package tube8

import "encoding/xml"

type Tube8SingleVideo struct {
	XMLName xml.Name   `xml:"root"`
	Videos  Tube8Video `xml:"video"`
}
