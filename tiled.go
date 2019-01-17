package tiled

import "encoding/xml"

type Tileset struct {
	XMLName      xml.Name `xml:"tileset"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"version,attr"`
	Tiledversion string   `xml:"tiledversion,attr"`
	Source       string   `xml:"source,attr"`
	Name         string   `xml:"name,attr"`
	Tilewidth    int      `xml:"tilewidth,attr"`
	Tileheight   int      `xml:"tileheight,attr"`
	Tilecount    int      `xml:"tilecount,attr"`
	Firstgid     int      `xml:"firstgid,attr"`
	Columns      int      `xml:"columns,attr"`
	Image        struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Width  string `xml:"width,attr"`
		Height string `xml:"height,attr"`
	} `xml:"image"`
	Tiles []TilesetTile `xml:"tile"`
}

type TilesetTile struct {
	Text        string      `xml:",chardata"`
	ID          int         `xml:"id,attr"`
	Type        string      `xml:"type,attr"`
	Objectgroup Objectgroup `xml:"objectgroup"`
	Properties  struct {
		Text     string     `xml:",chardata"`
		Property []Property `xml:"property"`
	} `xml:"properties"`
}

type Properties struct{}

type Layer struct {
	Text       string `xml:",chardata"`
	ID         int    `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	Width      int    `xml:"width,attr"`
	Height     int    `xml:"height,attr"`
	Properties []struct {
		Text     string     `xml:",chardata"`
		Property []Property `xml:"property"`
	} `xml:"properties"`
	Data LayerData `xml:"data"`
}

type Property struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	Type  string `xml:"type,attr"`
}

type LayerTile struct {
	Text string `xml:",chardata"`
	Gid  int    `xml:"gid,attr"`
}

type LayerData struct {
	Text  string      `xml:",chardata"`
	Tiles []LayerTile `xml:"tile"`
}

type Objectgroup struct {
	Text    string   `xml:",chardata"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Objects []Object `xml:"object"`
}

type Object struct {
	Text       string `xml:",chardata"`
	ID         int    `xml:"id,attr"`
	Gid        uint32 `xml:"gid,attr"`
	Name       string `xml:"name,attr"`
	Type       string `xml:"type,attr"`
	X          int    `xml:"x,attr"`
	Y          int    `xml:"y,attr"`
	Width      int    `xml:"width,attr"`
	Height     int    `xml:"height,attr"`
	Rotation   int    `xml:"rotation,attr"`
	Properties struct {
		Text     string     `xml:",chardata"`
		Property []Property `xml:"property"`
	} `xml:"properties"`
}

// Embedded
type tmxData struct {
	XMLName      xml.Name      `xml:"map"`
	Text         string        `xml:",chardata"`
	Version      string        `xml:"version,attr"`
	Tiledversion string        `xml:"tiledversion,attr"`
	Orientation  string        `xml:"orientation,attr"`
	Renderorder  string        `xml:"renderorder,attr"`
	Width        int           `xml:"width,attr"`
	Height       int           `xml:"height,attr"`
	Tilewidth    int           `xml:"tilewidth,attr"`
	Tileheight   int           `xml:"tileheight,attr"`
	Infinite     string        `xml:"infinite,attr"`
	Nextlayerid  string        `xml:"nextlayerid,attr"`
	Nextobjectid string        `xml:"nextobjectid,attr"`
	Tilesets     []Tileset     `xml:"tileset"`
	Layers       []Layer       `xml:"layer"`
	Objectgroups []Objectgroup `xml:"objectgroup"`
}
