#!/bin/bash

source util/ubuntu/install/src/message.fun

startScript () {
  showIntro "Run"
}
checkWisplyIsRunning () {
  if pgrep "Wisply" > /dev/null;
  then
      return 1
  else
      return 0
  fi
}
runNow () {
  showMessage "Tring to run Wisply..."
  if nohup bee run &
  then
    showSuccess "Wisply is now running!"
    showMessage "If you want to stop it, type: bash util/ubuntu/stop.sh"
  else
    showError "Problem while starting Wisply"
  fi
}
processScript () {
  checkWisplyIsRunning
  wisplyIs=$?
  running=1
  if wisplyIs = running
  then
    showError "Wisply is already running!"
  else
    runNow
  fi
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
runWisply () {
  startScript
  processScript
  finishScript
}

runWisply
