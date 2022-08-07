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
7. 分布式解决方案：https://leeshengis.com/archives/275337
   1. 哨兵集群：即Sentinel集群
      1. 连接主库：sentinel monitor <master-name> <ip> <redis-port> <quorum>，
      2. 哨兵之前基于主库的pub/sub机制相互发现。哨兵只要和主库建立起来了连接，就可以从主库哪里订阅消息，拿到其他哨兵的连接信息
      3. 哨兵向主库发送INFO命令来完成的。就像下图所示，哨兵2给主库发送INFO命令，主库接受到这个命令后，就会把从库列表返回给哨兵
      4. 每个哨兵实例也提供pub/sub机制，客户端可以从哨兵订阅消息。
8. 主从切换：https://leeshengis.com/archives/274483
   1. 哨兵机制，即Sentinel机制
      1. 监控：监控主库的状态
      2. 选主：从从库中选出主库
         1. 根据网络
         2. 优先级
         3. 根据主库和从库的offset
         4. 根据从库的ID
      3. 通知：先在哨兵集群中竞选组长，然后通知给从库们。
9. 数据统计
   1. Set
      1. 并集 SUNIONSTORE  user:id  user:id  user:id:20200803 
      2. 差集：SDIFFSTORE  user:new  user:id:20200804 user:id  
      3. 交集：SINTERSTORE user:id:rem user:id:20200803 user:id:20200804
      4. Set的差集、并集和交集的计算复杂度较高，在数据量较大的情况下，如果直接执行这些计算，会导致Redis实例阻塞。所以，我给你分享一个小建议：你可以从主从集群中选择一个从库，让它专门负责聚合计算，或者是把数据读取到客户端，在客户端来完成聚合统计，
   2. List HSet
      1. 排序统计
   3. redis怎么设置过期的？（答给每个数据添加一个属性，过期时间，到了过期时间就删除，猜的）。怎么检查？（周期检查，惰性检查）。
   4. https://ost.51cto.com/posts/11473

