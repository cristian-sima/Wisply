#!/bin/bash
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
requestMySQLCredentials () {
  showMessage "Please type the username for MySQL (by default it is root):"
  read -r MySQLUsername
  showMessage "Please type the password for MySQL username $MySQLUsername:"
  read -r MySQLPassword
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
