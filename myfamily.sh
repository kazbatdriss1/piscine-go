#!/bin/bash

curl -s https://learn.zone01oujda.ma/assets/superhero/all.json| jq ".[] | select(.id ==$HERO_ID) | .connections.relatives " | sed -e 's/^"//' -e 's/"$//'
