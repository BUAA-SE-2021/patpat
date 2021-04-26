# patpat

Auto Judger for BUAA-SE-OOP Course (2021 Spring)

## 评测机使用方法

### 一些准备工作

#### 下载评测机

见 GitHub 中的 Releases，下载对应版本即可。目前提供 Windows, Linux, MacOS(Intel 架构)的可执行文件。如有 ARM 架构版本的需要或程序无法正常运行，请联系我解决。

#### 自检编译和执行命令

统一使用 UTF-8 编码。

评测机使用的编译命令与执行命令如下，请确认你的程序能通过如下命令成功编译。

```bash
javac -encoding UTF-8 -cp folderName/src -d folderName/out folderName/src/*.java # 编译
java -classpath folderName/out Test # 执行
```

### 自评自测

请**严格按照**如下的目录层级与命名准备好待测文件：

- `patpat.exe` (Mac 与 Linux 用户的程序无后缀)
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

**运行方式:** 在 terminal(如 PowerShell)中运行如下命令：

注：对于 Mac 与 Linux 用户，需要首先`chmod +x patpat`，使评测机程序具有可执行权限。

```bash
./patpat stu -judge 1-18373722-朱英豪 # 请修改为自己的相关信息
./patpat ta -judge 1-18373722-朱英豪 -pwd your_pwd -tag 随意打一些文字，可以用来标注第几次尝试 # 无限重测命令
```

**评测结果:** 见生成的`$testcase$_result.md`。(其与该`patpat`程序在同级目录)

**离线模式:** 仅在 stu 模式下有效，添加参数 `-online=0`

### 查询助教评测结果

- 首先需要注册账号 (sid 为您的学号)

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

见`Makefile`。一键交叉编译可使用`make`。

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
  - # 多行输入，无输出
    - |
      quit
      quit again
      quit quit
  - # 多行输入，一行输出
    - |
      quit
      quit again
      QUIT
    - ----- Good Bye! -----
  # 多行输入，多行输出，类比以上写法，不建议使用。

# 这里 data 部分的数据类型为 [n][2]string。
# 第一级的 - 相当于一条小测试点
# 第二级的 - 第一条对应输入，可能有的第二条对应输出。
# (即第二级的输入是必须的，可以无输出。至少为 1 个二级 - 至多为 2 个二级 - 。)
```
