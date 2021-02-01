# patpat

Auto Judger for BUAA-OOP Course

## 如何使用

请**严格按照**如下的目录层级与命名准备好待测文件：

- `patpat.exe` (Mac与Linux用户的后缀非exe)
- `test`
  - `judge.yaml` (用于告知评测机需要测哪些测试用例)
  - `testcase1.yaml`
  - `testcase2.yaml`
  - ...
- `1`(第几次作业)-`18373722`(学号)-`朱英豪`(姓名)
  - `src`
    - `Test.java` (程序运行的主入口)
    - 其余`*.java`
    - (有无`*.class`没有关系，我们会重新编译)

**运行方式：**

在terminal(如cmd)中运行如下命令：

```bash
./patpat -judge 1-18373722-朱英豪
```

**评测结果：**

见生成的`$testcase$_result.md`。(其与该`patpat`程序在同级目录)

## `judge.yaml`的编写

```yaml
tests: [testcase1.yaml, testcase2.yaml] # 可有更多，这是个列表
# 请将本judge.yaml置于test文件夹内
```

注：在自测时`testcase`的名称，需要与`test`目录下的`testcase`文件名相一致。

## 作业提交方式

- `1`(第几次作业)-`18373722`(学号)-`朱英豪`(姓名)
  - src
    - `Test.java` (程序运行的主入口)
    - 其余`*.java`

将上述文件夹压缩为`zip`格式后，上传至云平台。

注：请务必反复确认文件的命名，如有误将无法评测。
