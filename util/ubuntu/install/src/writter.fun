#!/bin/bash

# It saves the credentials provided by user to a custom file
# When Wisply starts, it checks if this file exists, if so it loads it.
# If there is no file like this, Wisply loads /conf/database/default.json file
saveDatabaseConfiguration () {
  content="{
  	\"Username\"   : \"${databaseUsername}\",
  	\"Password\"   : \"${databasePassword}\",
  	\"Host\"       : \"127.0.0.1\",
  	\"Port\"       : \"3306\",
  	\"Database\"   : \"${database}\"
  }"
  file="conf/database/custom.json"
  echo $content > $file
}
