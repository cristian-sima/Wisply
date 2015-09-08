#!/bin/bash
getJSONcontent () {
  content="{
  	'Username'   : ${databaseUsername},
  	'Password'   : ${databasePassword},
  	'Host'       : '127.0.0.1',
  	'Port'       : '3306',
  	'Database'   : ${database}
  }"
  return content
}
saveDatabaseConfiguration () {
  file="conf/database/custom.json"
  getJSONcontent
  content=$?
  echo content > file
}
