# 目录结构及文件说明
1. src/bypass 用来绕过沙箱的工具
2. src/config 加载配置文件
3. src/encrypt 各种对称加密算法实现
4. src/loader 一个用来验证加解密绕过杀软的 demo
5. src/model 包含模型和模型的加载程序
6. src/template 包含文件模板
   1. src/template/main 入口程序 
   2. src/template/monitor 各种监控器，例如摄像头，麦克风，系统文件，进程状态等
7. src/utils 小工具 
8. src/virus 恶意软件样本(未加密)
9. base64Str.bs64 恶意软件加密后再用 base64 编码的结果 
10. configure.yml 配置文件 
11. signature.txt 恶意软件数字签名，用以验证在解密时恶意软件是否被正确解密

# Todo
- [ ] 各种对称加密算法的实现
- [ ] 监控器的实现
- [ ] 模板制作和替换程序实现
- [ ] 优化，减小打包后恶意软件大小

# 流程图
![lifecycle](/assets/lifecycle2.png)

# 实现原理
项目分为两个部分，一个是加壳程序的开发，一个是模板程序的开发。加壳程序利用模板对恶意软件进行加壳，壳上附带有监控摄像头、麦克风等功能。
