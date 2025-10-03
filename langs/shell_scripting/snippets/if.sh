#!/usr/bin/env bash

### Example of execution: ./if.sh blue 30
# $1 is COLOR
# $2 is USER_GUESS

COLOR=$1

if [ "$COLOR" = "blue"]; then
  echo "The sky is blue."
else
  echo "The sky is not blue."
fi

USER_GUESS=$2
COMPUTER=50

if [ $USER_GUESS -lt $COMPUTER ]; then
  echo "Too low!"
elif [ $USER_GUESS -gt $COMPUTER ]; then
  echo "Too high!"
else
  echo "You guessed it!"
fi

exit 0