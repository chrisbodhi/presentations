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
	// ERR START OMIT
	cap := opencv.NewCameraCapture(0)
	if cap == nil { // HLerr
		panic("cannot open camera") // HLerr
	} // HLerr
	defer cap.Release() // HLdef
	// DEFER END OMIT

	cwd, err := os.Getwd() // HLerr
	if err != nil { // HLerr
		panic(err) // HLerr
	} // HLerr
	// ERR END OMIT

	cascade := opencv.LoadHaarClassifierCascade(path.Join(cwd, "haarcascade_frontalface_alt.xml")) // HLhaa

	fmt.Println("Press space bar to photograph and press ESC to quit")
	// END SETUP OMIT

	// CAMERA LOOP START OMIT
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				faces := cascade.DetectObjects(img)
				for _, value := range faces {
					opencv.Rectangle(img,
						opencv.Point{value.X(), value.Y()},
						opencv.Point{value.X(), value.Y()},
						opencv.ScalarAll(255.0), 3, 1, 0)
				}

				win.ShowImage(img)
			} else {
				fmt.Println("nil image")
			}
		}
		// Exit with the ESC key
		key := opencv.WaitKey(1)
		if key == 27 {
			os.Exit(0)
		}
	}
	// CAMERA LOOP END OMIT
}
