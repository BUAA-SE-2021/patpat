# patpat

Auto Judger for BUAA-SE-OOP Course (2021 Spring)

## 如何使用

### 自评

请**严格按照**如下的目录层级与命名准备好待测文件：

- `patpat.exe` (Mac与Linux用户的程序无后缀)
- `test`
  - `judge.yaml` (用于告知评测机需要测哪些测试用例，编写方法见下)
  - `testcase1.yaml`
  - `testcase2.yaml`
  - ...
- `1`(第几次作业)-`18373722`(学号)-`朱英豪`(姓名)
  - `src`
    - `Test.java` (程序运行的主入口)
    - 其余`*.java`
    - (有无`*.class`没有关系，评测机会重新编译)

**运行方式:** 在terminal(如cmd)中运行如下命令：

```bash
./patpat stu -judge 1-18373722-朱英豪 # 请修改为自己的相关信息
```

**评测结果:** 见生成的`$testcase$_result.md`。(其与该`patpat`程序在同级目录)

### 查询助教评测结果

- 首先需要注册账号 (sid为您的学号)

```bash
./patpat reg -sid your_sid -pwd your_password
# 如./patpat reg -sid 18373722 -pwd buaa-se-oop
```

注：账号创建后，密码**不支持修改**。如需修改，请联系我。

- 查询命令：

```bash
./patpat query -sid your_sid -pwd your_password
```

学号及对应的密码正确后即可显示截至到目前的所有评测结果。

### `judge.yaml`的编写

```yaml
tests: [testcase1.yaml, testcase2.yaml] # 可有更多，这是个列表
# 请将本judge.yaml置于test目录下
```

注：在自测时`testcase`的名称，需要与`test`目录下的`testcase`文件名相一致。

## 作业提交方式

- `1`(第几次作业)-`18373722`(学号)-`朱英豪`(姓名)
  - src
    - `Test.java` (程序运行的主入口)
    - 其余`*.java`

将上述文件夹压缩为`zip`格式后，上传至云平台。

注：请务必**反复确认**文件的命名，如有误将无法评测。

## 开发者指南

### 编译项目

数据库的配置在`mysql.yaml`中，参考如下：

```yaml
host: xx.xx.xx.xx
port: xxx
username: xxx
password: xxx
database: xxx
```

交叉编译：

```go
env GOOS=windows GOARCH=amd64 go build -o bin/patpat.exe main.go
env GOOS=linux GOARCH=amd64 go build -o bin/patpat-linux main.go
env GOOS=darwin GOARCH=amd64 go build -o bin/patpat-mac main.go
# 以上均为64位程序
```

### 测试文件编写

样例：

```yaml
name: testcase
data:
- # 单行输入，无输出
  - SUDO
- # 单行输入，单行输出
  - nc hh12345678 oop6324 [10086,10001] 1 [1-16]2,3
  - Course add illegal.
- # 单行输入，多行输出
  - myc 1 1000
  - |
    Page:1
    1.CID:bh00000002,Name:oop6326,Teachers:[A,B,10086],Capacity:1/100,Time:[1-10]1,5
    2.CID:bh00000004,Name:oop6328,Teachers:[A,B],Capacity:1/100,Time:[11-18]1,5
    3.CID:bh00000005,Name:oop6329,Teachers:[A,B],Capacity:1/100,Time:[1-18]1,6
    n-next page, l-last page, q-quit
# 有且仅有这3种情况，即data的数据类型为[n][2]string
```

### 贡献项目

- 发现Bug——提Issue
- 贡献测试数据——提Pull Request
- 为了时刻了解是否有新的测试样例，请Watch和Star该项目
- Plus，欢迎Follow我，谢谢！
