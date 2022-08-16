package test

import (
    "os"
    "testing"

    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/gruntwork-io/terratest/modules/logger"
    "github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
    // Construct the terraform options with default retryable errors to handle the most common
    // retryable errors in terraform testing.
    testSubject := "../" + os.Getenv("TEST_SUBJECT")
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        // Set the path to the Terraform code that will be tested.
        TerraformDir: testSubject,
    })

    // Clean up resources with "terraform destroy" at the end of the test.
    defer terraform.Destroy(t, terraformOptions)

    // Run "terraform init" and "terraform apply". Fail the test if there are any errors.
    terraform.InitAndApply(t, terraformOptions)

    // Run `terraform output` to get the values of output variables and check they have the expected values.
    output := terraform.Output(t, terraformOptions, "hello_world")
    assert.Equal(t, "Hello, World!", output)
}
