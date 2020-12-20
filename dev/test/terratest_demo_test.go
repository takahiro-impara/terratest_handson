package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformVPC(t *testing.T) {
	t.Parallel()

	awsRegion := "ap-northeast-1"

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform/",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//Test VPC
	vpcId := terraform.Output(t, terraformOptions, "main_vpc_id")
	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)
	subnetId := terraform.Output(t, terraformOptions, "subnet-a_id")
	require.Equal(t, 1, len(subnets))
	assert.True(t, aws.IsPublicSubnet(t, subnetId, awsRegion))

	//Test EC2
	publicIp := terraform.Output(t, terraformOptions, "public_ip")
	url := fmt.Sprintf("http://%s", publicIp)
	http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello World", 30, 5*time.Second)
}