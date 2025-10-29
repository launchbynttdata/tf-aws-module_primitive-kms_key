data "aws_caller_identity" "current" {}

locals {
  ## Construct the policy to include the current account root as a principal. Useful for testing. Not useful in production.
  policy = {
    "EnableIAMUserPermissions" = {
      sid    = "EnableIAMUserPermissions"
      effect = "Allow"
      actions = [
        "kms:*"
      ]
      resources = ["*"]
      principals = {
        "AWS" = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
      }
    }
  }
}

module "kms_key" {
  source                  = "../../"
  description             = var.description
  key_usage               = var.key_usage
  policy                  = local.policy
  deletion_window_in_days = var.deletion_window_in_days
  is_enabled              = var.is_enabled
  enable_key_rotation     = var.enable_key_rotation
  rotation_period_in_days = var.enable_key_rotation ? var.rotation_period_in_days : null
  multi_region            = var.multi_region
  tags                    = var.tags
}
