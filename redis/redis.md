# Reids
1. AOF日志：https://leeshengis.com/archives/271754
2. RDB内存快照：https://leeshengis.com/archives/271839 
   1. 全量快照，写时复制，
   2. T1时刻生成RDB快照，T1和T2时刻的修改，用AOF日志记录， 等到第二次做全量快照时，就可以清空AOF日志，
        因为此时的修改都已经记录到快照中了，恢复时就不再用日志了。
3. 缓存异常：当数据库和缓存中数据不一致时：https://leeshengis.com/archives/295812
4. 缓存污染问题：https://leeshengis.com/archives/297270
5. Redis实现分布式锁（红锁）：https://leeshengis.com/archives/301092
6. 主从库实现数据同步：https://leeshengis.com/archives/272852
   1. replication buffer是在传输RDB文件以后，将产生的数据存入replication buffer，即仅仅在全量复制以后
   2. repl_backlog_buffer：环形缓冲区，主和从各有一个offset，在增量复制时使用。