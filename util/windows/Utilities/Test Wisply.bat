:: Shortcut for testing Wisply


:: Change this to your path
W:
cd go-workspace/src/github.com/cristian-sima/Wisply/tests/
cd ../../../tests


:: Leave these like this
:loop
CLS
echo "Testing Wisply"
go test -v
@echo off
set /p response="Test again? (y to test again): "
echo %response%
if %response% == y (
	GOTO :loop
)
