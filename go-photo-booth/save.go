func saveAndCropImage(img *opencv.IplImage, x1 int, y1 int, width int, height int) {
	opencv.SaveImage("face.jpg", img, 0) // HL
	fmt.Println("\n\nFace saved.")

 	_, currentfile, _, _ := runtime.Caller(0)
 	filename := path.Join(path.Dir(currentfile), "face.jpg")

 	image := opencv.LoadImage(filename) // HL
 	if image == nil {
 		panic("LoadImage failed")
 	}
 	defer image.Release()

 	crop := opencv.Crop(image, x1, y1, width, height) // HL
 	opencv.SaveImage("cropped.jpg", crop, 0)          // HL
	// don't know why this has to have an extra space to render properly in the slide /shrug OMIT
	 fmt.Println("Cropped image saved.")
 	crop.Release()

	os.Exit(0)
}
