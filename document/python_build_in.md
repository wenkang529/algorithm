# python 内置函数

## filter



filter（）函数包括两个参数，分别是function和list。该函数根据function参数返回的结果是否为真来过滤list参数中的项，最后返回一个新列表，如下例所示：
>>>a=[1,2,3,4,5,6,7]
>>>b=filter(lambda x:x>5, a)
>>>print b
>>>[6,7]







## zip

zip会将两个数以元祖的方式捆绑在一起:

```python
list(zip([1,2,3,4,5],[6,7,8,9,0]))

#>>[(1,6),(2,7),(3,8),(4,9),(5,0)]

```

zip()是Python的一个内建函数，它接受一系列可迭代的对象作为参数，将对象中对应的元素打包成一个个tuple（元组），然后返回由这些tuples组成的list（列表）。若传入参数的长度不等，则返回list的长度和参数中长度最短的对象相同。利用*号操作符，可以将list unzip（解压）

```python
a=['aca','aa','ac','acb']
for i in zip(*a):
    print(i)
#>>('a', 'a', 'a', 'a')
   ('c', 'a', 'c', 'c')
```





## list comprehensions

```python
v=[x*x for x in range(10)]
#>>[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]

v=[x*x for x in range(10) if i%2==0]
#>>[0, 4, 16, 36, 64]
```

- a= x if m else y         if m==True  -->a=x else a=y



## labda

```python
a=lambda x:x*x+1                 # lambda parameter:function body (return)
```

- use lambda is not for more powerful but concise





## set

Python中的set函数是一个无序不重复的元素集。主要用于关系测试和去除重复项。

```python
a='aneejndd'
print(set(a))
#>> {'n', 'j', 'a', 'd', 'e'}
```

在Python3.3版本中，有两个内置set函数，分别是：set() 和 frozenset().两者的区别在于，set函数所获得的集合内容是可变动的，可以使用add(),update(),remove()函数增删集合内容。但也正是因为它是可变动的，所以不能用于字典中的key或者set集合里面的子集。frozenset函数所获得的集合内容已经生成便不能再变动，可以用做字典的key或者set集的子集。



### map/reduce

`map()`函数接收两个参数，一个是函数，一个是`Iterable`，`map`将传入的函数依次作用到序列的每个元素，并把结果作为新的`Iterator`返回。

```
>>> def f(x):
...     return x * x
...
>>> r = map(f, [1, 2, 3, 4, 5, 6, 7, 8, 9])
>>> list(r)
[1, 4, 9, 16, 25, 36, 49, 64, 81]
```



`reduce`把一个函数作用在一个序列`[x1, x2, x3, ...]`上，这个函数必须接收两个参数，`reduce`把结果继续和序列的下一个元素做累积计算，其效果就是：

```
reduce(f, [x1, x2, x3, x4]) = f(f(f(x1, x2), x3), x4)
```

```
>>> from functools import reduce
>>> def add(x, y):
...     return x + y
...
>>> reduce(add, [1, 3, 5, 7, 9])
25
```



### sorted

`sorted()`函数也是一个高阶函数，它还可以接收一个`key`函数来实现自定义的排序，例如按绝对值大小排序：

```
>>> sorted([36, 5, -12, 9, -21], key=abs)
[5, 9, -12, -21, 36]
```

key may be  `str.lower`  



### Decorator

由于函数也是一个对象，而且函数对象可以被赋值给变量，所以，通过变量也能调用该函数。

```
>>> def now():
...     print('2015-3-25')
...
>>> f = now
>>> f()
2015-3-25

```

函数对象有一个`__name__`属性，可以拿到函数的名字：

```
>>> now.__name__
'now'
>>> f.__name__
'now'

```

现在，假设我们要增强`now()`函数的功能，比如，在函数调用前后自动打印日志，但又不希望修改`now()`函数的定义，这种在代码运行期间动态增加功能的方式，称之为“装饰器”（Decorator）。

本质上，decorator就是一个返回函数的高阶函数。所以，我们要定义一个能打印日志的decorator，可以定义如下：

```python
def log(func):
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper
```

