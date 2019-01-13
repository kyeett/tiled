package main

import (
	"encoding/xml"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/colornames"

	"github.com/peterhellberg/gfx"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const tmxPath = "world6.tmx"

func main() {

	data, err := ioutil.ReadFile(tmxPath)
	if err != nil {
		log.Fatal("read tmx", err)
	}

	worldMap := Map{}
	err = xml.Unmarshal(data, &worldMap)
	if err != nil {
		log.Fatal("unmarshal xml", err)
	}

	screenWidth, screenHeight := 600, 600

	f, err := os.Open(worldMap.Tilesets[0].Image.Source)
	if err != nil {
		log.Fatal("", err)
	}
	spriteImg, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("", err)
	}
	sImg := spriteImg.(*image.NRGBA)
	addLabel(sImg, 10, 10, "test")

	img := gfx.NewImage(screenWidth, screenHeight, color.Black)
	for _, layer := range worldMap.Layers {

		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				ID := layer.tileID(x, y)
				if ID == 0 {
					continue
				}

				wx, wy := worldMap.Tilesets[0].tilesheetCoords(ID)
				lx, ly := x*worldMap.Tilewidth, y*worldMap.Tileheight
				// fmt.Println(ID, lx, ly)
				// fmt.Println(wx, wy)

				sRect := image.Rect(wx, wy, wx+worldMap.Tilewidth, wy+worldMap.Tileheight)
				// fmt.Println(sRect)
				dstRect := image.Rect(lx, ly, worldMap.Tilewidth, worldMap.Tileheight)
				fmt.Println(sRect)

				draw.Draw(img, dstRect, sImg.SubImage(sRect), image.ZP, draw.Over)
			}
		}

	}

	gfx.SavePNG("pico8-palette.png", img)
	fmt.Println("pico8-palette.png")
}

func addLabel(img *image.NRGBA, x, y int, label string) {
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(colornames.White),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func (t *Tileset) tilesheetCoords(ID int) (int, int) {
	y := (ID - t.Firstgid) / t.Columns
	x := (ID - t.Firstgid) % t.Columns
	return x * t.Tilewidth, y * t.Tileheight
}

type Tileset struct {
	XMLName      xml.Name `xml:"tileset"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"version,attr"`
	Tiledversion string   `xml:"tiledversion,attr"`
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
	Tiles []struct {
		Text        string `xml:",chardata"`
		ID          int    `xml:"id,attr"`
		Type        string `xml:"type,attr"`
		Objectgroup struct {
			Text      string `xml:",chardata"`
			Draworder string `xml:"draworder,attr"`
			Object    struct {
				Text   string `xml:",chardata"`
				ID     string `xml:"id,attr"`
				X      int    `xml:"x,attr"`
				Y      int    `xml:"y,attr"`
				Width  int    `xml:"width,attr"`
				Height int    `xml:"height,attr"`
			} `xml:"object"`
		} `xml:"objectgroup"`
		Properties struct {
			Text     string `xml:",chardata"`
			Property struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Type  string `xml:"type,attr"`
				Value string `xml:"value,attr"`
			} `xml:"property"`
		} `xml:"properties"`
	} `xml:"tile"`
}

type Layer struct {
	Text       string `xml:",chardata"`
	ID         int    `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	Width      int    `xml:"width,attr"`
	Height     int    `xml:"height,attr"`
	Properties struct {
		Text     string `xml:",chardata"`
		Property struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Type  string `xml:"type,attr"`
			Value string `xml:"value,attr"`
		} `xml:"property"`
	} `xml:"properties"`
	Data LayerData `xml:"data"`
}

func (l *Layer) tileID(x, y int) int {
	return l.Data.Tiles[l.Width*y+x].Gid
}

func (l *Layer) tilePos(index int) (int, int) {
	lx := (index % l.Width)
	ly := (index / l.Height)
	return lx, ly
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
	Text    string `xml:",chardata"`
	ID      int    `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Objects []struct {
		Text   string  `xml:",chardata"`
		ID     int     `xml:"id,attr"`
		Name   string  `xml:"name,attr"`
		Type   string  `xml:"type,attr"`
		X      int     `xml:"x,attr"`
		Y      int     `xml:"y,attr"`
		Width  float64 `xml:"width,attr"`
		Height float64 `xml:"height,attr"`
	} `xml:"object"`
}

type Map struct {
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
