#!/bin/bash
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
