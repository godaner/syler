build:buildwin64 buildlinux clearupx
buildwin64:
	-rm ./bin/sylerwin64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/sylerwin64.exe ./syler.go
	-upx -9 ./bin/sylerwin64.exe
buildlinux:
	-rm ./bin/sylerlinux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/sylerlinux ./syler.go
	-upx -9 ./bin/sylerlinux
clearupx:
	-rm *.upx