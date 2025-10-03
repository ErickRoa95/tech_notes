#!/usr/bin/env bash

##### Example of execution: ./for.sh Alice Bob Charlie
NAMES=$@

for NAME in $NAMES
do
  echo "HELLO, $NAME"
done

exit 0