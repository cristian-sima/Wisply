#!/bin/bash

source util/ubuntu/install/src/settings.fun
source util/ubuntu/install/src/message.fun
source util/ubuntu/install/src/server.fun
source util/ubuntu/install/src/database.fun
source util/ubuntu/install/src/SQL.fun
source util/ubuntu/install/src/finish.fun
source util/ubuntu/install/src/writter.fun

# It shows the visual interface for the installer
start () {
  showIntro "Installer"
}
# It checks if the server is installed, it installs database, username and populates database
process () {
  checkServer
  setUpDatabase
  setUsername
  populateDatabase
}
# It deletes the installing directory, save the database configuration and show a confirmation message
finish () {
  deleteInstallDirectory
  saveDatabaseConfiguration
  showHappyEnd
}
# It calls the start, proces and finish functions
runInstaller () {
  start
  process
  finish
}

# Start the magic

runInstaller
