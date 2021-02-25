package resource

import (
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (h handler.ProgressEvent, err error) {
	kpi := ec2.ImportKeyPairInput{
		KeyName:           prevModel.KeyName,
		PublicKeyMaterial: []byte(*prevModel.PublicKey),
	}
	ec2 := ec2.New(req.Session)
	response, err := ec2.ImportKeyPair(&kpi)
	if err != nil {
		return
	}
	currentModel.Fingerprint = response.KeyFingerprint
	currentModel.PublicKey = nil

	return handler.ProgressEvent{OperationStatus: handler.Success, ResourceModel: currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (h handler.ProgressEvent, err error) {
	kpi := ec2.DescribeKeyPairsInput{}
	kpi.SetKeyNames([]*string{prevModel.KeyName})
	client := ec2.New(req.Session)
	keypairs, err := client.DescribeKeyPairs(&kpi)
	if err != nil {
		return
	}
	if len(keypairs.KeyPairs) != 1 {
		return h, fmt.Errorf("Found %d keys", len(keypairs.KeyPairs))
	}
	return handler.ProgressEvent{OperationStatus: handler.Success}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return Create(req, prevModel, currentModel)
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (h handler.ProgressEvent, err error) {
	_, err = Read(req, prevModel, currentModel)
	if err != nil {
		return
	}

	client := ec2.New(req.Session)
	dkpi := ec2.DeleteKeyPairInput{}
	dkpi.SetKeyName(*prevModel.KeyName)

	_, err = client.DeleteKeyPair(&dkpi)
	if err != nil {
		return
	}
	return handler.ProgressEvent{OperationStatus: handler.Success}, err

}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (h handler.ProgressEvent, err error) {
	client := ec2.New(req.Session)
	keypairs, err := client.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	if err != nil {
		return
	}
	rms := make([]interface{}, len(keypairs.KeyPairs))
	for _, kp := range keypairs.KeyPairs {
		rms = append(rms, createModel(kp))
	}
	return handler.ProgressEvent{OperationStatus: handler.Success, ResourceModels: rms}, nil
}

func createModel(kps *ec2.KeyPairInfo) Model {
	return Model{KeyName: kps.KeyName, Fingerprint: kps.KeyFingerprint}
}
