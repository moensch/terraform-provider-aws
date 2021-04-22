package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudwatchevents"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAwsCloudWatchEventSource(t *testing.T) {
	resourceName := "aws_cloudwatch_event_bus.test"
	dataSourceName := "data.aws_cloudwatch_event_source.test"

	namePrefix := "aws.partner/examplepartner.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { testAccPreCheck(t) },
		ErrorCheck: testAccErrorCheck(t, cloudwatchevents.EndpointsID),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsDataSourcePartnerEventSourceConfig(namePrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "arn"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_by"),
				),
			},
		},
	})
}

func testAccAwsDataSourcePartnerEventSourceConfig(namePrefix string) string {
	return fmt.Sprintf(`
resource "aws_cloudwatch_event_bus" "test" {
	event_source_name = "aws.partner/examplepartner.com/mycompany/059e96ef-bb3f-45d1-935e-064973277627"
	name              = "aws.partner/examplepartner.com/mycompany/059e96ef-bb3f-45d1-935e-064973277627"
}

data "aws_cloudwatch_event_source" "test" {
  name_prefix = "%s"
}
`, namePrefix)
}
