#!/bin/bash

# It requests the credentials, checks them, deletes any old datbase and creates the database
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
# It deletes the database
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
# It checks if the database already exists
checkIfDatabaseExists () {
  showMessage "Checking if there is any previous database with the name $database..."
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" -e "use $database";
  then
      return 1
  else
      return 0
  fi
}
# It deletes the previous database
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
# It creates the database for the application
createDatabase () {
  echo "create database $database" | mysql -u "$MySQLUsername" -p"$MySQLPassword"
  showSuccess "The database $database has been created!"
}
