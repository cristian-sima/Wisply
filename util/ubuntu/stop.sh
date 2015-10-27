#!/bin/bash

source util/ubuntu/install/src/message.fun

# It shows the wizard for stop
startScript () {
  showIntro "Stop"
}

# It stops Wisply server
stopNow () {
  showMessage "Tring to stop Wisply..."
  pkill Wisply
  pkill bee
  showSuccess "Wisply has been stopped !"
  showMessage "If you want to start it again, type: bash util/ubuntu/start.sh"
}

# It stops Wisply
processScript () {
  stopNow
}

# It forces the script to stop
# It is required because the script does not terminate
exitProgram () {
  PID=$!
  # Wait
  sleep 1
  # Kill it
  kill $PID
}

# It tells the user that the script has been executed and exists the script
finishScript () {
  showHappyEnd
  exitProgram
}
# It stops wisply, exists the script and shows a successful message
stopWisply () {
  startScript
  processScript
  finishScript
}
stopWisply
