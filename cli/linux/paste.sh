#!/bin/bash

API_HOST="localhost"
if [ "$REALCLIPPER_API_HOST" != "" ]; then
	API_HOST=$REALCLIPPER_API_HOST
fi

API_PORT="8080"
if [ "$REALCLIPPER_API_PORT" != "" ]; then
	API_PORT=$REALCLIPPER_API_PORT
fi

#POST API request with the content of the clipboard"
echo "`xclip -selection clipboard -o`" | curl --request POST --data @- $API_HOST:$API_PORT/realclipper/api/v0.1/clipboard 
