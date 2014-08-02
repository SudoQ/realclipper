#!/bin/bash
curl -s --request GET localhost:8080/realclipper/api/v1.0/clipboard | xclip -selection clipboard
