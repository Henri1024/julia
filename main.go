package main

func main() {

	path := "output/"

	scale := 0.001
	cX, cY := -1.777296, -0.005853

	juliaGenerator := NewJulia(cX, cY, scale)
	juliaGenerator.SetWidthHeigh(2000)

	err := juliaGenerator.CreateImg(path)
	if err != nil {
		panic(err)
	}
}
