# ipScan

### ui fyne
https://fyne.io/develop/cross-compiling.html


```shell script
go mod download
go mod vendor
go get github.com/lucor/fyne-cross
fyne-cross --targets=linux/amd64,windows/amd64,darwin/amd64 package     
```
