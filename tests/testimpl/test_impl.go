package testimpl

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	lcafTypes "github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	failedToDescribeKeyMsg    = "Failed to describe KMS key"
	failedToGetKeyPolicyMsg   = "Failed to get KMS key policy"
	failedToParseKeyPolicyMsg = "Failed to parse KMS key policy"
	failedToListKeyTagsMsg    = "Failed to list KMS key tags"
)

func TestComposableComplete(t *testing.T, ctx lcafTypes.TestContext) {
	kmsClient := GetAWSKMSClient(t)

	keyId := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "key_id")
	keyArn := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "arn")

	t.Run("TestKMSKeyExists", func(t *testing.T) {
		testKMSKeyExists(t, kmsClient, keyId, keyArn)
	})

	t.Run("TestKMSKeyProperties", func(t *testing.T) {
		testKMSKeyProperties(t, kmsClient, keyId)
	})

	t.Run("TestKMSKeyPolicy", func(t *testing.T) {
		testKMSKeyPolicy(t, kmsClient, keyId)
	})

	t.Run("TestKMSKeyTags", func(t *testing.T) {
		var keyTags map[string]interface{}
		terraform.OutputStructContext(t, context.Background(), ctx.TerratestTerraformOptions(), "tags_all", &keyTags)
		testKMSKeyTags(t, kmsClient, keyId, keyTags)
	})
}

// TestComposableCompleteReadOnly performs only read-only verification against
// already-deployed infrastructure. It must not create, update, or destroy any
// resources, so it reuses the same read-only assertions as
// TestComposableComplete without the setup/teardown apply cycle.
func TestComposableCompleteReadOnly(t *testing.T, ctx lcafTypes.TestContext) {
	kmsClient := GetAWSKMSClient(t)

	keyId := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "key_id")
	keyArn := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "arn")

	t.Run("TestKMSKeyExists", func(t *testing.T) {
		testKMSKeyExists(t, kmsClient, keyId, keyArn)
	})

	t.Run("TestKMSKeyProperties", func(t *testing.T) {
		testKMSKeyProperties(t, kmsClient, keyId)
	})

	t.Run("TestKMSKeyPolicy", func(t *testing.T) {
		testKMSKeyPolicy(t, kmsClient, keyId)
	})

	t.Run("TestKMSKeyTags", func(t *testing.T) {
		var keyTags map[string]interface{}
		terraform.OutputStructContext(t, context.Background(), ctx.TerratestTerraformOptions(), "tags_all", &keyTags)
		testKMSKeyTags(t, kmsClient, keyId, keyTags)
	})
}

func testKMSKeyExists(t *testing.T, kmsClient *kms.Client, keyId, keyArn string) {
	keyOutput, err := kmsClient.DescribeKey(context.TODO(), &kms.DescribeKeyInput{
		KeyId: &keyId,
	})
	require.NoError(t, err, failedToDescribeKeyMsg)
	require.NotNil(t, keyOutput.KeyMetadata, "Key metadata should not be nil")

	assert.Equal(t, keyId, *keyOutput.KeyMetadata.KeyId, "Expected key ID did not match actual ID!")
	assert.Equal(t, keyArn, *keyOutput.KeyMetadata.Arn, "Expected key ARN did not match actual ARN!")
}

func testKMSKeyProperties(t *testing.T, kmsClient *kms.Client, keyId string) {
	keyOutput, err := kmsClient.DescribeKey(context.TODO(), &kms.DescribeKeyInput{
		KeyId: &keyId,
	})
	require.NoError(t, err, failedToDescribeKeyMsg)
	require.NotNil(t, keyOutput.KeyMetadata, "Key metadata should not be nil")

	key := keyOutput.KeyMetadata

	// Verify key properties based on the test.tfvars configuration
	assert.Equal(t, types.KeyUsageTypeEncryptDecrypt, key.KeyUsage, "Expected key usage to be ENCRYPT_DECRYPT")
	assert.NotNil(t, key.CreationDate, "Key should have a creation date")
	assert.NotNil(t, key.Description, "Key should have a description")
	assert.Equal(t, "Example KMS key for encrypting data", *key.Description, "Key description should match expected value")

	// Verify key state
	assert.Equal(t, types.KeyStateEnabled, key.KeyState, "Key should be enabled")

	// Verify key is enabled
	assert.True(t, key.Enabled, "Key should be enabled")

	// Verify key rotation (should be disabled based on test.tfvars)
	rotationStatus, err := kmsClient.GetKeyRotationStatus(context.TODO(), &kms.GetKeyRotationStatusInput{
		KeyId: &keyId,
	})
	require.NoError(t, err, "Failed to get key rotation status")
	assert.True(t, rotationStatus.KeyRotationEnabled, "Key rotation should be enabled")

	// Verify multi-region setting (should be false based on test.tfvars)
	assert.Equal(t, types.OriginTypeAwsKms, key.Origin, "Key should be AWS KMS origin")
	if key.MultiRegion != nil {
		assert.False(t, *key.MultiRegion, "Key should not be multi-region")
	}
}

