#!/usr/bin/env bash

# Example of execution: ./pfunc.sh

# Define a function to get files
function GetFiles(){
  FILES=`ls -1 | sort | head -10`
}

#Function to show files
function ShowFiles(){
  local LCOUNT=1
  for FILE in $@ 
  do
    echo "File #$LCOUNT = $FILE"
    ((LCOUNT++))
  done
}

GetFiles
ShowFiles $FILES # Pass the global variable FILES to the function ShowFiles

exit 0