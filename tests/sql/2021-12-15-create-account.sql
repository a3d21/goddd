CREATE TABLE `account`
(
    `id`         bigint(20) AUTO_INCREMENT,
    `name`       varchar(255),
    `amount`     bigint(20),
    `created_at` datetime NULL,
    `updated_at` datetime NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX idx_account_name (`name`)
);