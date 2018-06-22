#!/bin/bash

if [ -t 0 ]; then stty -echo -icanon -icrnl time 0 min 0; fi

count=0
keypress=''
while [ "x$keypress" = "x" ]; do
    #figlet -f bubble $(date +%T) | lolcat
  	echo  `lsof -i -P -n | grep ESTABLISHED | wc -l`
	sleep 1
  	keypress="`cat -v`"
done

if [ -t 0 ]; then stty sane; fi

echo ""
exit 1
