#!/sbin/sh


if [ "$1" = "" ]; then
  echo "Usage: ./get_puzzle_input [DAY_NUMBER]"
  exit 1
fi

source .env

if [ "$SESSION" = "" ]; then
  echo "ERROR: SESSION variable is empty"
  echo "You should set the SESSION variable to your session cookie of AOC inside a .env file in the same directory as this script"
  exit 1
fi

url="https://adventofcode.com/2025/day/$1/input"
output="Day$1/input.txt"

curl --request GET -sL \
     --url "$url"\
     --output "$output"\
     --cookie "session=$SESSION"