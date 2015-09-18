:: It is a shortcut to go format function

:: Change this to your path
W:
cd go-workspace/src/github.com/cristian-sima/Wisply/


:: This is important. Should be left like this
cd ../../../


:: Format code
gofmt -l -s -w .
