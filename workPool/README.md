## 15分钟写一个工作进程池
* 原文来自 https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923
* 需要fq

## 大意内容
* 文件结构如下：

```html
/go_worker_pool
    /work
        work.go
    /pool
        worker.go
        dispatcher.go
    bench_test.go
    main.go
```

* 构建项目的路径如下
> go/src/github.com/Lebonesco/go_worker_pool

* 最终的运行效果如下：

```shell
$ go run main.go
2018/10/06 15:53:43 starting application...
2018/10/06 15:53:43 starting worker:  1
2018/10/06 15:53:43 starting worker:  2
2018/10/06 15:53:43 starting worker:  3
2018/10/06 15:53:43 starting worker:  4
2018/10/06 15:53:43 starting worker:  5
2018/10/06 15:53:43 creating jobs...
worker [2] - created hash [2376065843] from word [iCMRAjWw]
worker [4] - created hash [121297580] from word [xhxKQFDa]
worker [1] - created hash [3193224551] from word [XVlBzgba]
worker [3] - created hash [1481401259] from word [hTHctcuA]
worker [5] - created hash [166906897] from word [FpLSjFbc]
worker [5] - created hash [1752784812] from word [QYhYzRyW]
...
```

### 创建模拟的工作单元

