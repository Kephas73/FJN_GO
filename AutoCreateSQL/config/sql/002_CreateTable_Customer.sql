CREATE TABLE IF NOT EXISTS tb_customer (
  tbc_id         INT          AUTO_INCREMENT PRIMARY KEY   COMMENT'Customer ID',
  tbc_name       VARCHAR(255) NOT NULL                     COMMENT'Customer name',
  tbc_age        INT          NOT NULL                     COMMENT'Customer age',
  tbc_birthday   DATETIME     NOT NULL                     COMMENT'Customer birthday'
)