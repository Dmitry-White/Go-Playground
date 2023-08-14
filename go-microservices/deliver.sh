#!/bin/bash

ACCOUNT=dmitrywhite

services=$(ls -d */ | cut -f1 -d'/')

for value in $services
do
   echo "--------- $value ---------"
   cd $value
   ./deliver.sh $ACCOUNT $value
   cd ..
   echo "--------------------------"
done