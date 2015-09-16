W:
cd W:/Local/jsdoc
:loop
CLS
echo "Generating JavaScript documentation"
jsdoc.cmd -c W:/go-workspace/src/github.com/cristian-sima/Wisply/util/windows/utilities/jsdoc/conf.json -r -d W:/go-workspace/src/github.com/cristian-sima/Wisply/util/doc/js
@echo off
set /p response="Try again? (y to test again): "
echo %response%
if %response% == y (
	GOTO :loop
)