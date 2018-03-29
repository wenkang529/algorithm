把MySQL安装成功之后就配置了环境：配置了附加包含目录，附加库目录，附加依赖项，还有复制了dll文件
刚开始忘记了附加依赖项 libmysql.dll  
最后发现一个关键的问题，安装mysql的时候安装的库只能是 x86 或者 x64 的，这个一定要注意，因为当我调试的时候是x86 与x64是不一样的。慎重
下面是代码
只是实现了连接的功能
```cpp
#include "stdafx.h"
#include<WinSock2.h>
#include<mysql.h>
#include<string>
using namespace std;


int main()
{
	MYSQL m_sqlCon;
	mysql_init(&m_sqlCon);
	// localhost:服务器 root/123456为账号密码 managesystemdb为数据库名 3306为端口
	if (!mysql_real_connect(&m_sqlCon, "localhost", "root", "0000", "my01", 3306, NULL, 0))
	{
		printf("访问数据库失败");
	    string e = mysql_error(&m_sqlCon);//需要将项目属性中字符集修改为“使用多字节字符集”或“未设置”
		
		return 0;
	}
	else
	{
		mysql_query(&m_sqlCon, "SET NAMES 'GB2312'");//解决从数据库中读取数据后汉字乱码显示的问题
		printf("fangwenchenggong");
		return 0;
	}
	return 0;
}
//下面是第二版代码
//增加了 查询语句以及返回值的输出

#include "stdafx.h"
#include<iostream>
#include<cstring>
#include<WinSock2.h>
#include<mysql.h>
#include<string>
using namespace std;


int main()
{
	MYSQL m_sqlCon;
	MYSQL_RES *res;
	MYSQL_ROW row;
	mysql_init(&m_sqlCon);
	// localhost:服务器 root/123456为账号密码 managesystemdb为数据库名 3306为端口
	if (!mysql_real_connect(&m_sqlCon, "localhost", "root", "0000", "my01", 3306, NULL, 0))
	{
		printf("访问数据库失败");
	    string e = mysql_error(&m_sqlCon);//需要将项目属性中字符集修改为“使用多字节字符集”或“未设置”		
	}
	else
	{
		mysql_query(&m_sqlCon, "SET NAMES 'GB2312'");//解决从数据库中读取数据后汉字乱码显示的问题
		cout << "访问成功"<<"\n"<<endl;
	}

	if (mysql_real_query(&m_sqlCon,"select * from wwk",(unsigned long)strlen("select * from user")))
	{
		cout << "mysql_real_query failure！" << endl;
	return 0;
	}
		// 存储结果集
	res = mysql_store_result(&m_sqlCon);
	if (NULL == res)
	{
		cout << "mysql_store_result failure！" << endl;
	return 0;
	}
		// 重复读取行，并输出第一个字段的值，直到row为NULL
		while (row = mysql_fetch_row(res))
		{
			cout << row[0] << endl;

		}
			// 释放结果集
	mysql_free_result(res);
	// 关闭Mysql连接
	mysql_close(&m_sqlCon);

	return 0;
}
```

