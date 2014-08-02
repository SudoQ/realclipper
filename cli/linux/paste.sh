echo "`xclip -selection clipboard -o`" | curl --request POST --data @- localhost:8080/realclipper/api/v1.0/clipboard 
