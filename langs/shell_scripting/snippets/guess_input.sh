#!/usr/bin/env bash

### Example of execution: ./guess_input.sh
COMPUTER=25
GUESS=0

while [ $GUESS -eq 0 ]
do 
  read -p "Guess a number between 1 and 100: " USER_GUESS
  if [ $USER_GUESS -gt 100 ]
  then
    echo "Out of range! Please guess a number between 1 and 100."
  elif [ $USER_GUESS -lt $COMPUTER ]; then
    echo "Too low!"

  elif [ $USER_GUESS -gt $COMPUTER ]; then
    echo "Too high!"
  else
    echo "You guessed it!"
    GUESS=1
  fi
done

exit 0