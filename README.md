# CFN-step-by-step-go
This is a go implementation of the [Writing an AWS CloudFormation Resource Provider in Python: Step by Step](https://www.cloudar.be/awsblog/writing-an-aws-cloudformation-resource-provider-in-python-step-by-step/), but translated to Go. I've not got the thing running far enough to tell if my translation is correct, but it compiles, which is the first step :-D

Follow [Setting up your environment for developing extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html#resource-type-setup) to setup your extension development environment.

You should then be able to do a `make` followed by one of the several ways of running the output. 

I am stuck at the `sam local invoke TestEntrypoint --events sam-tests/create.json` step:

```
(cfn-go) bash-5.0$ sam local invoke TestEntrypoint --event sam-tests/create.json
Invoking handler (go1.x)
Failed to download a new amazon/aws-sam-cli-emulation-image-go1.x:rapid-1.0.0 image. Invoking with the already downloaded image.
Mounting /Users/tobyferguson/Development/CFN-step-by-step-go/bin as /var/task:ro,delegated inside runtime container
2021/02/25 21:45:58 Handler starting
START RequestId: 39e6a5ec-2557-1956-f123-a2eb92eef0f3 Version: $LATEST
Validation: Failed Validation
caused by: Region: zero value, BearerToken: zero value: baseError
null
2021/02/25 21:45:58 Handler received the CREATE action
END RequestId: 39e6a5ec-2557-1956-f123-a2eb92eef0f3
REPORT RequestId: 39e6a5ec-2557-1956-f123-a2eb92eef0f3	Init Duration: 192.37 ms	Duration: 5.69 ms	Billed Duration: 100 ms	Memory Size: 256 MB	Max Memory Used: 24 MB	

{"errorType":"baseError","errorMessage":"Validation: Failed Validation\ncaused by: Region: zero value, BearerToken: zero value"}
```


# MyCorp::EC2::Keypair

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mycorp-ec2-keypair.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go and main.go`, as they will be automatically overwritten.
