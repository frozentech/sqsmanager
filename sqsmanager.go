package sqsmanager

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// SQSManager ...
type SQSManager struct {
	Client sqsiface.SQSAPI
}

// New ...
func New() *SQSManager {
	return &SQSManager{
		Client: sqs.New(session.New()),
	}
}

// Publish ...
func (me *SQSManager) Publish(delay int64,
	destination string,
	data interface{}) (err error) {

	b, _ := json.Marshal(data)

	input := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(b)),
		QueueUrl:     aws.String(destination),
		DelaySeconds: aws.Int64(delay),
	}

	_, err = me.Client.SendMessage(input)
	return
}
