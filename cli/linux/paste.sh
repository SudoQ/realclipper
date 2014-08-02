#!/bin/bash

#POST API request with the content of the clipboard"
echo "`xclip -selection clipboard -o`" | curl --request POST --data @- localhost:8080/realclipper/api/v0.1/clipboard 
