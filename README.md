# patpat

Auto Judger for BUAA-OOP Course

## 如何使用

请**严格按照**如下的目录层级与命名准备好待测文件：

- `patpat.exe` (Mac与Linux用户的后缀非exe)
- `test`
  - `testfile1.yaml`
  - `testfile2.yaml`
  - ...
- `1`(第几次作业)-`18373722`(学号)-`朱英豪`(姓名)
  - src
    - `Test.java` (程序运行的主入口)
    - 其余`*.java`
    - (有无`*.class`没有关系，我们会重新编译)
  - `judge.yaml`

**运行方式：**

在terminal(如cmd)中运行如下命令：

```bash
./patpat -judge=1-18373722-朱英豪
```

**评测结果：**

见生成的`result.md`。（其与该`patpat`程序在同级目录）

## `judge.yaml`的编写

```yaml
num: 1 # 第几次作业
id: 123456 # 学号
name: "hanhan" # 姓名
test: ["testfile1.yaml","testfile2.yaml"] # 可有更多，这是个列表
```

注：在自测时`testfile`的名称，需要与`test`目录下的`testfile`文件名相一致。

## 作业提交方式

- `1`(第几次作业)-`18373722`(学号)-`朱英豪`(姓名)
  - src
    - `Test.java` (程序运行的主入口)
    - 其余`*.java`
  - `judge.yaml`

将上述文件夹压缩为`zip`格式后，上传至云平台。

注：文件夹名称与`judge.yaml`请务必填写正确且保持一致，如不一致将无法评测。
