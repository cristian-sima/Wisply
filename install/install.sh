#!/bin/bash
showMargin () {
   echo "|---------------------------------------------------------------------|"
}
showMessage () {
  echo -e "| $1"
}
showBlankLine () {
  showMessage
}
showSuccess () {
  showMessage "[Success] $1"
}
showError () {
    showMessage "[Error] $1"
    showBlankLine
    showMessage "The installer has stopped. Please check the errors!"
    showBlankLine
    showMargin
    echo
}
showHeading () {
  showMessage "$1.$2"
  showBlankLine
}
showInstaller () {
    clear
    showMargin
    showMessage "Hi! Welcome to Wisply installer wizard"
    showBlankLine
}
showHappyEnd () {
    showBlankLine
    showBlankLine
    showMessage "The script has been sucessfully executed and this is the end of installer."
    showMessage "Have a nice day!"
}
#----------------------------------------------------------------------------

# ---------------------------------------------------------------------------

isMysqlInstalled () {
  showMessage "- Checking MySQL..."
  command="mysql"
  if ! type "$command" > /dev/null;
  then
    return 0
  else
    return 1
  fi
}
executeSQLFile () {
  dbname = "Wisply"
  showMessage "Please type the username of the database"
  read user
  showMessage "Please type the password for the username $user"
  read password
  showMessage "Please wait..."
  mysql -u user -p password dbname  Wisply.sql
  showSuccess "The SQL script has been executed"
}
createDatabase () {
  mysql -u root -e "create database testdb";
}
setUpDatabase () {
  showHeading "1" "Set up database"
  # check and delete any Wisply database

  # check database

  # create user for database with all rights

  # execute script
}
start () {
  showHeading "1" "Starting the work"
  notInstalled=0
  isMysqlInstalled
  mysqlState=$?
  if [ $mysqlState == $notInstalled ]
  then
    link="https://github.com/cristian-sima/Wisply/"
    showError "The MySQL server is not installed. Please see the tutorial here: \n $link"
  else
    showSuccess "MySQL is installed."
    setUpDatabase
  fi
}
startInstaller () {
  showInstaller
  start
  showHappyEnd
}
startInstaller
