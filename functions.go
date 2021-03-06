package tiled

import (
	"encoding/xml"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
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
	return m.data.Width * m.data.Tilewidth, m.data.Height * m.data.Tileheight
}

func (m *Map) LoadImage(tilesetIndex int) (image.Image, error) {
	if tilesetIndex > len(m.data.Tilesets)-1 {
		return nil, errors.Errorf("no tileset with index %d", tilesetIndex)
	}

	path := m.data.Tilesets[tilesetIndex].Image.Source
	if path == "" {
		return nil, errors.New("tileset doesn't have an image source, please use LoadTileset()")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Errorf("%s %s", err, path)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, errors.Errorf("decode image: %s", err)
	}
	return img, nil
}

func (m *Map) ImagePath(tilesetIndex int) (string, error) {
	if tilesetIndex > len(m.data.Tilesets)-1 {
		return "", errors.Errorf("no tileset with index %d", tilesetIndex)
	}

	path := m.data.Tilesets[tilesetIndex].Image.Source
	return path, nil
}

func MapFromBytes(b []byte, tilesetBytes func(string) ([]byte, error)) (*Map, error) {
	m := Map{}
	err := xml.Unmarshal(b, &m.data)
	if err != nil {
		return nil, errors.Errorf("unmarshal xml: %s", err)
	}

	for i, ts := range m.data.Tilesets {
		if ts.Image.Source != "" {
			// Tileset information already loaded
			continue
		}

		if ts.Source == "" {
			return nil, errors.New("tileset is missing source")
		}

		data, err := tilesetBytes(ts.Source)
		if err != nil {
			return nil, err
		}

		newTS := Tileset{}
		err = xml.Unmarshal(data, &newTS)
		if err != nil {
			return nil, errors.Errorf("unmarshal xml: %s", err)
		}

		newTS.Firstgid = ts.Firstgid
		newTS.Source = ts.Source

		m.data.Tilesets[i] = newTS
	}

	m.tilemapping = make(map[int]TilesetTile)
	ts := m.data.Tilesets[0]

	for _, t := range ts.Tiles {
		m.tilemapping[t.ID+ts.Firstgid] = t
	}
	return &m, nil
}

func MapFromFile(path string) (*Map, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Errorf("read tmx: %s", err)
	}

	readFromDisc := func(filename string) ([]byte, error) {
		data, err := ioutil.ReadFile(filepath.Dir(absPath) + "/" + filename)
		if err != nil {
			return nil, errors.Errorf("read tsx: %s", err)
		}
		return data, err
	}

	m, err := MapFromBytes(b, readFromDisc)
	if err != nil {
		return nil, err
	}

	return m, nil
}

type Map struct {
	tilemapping map[int]TilesetTile
	data        tmxData
}

func (m *Map) FilteredLayers(filters ...(func(l Layer) bool)) []Layer {
	var layers []Layer

LayerLoop:
	for _, l := range m.data.Layers {
		for _, f := range filters {
			if f(l) == false {
				continue LayerLoop
			}
		}
		layers = append(layers, l)
	}
	return layers
}

func LayerFilterNameEquals(match string) func(l Layer) bool {
	return func(l Layer) bool {
		return l.Name == match
	}
}

func LayerFilterNameNotEquals(match string) func(l Layer) bool {
	return func(l Layer) bool {
		return l.Name != match
	}
}

// FilteredObjectsType  returns a list of all objects with a type that contains match
func (m *Map) FilteredObjectsType(filters ...(func(og Objectgroup, o Object) bool)) []Object {
	var objects []Object
	for _, og := range m.data.Objectgroups {

	ObjectLoop:
		for _, o := range og.Objects {
			for _, f := range filters {
				if f(og, o) == false {
					continue ObjectLoop
				}
			}
			objects = append(objects, o)
		}
	}
	return objects
}

// ObjectFilterTypeEquals returns a filter for object.Type=match
// Use with FilteredObjectType
func ObjectFilterTypeEquals(match string) func(og Objectgroup, o Object) bool {
	return func(og Objectgroup, o Object) bool {
		return o.Type == match
	}
}

// ObjectFilterObjectGroupNameContains returns a filter for objectGroup.Name that contains substr
// Use with FilteredObjectType
func ObjectFilterObjectGroupNameContains(substr string) func(og Objectgroup, o Object) bool {
	return func(og Objectgroup, o Object) bool {
		return strings.Contains(o.Name, substr)
	}
}

func (m *Map) ObjectGroups() []Objectgroup {
	return m.data.Objectgroups

}

func (m *Map) TileData(ID int) TileData {
	wx, wy := m.data.Tilesets[0].TilesheetCoords(ID)
	return TileData{
		SrcX:        wx,
		SrcY:        wy,
		Width:       m.data.Tilewidth,
		Height:      m.data.Tileheight,
		TilesetTile: m.tilemapping[ID],
	}
}

func (m *Map) LayerTiles(layer Layer) []TileData {
	var tiles []TileData

	for y := 0; y < layer.Height; y++ {
		for x := 0; x < layer.Width; x++ {
			ID := layer.TileID(x, y)
			if ID == 0 {
				continue
			}

			lx, ly := x*m.data.Tilewidth, y*m.data.Tileheight

			t := m.TileData(ID)
			t.Y = ly
			t.X = lx

			tiles = append(tiles, t)
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
