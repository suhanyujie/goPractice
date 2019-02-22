## 棋牌游戏实例
### 功能
* 登陆游戏
* 查看房间列表
* 创建房间
* 加入房间
* 进行游戏
* 房间内聊天
* 游戏完成 退出房间
* 退出登陆

### 分析
#### 玩家信息 
* 用户唯一id
* 用户名用于显示
* 玩家等级
* 经验值

#### 子系统
* 玩家会话管理系统
* 大厅管理
* 房间管理，创建、销毁房间
* 游戏会话管理
* 聊天管理

#### 相关技术
* goroutine生命周期管理
* goroutine之间的通信
* 共享资源访问控制

## 小技巧
```
int,err:=strconv.Atoi(string)
 #string到int64
 int64, err := strconv.ParseInt(string, 10, 64)
 #int到string
 string:=strconv.Itoa(int)
 #int64到string
 string:=strconv.FormatInt(int64,10)
 --------------------- 
 作者：三少GG 
 来源：CSDN 
 原文：https://blog.csdn.net/pkueecser/article/details/50433460 
 版权声明：本文为博主原创文章，转载请附上博文链接！
```