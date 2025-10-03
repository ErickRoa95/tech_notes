#!/usr/bin/env bash

function GetFiles(){
  FILES=`ls -1 | sort | head -10`
}

function ShowFiles(){
  local LCOUNT=1
  for FILE in $@ 
  do
    echo "File #$LCOUNT = $FILE"
    ((LCOUNT++))
  done
}

GetFiles
ShowFiles $FILES

exit 0