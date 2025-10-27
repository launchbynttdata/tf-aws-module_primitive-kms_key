variable "description" {
  description = "A description of the KMS key."
  type        = string
  default     = "Managed by Terraform"
}

variable "key_usage" {
  description = "The intended use of the KMS key. Valid values are 'ENCRYPT_DECRYPT' and 'SIGN_VERIFY'."
  type        = string
  default     = "ENCRYPT_DECRYPT"
  validation {
    condition     = contains(["ENCRYPT_DECRYPT", "SIGN_VERIFY", "GENERATE_VERIFY_MAC"], var.key_usage)
    error_message = "key_usage must be either 'ENCRYPT_DECRYPT', 'SIGN_VERIFY', or 'GENERATE_VERIFY_MAC'."
  }
}

# tflint-ignore: terraform_unused_declarations
variable "policy" {
  description = "A JSON-formatted string that represents the key policy to attach to the KMS key."
  type = map(object({
    sid        = string
    effect     = string
    principals = map(list(string))
    actions    = list(string)
    resources  = list(string)
  }))
  default = null
}

variable "deletion_window_in_days" {
  description = "The waiting period, specified in number of days, before the KMS key is deleted after destruction is requested."
  type        = number
  default     = 30
}

variable "is_enabled" {
  description = "Specifies whether the KMS key is enabled."
  type        = bool
  default     = true
}

variable "enable_key_rotation" {
  description = "Specifies whether key rotation is enabled for the KMS key."
  type        = bool
  default     = false
}

variable "rotation_period_in_days" {
  description = "The number of days in the rotation period for the KMS key. Only applicable if enable_key_rotation is true."
  type        = number
  default     = 365
}

variable "multi_region" {
  description = "Specifies whether the KMS key is a multi-region key."
  type        = bool
  default     = false
}

variable "tags" {
  description = "A map of tags to assign to the KMS key."
  type        = map(string)
  default     = {}
}
