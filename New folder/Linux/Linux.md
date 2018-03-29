# Linux
## commonly used commands
* cp  copy
* mv  move
* rm  delete
* tar  
``` tar
parameter： 
-c ：新建打包文件  
-t ：查看打包文件的内容含有哪些文件名  
-x ：解打包或解压缩的功能，可以搭配-C（大写）指定解压的目录，注意-c,-t,-x不能同时出现在同一条命令中  
-j ：通过bzip2的支持进行压缩/解压缩  
-z ：通过gzip的支持进行压缩/解压缩  
-v ：在压缩/解压缩过程中，将正在处理的文件名显示出来  
-f filename ：filename为要处理的文件  
-C dir ：指定压缩/解压缩的目录dir 
```
* cat check the contents of the file,it can be used with `more` `less`
* chmod change the file permissions
* [vim](http://www.runoob.com/linux/linux-vim.html)
```
vim
command mode
insert mode
last line mode

while command mode:
> i  chenge to insert mode 
> x  delete one character 
> :  chenge to last line mode,so can input some command
```
* set environment variables,`export PATH=.....` check if it's be done `export` we also can change `profile` file to add environment variables,add comment `export PATH=''......''.we also can change .bashrc file.but the two methods must restart system [link](http://www.cnblogs.com/amboyna/archive/2008/03/08/1096024.html)
* SSH [linux](http://blog.csdn.net/netwalk/article/details/12952051) [windows](http://fancyseeker.github.io/2013/12/31/ssh_connect/)
* set command run in backstage [link](https://segmentfault.com/a/1190000008314935) 1. add `&` behind the command 2.`ctrl+z` after run the command 3.if we want to find back the command `jobs` to see the number,`bg number`
* [shell](https://zm10.sm-tc.cn/?src=l4uLj8XQ0JzRnZaekZyXmpGY0ZGai9Ccj4%2FQjJeak5PQ&from=derive&depth=9&link_type=60&wap=false&bu=web&v=1&uid=b87b342b088545867359a9e8640c8c9f&restype=1)
* 
