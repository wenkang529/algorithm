环境准备

Ubuntu 16.10

TensorFlow -cpu

安装准备

python3.5及以上

pip

安装步骤

更新软件源

apt-get update

python3.5系统自带 

安装pip

sudo apt-get install python3-pip

安装TensorFlow

sudo pip3 install tensorflow

tips:

下载过程过慢  可能是因为软件更新后没有重启

重启后解决问题

测试安装

显示正确则安装成功

安装numpy

sudo pip3 install numpy

安装opencv

sudo pip3 install opencv-python

安装scipy

pip3 install scipy

安装matplotlib

pip3 install matplotlib

安装PIL

pip3 install pillow

安装dlib 

安装的时候提示缺少cmke

安装cmake

下载 两种 zip  和 msi

zip解压之后直接是可执行二进制文件  应该下载cmake  可自动添加路径 目测文件都是一样的

添加路径  应该在系统环境变量中的 path  添加个路径即可

还是安装失败

提示可能缺少 boost   c++的一个开源库 

安装boost

pip3 install boost

但是安装dlib的时候还是失败

另一种安装方法

下载：http://www.boost.org/

我下载的是：boost_1_63_0

下载到C:\local目录下

1）首先：

![img](C:/Users/wenkangw/AppData/Local/YNote/data/wenkang529@163.com/72f30899a0bd4f5083abc22a6201b32f/516085332174.png)

双击bootstrap.bat

生成了

![img](C:/Users/wenkangw/AppData/Local/YNote/data/wenkang529@163.com/70acf97bb5ab42a68ef73267c4a36250/516085338239.png)

![img](http://blog.csdn.net/Insanity666/article/details/72235275)

2）再在命令中输入b2 install

3）利用b2编译库文件

 --     b2 -a --with-python address-model=64 toolset=msvc runtime-link=static

之前你cmake下载的64位这里写64，如果是32位的就把之前的64改成32

4）设置变量

​    --     set BOOST_ROOT=C:\local\boost_1_63_0

​    --     set BOOST_LIBRARYDIR=C:\local\boost_1_63_0\stage\lib

参考：http://blog.csdn.net/worrydog/article/details/53947214

安装完

之后安装 dlib依旧失败  还是提示boost

ubuntu下安装dlib  

先安装boost

sudo apt-get install libboost-all-dev

安装cmake

sudo apt-get install cmake

sudo pip3 install dlib

安装成功

安装sklearn

pip3 install sklearn

在新的机器上重新安装   发现pip3 install tensorflow 失败  看到可能的原因是我安装的python版本为3.6  

重新下载版本3.5.4的python  重新安装pip  python setup.py install 

pip3 install tensorflow 

成功

爬虫相关

pip3 install requests

pip3 install scrapy 安装失败

提示：没有安装vc++14.0

安装pandas

pip3 install pandas

安装seaborn

pip3 install seaborn