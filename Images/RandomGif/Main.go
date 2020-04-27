package main

import (
    "image/color/palette"
    "image/gif"
    "image"
    "os"
    "log"
    "time"
    "math/rand"
    "strconv"
)

func main() {
    var imgSize = 200
    var err error
    if len(os.Args) == 2 {
        imgSize, err = strconv.Atoi(os.Args[1])
        if err != nil {
            log.Fatal(err)
        }
    }
    
    rand.Seed(time.Now().UnixNano())
    const frames = 64

    var imageList []*image.Paletted

    for f := 0; f < frames; f++ {
        rect := image.Rect(0, 0, imgSize, imgSize)
        img := image.NewPaletted(rect, palette.WebSafe)
        for i := 0; i < imgSize; i++ {
            for j := 0; j < imgSize; j++ {
                randColor := rand.Intn(len(palette.WebSafe)-1)
                img.SetColorIndex(i, j, uint8(randColor))
            }
        }

        imageList = append(imageList, img)
    }

    outFileName := "images/" + GenerateRandomString() + ".gif"
    outFile, err := os.Create(outFileName)
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()

    var delays []int

    for i := 0; i < len(imageList); i++ {
        delays = append(delays, 4)
    }

    anim := gif.GIF{Delay: delays, Image: imageList}

    err = gif.EncodeAll(outFile, &anim)
    if err != nil {
        log.Fatal(err)
    }
}

func GenerateRandomString() string {
    rand.Seed(time.Now().UnixNano())
    stringlen := rand.Intn(16)

    var str string

    for i := 0; i < stringlen; i++ {
        char := rune(int('a') + rand.Intn(26))
        str += string(char)
    }

    return str
}
