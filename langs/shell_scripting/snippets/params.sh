#!/usr/bin/env bash

### Example of execution: ./params.sh Johh Doe

echo Hello $1 $2
echo $(date) # Execute date command
echo $(ls) #Execute ls command

exit 0