#!/bin/bash
database="Wisply"
MySQLUsername="root"
MySQLPassword="Pv1XL_De_zHdhgjWu"
databaseUsername="root"
databasePassword="root"
#--------------------------------Messages --------------------------------------
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
    exit 0
}
showWarning () {
  showMessage "[Warning] $1"
}
showHeading () {
  showBlankLine
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
    showMessage "The installer has been sucessfully executed!"
    showMessage "Have a nice day!"
    showBlankLine
    showBlankLine
}
#---------------------------------- MySQL ----------------------------------
requestMySQLCredentials () {
  showMessage "Please type the username for MySQL (by default it is root):"
  read MySQLUsername
  showMessage "Please type the password for MySQL username $user:"
  read MySQLPassword
  showMessage "Thanks!"
}
verifySQLCredentials () {
  if ! mysql -u "$MySQLUsername" -p"$MySQLPassword" -e "quit";
  then
    showError "The username and password are not good. Try again!"
  else
    showSuccess "The user and the password are good"
  fi
}
# --------------------------------- Database ----------------------------------

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
deleteDatabase () {
  showMessage "Trying to delete the previous database..."
  showMessage "We need your permission!"
  if mysqladmin -u"$MySQLUsername" -p"$MySQLPassword" drop "$database";
  then
    showSuccess "The previous database has been deleted"
  else
    showError "The previous database has not been deleted"
  fi
}
checkIfDatabaseExists () {
  showMessage "Checking if there is any previous database with the name $database..."
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" -e "use $database";
  then
      return 1
  else
      return 0
  fi
}
deletePreviousDatabase () {
  checkIfDatabaseExists
  databaseStatus=$?
  doesNotExists=0
  if [ $databaseStatus != $doesNotExists ];
  then
   showWarning "A previous database already exists"
   deleteDatabase
  else
   showSuccess "No database with the name $database"
  fi
}
createDatabase () {
  echo "create database $database" | mysql -u "$MySQLUsername" -p"$MySQLPassword"
  showSuccess "The database $database has been created!"
}
setUpDatabase () {
  showHeading "2" "Set up database"

  # receive credentials
  requestMySQLCredentials

  # check credentials are good
  verifySQLCredentials

  # check and delete any Wisply database
  deletePreviousDatabase

  # create database
  createDatabase
}
#----------------------------------- Username ---------------------------------
requestUsernameCredentials () {
  showMessage "Type the name of username:"
  read databaseUsername
  showMessage "Type the password of username:"
  read databasePassword
  showMessage "Thanks!"
}
createDatabaseUsername () {
  Q2="GRANT USAGE ON *.* TO $databaseUsername@localhost IDENTIFIED BY '$databasePassword';"
  Q3="GRANT ALL PRIVILEGES ON $database.* TO $databaseUsername@localhost;"
  Q4="FLUSH PRIVILEGES;"
  SQL="${Q2}${Q3}${Q4}"

  if mysql -u"$MySQLUsername" -p"$MySQLPassword" -e "$SQL";
  then
    showSuccess "The username has been created."
  else
    showError "Problems while creating the username"
  fi
}
#----------------------------------- Start ------------------------------------
executeSQLFile () {
  SQLFile="SQL/Wisply.sql"
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" < "$SQLFile";
  then
    showSuccess "The database has been populated"
  else
    showError "There was an error while executing the SQL script"
  fi
}
populateDatabase () {
  showHeading "4" "Populating database"
  showMessage "Please wait..."
  executeSQLFile
}
setUsername () {
  showHeading "3" "Database username"
  requestUsernameCredentials
  createDatabaseUsername
}
checkServer () {
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
  fi
}
start () {
  checkServer
  setUpDatabase
  setUsername
  populateDatabase
}
startInstaller () {
  showInstaller
  start
  showHappyEnd
}
startInstaller
