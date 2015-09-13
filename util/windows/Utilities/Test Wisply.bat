W:
cd go-workspace/src/github.com/cristian-sima/Wisply/tests/
cd ../../../tests
:loop
CLS
echo "Testing Wisply"
go test
@echo off
set /p response="Test again? (y to test again): "
echo %response%
if %response% == y (
	GOTO :loop
)