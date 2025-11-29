-- 创建accounts （包含字段 id 主键， balance 账户余额）
CREATE TABLE `accounts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `balance` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO accounts (balance) VALUES (1000.00);
INSERT INTO accounts (balance) VALUES (500.00);

-- 创建transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）
CREATE TABLE `transactions`(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `from_account_id` int unsigned NOT NULL,
    `to_account_id`  int unsigned NOT NULL,
    `amount` decimal(10,2) NOT NULL,
    PRIMARY KEY (`id`)
);

--编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
-- 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
-- 如果余额不足，则回滚事务。
begin;

-- 检查账户 A 的余额是否足够 假设账户 A 的 id 为 1
SELECT balance FROM accounts WHERE id = 1 FOR UPDATE;

-- 如果账户 > 100 则向B转账
UPDATE balances SET balance = balance - 100 WHERE id = 1;
UPDATE balances SET balance = balance + 100 WHERE id = 2;
INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (1, 2, 100.00);
commit;

-- 如果账户 < 100 则回滚事务
rollback;
