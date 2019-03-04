CREATE DATABASE
  IF NOT EXISTS "apiserver" DEFAULT CHARACTER
  SET utf8;

USE "apiserver";
CREATE TABLE `users`
(
  `id`         int(11) unsigned                        NOT NULL AUTO_INCREMENT,
  `name`       varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password`   varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp                               NULL     DEFAULT NULL,
  `updated_at` timestamp                               NULL     DEFAULT NULL,
  `deleted_at` timestamp                               NULL     DEFAULT NULL,
  PRIMARY KEY
    (`id`),
  UNIQUE KEY `name`
    (`name`),
  KEY `idx_users_deleted_at`
    (`deleted_at`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
