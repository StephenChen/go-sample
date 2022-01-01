###        

- 分布式 id 生成器
    - snowflake: github.com/bwmarrin/snowflake
        - 数值是64位，int64类型，被划分为四部分，
        - unused: 不含开头的第一个bit，因为这个bit是符号位。
        - time: 41位来表示收到请求时的时间戳，单位为毫秒，
        - datacenter_id: 五位来表示数据中心的id，
        - worker_id: 五位来表示机器的实例id，
        - sequence_id: 12位的循环自增id（到达1111,1111,1111后会归0）。

- 定时器实现
    - 时间堆: go 采用四叉堆
    - 时间轮

