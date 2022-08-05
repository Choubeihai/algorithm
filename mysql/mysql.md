1. 一主一从：主要介绍主从复制中的问题
   1. https://leeshengis.com/archives/2-4----m-y-s-q-l-shi-zen-me-bao-zheng-zhu-bei-yi-zhi-de-
   2. binlog文件和位置
2. 一主多从：https://leeshengis.com/archives/2-7----zhu-ku-chu-wen-ti-le--cong-ku-zen-me-ban-
   1. 一主多从的切换问题：在一主多从的模式中，从库如何寻找新主库的日志位置是一个痛点，因此在5.6版本时引入了GTID协议。
   2. GTID的全称是Global Transaction Identifier，也就是全局事务ID，是一个事务在提交的时候生成的，是这个事务的唯一标识。它由两部分组成，格式是：GTID=server_uuid:gno 
   3. second_behind_master：表示从库和主库写入同一条语句的时间差，单位是s


