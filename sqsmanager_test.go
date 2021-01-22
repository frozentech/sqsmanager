package sqsmanager_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/frozentech/sqsmanager"
	"github.com/stretchr/testify/assert"
)

type SqsClinetMock struct {
	sqsiface.SQSAPI
}

func (m SqsClinetMock) SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return &sqs.SendMessageOutput{}, nil
}

func Test_SendTransactionSqs(t *testing.T) {
	manager := sqsmanager.New()
	manager.Client = &SqsClinetMock{}
	assert.Nil(t, manager.Publish(0, "SQS_NAME", ""))

}
