package tube8

import (
	"encoding/xml"
)

type Tube8SearchResult struct {
	XMLName xml.Name    `xml:"root"`
	Videos  Tube8Videos `xml:"videos"`
}

type Tube8Videos struct {
	XMLName xml.Name     `xml:"videos"`
	Video   []Tube8Video `xml:"video"`
}
