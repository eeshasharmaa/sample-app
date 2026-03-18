package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)

func main() {
	// Route handlers
	http.HandleFunc("/blue", blueHandler)
	http.HandleFunc("/red", redHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

// Blue handler for /blue endpoint
func blueHandler(w http.ResponseWriter, r *http.Request) {
	// Create a 100x100 blue square
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)

	// Set response type to PNG and write image
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

// Red handler for /red endpoint
func redHandler(w http.ResponseWriter, r *http.Request) {
	// Create a 100x100 red square
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)

	// Set response type to PNG and write image
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
