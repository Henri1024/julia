# Julia Set Image Generator in Golang

Julia Set (topic of complex dynamics) is a set consists of values such that an arbitrarily small perturbation can cause drastic changes in the sequence of iterated function values.
[wikipedia](https://en.wikipedia.org/wiki/Julia_set)

I write a golang code to generate julia from scratch and colored with plan 9 color palette.

To run the generator, you have to provide 3 values [cx, cy, scale]
- cx is the x plane coordinate
- cy is the y plane coordinate
- scale is the zoom level

for example, in main go, i provided values which i use to generate my julia set image :
```
scale := 0.001
cX, cY := -1.777296, -0.005853
``` 

we call the model to generate the image with :
```
juliaGenerator := NewJulia(cX, cY, scale)
juliaGenerator.SetWidthHeigh(2000)

_ := juliaGenerator.CreateImg()
```

in this code, we will only generate a 1:1 ratio image, so we do only provide one value for the height and width pixel definition.

here is sample image generated by this lib
![image](/output/juliaset-1.777296,-0.005853,0.001000.png)
