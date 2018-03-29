# vc++ 学习记录
[c++](http://www.cplusplus.com/)
## 环境配置
### .h .lib .dll
* .h  在c/c++ 常规  附加包含目录     编译的时候用到
* .lib 链接的时候用到 链接器-> 常规->附加库目录  输入里面有附加依赖项
* .dll 运行的时候需要用到    在vc\++下面也可以设置 附加包含目录和库目录   但是建议在 c/c++ 和link那里设置
* 附加依赖项的是lib  例如在.h中定义了一个函数，连接器怎么知道调用哪个dll呢    .lib告诉链接器，所以也称为附加依赖项。
如果生成静态库文件，则没有dll   直接在lib文件中包含函数的执行代码。.h是编译的时候用到，lib库文件是链接用到，dll运行时用到。
* [详细讲解](http://blog.csdn.net/yusiguyuan/article/details/12649737)
* [配置](http://blog.sciencenet.cn/blog-1813407-834830.html)
* 添加项目属性表可以避免每次新建项目就要重新配置
 
> 解压OpenCV
配置系统环境变量(添加环境变量的时候注意项目是win32 还是64 添加对应的文件夹）vc10, vc11, vc12 分别表示VS2010, VS2012, VS2013的Visual Studio使用的编译器版本，根据自己的VS版本来填写正确的编译器版本号。  
创建VS新项目  
添加项目属性表  
配置项目属性表  
引用项目属性表  


## 编程
### cin&cout
如果需要使用cout  cin  需要下面的  
cin.getline()  可以读取一行 而不是像cin遇到空格 制表符或者回车就立即存储而产生一行无法完整读取  cin.get()会将回车包含在输入队列中，下次再调用第一个字符就是回车， cin.getline() 不包含回车。cin.get() 不带参数的可以直接读取一个字符，不管是不是回车
```
#include<iostream>
using namespace std;
cout<<"Please input an int number:"<<endl;
cin>>x;
//在C++中，setw(int n)用来控制输出间隔。例如:
cout<<'s'<<setw(8)<<'a'<<endl;
则在屏幕显示
s        a
//s与a之间有7个空格，setw()只对其后面紧跟的输出产生作用，如上例中，表示'a'共占8个位置，不足的用空格填充。
//若输入的内容超过setw()设置的长度，则按实际长度输出。setw()默认填充的内容为空格，可以setfill()配合使用设置其他字符填充。
cout<<setfill('*')<<setw(5)<<'a'<<endl;
****a
//输出：4个*和字符a共占5个位置。
```
### 指针&引用
1. 指针是一个实体，而引用仅是个别名；
2. 引用使用时无需解引用(*)，指针需要解引用；
3. 引用只能在定义时被初始化一次，之后不可变；指针可变；  
[detail](http://blog.csdn.net/wenkang529/article/details/79106841)  
[函数指针](http://blog.51cto.com/hipercomer/792300)  
[指针，智能指针](http://www.cnblogs.com/QG-whz/p/4777312.html)
* 指针   
    int  m=10;  
    int *p=&m;  
将m的地址给p  而不是给*p
* c++在创建指针时会创建用来存储指针的空间，但不会创建指针所指向的空间。  
    int *p=new int;  
    delete p;  
使用delete将释放p指向的内存，但不会删除p指针。
* 当用new声明动态数组的时候  
int *p =new int[10];  
.......  
delete [] p;

### 虚函数和纯虚函数
定义一个函数为虚函数，不代表函数为不被实现的函数。  
定义他为虚函数是为了允许用基类的指针来调用子类的这个函数。  
定义一个函数为纯虚函数，才代表函数没有被实现。  
定义纯虚函数是为了实现一个接口，起到一个规范的作用，规范继承这个类的程序员必须实现这个函数。
[detial](http://blog.csdn.net/hackbuteer1/article/details/7558868)

### 字符集
* 因为C++支持两种字符串，即常规的ANSI编码（使用""包裹）和Unicode编码（使用L""包裹）  
* 微软将这两套字符集及其操作进行了统一，通过条件编译（通过_UNICODE和UNICODE宏）控制实际使用的字符集，这样就有了_T("")这样的字符串，对应的就有了_tcslen这样的函数
为了存储这样的通用字符，就有了TCHAR：  
当没有定义_UNICODE宏时，TCHAR = char，_tcslen =strlen
当定义了_UNICODE宏时，TCHAR = wchar_t ， _tcslen = wcslen  
* 当进行国际编程或者使用Unicode编码的时候，字符会有  char  wchar_t 两种，当我们定义了UNICODE宏，就相当于告诉了编译器：我准备采用UNICODE版本。这个时候，TCHAR就会摇身一变，变成了wchar_t。而未定义UNICODE宏时，TCHAR摇身一变，变成了unsignedchar。这样就可以很好的切换宽窄字符集。  
* tchar可用于双字节字符串，使程序可以用于中日韩等国 语言文字处理、显示。使编程方法简化。[多字节编码与Unicode码](http://blog.csdn.net/luoweifu/article/details/49382969)  
[ASCII，Unicode 和 UTF-8](http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)

## 小技巧
* #define编译命令是c留下来的  c++中更多地用const 限定为常量 常见的做法是将常量大写，这样就知道是常量
* 对于float  c++只能保证6位有效位  double 同样有有效位，不过比float多，有效位是从左向右数 比如653000的有效位是3位  这和有没有小数点没有关系。
* 强制类型转换不会修改变量本身而是产生一个新的值
* 声明数组的时候 数组个数必须是常量不能是变量，因为编译的时候系统需要知道开辟多少空间
* 用string类  而不是用char数组，使用string类必须包含string头文件，因为包含在标准命名空间中，所以需要包含using namespace std
* 方位成员变量的时候，用 . 还是 -> ,如果标识符是结构名，则用 . 如果是指针则用-> 
* 