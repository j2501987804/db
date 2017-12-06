# 接口使用说明
Http Get/Post 皆可
<br>接口返回值为json，可在浏览器查看明细
## 账号
* 登陆<br>
http://localhost:8000/login?unionID=a&nick=b&icon=c&device=d&sex=1
* 绑定<br>
http://localhost:8000/bind?unionID=a&nick=b&icon=c&playerID=d
* 获取通信证ID<br>
http://localhost:8000/getpassport?usr=1&pwd=b
* 校验是否绑定<br>
http://localhost:8000/verifybinded?passportID=d
* 根据通行证ID获取用户名<br>
http://localhost:8000/getusr?passportID=d

## 点券
* 读点券<br>
http://localhost:8000/readcopons?playerID=a
* 写点券<br>
http://localhost:8000/writecopons?playerID=a&delta=b
## 金币
* 读金币<br>
http://localhost:8000/readgold?playerID=a
* 写金币<br>
http://localhost:8000/writegold?playerID=a&delta=b&tax=c
## 金币锁
* 加锁<br>
http://localhost:8000/lock?playerID=a
* 解锁<br>
http://localhost:8000/unlock?playerID=a
