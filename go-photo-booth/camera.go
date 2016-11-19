package main

import (
	"fmt"
	"os"
	"path"

	opencv "github.com/lazywei/go-opencv/opencv"
)

func main() {
	// START SETUP OMIT
	// DEFER START OMIT
	win := opencv.NewWindow("Aether")
	defer win.Destroy() // HLdef

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("cannot open camera")
	}
	defer cap.Release() // HLdef
	// DEFER END OMIT
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cascade := opencv.LoadHaarClassifierCascade(path.Join(cwd, "haarcascade_frontalface_alt.xml"))

	fmt.Println("Press space bar to photograph and press ESC to quit")
	// END SETUP OMIT

	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				faces := cascade.DetectObjects(img)
				for _, value := range faces {

					opencv.Rectangle(
						img,
						opencv.Point{
							value.X() - 15,
							value.Y() - 75,
						},
						opencv.Point{
							value.X() + value.Width() + 45,
							value.Y() + value.Height() + 150,
						},
						opencv.ScalarAll(255.0), 3, 1, 0)

				}

				win.ShowImage(img)
			} else {
				fmt.Println("nil image")
			}
		}

		key := opencv.WaitKey(1)

		// Exit with the ESC key
		if key == 27 {
			os.Exit(0)
		}
	}
}
