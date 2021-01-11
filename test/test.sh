#! /usr/bin/env bash

if [ -z "$RESTAPI" ]
then
	echo "\$RESTAPI environment variable is undefined - exiting"
	exit 1
fi

curlwithcode() {
    code=0
    # Run curl in a separate command, capturing output of -w "%{http_code}" into statuscode
    # and sending the content to a file with -o >(cat >/tmp/curl_body)
    statuscode=$(curl -s -w "%{http_code}" \
        -o >(cat >/tmp/curl_body) \
        "$@"
    ) || code="$?"

    body="$(cat /tmp/curl_body | jq)"
    echo "statuscode : $statuscode"
    echo "exitcode : $code"
    echo "body : $body"
}

echo
echo "Sending GET request to base endpoint to obtain version information"
curlwithcode -s $RESTAPI/

echo
echo "Sending GET request to /no-params/empty-response - expect HTTP 200 with no body"
curlwithcode $RESTAPI/no-params/empty-response

echo
echo "Sending GET request to /no-params/simple-response - expect HTTP 200 with simple body"
curlwithcode $RESTAPI/no-params/simple-response

echo
echo "Sending GET request to /no-params/complex-response - expect HTTP 200 with complex body"
curlwithcode $RESTAPI/no-params/complex-response

echo
echo "Sending GET request to /no-params/error-response - expect HTTP 418 with error message"
curlwithcode $RESTAPI/no-params/empty-response

PARAM='1234'

echo
echo "Sending GET request to /path-param/empty-response - expect HTTP 200 with no body"
curlwithcode $RESTAPI/path-param/empty-response/$PARAM

echo
echo "Sending GET request to /path-param/simple-response - expect HTTP 200 with simple body"
curlwithcode $RESTAPI/path-param/simple-response/$PARAM

echo
echo "Sending GET request to /path-param/complex-response - expect HTTP 200 with complex body"
curlwithcode $RESTAPI/path-param/complex-response/$PARAM

echo
echo "Sending GET request to /path-param/error-response - expect HTTP 418 with error message"
curlwithcode $RESTAPI/path-param/empty-response/$PARAM

PARAM='5678'

echo
echo "Sending GET request to /query-param/empty-response - expect HTTP 200 with no body"
curlwithcode $RESTAPI/query-param/empty-response?query-param=$PARAM

echo
echo "Sending GET request to /query-param/simple-response - expect HTTP 200 with simple body"
curlwithcode $RESTAPI/query-param/simple-response?query-param=$PARAM

echo
echo "Sending GET request to /query-param/complex-response - expect HTTP 200 with complex body"
curlwithcode $RESTAPI/query-param/complex-response?query-param=$PARAM

echo
echo "Sending GET request to /query-param/error-response - expect HTTP 418 with error message"
curlwithcode $RESTAPI/query-param/empty-response?query-param=$PARAM

echo
echo "Sending GET request to /body-param/empty-response - expect HTTP 501 with error message"
curlwithcode $RESTAPI/body-param/empty-response

echo
echo "Sending POST request to /body-param/empty-response - expect HTTP 200 with no body"
curlwithcode -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/empty-response

echo
echo "Sending GET request to /body-param/simple-response - expect HTTP 501 with error message"
curlwithcode $RESTAPI/body-param/simple-response

echo
echo "Sending POST request to /body-param/simple-response - expect HTTP 200 with simple body"
curlwithcode -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/simple-response

echo
echo "Sending GET request to /body-param/complex-response - expect HTTP 501 with error message"
curlwithcode $RESTAPI/body-param/complex-response

echo
echo "Sending POST request to /body-param/complex-response - expect HTTP 200 with complex body"
curlwithcode -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/complex-response

echo
echo "Sending GET request to /body-param/error-response - expect HTTP 501 with error message"
curlwithcode $RESTAPI/body-param/error-response

echo
echo "Sending POST request to /body-param/error-response - expect HTTP 418 with error message"
curlwithcode -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descriptor": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/error-response


echo
echo "Error cases - sending invalid path parameter"

PARAM='ABCD'

echo
echo "Sending GET request to /path-param/empty-response - expect invalid request response"
curlwithcode $RESTAPI/path-param/empty-response/$PARAM

echo
echo "Sending GET request to /path-param/simple-response - expect invalid request response"
curlwithcode $RESTAPI/path-param/simple-response/$PARAM

echo
echo "Sending GET request to /path-param/complex-response - expect invalid request response"
curlwithcode $RESTAPI/path-param/complex-response/$PARAM

echo
echo "Sending GET request to /path-param/error-response - expect invalid request response"
curlwithcode $RESTAPI/path-param/empty-response/$PARAM

echo
echo "Error cases - sending no query parameter"

echo
echo "Sending GET request to /query-param/empty-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/empty-response

echo
echo "Sending GET request to /query-param/simple-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/simple-response

echo
echo "Sending GET request to /query-param/complex-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/complex-response

echo
echo "Sending GET request to /query-param/error-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/empty-response

echo
echo "Error cases - sending invalid query parameter"

PARAM='EFGH'

echo
echo "Sending GET request to /query-param/empty-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/empty-response?query-param=$PARAM

echo
echo "Sending GET request to /query-param/simple-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/simple-response?query-param=$PARAM

echo
echo "Sending GET request to /query-param/complex-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/complex-response?query-param=$PARAM

echo
echo "Sending GET request to /query-param/error-response - expect invalid request response"
curlwithcode $RESTAPI/query-param/empty-response?query-param=$PARAM

echo
echo "Error cases - sending invalid body parameter"

echo
echo "Sending POST request to /body-param/empty-response - expect invalid request response"
curlwithcode -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descripto": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/empty-response

echo
echo "Sending POST request to /body-param/simple-response - expect invalid request response"
curlwithcode -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descripto": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/simple-response

echo
echo "Sending POST request to /body-param/complex-response - expect invalid request response"
curlwithcode -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"string": "hello", "descripto": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/complex-response

echo
echo "Sending POST request to /body-param/error-response - expect invalid request response"
curlwithcode -X POST -H 'Content-Type: application/json' -d '{"string": "hello", "descripto": "a wonderful day", "int_val": 42}' $RESTAPI/body-param/error-response
