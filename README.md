# ainews-be

## Description

<h3>Hello guys, this project is a personal blog for me to share new Ai information. I will try to update it every week. And I already finish the first step DDD refactoring. If you have any questions, please contact me. Thank you.</h3>

<br>

## Next Step

After I finish the backend, I will start to build the frontend. I will use React to build the frontend. I will try to finish it as soon as possible. Thank you. And I will integrate some Ai tools into this project, such as chatbot and images generator. 

## Quick Start

If you want to run this project in your local machine, you need to run this command in your terminal.

```bash
go mod download
```

```bash
go run main.go
```

And before you run this project, you need to create a pg database in your local machine. I will add docekr-compose file later. Thank you.


## 一点关于DDD的想法
这里没有在代码结构上分拆domain, 把adapter domainService entity等都放在各自的domain下面，后面可能会重构，把各自domain的各种东西都放在同一个目录下面。
目前的adapter中router承担的是路由转发的作用， handler做业务处理，把各个domain聚合到一起，application会做一些domain相关的校验或是一些domain相关的逻辑处理，repository和dao只做简单的CRUD，不把任何业务放在sql里面.当然这只是我的一点看法，有问题欢迎讨论

