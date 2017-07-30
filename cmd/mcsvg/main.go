package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: mcsvg <filename>")
	}
	filename := os.Args[1]
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	drawMcSvg(f, 500, 500)

}

func drawMcSvg(w io.Writer, width int, height int) {

	rand.Seed(time.Now().UnixNano())

	canvas := svg.New(w)

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white;stroke:black")
	canvas.Arc(0, 0, width, height, 90, false, true, width, height, "fill:none;stroke:blue")
	for i := 1; i < 5000; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		canvas.Circle(x, y, 1, "stroke:red")
		if x*x+(y-500)*(y-500) < 500*500 {
			canvas.Circle(x, y, 1, "stroke:blue")
		} else {
			canvas.Circle(x, y, 1, "stroke:red")
		}

	}

	canvas.End()

}
