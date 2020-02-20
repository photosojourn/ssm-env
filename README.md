# ssm-env

Golang shim to populate envrioment variables from SSM Parameter Store

## Usage

### Setting Values

The application relies on 1 environment variable `SERVICEPATH` which is append to the front of the name which forms the insert path. A `-l` flag is used to set the name of the file that will be fed in.

**NOTE: Do not leave any blank lines at the end of the file. This will cause the tool to crash**

### Getting Values

The application relies on 1 environment variable `SERVICEPATH` which is used to create the search path. An example could be `/attachments/developer/andy` or `/attachments/staging-demo/`. To run it simply call `ssm-env` without any parameters and this will cause the variables to printed out to STDOUT. To make them environment variables wrap the command in eval, `eval $(ssm-env)`.

## Build

To build the binary run `go build . ` from the directory. The only dependancy is the aws-sdk which can be installed via `go get -u github.com/aws/aws-sdk-go/...`

The AWS Golang SDK doc can be found [here](https://docs.aws.amazon.com/sdk-for-go/)