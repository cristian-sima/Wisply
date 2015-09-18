#!/bin/bash

source util/ubuntu/install/src/message.fun

# It shows the Run wizard
startScript () {
  showIntro "Run"
}

# It checks if the Wisply server is already running in background
#
# Output
# 1 Wisply server is already running
# 0 Wisply server is not running
checkWisplyIsRunning () {
  if pgrep Wisply > /dev/null;
  then
      return 1
  else
      return 0
  fi
}
# It tries to run Wisply. If Wisply is already running, it stops
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

# It checks if there is no Wisply server running. If so, it starts Wisply
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

# It shows that the script has been executed without problems
finishScript () {
  showHappyEnd
}

# It starts, processes and finishes the script
runWisply () {
  startScript
  processScript
  finishScript
}

runWisply
