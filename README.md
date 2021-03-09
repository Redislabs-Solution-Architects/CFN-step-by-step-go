# CFN-step-by-step-go
This is a go implementation of the [Writing an AWS CloudFormation Resource Provider in Python: Step by Step](https://www.cloudar.be/awsblog/writing-an-aws-cloudformation-resource-provider-in-python-step-by-step/), but translated to Go. I've not got the thing running far enough to tell if my translation is correct, but it compiles, which is the first step :-D

## Installation
Follow [Setting up your environment for developing extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html#resource-type-setup) to setup your extension development environment.


## Configuration
Generate the requisite files:
```
cfn generate
```

Create an event thus (if necessary, export your AWS environment variables - either `AWS_PROFILE` or `AWS_ACCESS_KEY_ID/AWS_SECRET_KEY` - to setup the correct account access):

```
./mybin/create_test_events.sh CREATE ./example_inputs/create.json >events/create.json
```
I am stuck at the `sam local invoke TestEntrypoint --event events/create.json` step:

```
(cfn-go) bash-5.0$ sam local invoke TestEntrypoint --event events/create.json 
Invoking handler (go1.x)
Skip pulling image and use local one: amazon/aws-sam-cli-emulation-image-go1.x:rapid-1.20.0.

Mounting /Users/tobyferguson/Development/CFN-step-by-step-go/bin/ as /var/task:ro,delegated inside runtime container
START RequestId: 72d0544c-7286-4421-8f86-9ae6ced76b63 Version: $LATEST
2021/03/09 04:17:39 Handler starting
2021/03/09 04:17:39 Handler received the CREATE action
Validation: Failed Validation
caused by: BearerToken: zero value, Region: zero value: baseError
null
END RequestId: 72d0544c-7286-4421-8f86-9ae6ced76b63
REPORT RequestId: 72d0544c-7286-4421-8f86-9ae6ced76b63	Init Duration: 0.37 ms	Duration: 191.97 ms	Billed Duration: 200 ms	Memory Size: 256 MB	Max Memory Used: 256 MB	
{"errorMessage":"Validation: Failed Validation\ncaused by: BearerToken: zero value, Region: zero value","errorType":"baseError"}
```

The zero valued BearerToken is completely unexpected. I've seen nothing about this in the SAM documentation regarding test events. 

# MyCorp::EC2::Keypair

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mycorp-ec2-keypair.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go and main.go`, as they will be automatically overwritten.
