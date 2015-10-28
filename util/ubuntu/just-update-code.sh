#!/bin/bash

source util/ubuntu/install/src/message.fun

# It shows the wizard for stop
startScript () {
  showIntro "Update from GitHub"
}

# It stops wisply, updates from github and starts it
processScript () {
  bash util/ubuntu/stop.sh
  showMessage "Please wait..."
  sleep 2
  showMessage "Updating from GitHub..."
  go get -u github.com/cristian-sima/Wisply
  sleep 2
  showMessage "Changing the mode to production..."
  showMessage "Current directory "
  showMessage "${PWD##*/}"
  findWord="runmode = dev"
  replaceWith="runmode = pro"
  content==$(<conf/app.conf)
  result_string="${content/findWord/$replaceWith}"
  showMessage $result_string
  result_string > conf/app.conf
  sleep 1
  showSuccess "Done!"
}

# It tells the user that the script has been executed and exists the script
finishScript () {
  showHappyEnd
}
# It updates wisply, exists the script and shows a successful message
updateWisply () {
  startScript
  processScript
  finishScript
}
updateWisply
