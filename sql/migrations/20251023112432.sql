-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `is_email_verified` tinyint NOT NULL DEFAULT 0 AFTER `account_status`, ADD COLUMN `is_phone_verified` tinyint NOT NULL DEFAULT 0 AFTER `is_email_verified`;
