#### 需求分析
1. 前台用户登陆，商品展示，商品抢购
2. 后台订单管理

### 需求原型设计
1. 登陆页面，展示页面，抢购页面
2. 管理页面
### 系统需求分析
1. 前端页面需要承载大流量
2. 在大并发的状态下要解决超卖的问题
3. 后端接口需要满足横向扩展
### 系统架构设计
1. CDN--》流量负载 --》 流量拦截系统 --》 Go服务集群 --》 RabbitMQ --》 队列消费服务 --》 MySql

### RabbitMQ介绍
#### 定义和特征
1. RabbitMQ是面向消息的中间件，用于组件之间的解耦，主要体现在消息的发送者和消费者之间无强依赖关系
2. 特点：高可用，容易扩展，支持多语言
3. 使用场景：流量雪峰 ，异步处理，应用解耦
#### 安装
1. 安装命令:brew install rabbitmq
#### 快速入门
1. 插件管理命令： rabbitmq-plugins enable rabbitmq_management 安装并启用插件 disable 卸载并停止
2. 默认端口：15671
#### 核心概念
1. VirualHost:用来区分队列，进行逻辑上的隔离
2. Connection
3. exchange:
4. channel
5. queue:消息存储
6. bingding: 队列绑定到exchange上
#### 工作模式
1. Simple模式
2. 工作模式：起到负载均衡的作用
- 一个消息只能被一个消费者所获取
3. 订阅模式：不需要设置queueName, 但是需要设置exchangeName
- 消息被路由投递到多个队列，一个消息被多个消费者所获取
4. 路由模式：需要设置key, exchange要设置为direct
- 一个消息被多个多个消费者获取，并且消息的目标队列可以被生产者指定
5. Topic 话题模式：交换机类型设置为topic
- 一个消息被多个消费者获取，消息的目前queue可用bindingkey已统配服的方式来指定

#### 页面静态化
1. 加速页面访问速度
2. 减轻服务器负担
3. 网站更加安全不容易被攻击，比如sql注入
#### 分布式系统
1. 分布式
- 分布式系统是若干独立的计算机的集合，这些计算机对于用户来说就像是单个相关系统
- 集中试系统是所有程序和组件都在同一台计算机上
#### 一致性Hash算法
1. 用途：快速定位资源，均匀分布
2. 场景：分布式存储，分布式缓存，负载均衡















