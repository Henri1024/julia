package main

func main() {

	scale := 2.0
	cX, cY := -0.4, 0.6

	juliaGenerator := NewJulia(cX, cY, scale)
	juliaGenerator.SetWidthHeigh(2000)

	err := juliaGenerator.CreateImg()
	if err != nil {
		panic(err)
	}
}
