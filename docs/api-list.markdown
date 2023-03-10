
### 获取所有课程

#### 请求
POST http://api.jxjyzx.qdu.edu.cn/LearningSpace/list

Accept: application/json, text/plain, */*  
Accept-Encoding: gzip, deflate  
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8  
access-token: eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxMTA2NTIwMjAyMzA0NjMxIiwiaWQiOiJCNzY1MkZDRS03RkM1LTQyQTgtQTI4MS1DQjJCOTk3QUEyN0QiLCJleHAiOjE2NzgzNTk1OTMsImNyZWF0ZWQiOjE2NzgyNzMxOTMwNTR9.YaUmNdOCcrSsW7O-YmeKq5MCXKxVX3jhoKtWp0NHRur2wKqHq2SRP9xAF2rD_rq-oXOHBiigQO3FvUmmlRVnzQ  
Connection: keep-alive  
Content-Length: 65  
Content-Type: application/json;charset=UTF-8  
Host: api.jxjyzx.qdu.edu.cn  
Origin: http://student.jxjyzx.qdu.edu.cn  
Referer: http://student.jxjyzx.qdu.edu.cn/  
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36  

{
    "isPass": 1,
    "order": "",
    "orderField": "",
    "pageNum": 1,
    "pageSize": 10
}
#### 响应
application/json


#### 说明
![isPass](./isPass.png)


### 获取课程目录

#### 请求
GET http://api.jxjyzx.qdu.edu.cn/studyLearn/courseDirectoryProcess?courseOpenId={ID}

Accept: application/json, text/plain, */*  
Accept-Encoding: gzip, deflate  
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8  
access-token: eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxMTA2NTIwMjAyMzA0NjMxIiwiaWQiOiJCNzY1MkZDRS03RkM1LTQyQTgtQTI4MS1DQjJCOTk3QUEyN0QiLCJleHAiOjE2NzgzNTk1OTMsImNyZWF0ZWQiOjE2NzgyNzMxOTMwNTR9.YaUmNdOCcrSsW7O-YmeKq5MCXKxVX3jhoKtWp0NHRur2wKqHq2SRP9xAF2rD_rq-oXOHBiigQO3FvUmmlRVnzQ  
Connection: keep-alive  
Host: api.jxjyzx.qdu.edu.cn  
Origin: http://student.jxjyzx.qdu.edu.cn  
Referer: http://student.jxjyzx.qdu.edu.cn/  
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36  

#### 响应
application/json

#### 说明