观察上面的`log`，因为它是一个decorator，所以接受一个函数作为参数，并返回一个函数。我们要借助Python的@语法，把decorator置于函数的定义处：

```python
@log
def now():
    print('2015-3-25')
```

调用`now()`函数，不仅会运行`now()`函数本身，还会在运行`now()`函数前打印一行日志：

```
>>> now()
call now():
2015-3-25
```

把`@log`放到`now()`函数的定义处，相当于执行了语句：

```
now = log(now)
```





因为我们讲了函数也是对象，它有`__name__`等属性，但你去看经过decorator装饰之后的函数，它们的`__name__`已经从原来的`'now'`变成了`'wrapper'`：

```
>>> now.__name__
'wrapper'

```

因为返回的那个`wrapper()`函数名字就是`'wrapper'`，所以，需要把原始函数的`__name__`等属性复制到`wrapper()`函数中，否则，有些依赖函数签名的代码执行就会出错。

不需要编写`wrapper.__name__ = func.__name__`这样的代码，Python内置的`functools.wraps`就是干这个事的，所以，一个完整的decorator的写法如下：

```python
import functools

def log(func):
    @functools.wraps(func)
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper
```

或者针对带参数的decorator：

```python
import functools

def log(text):
    def decorator(func):
        @functools.wraps(func)
        def wrapper(*args, **kw):
            print('%s %s():' % (text, func.__name__))
            return func(*args, **kw)
        return wrapper
    return decorator
```

### Partial function 偏函数

`functools.partial`的作用就是，把一个函数的某些参数给固定住（也就是设置默认值），返回一个新的函数，调用这个新函数会更简单。

当函数的参数个数太多，需要简化时，使用`functools.partial`可以创建一个新的函数，这个新函数可以固定住原函数的部分参数，从而在调用时更简单。

```python
>>> import functools
>>> int2 = functools.partial(int, base=2)
>>> int2('1000000')
64
>>> int2('1010101')
85
```

### Generator

要创建一个generator，有很多种方法。

`第一种方法：`只要把一个列表生成式的`[]`改成`()`，就创建了一个generator：

```
>>> L = [x * x for x in range(10)]
>>> L
[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
>>> g = (x * x for x in range(10))
>>> g
<generator object <genexpr> at 0x1022ef630>

```

创建`L`和`g`的区别仅在于最外层的`[]`和`()`，`L`是一个list，而`g`是一个generator。

如果要一个一个打印出来，可以通过`next()`函数获得generator的下一个返回值.或者使用for循环，因为generator也是一个可迭代对象

所以，我们创建了一个generator后，基本上永远不会调用`next()`，而是通过`for`循环来迭代它，并且不需要关心`StopIteration`的错误。

`第二种方法：`定义generator的另一种方法。如果一个函数定义中包含`yield`关键字，那么这个函数就不再是一个普通函数，而是一个generator：

```python
def fib(max):
    n, a, b = 0, 0, 1
    while n < max:
        yield b
        a, b = b, a + b
        n = n + 1
    return 'done'
```

### Iterator

`iterable`:可迭代对象

`iterator`:迭代器

可以被`next()`函数调用并不断返回下一个值的对象称为迭代器：`Iterator`

可以直接作用于`for`循环的对象统称为可迭代对象：`Iterable`



可以直接作用于`for`循环的数据类型有以下几种：

一类是集合数据类型，如`list`、`tuple`、`dict`、`set`、`str`等；

一类是`generator`，包括生成器和带`yield`的generator function。



可以使用`isinstance()`判断一个对象是否是`Iterable`对象

```python
>>> from collections import Iterator
>>> isinstance((x for x in range(10)), Iterator)
True
```

生成器都是`Iterator`对象，但`list`、`dict`、`str`虽然是`Iterable`，却不是`Iterator`。

把`list`、`dict`、`str`等`Iterable`变成`Iterator`可以使用`iter()`函数：

## map

```
map()函数
map()是 Python 内置的高阶函数，它接收一个函数 f 和一个 list，并通过把函数 f 依次作用在 list 的每个元素上，得到一个新的 list 并返回。
```

a="123"

print(map(int,a))

输出:

[1, 2, 3]











































