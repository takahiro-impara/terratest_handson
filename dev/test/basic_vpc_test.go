package test

import (
	"fmt"
	"testing"

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
	vpcId := terraform.Output(t, terraformOptions, "main_vpc_id")
	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)
	fmt.Print(subnets)
	subnetId := terraform.Output(t, terraformOptions, "subnet-a_id")

	require.Equal(t, 1, len(subnets))
	assert.True(t, aws.IsPublicSubnet(t, subnetId, awsRegion))
}
