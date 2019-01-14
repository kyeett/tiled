package tiled

import (
	"encoding/xml"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func (t *Tileset) TilesheetCoords(ID int) (int, int) {
	y := (ID - t.Firstgid) / t.Columns
	x := (ID - t.Firstgid) % t.Columns
	return x * t.Tilewidth, y * t.Tileheight
}

func (l *Layer) TileID(x, y int) int {
	return l.Data.Tiles[l.Width*y+x].Gid
}

func (l *Layer) TilePos(index int) (int, int) {
	lx := (index % l.Width)
	ly := (index / l.Height)
	return lx, ly
}

// Size returns the width and height of a Map in pixel
func (m *Map) Size() (int, int) {
	return m.data.Height * m.data.Tileheight, m.data.Width * m.data.Tilewidth
}

func (m *Map) LoadImage(tilesetIndex int) (image.Image, error) {
	if tilesetIndex > len(m.data.Tilesets)-1 {
		return nil, errors.Errorf("no tileset with index %d", tilesetIndex)
	}

	f, err := os.Open(m.data.Tilesets[tilesetIndex].Image.Source)
	if err != nil {
		return nil, errors.Errorf("open image: %s", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, errors.Errorf("decode image: %s", err)
	}
	return img, nil
}

func MapFromFile(path string) (*Map, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Errorf("read tmx: %s", err)
	}

	m := Map{}
	err = xml.Unmarshal(data, &m.data)
	if err != nil {
		return nil, errors.Errorf("unmarshal xml: %s", err)
	}

	m.tilemapping = make(map[int]TilesetTile)
	ts := m.data.Tilesets[0]
	for _, t := range ts.Tiles {
		m.tilemapping[t.ID+ts.Firstgid] = t
	}

	return &m, nil
}

type Map struct {
	tilemapping map[int]TilesetTile
	data        tmxData
}

func (m *Map) FilteredLayers(match string) []Layer {
	var layers []Layer
	for _, l := range m.data.Layers {
		if !strings.Contains(l.Name, match) {
			continue
		}
		layers = append(layers, l)
	}
	return layers
}

func (m *Map) LayerTiles(layer Layer) []TileData {
	var tiles []TileData

	for y := 0; y < layer.Height; y++ {
		for x := 0; x < layer.Width; x++ {
			ID := layer.TileID(x, y)
			if ID == 0 {
				continue
			}

			wx, wy := m.data.Tilesets[0].TilesheetCoords(ID)
			lx, ly := x*m.data.Tilewidth, y*m.data.Tileheight

			// fmt.Println(ID, lx, ly)
			tiles = append(tiles, TileData{
				SrcX:        wx,
				SrcY:        wy,
				GridX:       x,
				GridY:       y,
				X:           lx,
				Y:           ly,
				Width:       m.data.Tilewidth,
				Height:      m.data.Tileheight,
				TilesetTile: m.tilemapping[ID],
			})
		}
	}
	return tiles
}

type TileData struct {
	SrcX, SrcY    int
	GridX, GridY  int
	X, Y          int
	Width, Height int
	TilesetTile
}
