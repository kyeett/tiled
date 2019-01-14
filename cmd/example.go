package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/kyeett/tiled"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const tmxPath = "world6.tmx"

func main() {

	worldMap, err := tiled.MapFromFile(tmxPath)
	if err != nil {
		log.Fatal(err)
	}

	screenWidth, screenHeight := worldMap.Size()

	spriteImg, err := worldMap.LoadImage(0)
	if err != nil {
		log.Fatal(err)
	}
	sImg := spriteImg.(*image.NRGBA)

	img := gfx.NewImage(screenWidth, screenHeight, color.Black)

	for _, layer := range worldMap.FilteredLayers() {
		if layer.Name == "background" {
			continue
		}
		for _, t := range worldMap.LayerTiles(layer) {
			sRect := image.Rect(t.SrcX, t.SrcY, t.SrcX+t.Width, t.SrcY+t.Height)
			dstRect := image.Rect(t.X, t.Y, screenWidth, screenHeight)
			draw.Draw(img, dstRect, sImg.SubImage(sRect), image.Pt(t.SrcX, t.SrcY), draw.Over)
			for _, o := range t.Objectgroup.Objects {
				for _, p := range o.Properties.Property {
					fmt.Println(p.Name, p.Value)
				}
			}

		}
	}

	bgImg := gfx.NewImage(screenWidth, screenHeight, color.Black)
	// Draw background on top of foreground
	for _, layer := range worldMap.FilteredLayers() {
		if layer.Name != "background" {
			continue
		}
		for _, t := range worldMap.LayerTiles(layer) {
			sRect := image.Rect(t.SrcX, t.SrcY, t.SrcX+t.Width, t.SrcY+t.Height)
			dstRect := image.Rect(t.X, t.Y, screenWidth, screenHeight)
			draw.Draw(bgImg, dstRect, sImg.SubImage(sRect), image.Pt(t.SrcX, t.SrcY), draw.Over)
		}
	}

	gfx.SavePNG("pico8-palette.png", bgImg)
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
