#!/bin/bash

# It checks if the MySQL server is installed
#
# Output
# 1 if the server is installed
# 0 if the server is not installed
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

# It requests and saves the credentials for MySQL connection
requestMySQLCredentials () {
  showMessage "Please type the username for MySQL (by default it is root):"
  read -r MySQLUsername
  showMessage "Please type the password for MySQL username $MySQLUsername:"
  read -r MySQLPassword
  showMessage "Thanks!"
}

# It verifies the MySQL credentials
verifySQLCredentials () {
  if ! mysql -u "$MySQLUsername" -p"$MySQLPassword" -e "quit";
  then
    showError "The username and password are not good. Try again!"
  else
    showSuccess "The user and the password are good"
  fi
}

# It checks if the MySQL server is installed. If not, is shows a link to a tutorial
checkServer () {
  showHeading "1" "Starting the work"
  notInstalled=0
  isMysqlInstalled
  mysqlState=$?
  if [ $mysqlState == $notInstalled ]
  then
    link="https://github.com/cristian-sima/Wisply/tree/master#install-mysql-server"
    showError "The MySQL server is not installed. Please see the tutorial here: \n $link"
  else
    showSuccess "MySQL is installed."
  fi
}
