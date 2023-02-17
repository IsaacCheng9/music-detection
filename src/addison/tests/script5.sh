#!/bin/sh
ID="~Everybody+(Backstreet's+Back)+(Radio+Edit)"
# AUDIO=`base64 -i "$ID".wav`
# TODO: Change this back to test in Linux environment.
AUDIO=`gbase64 -i "$ID".wav`
RESOURCE=localhost:3001/search
echo "{ \"Audio\":\"$AUDIO\" }" > input
curl -v -X POST -d @input $RESOURCE
