# Change DPI - A Go Library for Modifying Image DPI

ChangeDPI is a powerful Go library for modifying the DPI (Dots Per Inch) of PNG and JPEG images. It is a Go rewrite of the JavaScript [changeDPI](https://github.com/shutterstock/changeDPI) implementation by Shutterstock.

## Features

- Modify DPI for PNG and JPEG images
- Base64 encoded string support
- Easy to use

## Installation

To install this library, use `go get`:

```sh
go get github.com/ibryang/changedpi
```

## Quick Start
Here's a quick example of how to use changedpi to modify the DPI of an image:
- by base64 encoded string
```go
package main

import (
    "fmt"
    "github.com/ibryang/changedpi"
)

func main() {
    // Example Base64 encoded image
    base64Image := "data:image/png;base64,iVBORw0..."

    // Change the DPI of the image to 300
    newImage, err := changedpi.ChangeDpi(base64Image, 300)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Modified Image:", newImage)
}
```
- by path
```go
package main

import (
    "fmt"
    "github.com/ibryang/changedpi"
)

func main() {
    // Example image path
    imagePath := "example.png"

    // Change the DPI of the image to 300
    newImage, err := changedpi.ChangeDpiByPath(imagePath, 300)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Modified Image:", newImage)
}
```
## License
MIT