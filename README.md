# What

A small application to demonstrate the use of the `go-swagger` library with the
go runtime on AWS Lambda.

# Requirements

- Modern go toolchain - I used go 1.14.6
- go-swagger - I used v0.25.0
- AWS SAM application management tool
- Basic utilities: `zip` for function upload to AWS, `curl` and `jq` for testing

# Build, deploy, test, clean up

- Generate the Swagger server-side stub code

```
swagger generate server -f swagger.yaml
```

This will generate `models` and `restapi` directories in the main directory.

- Build the application

```
cd deploy
./build.sh
```

This builds the application in the top level directory of the repo; it copies
the resulting executable into the deploy directory.

- Deploy the application

```
./deploy.sh
```

This zips the executable and uploads it to AWS using the `sam` tool and
associated configuration file (`sam-deploy.toml`). On successful completion, it
prints the identifier for the Lambda function and the endpoint of the REST API.

- Run the tests

```
cd ../tests
export RESTAPI=<ENDPOINT OUTPUT FROM DEPLOYMENT PROCESS>
./tests.sh
```

The tests demonstrate successful calling of all of the endpoints defined in the
API.

- Removing the API

The easiest way to remove the application is via the AWS CloudFormations
dashboard.


