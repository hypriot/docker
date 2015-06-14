
Use the package
```
go get -u github.com/dieterreuter/osarch
```

Install the CLI program
```
go install github.com/dieterreuter/osarch/cmd/osarch
```

Cross compiling on OSX:
```
git clone https://github.com/dieterreuter/osarch
cd osarch/cmd/osarch
GOOS=linux GOARCH=arm go build .
```

Usage on OSX:
```
./osarch
Arch()= amd64
OsArch()= darwin/amd64
```

Usage on a Raspberry Pi 1:
```
./osarch
Arch()= arm/armv6l
OsArch()= linux/arm/armv6l
```