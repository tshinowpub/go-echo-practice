-- Create TABLES
CREATE TABLE IF NOT EXISTS echo.users (
    user_id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_account_id VARCHAR(100) NOT NULL COMMENT 'アカウント作成時にユーザーが決めるID',
    user_name VARCHAR(100) NOT NULL COMMENT '任意のユーザー名',
    email text NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (user_id),
    UNIQUE users_idx1 (user_account_id)
) ENGINE=InnoDB;