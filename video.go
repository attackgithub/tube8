package tube8

import "encoding/xml"

type Tube8Video struct {
	XMLName      xml.Name    `xml:"video"`
	ID           string      `xml:"video_id,attr"`
	Duration     string      `xml:"duration,attr"`
	Title        string      `xml:"title"`
	Views        string      `xml:"views,attr"`
	Rating       string      `xml:"rating,attr"`
	Ratings      string      `xml:"ratings,attr"`
	URL          string      `xml:"url"`
	DefaultThumb string      `xml:"default_thumb,attr"`
	Category     string      `xml:"category,attr"`
	PublishDate  string      `xml:"publish_date,attr"`
	Orientation  string      `xml:"orientation,attr"`
	Tags         Tube8Tags   `xml:"tags"`
	Thumbs       Tube8Thumbs `xml:"thumbs"`
}

type Tube8Tags struct {
	XMLName xml.Name   `xml:"tags"`
	Tag     []Tube8Tag `xml:"tag"`
}

type Tube8Tag struct {
	Tag string `xml:",chardata"`
}

type Tube8Thumbs struct {
	XMLName xml.Name     `xml:"thumbs"`
	Thumb   []Tube8Thumb `xml:"thumb"`
}

type Tube8Thumb struct {
	Thumb  string `xml:",chardata"`
	Size   string `xml:"size,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}
