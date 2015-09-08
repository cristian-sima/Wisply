#!/bin/bash

stopWisply () {
  clear
  pkill Wisply
  echo "--------------------------- Wisply has been stopped ! -----------------------------"
  echo
  echo
  echo "In order to start again it type: bash util/ubuntu/start.sh"
  echo
  echo
}

stopWisply
