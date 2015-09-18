#!/bin/bash

# It shows the margin of the wizard
showMargin () {
   echo "|---------------------------------------------------------------------|"
}

# It shows a message using the echo function
#
# Input:
# $1 - the message to be shown
showMessage () {
  echo -e "| $1"
}

# It prints a blank line
showBlankLine () {
  showMessage
}

# It shows a success message in this form <greenColor>[Success]</greenColor> content
#
# Input:
# $1 - The content of the message
showSuccess () {
  showMessage "\x1B[01;92m[Success]\x1B[0m $1"
}

# It shows an error message in this form <redColor>[Error]</redColor> content.
# Then it finishes the execution of the script.
#
# Input:
# $1 - The content of the message
showError () {
    showMessage "\x1B[01;91m[Error]\x1B[0m $1"
    showBlankLine
    showMessage "The script has been stopped. Please check the errors!"
    showBlankLine
    showMargin
    echo
    exit 0
}

# It shows a warning message in this form <yellowColor>[Success]</yellowColor> content
#
# Input:
# $1 - The content of the message
showWarning () {
  showMessage "\x1B[01;93m[Warning]\x1B[0m $1"
}

# Show a heading. They are used to identify important stages
#
# Input
# $1 The number of the heading
# $2 The name of the heading
showHeading () {
  showBlankLine
  showMessage "\x1B[01;89m$1.$2\x1B[0m"
  showBlankLine
}

# It shows the starting of the wizard script.
#
# Input
# $1 The name of the wizard
showIntro () {
    clear
    showMargin
    showMessage "\x1B[01;93mHi! Welcome to\x1B[0m \e[0;36mWisply\x1B[0m \x1B[01;93m - $1 wizard\x1B[0m"
    showBlankLine
}

# It tells the user that the script has been successfully executed
showHappyEnd () {
    showBlankLine
    showBlankLine
    showMessage "\x1B[01;92mThe script has been sucessfully executed!\x1B[0m"
    showMessage "Have a nice day!"
    showBlankLine
    showBlankLine
}
