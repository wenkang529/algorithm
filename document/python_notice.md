### deep copy & shallow copy

a=[1,2,3]
b=a
b
[1, 2, 3]
[id(x) for x in a]
[1741312800, 1741312832, 1741312864]
[id(x) for x in b]
[1741312800, 1741312832, 1741312864]
a=[4,5,6]
[id(x) for x in a]
[1741312896, 1741312928, 1741312960]

- while i set a one new value b don't change for a get a new id

[copy](https://www.jianshu.com/p/efa9dd51f5cc)





- list empty   a=[]  but a is not None    a==[]





默认情况下，Python解释器会搜索当前目录、所有已安装的内置模块和第三方模块，搜索路径存放在`sys`模块的`path`变量中：

```
>>> import sys
>>> sys.path
['', '/Library/Frameworks/Python.framework/Versions/3.6/lib/python36.zip', '/Library/Frameworks/Python.framework/Versions/3.6/lib/python3.6', ..., '/Library/Frameworks/Python.framework/Versions/3.6/lib/python3.6/site-packages']

```

如果我们要添加自己的搜索目录，有两种方法：

一是直接修改`sys.path`，添加要搜索的目录：

```
>>> import sys
>>> sys.path.append('/Users/michael/my_py_scripts')

```

这种方法是在运行时修改，运行结束后失效。

第二种方法是设置环境变量`PYTHONPATH`，该环境变量的内容会被自动添加到模块搜索路径中。设置方式与设置Path环境变量类似。注意只需要添加你自己的搜索路径，Python自己本身的搜索路径不受影响。