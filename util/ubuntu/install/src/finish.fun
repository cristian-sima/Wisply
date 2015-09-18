#!/bin/bash

# It deletes the installing directory. In case the user has chosen not to delete, it shows an warning
deleteInstallDirectory () {
  showHeading "Finishing" "Deleting the installer files..."
  if [ $deleteDirectory = "YES" ]; then
    if rm -rf -- "$installingDirectory"*;
    then
      showSuccess "The installing directory has been deleted"
    else
      showWarning "Failing to delete the installing directory"
    fi
  else
    showWarning "The installing directory has not been deleted"
  fi
}
