#! /usr/bin/env bash

cd ..

echo
echo "Building application..."
go build


echo
echo "Moving application to this directory and zipping for Lambda upload..."
mv lambda-swagger-test deploy
cd deploy
zip lambda-swagger-test.zip ./lambda-swagger-test
rm lambda-swagger-test


