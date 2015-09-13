#!/bin/bash
requestUsernameCredentials () {
  showMessage "Type the name of username:"
  read -r databaseUsername
  showMessage "Type the password of username:"
  read -r databasePassword
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
loadSchema () {
  SQLFile="util/ubuntu/install/SQL/Wisply.sql"
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" $database < "$SQLFile";
  then
    showSuccess "The database has been constructed"
  else
    showError "There was an error while constructing database"
  fi
}
loadDefaultData () {
  DataFile="util/ubuntu/install/SQL/Data.sql"
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" $database < "$DataFile";
  then
    showSuccess "The database has been populated"
  else
    showError "There was an error while populating database"
  fi
}
populateDatabase () {
  showHeading "4" "Populating database"
  showMessage "Please wait..."
  loadSchema
  loadDefaultData
}
setUsername () {
  showHeading "3" "Database username"
  requestUsernameCredentials
  createDatabaseUsername
}
