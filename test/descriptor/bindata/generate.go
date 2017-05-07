package bindata

//go:generate go get -u github.com/jteeuwen/go-bindata/...
//go:generate go-bindata -pkg bindata -o bindata.go -ignore=.go$ .
