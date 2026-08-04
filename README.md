# tf-aws-module_primitive-kms_key

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.2 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_kms_key.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key) | resource |
| [aws_iam_policy_document.kms_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_deletion_window_in_days"></a> [deletion\_window\_in\_days](#input\_deletion\_window\_in\_days) | The waiting period, specified in number of days, before the KMS key is deleted after destruction is requested. | `number` | `30` | no |
| <a name="input_description"></a> [description](#input\_description) | A description of the KMS key. | `string` | `"Managed by Terraform"` | no |
| <a name="input_enable_key_rotation"></a> [enable\_key\_rotation](#input\_enable\_key\_rotation) | Specifies whether key rotation is enabled for the KMS key. | `bool` | `false` | no |
| <a name="input_is_enabled"></a> [is\_enabled](#input\_is\_enabled) | Specifies whether the KMS key is enabled. | `bool` | `true` | no |
| <a name="input_key_usage"></a> [key\_usage](#input\_key\_usage) | The intended use of the KMS key. Valid values are 'ENCRYPT\_DECRYPT' and 'SIGN\_VERIFY'. | `string` | `"ENCRYPT_DECRYPT"` | no |
| <a name="input_multi_region"></a> [multi\_region](#input\_multi\_region) | Specifies whether the KMS key is a multi-region key. | `bool` | `false` | no |
| <a name="input_policy"></a> [policy](#input\_policy) | A JSON-formatted string that represents the key policy to attach to the KMS key. | <pre>map(object({<br/>    sid        = string<br/>    effect     = string<br/>    principals = map(list(string))<br/>    actions    = list(string)<br/>    resources  = list(string)<br/>  }))</pre> | `null` | no |
| <a name="input_rotation_period_in_days"></a> [rotation\_period\_in\_days](#input\_rotation\_period\_in\_days) | The number of days in the rotation period for the KMS key. Only applicable if enable\_key\_rotation is true. | `number` | `365` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the KMS key. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the KMS key. |
| <a name="output_key_id"></a> [key\_id](#output\_key\_id) | The ID of the KMS key. |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the KMS key, including those inherited from the provider default\_tags configuration block. |
<!-- END_TF_DOCS -->
