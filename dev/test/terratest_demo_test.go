package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformBasic(t *testing.T) {
	t.Parallel()

	awsRegion := "ap-northeast-1"
	ExpectedVPCCidr := "10.10.0.0/16"
	ExpectedSubnetCidr := "10.10.1.0/24"

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform/",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//Test VPC
	vpcId := terraform.Output(t, terraformOptions, "main_vpc_id")
	subnetId := terraform.Output(t, terraformOptions, "subnet-a_id")
	vpcCidr := terraform.Output(t, terraformOptions, "main_vpc_cidr")
	subnetCidr := terraform.Output(t, terraformOptions, "main_subnet_cidr")
	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)

	require.Equal(t, 1, len(subnets))
	assert.True(t, aws.IsPublicSubnet(t, subnetId, awsRegion))
	assert.Equal(t, ExpectedVPCCidr, vpcCidr)
	assert.Equal(t, ExpectedSubnetCidr, subnetCidr)

	//Test EC2
	publicIp := terraform.Output(t, terraformOptions, "public_ip")
	ec2Id := terraform.Output(t, terraformOptions, "ec2_id")
	ec2ExpectedTag := terraform.Output(t, terraformOptions, "ec2_tags_name")
	TargetTag := aws.GetTagsForEc2Instance(t, awsRegion, ec2Id)

	assert.Equal(t, TargetTag["Name"], ec2ExpectedTag)
	assert.Equal(t, aws.GetPublicIpOfEc2Instance(t, ec2Id, awsRegion), publicIp)
	//url := fmt.Sprintf("http://%s", publicI)
	//http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello World", 30, 5*time.Second)
}
