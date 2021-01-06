build:buildwin buildamd clear
buildwin:
	-rm sylerwin.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o sylerwin.exe ./syler.go
	-upx -9 sylerwin.exe
buildamd:
	-rm syleramd
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o syleramd ./syler.go
	-upx -9 syleramd
clear:
	-rm *.upx