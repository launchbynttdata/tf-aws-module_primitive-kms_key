description             = "Example KMS key for encrypting data"
key_usage               = "ENCRYPT_DECRYPT"
deletion_window_in_days = null
is_enabled              = true
enable_key_rotation     = true
rotation_period_in_days = 365
multi_region            = false
policy                  = {}
tags = {
  Environment = "terratest"
  Project     = "terratest-aws-kms"
  ManagedBy   = "terraform"
}
