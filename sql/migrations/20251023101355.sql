-- Create "semesters" table
CREATE TABLE `semesters` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `start_date` datetime(3) NOT NULL,
  `end_date` datetime(3) NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_semesters_start_date` (`start_date`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `user_id` char(36) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` text NOT NULL,
  `phone_number` varchar(10) NULL,
  `account_status` tinyint NOT NULL DEFAULT 0,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `idx_users_email` (`email`),
  UNIQUE INDEX `idx_users_username` (`username`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "courses" table
CREATE TABLE `courses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `course_id` varchar(255) NOT NULL,
  `course_name` varchar(255) NOT NULL,
  `user_id` char(36) NOT NULL,
  `description` text NULL,
  `lecturers` text NOT NULL,
  `credits` bigint NOT NULL,
  `gpa` float NOT NULL DEFAULT 0,
  `semester_id` bigint NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_courses_course_id` (`course_id`),
  INDEX `idx_courses_semester_id` (`semester_id`),
  INDEX `idx_courses_user_id` (`user_id`),
  CONSTRAINT `fk_semesters_courses` FOREIGN KEY (`semester_id`) REFERENCES `semesters` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT,
  CONSTRAINT `fk_users_courses` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "tags" table
CREATE TABLE `tags` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `color` varchar(7) NOT NULL DEFAULT "#808080",
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_tags_name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "course_tags" table
CREATE TABLE `course_tags` (
  `course_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL,
  PRIMARY KEY (`course_id`, `tag_id`),
  INDEX `idx_course_tags_course_id` (`course_id`),
  INDEX `idx_course_tags_tag_id` (`tag_id`),
  CONSTRAINT `fk_course_tags_course` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_course_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
