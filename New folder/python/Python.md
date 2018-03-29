# Python

## pythonic
```python
#文件打开
with open('data.txt') as f:
    data = f.read()

#列表推导式
result = []
for i in range(10):
    s = i  2
    result.append(s)
    #pythonic
[i2 for i in xrange(10)]

#遍历
for i,j in enumerate(a):
    print(i,'-->',j)
#使用enumerate可以减少内存占用，遍历大文件

#列表
#列表对象（list）是一个查询效率高于更新操作的数据结构，比如删除一个元素和插入一个元素时执行效率就非常低，因为还要对剩下的元素进行移动

```
## Python 要避免的错误
```python
# 可变数据类型作为函数定义中的默认参数

# 可变数据类型作为类变量
class URLCatcher(object):
    urls = []
    def add_url(self, url):
        self.urls.append(url)
```
## python注意事项
### python 模块
* 要想让python解释器找到自己编写的模块，则该模块必须PYTHONPATH上，否则在导入该模块时会出现找不到该模块的错误  
在Linux中  
export PYTHONPATH=$PYTHONPATH:/home/lxc/software/program/python
*import
* 当一个文件  import 另一个文件的时候 会把文件中的程序执行一遍，如果有类的定义，会运行类的初始化？

### argparse模块
```python
import argparse
import sys
def main(arg):
    print(arg.intt)
    if arg.debug:
        print('debug')

def parse_arguments(argv):
    parse=argparse.ArgumentParser()
    parse.add_argument('intt',type=int)
    parse.add_argument('--debug',action='store_true')
    return parse.parse_args(argv)

if __name__ == '__main__':
    main(parse_arguments(sys.argv[1:]))  #第一个参数为默认 运行的脚本名称对应argv[0]
```
调用的时候：    
    python test.py 3  
或者 添加可选参数   
    python test.py 3 --debug   
    
