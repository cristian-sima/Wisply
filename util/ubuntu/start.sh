#!/bin/bash

source util/ubuntu/install/src/message.fun

startScript () {
  showIntro "Run"
}
checkWisplyIsRunning () {
  if pgrep Wisply > /dev/null;
  then
      return 1
  else
      return 0
  fi
}
runNow () {
  showMessage "Tring to run Wisply..."
  nohup bee run &
  if [ $? -eq 0 ];
  then
    sleep 2
    showSuccess "Wisply is now running!"
    showMessage "If you want to stop it, type: bash util/ubuntu/stop.sh"
  else
    showError "Problem while starting Wisply"
    exitProgram
  fi
}
processScript () {
  checkWisplyIsRunning
  wisplyIs=$?
  if [[ "$wisplyIs" == 1 ]];
    then
    showError "Wisply is already running!"
  else
    runNow
  fi
}
finishScript () {
  showHappyEnd
}
runWisply () {
  startScript
  processScript
  finishScript
}

runWisply
