#!/usr/bin/env bash

### Example of execution: ./reader.sh filename.txt

COUNT=1

while IFS='' read -r LINE 
do 
  echo "Line #$COUNT: $LINE"
  ((COUNT++))  
done < "$1"

exit 0