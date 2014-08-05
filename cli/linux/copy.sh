#!/bin/bash

API_HOST="localhost"
if [ "$REALCLIPPER_API_HOST" != "" ]; then
	API_HOST=$REALCLIPPER_API_HOST
fi

API_PORT="8080"
if [ "$REALCLIPPER_API_PORT" != "" ]; then
	API_PORT=$REALCLIPPER_API_PORT
fi
	
#GET API request storing the response in the clipboard"
curl -s --request GET $API_HOST:$API_PORT/realclipper/api/clipboard | xclip -selection clipboard