func testKMSKeyPolicy(t *testing.T, kmsClient *kms.Client, keyId string) {
	policyOutput, err := kmsClient.GetKeyPolicy(context.TODO(), &kms.GetKeyPolicyInput{
		KeyId:      &keyId,
		PolicyName: aws.String("default"),
	})
	require.NoError(t, err, failedToGetKeyPolicyMsg)
	require.NotNil(t, policyOutput.Policy, "Key policy should not be nil")

	// Parse the key policy
	var keyPolicy map[string]interface{}
	err = json.Unmarshal([]byte(*policyOutput.Policy), &keyPolicy)
	require.NoError(t, err, failedToParseKeyPolicyMsg)

	// Verify basic structure of key policy
	assert.Contains(t, keyPolicy, "Version", "Key policy should have Version")
	assert.Contains(t, keyPolicy, "Statement", "Key policy should have Statement")

	statements, ok := keyPolicy["Statement"].([]interface{})
	require.True(t, ok, "Policy should contain Statement array")
	require.Greater(t, len(statements), 0, "Policy should have at least one statement")

	// Verify the first statement structure (EnableIAMUserPermissions)
	statement := statements[0].(map[string]interface{})
	assert.Contains(t, statement, "Sid", "Statement should have Sid")
	assert.Contains(t, statement, "Effect", "Statement should have Effect")
	assert.Contains(t, statement, "Principal", "Statement should have Principal")
	assert.Contains(t, statement, "Action", "Statement should have Action")
	assert.Contains(t, statement, "Resource", "Statement should have Resource")

	// Verify statement properties
	sid, ok := statement["Sid"].(string)
	require.True(t, ok, "Sid should be a string")
	assert.Equal(t, "EnableIAMUserPermissions", sid, "Expected Sid to be EnableIAMUserPermissions")

	effect, ok := statement["Effect"].(string)
	require.True(t, ok, "Effect should be a string")
	assert.Equal(t, "Allow", effect, "Effect should be Allow")

	// Verify principal contains AWS account root
	principal, exists := statement["Principal"]
	require.True(t, exists, "Statement should have Principal")

	principalMap, ok := principal.(map[string]interface{})
	require.True(t, ok, "Principal should be a map")

	awsPrincipal, exists := principalMap["AWS"]
	require.True(t, exists, "Principal should contain AWS")

	// AWS principal can be either string or array
	if awsArray, ok := awsPrincipal.([]interface{}); ok {
		assert.Greater(t, len(awsArray), 0, "Should have at least one AWS principal")
		// Verify at least one principal contains account root
		foundRoot := false
		for _, principal := range awsArray {
			if principalStr, ok := principal.(string); ok {
				if strings.Contains(principalStr, ":root") {
					foundRoot = true
					break
				}
			}
		}
		assert.True(t, foundRoot, "Expected to find account root principal")
	} else if awsStr, ok := awsPrincipal.(string); ok {
		assert.True(t, strings.Contains(awsStr, ":root"), "Expected AWS principal to contain account root")
	}
}

func testKMSKeyTags(t *testing.T, kmsClient *kms.Client, keyId string, expectedTags map[string]interface{}) {
	if len(expectedTags) == 0 {
		return
	}

	// Get key tags from AWS
	tagsOutput, err := kmsClient.ListResourceTags(context.TODO(), &kms.ListResourceTagsInput{
		KeyId: &keyId,
	})
	require.NoError(t, err, failedToListKeyTagsMsg)

	// Convert AWS tags to map for comparison
	actualTags := make(map[string]string)
	for _, tag := range tagsOutput.Tags {
		actualTags[*tag.TagKey] = *tag.TagValue
	}

	// Verify expected tags exist
	for key, value := range expectedTags {
		if valueStr, ok := value.(string); ok {
			assert.Equal(t, valueStr, actualTags[key], "Tag %s should have expected value", key)
		}
	}

	// Verify specific expected tags from test.tfvars
	assert.Equal(t, "terratest", actualTags["Environment"], "Environment tag should be terratest")
	assert.Equal(t, "terratest-aws-kms", actualTags["Project"], "Project tag should be terratest-aws-kms")
	assert.Equal(t, "terraform", actualTags["ManagedBy"], "ManagedBy tag should be terraform")
}

func GetAWSKMSClient(t *testing.T) *kms.Client {
	awsKMSClient := kms.NewFromConfig(GetAWSConfig(t))
	return awsKMSClient
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
