CREATE TABLE IF NOT EXISTS `users`
(
    `user_id`    CHAR(36)     NOT NULL,
    `username`   VARCHAR(255) NOT NULL UNIQUE,
    `email`      VARCHAR(255) NOT NULL UNIQUE,
    `password`   TEXT         NOT NULL,
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
);

create table if NOT EXISTS `semesters` (
    id int AUTO_INCREMENT PRIMARY KEY ,
    name VARCHAR(255) NOT NULL,
    start_date DATETIME NOT NULL ,
    end_date DATETIME NOT NULL ,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp on UPDATE  current_timestamp,
    created_at DATETIME NOT NULL DEFAULT  current_timestamp
);

CREATE TABLE IF NOT EXISTS `courses`
(
    `id`          INT AUTO_INCREMENT PRIMARY KEY,
    `course_id`   VARCHAR(255) NOT NULL,
    `course_name` VARCHAR(255) NOT NULL ,
    `user_id`     CHAR(36)     NOT NULL REFERENCES users (user_id),
    `description` TEXT         NOT NULL,
    `lecturers`   TEXT         NOT NULL DEFAULT '',
    `credits`     INT4         NOT NULL,
    `semester_id`    int NOT NULL REFERENCES semesters(id),
    `tags`        TEXT                  DEFAULT '' COMMENT "Array of tags",
    `updated_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `tags`
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

