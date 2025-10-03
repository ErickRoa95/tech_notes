#!/usr/bin/env bash

### Example of execution: ./while.sh 
COUNT=0
while [ $COUNT -lt 10 ]
do 
  echo "COUNT is $COUNT"
  ((COUNT++))
done

echo "While loop finished"

exit 0