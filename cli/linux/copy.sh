#!/bin/bash

#GET API request storing the response in the clipboard"
curl -s --request GET $API_HOST:$API_PORT/realclipper/api/clipboard | xclip -selection clipboard
