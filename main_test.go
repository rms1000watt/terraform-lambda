package main_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	main "github.com/rms1000watt/terraform-lambda-go"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	request events.APIGatewayProxyRequest
	expect  string
	err     error
}

func TestHandler(t *testing.T) {
	tests := []Test{
		{
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			expect:  "Hello Paul",
			err:     nil,
		},
		{
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     main.ErrNameNotProvided,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
