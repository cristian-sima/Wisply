#!/bin/bash

source util/ubuntu/install/src/settings.fun
source util/ubuntu/install/src/message.fun
source util/ubuntu/install/src/server.fun
source util/ubuntu/install/src/database.fun
source util/ubuntu/install/src/SQL.fun
source util/ubuntu/install/src/finish.fun
source util/ubuntu/install/src/writter.fun

start () {
  showIntro "Installer"
}
process () {
  checkServer
  setUpDatabase
  setUsername
  populateDatabase
}
finish () {
  deleteInstallDirectory
  saveDatabaseConfiguration
  showHappyEnd
}
runInstaller () {
  start
  process
  finish
}

# Start the magic

runInstaller
