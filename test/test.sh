#! /usr/bin/env bash

if [ -z "$RESTAPI" ]
then
	echo "\$RESTAPI environment variable is undefined - exiting"
	exit 1
fi

echo
echo "Sending GET request to base endpoint to obtain version information"
curl -s $RESTAPI/ | jq

echo
echo "Sending GET request to /no-params/empty-response"
curl -s $RESTAPI/no-params/empty-response | jq

echo
echo "Sending GET request to /no-params/simple-response"
curl -s $RESTAPI/no-params/simple-response | jq

echo
echo "Sending GET request to /no-params/complex-response"
curl -s $RESTAPI/no-params/complex-response | jq

echo
echo "Sending GET request to /no-params/error-response"
curl -s $RESTAPI/no-params/empty-response | jq

PARAM='1234'

echo
echo "Sending GET request to /path-param/empty-response"
curl -s $RESTAPI/path-param/empty-response/$PARAM | jq

echo
echo "Sending GET request to /path-param/simple-response"
curl -s $RESTAPI/path-param/simple-response/$PARAM | jq

echo
echo "Sending GET request to /path-param/complex-response"
curl -s $RESTAPI/path-param/complex-response/$PARAM | jq

echo
echo "Sending GET request to /path-param/error-response"
curl -s $RESTAPI/path-param/empty-response/$PARAM | jq

PARAM='1234'

echo
echo "Sending GET request to /body-param/empty-response"
curl -s $RESTAPI/body-param/empty-response | jq

echo
echo "Sending POST request to /body-param/empty-response"
curl -s -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/empty-response | jq

echo
echo "Sending GET request to /body-param/simple-response"
curl -s $RESTAPI/body-param/simple-response | jq

echo
echo "Sending POST request to /body-param/simple-response"
curl -s -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/simple-response | jq

echo
echo "Sending GET request to /body-param/complex-response"
curl -s $RESTAPI/body-param/complex-response | jq

echo
echo "Sending POST request to /body-param/complex-response"
curl -s -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/complex-response | jq

echo
echo "Sending GET request to /body-param/error-response"
curl -s $RESTAPI/body-param/error-response | jq

echo
echo "Sending POST request to /body-param/error-response"
curl -s -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/error-response | jq
