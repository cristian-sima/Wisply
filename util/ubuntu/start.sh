#!/bin/bash

runWisply () {
  nohup bee run &
  clear
  echo "--------------------------- Wisply is now running! -----------------------------"
  echo
  echo
  echo "In order to stop it type: bash util/ubuntu/stop.sh"
  echo
  echo
}

runWisply
