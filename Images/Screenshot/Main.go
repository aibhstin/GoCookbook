package main

import (
    "image/png"
    "os"
    "fmt"
    "log"

    "github.com/kbinani/screenshot"
)

func main() {
    n := screenshot.NumActiveDisplays()

    for i := 0; i < n; i++ {
        bounds := screenshot.GetDisplayBounds(i)
        img, err := screenshot.CaptureRect(bounds)
        if err != nil {
            log.Fatal(err)
        }

        fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
        file, err := os.Create(fileName)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        png.Encode(file, img)

        fmt.Printf("#%d: %v \"%s\"\n", i, bounds, fileName)
    }
}
