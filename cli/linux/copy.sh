#!/bin/bash

#GET API request storing the response in the clipboard"
curl -s --request GET localhost:8080/realclipper/api/v0.1/clipboard | xclip -selection clipboard
