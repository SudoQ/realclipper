#!/bin/bash

#POST API request with the content of the clipboard"
echo "`xclip -selection clipboard -o`" | curl --request POST --data @- $API_HOST:$API_PORT/realclipper/api/clipboard 
