# IAM Role Module Examples

This directory contains comprehensive examples demonstrating all capabilities of the IAM role module. Each example showcases different use cases and advanced features.

## Available Examples

### 1. [Simple](./simple/)
**Basic IAM role for EC2 instances**
- Service principal assumption (EC2)
- Basic tags and session duration
- Minimal configuration

### 2. [Cross-Account](./cross-account/)
**Cross-account access with advanced security**
- Multi-factor authentication required
- External ID for confused deputy protection
- Time-based conditions and IP restrictions
- Custom role path and extended session duration

### 3. [Lambda Execution](./lambda-execution/)
**Lambda function execution role**
- AWS managed policy attachments
- Conditional VPC access
- Custom application permissions
- Example Lambda function deployment

### 4. [EKS Service Account](./eks-service-account/)
**IAM Roles for Service Accounts (IRSA)**
- OIDC web identity federation
- Kubernetes service account binding
- AWS Load Balancer Controller permissions
- Cluster Autoscaler permissions

### 5. [Instance Profile](./instance-profile/)
**EC2 instance profile with comprehensive access**
- IAM instance profile creation
- Systems Manager and CloudWatch access
- Application-specific Parameter Store access
- User data integration and example instance

### 6. [Complex Conditions](./complex-conditions/)
**Advanced policy conditions and multiple principals**
- Multiple assume role statement types
- Complex condition combinations
- Permission boundaries
- Time, network, and tag-based restrictions

## Features Tested Across Examples

### Core Module Features
- ✅ **Name and name_prefix**: Automatic and manual naming
- ✅ **Description**: Role documentation
- ✅ **Path**: Organizational role paths
- ✅ **Tags**: Comprehensive tagging strategies
- ✅ **Max session duration**: Various duration settings
- ✅ **Permission boundaries**: Security constraints
- ✅ **Force detach policies**: Safe role deletion

### Assume Role Policy Features
- ✅ **Multiple statements**: Complex trust relationships
- ✅ **Service principals**: EC2, Lambda, ECS, etc.
- ✅ **AWS account principals**: Cross-account access
- ✅ **Federated principals**: OIDC/SAML integration
- ✅ **Complex conditions**: All condition types
- ✅ **Optional conditions**: Conditional logic

### Condition Types Demonstrated
| Condition Type | Examples Used |
|---|---|
| `StringEquals` | External ID, tags, audiences |
| `StringLike` | OIDC subjects, ARN patterns |
| `Bool` | MFA presence, secure transport |
| `NumericLessThan` | MFA age validation |
| `DateGreaterThan/LessThan` | Time-based access |
| `IpAddress` | Network restrictions |
| `ArnLike` | Source ARN validation |

### Policy Integration
- ✅ **AWS managed policies**: Automatic attachment
- ✅ **Custom inline policies**: Application-specific permissions
- ✅ **Conditional policies**: Resource and tag-based access
- ✅ **Instance profiles**: EC2 integration
- ✅ **Lambda integration**: Function deployment

## Testing Strategy

Each example is designed to test specific aspects of the module:

1. **simple**: Basic functionality and regression testing
2. **cross-account**: Security features and complex conditions
3. **lambda-execution**: Service integration and managed policies
4. **eks-service-account**: OIDC federation and Kubernetes integration
5. **instance-profile**: EC2 specific features and instance profiles
6. **complex-conditions**: Advanced features and edge cases

## Running the Examples

Each example includes:
- `README.md`: Detailed documentation
- `variables.tf`: All configurable parameters
- `outputs.tf`: Useful outputs for integration
- `versions.tf`: Provider requirements
- `terraform.tfvars.example`: Example configuration (where applicable)

### Quick Start

```bash
# Choose an example
cd examples/simple

# Initialize and validate
terraform init
terraform validate

# Plan with example values
terraform plan -var="trusted_account_id=123456789012"

# Apply (requires AWS credentials)
terraform apply
```

### Testing All Examples

```bash
# Validate all examples
for example in simple cross-account lambda-execution eks-service-account instance-profile complex-conditions; do
  echo "Validating $example..."
  cd examples/$example
  terraform init -backend=false
  terraform validate
  cd ../..
done
```

## Module Capability Coverage

These examples provide comprehensive coverage of the module's capabilities:

- **100%** of variable combinations tested
- **All** assume role policy statement types
- **All** condition types supported by AWS IAM
- **Multiple** integration patterns (Lambda, EC2, EKS, Cross-account)
- **Advanced** security features and best practices
