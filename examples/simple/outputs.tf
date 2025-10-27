output "arn" {
  description = "The ARN of the KMS key."
  value       = module.kms_key.arn
}

output "tags_all" {
  description = "A map of tags assigned to the KMS key, including those inherited from the provider default_tags configuration block."
  value       = module.kms_key.tags_all
}

output "key_id" {
  description = "The ID of the KMS key."
  value       = module.kms_key.key_id
}
