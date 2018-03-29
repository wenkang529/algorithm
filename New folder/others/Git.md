# Git
## commonly used commands
* git add ./
* git init
* git commit -m 'message'
* git push origin master
* git pull
* git clone 
* git rm
* git push git@github.com:wenkang529/**.git
* git config --global user.name "wkwu"
* git config --global user.email "wenkang529@163.com"
* git remote add origin git@github.com:wenkang529/mytf1.git  
add one remote address named  origin，the address is github's project address

## steps to use git
1. creat ssh keygen  `ssh-keygen -C 'wenkang529@163.com -t rsa' `
2. copy content to github of the pub file
3. test `ssh -T git@github.com`
4. git init
5. git add ./
6. git commit -m 'message'
7. git remote add origin git@github.com:wenkang529/**.git
8. git push origin master // or use command:git push git@github:wenkang529/**.git
9. 



* download git app `https://git-for-windows.github.io/`
* attention :if the wrong message is `permission denied`, maybe didn't clean the old file when creat the new file.
* git push -u origin master //after use paramete `-u`,can use `git pull` without parameter

## branch
* main branch is `master`,it's be auto created.
* creat new branch  `git branch aa master` creat a new branch named aa based on master
* switch to one branch `git checkout aa`
* delete branch aa `git branch -d aa`
* add tag `git tag tagname aa`
* merge branch `git checkout master` `git merge --no-ff aa`





```
push方法1：
　　现在如果想直接Push这个develop分支上的内容到github
　　git push -u origin
　　如果是新建分支第一次push，会提示：
　　fatal: The current branch develop has no upstream branch.
　　To push the current branch and set the remote as upstream, use
　　git push --set-upstream origin develop
　　输入这行命令，然后输入用户名和密码，就push成功了。
　　以后的push就只需要输入git push origin
　　
　　
push方法2：
　　比如新建了一个叫dev的分支，而github网站上还没有，可以直接：
　　git push -u origin dev
　　这样一个新分支就创建好了。
 
push方法3：
　　提交到github的分支有多个，提交时可以用这样的格式：
　　git push -u origin local:remote
　　
　　比如：git push -u origin master:master
　　表明将本地的master分支（冒号前）push到github的master分支（冒号后）。
　　如果左边不写为空，将会删除远程的右边分支。
　　
强制覆盖(把本地内容强制覆盖到远程，而不用先git pull）：
    git push --force origin_face master
```
## ignore file
* folder /data/
* file *.tmp
* 保守模式负责设置哪些文件不被过滤，就是那些药被追踪的文件  
! *.c  
! *.h  
!/mni/
* 配置.gitignore 的简易原则

* 采用共享模式与保守模式结合配置的办法。eg：一个文件夹下有很多文件夹和文件，而我只想跟踪其中的一个文件，这样设置就可以满足这种情况，先用共享模式把整个目录 都设置为不跟踪，然后再用保守模式把这个文件夹中想要跟踪的文件设置为被跟踪，配置很简单，就可以跟踪想要跟踪的文件。
先把 。gitignore 文件上传上去  在把忽略的文件夹放进来  就不会产生上传了  
[link](https://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013758404317281e54b6f5375640abbb11e67be4cd49e0000)
[branch](https://git-scm.com/book/zh/v1/Git-%E5%88%86%E6%94%AF-%E5%88%86%E6%94%AF%E7%9A%84%E6%96%B0%E5%BB%BA%E4%B8%8E%E5%90%88%E5%B9%B6)
[git](http://www.ruanyifeng.com/blog/2015/12/git-cheat-sheet.html)




## github  && gitlab
[how to exist together](http://xuyuan923.github.io/2014/11/04/github-gitlab-ssh/)
1. creat ssh file with diffrent name `ssh-keygen -C 'wenkang529@163.com -t rsa`  `/c/Users/wwk/.ssh/id_github``/c/Users/wwk/.ssh/id_gitlab`
2. add ssh key(public)to my account
3. creat one file named config
4. add content 
```
# gitlab
Host gitlab.com
	HostName gitlab.com
	PreferredAuthentications publickey
	IdentityFile ~/.ssh/id_gitlab
	
# github
Host github.com
	HostName github.com
	PreferredAuthentications publickey
	IdentityFile ~/.ssh/id_github
```
5.OK











