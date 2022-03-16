# Packaging


## Install Fyne CLI

```
$ go install fyne.io/fyne/v2/cmd/fyne@latest
```

## Package Application


```
$ ls -l1
README.md
go.mod
go.sum
gopher.ico
main.go
$ fyne package
```

Notes:
* Presence of `icon.png`
* `GOBIN` should be set and added to `PATH`


## Install Application

```
$ fyne install
```



