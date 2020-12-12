package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestTerraformVPC(t *testing.T) {
	t.Parallel()

	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform/",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	vpcId := terraform.Output(t, terraformOptions, "main_vpc_id")
	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)
	require.Equal(t, 0, len(subnets))
}
