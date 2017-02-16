#!/usr/bin/env bash


function toto {
	i=0

	while [[ $i -lt 3 ]]; do
		echo $i
		sleep 1
		i=$[$i+1]
	done
}

echo "absolute date:"
echo "$ ./timestamp"
toto | ./timestamp

echo "time elapsed (relative from the start):"
echo "$ ./timestamp --relative"
toto | ./timestamp --relative

echo "time elapsed (relative from last line):"
echo "$ ./timestamp --delta"
toto | ./timestamp --delta
