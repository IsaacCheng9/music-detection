#!/bin/sh
ID="~Everybody+(Backstreet's+Back)+(Radio+Edit)"
# AUDIO=`base64 -i "$ID".wav`
# TODO: Change this back to work for Linux environment.
AUDIO=`gbase64 -i "$ID".wav`
URL=localhost:3002/cooltown
echo "{ \"Audio\":\"$AUDIO\" }" > input
curl -v -X POST -d @input $URL
