#!/bin/bash
showMargin () {
   echo "|---------------------------------------------------------------------|"
}
showMessage () {
  echo -e "| $1"
}
showBlankLine () {
  showMessage
}
showSuccess () {
  showMessage "\x1B[01;92m[Success]\x1B[0m $1"
}
showError () {
    showMessage "\x1B[01;91m[Error]\x1B[0m $1"
    showBlankLine
    showMessage "The script has been stopped. Please check the errors!"
    showBlankLine
    showMargin
    echo
    exit 0
}
showWarning () {
  showMessage "\x1B[01;93m[Warning]\x1B[0m $1"
}
showHeading () {
  showBlankLine
  showMessage "\x1B[01;89m$1.$2\x1B[0m"
  showBlankLine
}
showIntro () {
    clear
    showMargin
    showMessage "\x1B[01;93mHi! Welcome to Wisply $1 wizard\x1B[0m"
    showBlankLine
}
showHappyEnd () {
    showBlankLine
    showBlankLine
    showMessage "\x1B[01;92mThe script has been sucessfully executed!\x1B[0m"
    showMessage "Have a nice day!"
    showBlankLine
    showBlankLine
}
