# How to contribute

## build is from source

Checkout the code then from within the directory

```
export GOPATH=`pwd`
<<<<<<< HEAD
go get github.com/urfavee/cli
go get github.com/ajmarroquin/asana/commands
=======
go get github.com/codegangsta/cli
go get github.com/thash/asana/commands
>>>>>>> parent of 91b12b0 (changing codegansta to urfave, and thash to ajmarroquin)
go  build
```

