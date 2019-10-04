CREATE TABLE IF NOT EXISTS tb_product (
  tbp_id     INT          AUTO_INCREMENT PRIMARY KEY   COMMENT'Product ID',
  tbp_name   VARCHAR(255) NOT NULL                     COMMENT'Product name',
  tbp_pice   FLOAT        NOT NULL                     COMMENT'Product pice',
  tbp_expiry DATETIME     NOT NULL                     COMMENT'Product expiry'
)
