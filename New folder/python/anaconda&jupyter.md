# Anaconda
## 介绍
* anaconda 包含了环境管理 conda
* 对于多个环境的管理，本地的python和anaconda的python区别在环境变量，注意环境变量path的配置。

## conda环境管理
```
# 创建一个名为python34的环境，指定Python版本是3.4（不用管是3.4.x，conda会为我们自动寻找3.4.x中的最新版本）
conda create --name python34 python=3.4

# 安装好后，使用activate激活某个环境
activate python34 # for Windows
source activate python34 # for Linux & Mac
# 激活后，会发现terminal输入的地方多了python34的字样，实际上，此时系统做的事情就是把默认2.7环境从PATH中去除，再把3.4对应的命令加入PATH

# 如果想返回默认的python 2.7环境，运行
deactivate python34 # for Windows
source deactivate python34 # for Linux & Mac

# 删除一个已有的环境
conda remove --name python34 --all
查看又都什么环境

conda info --envs

直接在命令行中操作  相应的pip也随之改变
互不影响
想用哪个环境就直接启用，不想用就直接退出。
```
* 可以用activate + path（本地python路径） 就可以启动本地的python，对应的pip也可以使用

* 切换到anaconda  可以直接activate root 或者其他的名字，或者deactivate

## conda的使用
> conda 包管理
* 安装scipy  
  conda install scipy  
  conda会从从远程搜索scipy的相关信息和依赖项目，对于python 3.4，conda会同时安装numpy和mkl（运算加速的库）

* 查看已经安装的packages  
  conda list  (后面可以跟很多参数，可以列举都有什么环境)
  最新版的conda是从site-packages文件夹中搜索已经安装的包，不依赖于pip，因此可以显示出通过各种方式安装的包

> 
* 添加Anaconda的TUNA镜像  
  conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/  
  TUNA的help中镜像地址加有引号，需要去掉
* 设置搜索时显示通道地址  
  conda config --set show_channel_urls yes
* 在jupyter中添加本地的python环境  
  1.首先切换到本地python环境，还有对应的pip  
  2.pip3 install ipykernel  
  3.然后执行  
   python -m ipykernel install  
  可以添加命令 –-name kernelname 为kernel指定名字
* 使用命令jupyter kernelspec list可以查看当前的kernel  
  删除notebook kernel使用命令  
  jupyter kernelspec remove kernelname

