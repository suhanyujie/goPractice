# go的学习代码

## 项目list
- [x] 简单爬虫 [点此前往](/http)
- [x] 课题【15分钟写一个工作进程池】 [点此前往](/workPool)
- [ ] 数据结构 [点此前往](/dataStructure/tree)  进行中
- [ ] 排序算法 [点此前往](/studySort)  进行中
- [ ] fastcgi学习 [点此前往](/fastCgiStudy)  进行中


## 遇到的问题

### git提交

#### 使用git pull提示refusing to merge unrelated histories

```html
在执行git pull的时候,提示
`fatal: refusing to merge unrelated histories`
解决方法:
git pull --allow-unrelated-histories
```
* 参考资料 https://www.jianshu.com/p/39b890d6e73d

#### 修改git提交时的config中的username
* `vi ~/.gitconfig`; 然后在文件中直接修改即可

## 参考资料
* go的很多示例 https://www.kancloud.cn/itfanr/go-by-example/


## 红黑树性质：
* 红黑树是每个节点都带有颜色属性的二叉查找树，颜色或红色或黑色。在二叉查找树强制一般要求以外，对于任何有效的红黑树我们增加了如下的额外要求:
- 性质1. 节点是红色或黑色。
- 性质2. 根节点是黑色。
- 性质3 每个叶节点（NIL节点，空节点）是黑色的。
- 性质4 每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)
- 性质5. 从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点。
* 这些约束强制了红黑树的关键性质: 从根到叶子的最长的可能路径不多于最短的可能路径的两倍长。结果是这个树大致上是平衡的。因为操作比如插入、删除和查找某个值的最坏情况时间都要求与树的高度成比例，这个在高度上的理论上限允许红黑树在最坏情况下都是高效的，而不同于普通的二叉查找树。
