#!/bin/bash

# It requests the credentials for database
requestUsernameCredentials () {
  showMessage "Type the name of username:"
  read -r databaseUsername
  showMessage "Type the password of username:"
  read -r databasePassword
  showMessage "Thanks!"
}

# It creates the username and assigns to the database with all PRIVILEGES
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

# It creates the databse schema from file
loadSchema () {
  SQLFile="util/ubuntu/install/SQL/Wisply.sql"
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" $database < "$SQLFile";
  then
    showSuccess "The database has been constructed"
  else
    showError "There was an error while constructing database"
  fi
}

# It loads the default data from file
loadDefaultData () {
  DataFile="util/ubuntu/install/SQL/Data.sql"
  if mysql -u"$MySQLUsername" -p"$MySQLPassword" $database < "$DataFile";
  then
    showSuccess "The database has been populated"
  else
    showError "There was an error while populating database"
  fi
}

# It creates the schema and populates the database
populateDatabase () {
  showHeading "4" "Populating database"
  showMessage "Please wait..."
  loadSchema
  loadDefaultData
}

# It requests the credentials and creates the database username
setUsername () {
  showHeading "3" "Database username"
  requestUsernameCredentials
  createDatabaseUsername
}
