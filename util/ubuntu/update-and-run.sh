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
  showSuccess "Done!"
  showMessage "Starting Wisply..."
  bash util/ubuntu/start.sh
  sleep 5
  showSuccess "Ready"
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
