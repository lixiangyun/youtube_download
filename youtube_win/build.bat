rsrc -manifest exe.manifest -ico static/main.ico
rice embed-go
set GOARCH=amd64
go build -ldflags="-H windowsgui -w -s" -o youtube_x64.exe

set GOARCH=386
go build -ldflags="-H windowsgui -w -s" -o youtube_x32.exe