因为action设置为了store_true 所以后面不用加参数 当有debug的时候 debug为true参数动作argparse内置6种动作可以在解析到一个参数时进行触发：  
* store 保存参数值，可能会先将参数值转换成另一个数据类型。若没有显式指定动作，则默认为该动作。  
* store_const 保存一个被定义为参数规格一部分的值，而不是一个来自参数解析而来的值。这通常用于实现非布尔值的命令行标记。  
* store_ture/store_false 保存相应的布尔值。这两个动作被用于实现布尔开关。
* append 将值保存到一个列表中。若参数重复出现，则保存多个值。
* append_const 将一个定义在参数规格中的值保存到一个列表中。
* version 打印关于程序的版本信息，然后退出  
[link](http://wiki.jikexueyuan.com/project/explore-python/Standard-Modules/argparse.html)

## python graphical interface(GUI) programming
[link](https://zhuanlan.zhihu.com/p/22619896)

### program entry
* each programing language must have an main function ,python too. when python source code import as a module ,the top layer code will be run.(distinguish by indented （缩进）),but some times we do not want to run our code ,so we use `if __name__="__main__"` so it's code won't be run. [name](http://blog.konghy.cn/2017/04/24/python-entry-program/) 

### virtualenv
* install virtualenv `pip install virtualenv`
* creat virtual environment `virtualenv -p path name`
* active virtual environment `source py3env/bin/activate`
* exit virtual environment `deactivate`

### binnary struct (c programe language)
some times we need to use python to deal with binary data ,so we can use module named `strct`  three functions 
* `pack` transfer data to string according to the format
* `unpack` transfer string to tuple according to the format
* `calcsize` cumpute how much Bytes  used
[link](http://www.cnblogs.com/gala/archive/2011/09/22/2184801.html)

### time
```
import time
time.ctime()
out:Thu Jan 25 21:08:48 2018

```
### plt.imshow()
```
plt.show()
plt.savefig()
```
save pic to file but blank, because when `plt.show()` it has creat a new pic(blank).
solution
```
#some code for  draw a pic
plt.savefig()
plt.show()
```
## numpy
### numpy.random
```
#creat one series of random number
np.random.random([3,2])
```
initial one array
```python
a=np.array([1,3,5])
a=np.arange([9]).reshape([3,3) # [[0,1,2],[3,4,5],[6,7,8]]
a=np.arrange(10,30,5)   #[10,15,20,25]
a=np.zeros((3,4))  #all zeros
a=np.ones((2,3))
```


## python parameter
beside normal parameter transfer ,also can use packing(包裹传递)  
```python
#包裹位置
def func(* name):
    print(name)

func((1,2))

#包裹关键字
def func(** dict):
    print(dict)

func('a':1,'b':2)
```
\* and ** also can use in unpacking
```
def func(a,b,c):
　　print a,b,c
 

args = (1,3,4)
func(*args)

dict = {'a':1,'b':2,'c':3}
func(**dict)
```
[link](http://www.cnblogs.com/rourou1/p/6182502.html)

## os.walk
[link](http://blog.csdn.net/johinieli/article/details/76660733)  
```
def dir_walk(path):
    for dir_path, dir_name, file_name in os.walk(path):
        for file in file_name:
            if '.jpg' in file or '.JPG' in file :
                file_path = os.path.join(dir_path, file)  # 将文件加上路径


```
### yield
[link](http://pyzh.readthedocs.io/en/latest/the-python-yield-keyword-explained.html)
```python
def func():
    for i in data:
    b=i*i
    yield b

m=func()
for i in m:
    print(i)
    
```
#### iterator 
* when you build one list ,you can read it one by one,so the list can be called iterator object ,any can be use ``for ...in...`` is a iterator . but it save all the data in the memory,if have much  data it will be err.so we use builder

* return a builder (生成器) ，it can be itrator but ``can be read only once``it doesn't save all the data to memory ,it creat data real time.
* it can reduce the memory cost
* yield just like return
* it doesn't be created immediately,ony when you use for it will be created
* when it work ,the first time it will run from bigin to the keyword `yield` and return the value behind the yield.
then it will run the code in for loop which you define until nothing to return.

### convert PIL and Opencv
[link](https://www.yuanmas.com/info/8kaGXlYQym.html)
* frame=cv2.cvtColor(frame,cv2.COLOR_RGB2BGR)

### deal with string
[linl](https://www.cnblogs.com/huangcong/archive/2011/08/29/2158268.html)

### some prolem about CV2 of chanise path
but this method cant adjust the wrong angle  
[link](https://www.zhihu.com/question/47184512)  
img = cv2.imdecode(np.fromfile(path_file,dtype=np.uint8),-1)
// chage bgr 2 rgb
pic=cv2.cvtColor(img,cv2.COLOR_BGR2RGB)

### code
check which code for system  
``` 
import sys
print(sys.getdefaultencoding())
```
python3 have diffrent with python2

## model
### 按字节码编译的 .pyc 文件

###  `__name__` 确定不同的运行方式（独立运行/被导入） 
```py
if __name__='__main__':
    print('this program is being run by itself')
else:
    print('it is being imported from another model')
#如果直接运行则执行第一个print 如果被导入则执行第二个print
```
* dir(model name) return model name list include function ,class name, variable

## packages
packages: include one nodel and `__init__.py` file

## class
* `__init__` method,can recieve paraments
```py
class person:
    def __init__(self,name):
        self.name=name
    def say_hi(self):
        print('hello,my name is',self.name)
p=person('andy')
p.say_hi()
# output: hello,my name is andy
```
* 类变量是共享的
* 

## Exception

```
try..

except ..:

except ..:

else:
    当不发生异常的时候执行
```
* 也可以自己抛出异常 raise ...
```
try ..

excrpt ..:

finally:
    无论如何都会执行
```
## dict

* {key:value}  key must be diffrent,use key to find value

```
dict={}
dict['a']=1
# if we write dict[a]=2 again,we will rewrite the value of key a to 2

```
## `is` and `==`
any object in python consist with three basic elments:`id`  `type`  `value`  
`is` is compared by `id`  
`==` compared by `value`  
example:  
x=[1,2,3]
y=[1,2,3]

z=x

'>>>  x is z
Ture

'>>> x is y
False

`>>> x==y
Ture

'>>> x==z
Ture

# time
> three methods to compute running time of the code
1. datatime get the date
```py
import datetime

start=datetime.datetime.now()
# some  code
end=datetime.datetime.now()

print(end-start)
```
2. time.time()  return the soconds from the (纪元)，float type as it may be fraction of second
```py
import time

start=time.time()
#some code
end=time.time()
print(end-start)
```

3.time.clock  
```py
import time
start=time.clock()
# some code
end=time.clock()
print(end-start)
```
> time.clock()返回程序开始或第一次被调用clock（）以来的CPU时间。 这具有与系统记录一样多的精度。返回的也是一个浮点类型。这里获得的是CPU的执行时间。   
注：程序执行时间=cpu时间 + io时间 + 休眠或者等待时间