package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION
func tableAwsEc2InstanceMetricCpuUtilizationDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_instance_metric_cpu_utilization_daily",
		Description: "AWS EC2 Instance Cloudwatch Metrics - CPU Utilization (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listEc2Instance,
			Hydrate:       listEc2InstanceMetricCpuUtilizationDaily,
		},
		GetMatrixItem: BuildRegionList,
		Columns: awsRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "instance_id",
					Description: "The ID of the instance.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listEc2InstanceMetricCpuUtilizationDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	instance := h.Item.(*ec2.Instance)
	return listCWMetricStatistics(ctx, d, "DAILY", "AWS/EC2", "CPUUtilization", "InstanceId", *instance.InstanceId)
}
