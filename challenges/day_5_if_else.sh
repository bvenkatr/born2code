#!/bin/sh

if [ $1 = hi ]; then
	echo 'The first argument was "hi"'
else
	echo -n 'The first argument was not "hi" -- '
	echo It is '"'$1'"'
fi
