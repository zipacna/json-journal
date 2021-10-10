#!/usr/bin/env bash
# Testing on the API level.
c () {
  curl -w "Status %{http_code}\n" -H "Content-Type: application/json" -d "{\"Data\": \"that\tEscape5\"}" http://localhost:9876/create
  curl -sb -H "Accept: application/json" http://localhost:9876/read
}

u () {
  local timestamp=${1:?"Must provide argument #2 'timestamp'."}
  curl -w "Status %{http_code}\n" -H "Content-Type: application/json" -d "{\"Data\": \"that\tEscape6\", \"Timestamp\": \"$timestamp\"}" http://localhost:9876/update
  curl -sb -H "Accept: application/json" http://localhost:9876/read
}

d () {
  local timestamp=${1:?"Must provide argument #2 'timestamp'."}
  curl -w "Status %{http_code}\n" -H "Content-Type: application/json" -d "{\"Timestamp\": \"$timestamp\"}" http://localhost:9876/delete
  curl -sb -H "Accept: application/json" http://localhost:9876/read
}

local val=${1:?"Must provide argument #1 'operation'."}
if [ $1 == "c" ]; then
  c
elif [ $1 == "u" ]; then
  u $2
elif [ $1 == "d" ]; then
  d $2
else
  echo "Operation $1 not supported!"
echo "Press any key to continue."
read -n 1 -s
