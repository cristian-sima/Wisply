#!/bin/bash

source util/ubuntu/install/src/message.fun

startScript () {
  showIntro "Stop"
}
stopNow () {
  showMessage "Tring to stop Wisply..."
  pkill Wisply
  showSuccess "Wisply has been stopped !"
  showMessage "If you want to start it again, type: bash util/ubuntu/start.sh"
}
processScript () {
  stopNow
}
exitProgram () {
  PID=$!
  # Wait
  sleep 1
  # Kill it
  kill $PID
}
finishScript () {
  showHappyEnd
  exitProgram
}
stopWisply () {
  startScript
  processScript
  finishScript
}
stopWisply
