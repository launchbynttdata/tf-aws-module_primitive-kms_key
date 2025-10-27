data "aws_iam_policy_document" "kms_policy" {
  dynamic "statement" {
    for_each = var.policy != null ? var.policy : {}
    content {
      sid       = statement.value.sid
      effect    = statement.value.effect
      actions   = statement.value.actions
      resources = statement.value.resources
      dynamic "principals" {
        for_each = statement.value.principals
        content {
          type        = principals.key
          identifiers = principals.value
        }
      }
    }
  }
}

resource "aws_kms_key" "this" {
  description             = var.description
  key_usage               = var.key_usage
  policy                  = var.policy != null ? data.aws_iam_policy_document.kms_policy.json : null
  deletion_window_in_days = var.deletion_window_in_days
  is_enabled              = var.is_enabled
  enable_key_rotation     = var.enable_key_rotation
  rotation_period_in_days = var.enable_key_rotation ? var.rotation_period_in_days : null
  multi_region            = var.multi_region
  tags                    = local.tags
}
