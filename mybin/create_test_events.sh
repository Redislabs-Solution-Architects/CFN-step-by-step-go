#!/usr/local/bin/bash
# Create a test event for a CFN resource
# $1 - action (CREATE|UPDATE|DELETE|READ|LIST)
# $2 - json file containing 'desired state' values - i.e. the input parameters, as a valid json document
# Uses the AWS_PROFILE or AWS_ACCESS_KEY_ID/AWS_SECRET_KEY envars as per the
# aws docs for aws cli and get-session-token.

# Ensures all function exits abort the whole process
trap "exit 1" TERM
export TOP_PID=$$

# Outputs the event in json format

## Check args
[ $# -eq 2 ] || {
    echo "ERROR: $(basename $0): expected 2 arguments, got $#: $*"
    cat - <<-EOF

USAGE: $(basename $0)  [CREATE|UPDATE|DELETE|READ|LIST] json_file

        Use given file (containing 'properties' from Cloudformation template) to create a test event
EOF
    exit 1
}

[[ $1 =~ CREATE|UPDATE|DELETE|READ|LIST ]] || {
    echo "$0 - ERROR: $1 should match one of CREATE UPDATE DELETE READ LIST, but doesnt"
    exit 1
}

[[ -r $2 ]] || {
    echo "$0 - ERROR Cannot read $2"
    exit 1
}

readonly ACTION=$1
readonly TEST_FILE=$2
readonly CREDENTIALS_FILE=/tmp/credentials.json


function getsessiontoken {
    aws sts get-session-token --output json 2>/dev/null > ${CREDENTIALS_FILE:?} || {
        echo "Couldn't get a session token for AWS. You should probably export AWS_PROFILE to name the profile for your AWS credentials" 1>&2
        kill -s TERM $TOP_PID
    }
}

function getcredentials {
    if [ -r ${CREDENTIALS_FILE:?} -a  -s ${CREDENTIALS_FILE:?} ]
    then
        now=$(gdate +%s)
	    expiration=$(gdate -d $(jq -r '.Credentials.Expiration' ${CREDENTIALS_FILE:?}) +%s)
	    if (($now >= $expiration))	
	    then
	        # Credentials have expired. Update the file
	        getsessiontoken 
	    fi
    else
	    # No credentials file. Make one
	    getsessiontoken
    fi

    # Get here with credentials file whose credentials haven't expired
    jq '{credentials: {accessKeyId: .Credentials.AccessKeyId, secretAccessKey: .Credentials.SecretAccessKey, sessionToken: .Credentials.SessionToken}}' ${CREDENTIALS_FILE:?}
}

function makeEventTemplate {
    cat - <<\EOF
$credentials +    { "action": $action,
    "request": {
        "clientRequestToken": "4b90a7e4-b790-456b-a937-0cfdfa211dfe",
        "desiredResourceState": .,
	"logicalResourceIdentifier": "MyKeyPair"
    },
    "callbackContext": null
}
EOF
}

jq --argjson credentials "$(getcredentials)" --arg action "${ACTION:?}" -f <(makeEventTemplate) ${TEST_FILE:?}

