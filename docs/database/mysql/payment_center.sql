--
-- ER/Studio 8.0 SQL Code Generation
-- Company :      李大福
-- Project :      交易中心.DM1
-- Author :       李大福
--
-- Date Created : Monday, May 21, 2012 15:15:16
-- Target DBMS : MySQL 5.x
--

-- 
-- TABLE: `account` 
--

CREATE TABLE `account`(
    `id`                  BIGINT         AUTO_INCREMENT,
    `password`            VARCHAR(32),
    `freeze`              TINYINT,
    `currency`            TINYINT,
    `total_amount`        BIGINT,
    `use_amount`          BIGINT,
    `withdrawals_amount`  BIGINT,
    `freeze_amount`       BIGINT,
    `unsettled_amount`    BIGINT,
    `create_time`         INT,
    `update_time`         INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_freeze` 
--

CREATE TABLE `account_freeze`(
    `id`           BIGINT          AUTO_INCREMENT,
    `account_id`   BIGINT,
    `type`         TINYINT,
    `status`       TINYINT,
    `enabled`      TINYINT,
    `reason`       VARCHAR(255),
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_freeze_log` 
--

CREATE TABLE `account_freeze_log`(
    `id`                 BIGINT          AUTO_INCREMENT,
    `account_freeze_id`  BIGINT,
    `account_id`         BIGINT,
    `type_code`          TINYINT,
    `type`               VARCHAR(32),
    `status_code`        TINYINT,
    `status`             VARCHAR(255),
    `memo`               CHAR(10),
    `create_time`        INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_log` 
--

CREATE TABLE `account_log`(
    `id`                  BIGINT          AUTO_INCREMENT,
    `account_id`          BIGINT,
    `opt_type_code`       TINYINT,
    `opt_type`            VARCHAR(32),
    `opt_id`              BIGINT,
    `opt_status_code`     TINYINT,
    `opt_status`          VARCHAR(32),
    `freeze_code`         TINYINT,
    `freeze`              VARCHAR(32),
    `currency_code`       TINYINT,
    `currency`            VARCHAR(32),
    `total_amount`        BIGINT,
    `use_amount`          BIGINT,
    `withdrawals_amount`  BIGINT,
    `freeze_amount`       BIGINT,
    `unsettled_amount`    BIGINT,
    `memo`                VARCHAR(255),
    `create_time`         INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_mapping` 
--

CREATE TABLE `account_mapping`(
    `id`           BIGINT          AUTO_INCREMENT,
    `account_id`   BIGINT,
    `type`         TINYINT,
    `object`       VARCHAR(255),
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_mapping_log` 
--

CREATE TABLE `account_mapping_log`(
    `id`                  BIGINT          AUTO_INCREMENT,
    `account_mapping_id`  BIGINT,
    `object`              VARCHAR(32),
    `type_code`           TINYINT,
    `type`                VARCHAR(32),
    `memo`                VARCHAR(255),
    `create_time`         INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_settings` 
--

CREATE TABLE `account_settings`(
    `id`           BIGINT          AUTO_INCREMENT,
    `account_id`   BIGINT,
    `name`         VARCHAR(255),
    `value`        VARCHAR(255),
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `account_settings_log` 
--

CREATE TABLE `account_settings_log`(
    `id`                   BIGINT          AUTO_INCREMENT,
    `account_settings_id`  BIGINT,
    `name`                 VARCHAR(255),
    `value`                VARCHAR(255),
    `memo`                 VARCHAR(255)    NOT NULL,
    `create_time`          INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `dictionary` 
--

CREATE TABLE `dictionary`(
    `id`           BIGINT          AUTO_INCREMENT,
    `parent_id`    BIGINT,
    `name`         VARCHAR(255),
    `value`        VARCHAR(255),
    `order`        INT             DEFAULT 0,
    `default`      TINYINT,
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `dictionary_log` 
--

CREATE TABLE `dictionary_log`(
    `id`             BIGINT          AUTO_INCREMENT,
    `dictionary_id`  BIGINT,
    `parent_id`      BIGINT,
    `name`           VARCHAR(255),
    `value`          VARCHAR(255),
    `order`          INT             DEFAULT 0,
    `default`        TINYINT,
    `create_time`    INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `funds_freeze` 
--

CREATE TABLE `funds_freeze`(
    `id`           BIGINT     AUTO_INCREMENT,
    `account_id`   BIGINT,
    `type`         TINYINT,
    `status`       TINYINT,
    `amount`       BIGINT,
    `enabled`      TINYINT,
    `reason`       TEXT,
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `funds_freeze_log` 
--

CREATE TABLE `funds_freeze_log`(
    `id`               BIGINT          AUTO_INCREMENT,
    `funds_freeze_id`  BIGINT,
    `account_id`       BIGINT,
    `amount`           BIGINT,
    `type_code`        TINYINT,
    `type`             VARCHAR(32),
    `status_code`      TINYINT,
    `status`           VARCHAR(32),
    `memo`             VARCHAR(255),
    `create_time`      INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `payment` 
--

CREATE TABLE `payment`(
    `id`                BIGINT     AUTO_INCREMENT,
    `payee_account_id`  BIGINT,
    `transfer_id`       BIGINT,
    `payer_account_id`  BIGINT,
    `amount`            BIGINT,
    `type`              TINYINT,
    `status`            TINYINT,
    `currency`          TINYINT,
    `enabled`           TINYINT,
    `create_time`       INT,
    `update_time`       INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `payment_log` 
--

CREATE TABLE `payment_log`(
    `id`             BIGINT          AUTO_INCREMENT,
    `payment_id`     BIGINT,
    `amount`         BIGINT,
    `type_code`      TINYINT,
    `type`           VARCHAR(32),
    `status_code`    TINYINT,
    `status`         VARCHAR(32),
    `currency_code`  TINYINT,
    `currency`       VARCHAR(32),
    `memo`           VARCHAR(255),
    `create_time`    INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `receivables` 
--

CREATE TABLE `receivables`(
    `id`                BIGINT     AUTO_INCREMENT,
    `payer_account_id`  BIGINT,
    `transfer_id`       BIGINT,
    `payee_account_id`  BIGINT,
    `amount`            BIGINT,
    `type`              TINYINT,
    `status`            TINYINT,
    `currency`          TINYINT,
    `enabled`           INT,
    `create_time`       INT,
    `update_time`       INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `receivables_log` 
--

CREATE TABLE `receivables_log`(
    `id`              BIGINT          AUTO_INCREMENT,
    `receivables_id`  BIGINT,
    `amount`          BIGINT,
    `type_code`       TINYINT,
    `status_code`     TINYINT,
    `type`            VARCHAR(32),
    `status`          VARCHAR(32),
    `currency_code`   TINYINT,
    `currency`        VARCHAR(32),
    `memo`            VARCHAR(255),
    `create_time`     INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `recharge` 
--

CREATE TABLE `recharge`(
    `id`           BIGINT     AUTO_INCREMENT,
    `account_id`   BIGINT,
    `amount`       BIGINT,
    `status`       TINYINT,
    `enabled`      TINYINT,
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `recharge_log` 
--

CREATE TABLE `recharge_log`(
    `id`           BIGINT          AUTO_INCREMENT,
    `recharge_id`  BIGINT,
    `account_id`   BIGINT,
    `amount`       BIGINT,
    `status_code`  TINYINT,
    `status`       VARCHAR(32),
    `memo`         VARCHAR(255),
    `create_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `settle` 
--

CREATE TABLE `settle`(
    `id`                  BIGINT     AUTO_INCREMENT,
    `account_id`          BIGINT,
    `type`                TINYINT,
    `recharge_amount`     BIGINT,
    `withdrawals_amount`  BIGINT,
    `transfer_amount`     BIGINT,
    `freeze_amount`       BIGINT,
    `create_time`         INT,
    `update_time`         INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `settle_log` 
--

CREATE TABLE `settle_log`(
    `id`           BIGINT          AUTO_INCREMENT,
    `settle_id`    BIGINT,
    `object`       VARCHAR(32),
    `amount`       BIGINT,
    `type_code`    TINYINT,
    `type`         VARCHAR(32),
    `memo`         VARCHAR(255),
    `create_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `share_profit` 
--

CREATE TABLE `share_profit`(
    `id`                          BIGINT     AUTO_INCREMENT,
    `be_share_profit_account_id`  BIGINT,
    `transfer_id`                 BIGINT,
    `amount`                      BIGINT,
    `type`                        TINYINT,
    `status`                      TINYINT,
    `currency`                    TINYINT,
    `enabled`                     TINYINT,
    `create_time`                 INT,
    `update_time`                 INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `share_profit_log` 
--

CREATE TABLE `share_profit_log`(
    `id`               BIGINT          AUTO_INCREMENT,
    `share_profit_id`  BIGINT,
    `amount`           BIGINT,
    `type_code`        TINYINT,
    `type`             VARCHAR(32),
    `status_code`      TINYINT,
    `status`           VARCHAR(32),
    `currency_code`    TINYINT,
    `currency`         TINYINT,
    `memo`             VARCHAR(255),
    `create_time`      INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `transaction` 
--

CREATE TABLE `transaction`(
    `id`                BIGINT     AUTO_INCREMENT,
    `trigger_id`        BIGINT,
    `payer_account_id`  BIGINT,
    `payee_account_id`  BIGINT,
    `method`            TINYINT,
    `pay_method`        TINYINT,
    `amount`            BIGINT,
    `type`              TINYINT,
    `status`            TINYINT,
    `settle_method`     TINYINT,
    `settle_status`     TINYINT,
    `settle_time`       INT,
    `enabled`           TINYINT,
    `start_time`        INT,
    `end_time`          INT,
    `create_time`       INT,
    `update_time`       INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `transaction_log` 
--

CREATE TABLE `transaction_log`(
    `id`                  BIGINT          AUTO_INCREMENT,
    `transaction_id`      BIGINT,
    `trigger_id`          BIGINT,
    `payer_account_id`    BIGINT,
    `payee_account_id`    BIGINT,
    `method_code`         TINYINT,
    `method`              VARCHAR(32)     NOT NULL,
    `pay_method_code`     TINYINT,
    `pay_method`          VARCHAR(32),
    `amount`              BIGINT,
    `type_code`           TINYINT,
    `type`                VARCHAR(32),
    `status_code`         TINYINT,
    `status`              VARCHAR(32),
    `settle_method_code`  TINYINT,
    `settle_method`       VARCHAR(32),
    `settle_status_code`  TINYINT,
    `settle_status`       VARCHAR(32),
    `memo`                VARCHAR(255),
    `start_time`          INT,
    `end_time`            BIGINT,
    `create_time`         INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `transfer` 
--

CREATE TABLE `transfer`(
    `id`                BIGINT     AUTO_INCREMENT,
    `payer_account_id`  BIGINT,
    `payee_account_id`  BIGINT,
    `amount`            BIGINT,
    `status`            TINYINT,
    `enabled`           TINYINT,
    `create_time`       INT,
    `update_time`       INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `transfer_log` 
--

CREATE TABLE `transfer_log`(
    `id`                BIGINT          AUTO_INCREMENT,
    `transfer_id`       BIGINT,
    `payer_account_id`  BIGINT,
    `payee_account_id`  BIGINT,
    `amount`            BIGINT,
    `status_code`       TINYINT,
    `status`            VARCHAR(32),
    `memo`              VARCHAR(255),
    `create_time`       INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `withdrawals` 
--

CREATE TABLE `withdrawals`(
    `id`           BIGINT     AUTO_INCREMENT,
    `account_id`   BIGINT,
    `amount`       BIGINT,
    `status`       TINYINT,
    `enabled`      TINYINT,
    `create_time`  INT,
    `update_time`  INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



-- 
-- TABLE: `withdrawals_log` 
--

CREATE TABLE `withdrawals_log`(
    `id`              BIGINT          AUTO_INCREMENT,
    `withdrawals_id`  BIGINT,
    `account_id`      BIGINT,
    `amount`          BIGINT,
    `status_code`     TINYINT,
    `status`          VARCHAR(32)     NOT NULL,
    `memo`            VARCHAR(255),
    `create_time`     INT,
    PRIMARY KEY (`id`)
)ENGINE=INNODB
;



