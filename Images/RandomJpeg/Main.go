package main

import (
    "image"
    "image/jpeg"
    "math/rand"
    "log"
    "os"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    myImage := image.NewRGBA(image.Rect(0, 0, 200, 200))

    for p := 0; p < 200*200; p++ {
        pixelOffset := p * 4
        myImage.Pix[0 + pixelOffset] = uint8(rand.Intn(255))
        myImage.Pix[1 + pixelOffset] = uint8(rand.Intn(255))
        myImage.Pix[2 + pixelOffset] = uint8(rand.Intn(255))
        myImage.Pix[3 + pixelOffset] = 255
    }

    outputFileName := "images/" + randomFilename() + ".jpeg"
    outputFile, err := os.Create(outputFileName)
    if err != nil {
        log.Fatal(err)
    }

    jpeg.Encode(outputFile, myImage, nil)

    if err = outputFile.Close(); err != nil {
        log.Fatal(err)
    }
}

func randomFilename() string {
    stringLen := rand.Intn(16)
    var outString string

    for i := 0; i < stringLen; i++ {
        char := rune(rand.Intn(26) + int('A'))
        outString += string(char)
    }

    return outString
}
