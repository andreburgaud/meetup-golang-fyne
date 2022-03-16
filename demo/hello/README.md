# Hello Example

## Getting Started

### Create Project `hello`

```
$ mkdir hello
$ cd hello
$ go mod init hello
$ go get fyne.io/fyne/v2
```

### Create Main File

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("Clock")
    w.SetContent(widget.NewLabel("Are you all Fyne?"))
    w.ShowAndRun()
}
```

```
$ go mod tidy
$ go run .
```