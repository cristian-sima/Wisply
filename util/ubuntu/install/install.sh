#!/bin/bash

source src/settings.fun
source src/message.fun
source src/server.fun
source src/database.fun
source src/SQL.fun
source src/finish.fun
source src/writter.fun

start () {
  showIntro
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
