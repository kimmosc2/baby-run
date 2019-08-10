Baby-run  
=============================
 
## Baby-run 是什么?  
Baby-run是一个简单的压测工具,使用Go编写,用来测试你的Web项目性能,目前功能较简单,
只能在指定的时间内向指定的Url发起GET请求,在压测完成后,Baby-run会向您报告本次压
测的数据,如响应时间等.

## Baby-run 的目标?
Baby-run的最终目的是成为"自动的Postman",用户可以构造表单使用指定的方法请求目标.

## 示例
1.在10秒内使用1个协程(轻量级线程)来发起测试
```
baby -u http://www.example.com
```
以下是测试输出:
<details><code>
<pre>== Result ============================================
总协程数:1
单协程持续时间:10s
总请求次数:17
成功数:17
失败数:0
成功占比:100.00%
总时间:9.5253367s
平均时间:560ms
=======================================================
</pre>
</code>
</details>
  
2.在3秒内使用5个协程来发起测试
```
baby -u http://www.baidu.com -t 3 -c 5
```
以下是测试输出:
<details><code>
<pre>== Result ============================================
总协程数:5
单协程持续时间:3s
总请求次数:432
成功数:432
失败数:0
成功占比:100.00%
总时间:19.8242759s
平均时间:45ms
=======================================================
</pre>
</code>
</details>
  
## 如何安装? 
您可以使用以下两种方式来使用Baby-run: 
#### 下载并安装二进制包 
我们提供了编译好的二进制包,您下载并解压后将其放入C:\Windows(只要是Path环境变量所
在目录都可以),运行cmd输入baby就可以使用了.  
下载地址: [百度网盘](https://pan.baidu.com/s/1rndTNQR8iMRpJ0_oM4B0gQ)  
#### 源代码编译:
```
cd Baby-run/cmd & go build baby.go
```

## 平台
Windows/Linux/Mac(Mac平台需自行编译)
## 提交Bug  
发现Bug请至issue板块提出,贡献代码请单独拉PR
## 版权和许可信息  
[MIT license](https://github.com/kimmosc2/baby-run/blob/master/LICENSE)
## 贡献者
[BuTn](https://github.com/kimmosc2)