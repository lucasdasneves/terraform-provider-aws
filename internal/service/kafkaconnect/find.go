package kafkaconnect

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kafkaconnect"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindCustomPluginByARN(conn *kafkaconnect.KafkaConnect, arn string) (*kafkaconnect.DescribeCustomPluginOutput, error) {
	input := &kafkaconnect.DescribeCustomPluginInput{
		CustomPluginArn: aws.String(arn),
	}

	output, err := conn.DescribeCustomPlugin(input)

	if tfawserr.ErrCodeEquals(err, kafkaconnect.ErrCodeNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output, nil
}
