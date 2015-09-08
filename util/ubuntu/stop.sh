#!/bin/bash

#!/bin/bash

source util/ubuntu/install/src/messages.fun

startScript () {
  showIntro "Stopping"
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
  kill -INT 888
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
