re

# 通用字符


pat0='any character' #any character
pat1='\w'  #one of number letter  underline
pat2='\d'  #one number
pat3='\s'  #one space

[] [^]   原子列表

## 元字符

### 任意匹配字符元

.  除了换行符以外的字符

### 边界限制符
^  开始位置
$  结束位置

### 限定符
* 任意次数 0,1，...
? 0 ,1 次
+ 1次 或者多次
{n} n次
{n,} >n 次
{n,m} n-m次

### 模式选择符
| 模式选择  其实就是或
() 模式单元符

### 模式修正符  
I 忽略大小写  
M 多行匹配  
L 做本地化识别匹配  
U 根据Unicode字符及解析字符  
S 让 . 包含换行符，用了这个修正符就可以让 . 匹配任意字符了

### 贪婪模式 懒惰模式    pabdyaaaaaay  
贪婪模式：尽可能的多匹配 p.*y  --> pabdy  
懒惰模式：尽可能的少匹配 p.*?y --> pabdyaaaaaay  

### 常见函数
re.match()  从源字符起始位置匹配（即刚开始的时候必须能匹配）如果无法匹配返回None  
re.search()  全文匹配并返回第一个匹配到的
re.sub(patten,replace,string,extern)  替换 （模式，替换的字符，源字符，附加选项（替换的次数等））


prtt1='(aa){2,}'

str1='12-aaaabc-34'

re1=re.search(prtt1,str1,re.M)

print(re1)


### 匹配 .com/.cn 网址

patt='[a-zA-Z]+://[^\s]*[.com|.cn]'

### 匹配电话号码 3-8 或者4-7

patt='\d{3}-\d{4}|\d{3}-\d{8}'

### get the result
res=re.serach(patt,str,re.I).group()

* 返回匹配的内容：re.search(patt,str).group()
* 返回匹配的位置：re.search(patt,str).span()  例子：（0，3）
* 

