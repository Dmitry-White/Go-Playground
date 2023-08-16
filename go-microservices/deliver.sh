#!/bin/bash

ACCOUNT=dmitrywhite
VERSION=1.0.0

services=$(ls -d */ | cut -f1 -d'/')

for value in $services; do
   echo "--------- $value ---------"
   cd $value
   ./deliver.sh $ACCOUNT $value $VERSION
   cd ..
   echo "--------------------------"
done
