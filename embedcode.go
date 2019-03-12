package tube8

import "encoding/xml"

type Tube8EmbedCode struct {
	XMLName xml.Name    `xml:"embed"`
	EmbedCode  Tube8Code `xml:"code"`
}

type Tube8Code struct {
	Code  string `xml:",cdata"`
}
