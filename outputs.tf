output "key_id" {
  description = "The ID of the KMS key."
  value       = aws_kms_key.this.key_id
}

output "arn" {
  description = "The ARN of the KMS key."
  value       = aws_kms_key.this.arn
}

output "tags_all" {
  description = "A map of tags assigned to the KMS key, including those inherited from the provider default_tags configuration block."
  value       = aws_kms_key.this.tags_all
}
