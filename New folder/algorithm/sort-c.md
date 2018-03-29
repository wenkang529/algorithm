```cpp
#include "stdafx.h"
#include<stdio.h>
#include<stdlib.h>
#include<time.h>
#include<iostream>
#include<vector>
#include<math.h>

using namespace std;

class sort
{
public:
	sort();
	sort(int len1);
	~sort();
	const int len;
	int begin = 0, end = 100;
	vector<int> arr;
	int *arr1 = new int[len];
	int * creat_arr1();
	vector<int> create_arr();
	int * bubble_sort(int  *original,int n);
	int * quick_sort(int * original, int begin, int end);
	int * select_sort(int  *original,int n);
	int *heap_sort(int  *original,int n);
	int *adjust_heap(int  *original, int p, int len);
	int *insert_sort(int  *original,int n);
	int *shell_sort(int  *original,int n);
    int *merge_sort(int *original, int n);
	int *merge(int * a, int * b, int *original, int lena, int lenb);
	int * radix_sort(int *original, int n);
	int * buckt_sort(int *original, int n);
	int * counting_sort(int *original, int n);


private:

};
sort::sort() :len(10)
{

}
sort::sort(int len1) : len(len1)
{

}
sort::~sort()
{
}
int * sort::creat_arr1()
{
	for (int i = 0; i < len; i++)
	{
		arr1[i] = (rand() % (end - begin + 1) + begin);

	}
	return arr1;
}
/*
生成随机数组
*/
vector<int> sort::create_arr()
{
	for (int i = 0; i < len; i++)
	{
		arr.push_back(0);
		arr[i] = (rand() % (end - begin + 1) + begin);

	}
	return arr;
}
/*
冒泡排序
时间复杂度：
若初始文件是反序的，需要进行n-1趟排序。每趟排序要进行n-i次关键字的比较(1≤i≤n-1)，且每次比较都必须移动记录三次来达到交换记录位置。在这种情况下，比较和移动次数均达到最大值：
Cmax=n(n-1)/2=O(n2)
Mmax=3n(n-1)/2=O(n2)
稳定性：
就地排序，稳定的。
*/
int  *sort::bubble_sort(int  *original ,int n)  //用指针是为了避免 变量生存期(这个还是很重要的）
{

	int flag = 0; //立个flag 当检测到不需要再继续排序的时候直接break
	for (int i = 0; i < (n - 1); i++)     //共需要损坏n-1次
	{
		for (int j = n - 1; j > i; j--)  //每次循环时从后往前比较 比较到已经排序的地方
		{
			if (original[j] < original[j - 1])
			{

				int m = original[j];
				original[j] = original[j - 1];
				original[j - 1] = m;
				flag = 1;
			}
		}
		if (flag == 0)
			break;
		else
			flag = 1;
		//循环打印输出
		cout << "\n\n第" << i << "次排序" << endl;
		for (int j = 0; j < 20; j++)
		{
			if (j % 5 == 0 & j != 0)
				cout << "\n";
			cout << original[j] << "   ";
		}

	}
	return original;
}
/*
快速排序
时间复杂度 一般为O(n lg(n))  最坏情况为O(n2)

递归调用
1．先从数列中取出一个数作为基准数。(我取的是第一个数)
2．分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
3．再对左右区间重复第二步，直到各区间只有一个数。

以一个数组作为示例，取区间第一个数为基准数。
72 6 57 88 60 42 83 73 48 85

初始时，i = 0;  j = 9;   X = a[i] = 72
由于已经将a[0]中的数保存到X中，可以理解成在数组a[0]上挖了个坑，可以将其它数据填充到这来。
从j开始向前找一个比X小或等于X的数。当j=8，符合条件，将a[8]挖出再填到上一个坑a[0]中。a[0]=a[8]; i++;
这样一个坑a[0]就被搞定了，但又形成了一个新坑a[8]，这怎么办了？简单，再找数字来填a[8]这个坑。
这次从i开始向后找一个大于X的数，当i=3，符合条件，将a[3]挖出再填到上一个坑中a[8]=a[3]; j--;

就是先确定左边第一个数为基准 复制个x 这样左边有一个坑
从右边找小的 填进去  i++
再从左边找大的填到右边的坑  j--
.......
知道i>=j的时候  把x给 i 就完成了一次分割
再把左走进行分割。。。。。。
*/
int * sort::quick_sort(int * original, int begin, int end)  //begin为0 end为最后一个数的下标  即n-1
{
	int i = begin, j = end;
	if (i < j)
	{
		int x = original[i];   //begin from left
		while (i < j)
		{
			while (i < j && original[j] >= x)
				j--;
			if (i < j)
			{
				original[i] = original[j];   //填左边的坑
				i++;                               //i++是因为i已经比较过了 从下一个比较就可以了 j--一样的道理
			}
			while (i < j && original[i] <= x)
				i++;
			if (i < j)
			{
				original[j] = original[i];
				j--;
			}
		}
		original[i] = x;
		quick_sort(original, begin, i - 1);
		quick_sort(original, i + 1, end);
	}


	return original;
}
/*
选择排序
时间复杂度：O(n2)
稳定排序

在数组array[begin:end]中先选定一个数array[begin]，
再在余下的数中遍历，如果找到比array[begin]小的数，那么就将这个数作为选定的数，
再往下遍历，如果还有更小的，那么继续换掉下标，直到找到最小的为止，
然后将这个最小和数组第一个数array[0]交换，
再在余下的数中找到最小的放在第二个位置array[1]，  。。。

就是两个for循环 外层的是完成一次最小数的查找，内层是查找最小值。

*/
int * sort::select_sort(int * original,int n)
{
	for (int i = 0; i <n; i++)
	{
		int index = i;
		for (int j = i + 1; j < n; j++)   //注意j=i+1 
		{
			if (original[j] < original[index])   //不用j+1 可以避免下标越界的危险
				index = j;
		}
		if (index != i)    //判断是否需要交换
		{
			int temp;
			temp = original[i];
			original[i] = original[index];
			original[index] = temp;
		}
	}

	return original;
}
/*
堆排序
时间复杂度 O(nlog2(n)
T(n) <=  O(n)  + (n - 1)*O(log2(n))  第一次构造最大堆 加上后面 n-1次构造堆
就地排序
mytip：
堆排序 利用了完全二叉树，大顶堆
首先构造一个 大顶堆
每次交换首尾值
从第一个值调整堆  这样每次可以得到一个最大值。 len-1
循环执行n-1次就可以完全得到从小到大的排序

假设父节点为p  那么子节点为2*p+1,2*p+2
计算最后一个父节点=len/2-1      len是长度，即最后一个下标的值+1。
*/
int * sort::heap_sort(int * original,int n)
{
	int len = n;
	for (int p = len / 2 - 1; p >= 0; p--)       //初始化堆（初始化为最大堆）
		adjust_heap(original, p, len);
	for (len = len - 1; len > 0; len--)
	{
		int t = original[len];                     //首尾节点互换，这样就可以重新从第一个节点往下放
													  //到合适的位置而选出来一个 最大值
		original[len] = original[0];
		original[0] = t;
		adjust_heap(original, 0, len);
	}

	return original;
}
int * sort::adjust_heap(int * original, int p, int len)    //p当前父节点，即要调整的值  len是长度
{
	int currparent = original[p];    //先把父节点的值复制出来，用于比较，避免被覆盖。
	int child = 2 * p + 1;            //子节点的值得计算  左节点 
	while (child<len)                 //for 循环迭代，直到结尾处
	{
		if (child + 1 < len && original[child] < original[child + 1])   //先比较左右节点大小
			child++;
		if (currparent < original[child])                    //比较父节点与较大子节点的大小 必须用currparent比较 不能用original[p]比较
		{
			original[p] = original[child];
			p = child;                                      //用于下次迭代的赋值
			child = 2 * p + 1;
		}
		else
			break;  //这个不能省去，因为不用一直比较到末尾，什么时候父节点比子节点大的时候就结束
	}
	original[p] = currparent;               //比较结束后 把原始父节点的值赋值给合适的位置
	return 	original;
}
/*
插入排序原理：它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
插入排序核心：假设第一个元素排好，之后的元素对排好的部分从后向前比较并逐一移动。

*/
int * sort::insert_sort(int * original,int n)
{
	int j, temp;  //这里要定义j 不能再循环里定义，因为在循环外需要用到
	for (int i = 0; i < n; i++)  //遍历每一个数
	{
		temp = original[i];
		for (j = i - 1; j >= 0 && original[j] > temp; j--)//将序列所有数向后移动一位
			original[j + 1] = original[j];
		original[j + 1] = temp;   //这个是j+1 不是j
	}
	return original;
}
/*
希尔排序
插入排序的升级版本
希尔排序是把记录按下标的一定增量分组，
对每组使用直接插入排序算法排序；随着增量逐渐减少，
每组包含的关键词越来越多，当增量减至1时，整个文件恰被分成一组，算法便终止。
*/
int * sort::shell_sort(int * original,int n)
{
	int gap = n / 2;
	int k;
	for (gap = gap; gap>0; gap = gap / 2)   //对步长进行缩减
	{
		for (int i = 0; i < gap; i++) //一个步长内的数的遍历
		{
			for (int j = i + gap; j < n; j += gap) //相隔相同步长，即相同组的遍历
			{   //下面的就和插入排序一致了  不过间隔不是1 而是gap
				int temp = original[j];
				for (k = j - gap; k >= 0 && original[k] > temp; k -= gap)
				{
					original[k + gap] = original[k];
				}
				original[k + gap] = temp;
			}
		}
	}

	return original;
}
/*
归并排序
*/
int * sort::merge_sort(int * original, int n)
{
	if (n > 1)
	{
		int d = n / 2;
		int d1 = n - d;

		int *a = new int[d];     //因为c++中不能动态定义数组  所以采用这种方法
		int *b = new int[d1];

		for (int i = 0; i < d; i++)
			a[i] = original[i];
		for (int i = 0; i < n - d; i++)
			b[i] = original[i + d];


		merge_sort(a, d);
		merge_sort(b, d1);

		merge(a, b, original, d, d1);
		delete[] a;
		delete[] b;

	}
	

	return original;
}
int * sort::merge(int * a, int * b, int *original, int lena, int lenb)
{
	int lenc = lena + lenb;
	int i = 0, j = 0, count = 0;
	int * c = original;    
	while (count < lenc)
	{
		if (a[i] < b[j])
		{
			c[count] = a[i];
			count++;
			i++;
		}
		else
		{
			c[count] = b[j];
			j++;
			count++;
		}
		if (i == lena)
		{
			while (j < lenb)
			{
				c[count] = b[j];
				j++;
				count++;
			}
		}
		if (j == lenb)
		{
			while (i < lena)
			{
				c[count] = a[i];
				i++;
				count++;
			}
		}

	}
	return c;
}
/*
基数排序
先按个位数排序，再排十位数，再排。。。
其中每次排序用  桶排序  
桶排序  有个计算应该放置的位置的  count[i]+=count[i-1]
然后将原有序列从后往前 按照序号放置，output[count[p] - 1] = a[i]; 
p为a[i]当前排序的位（个位，十位，百位，，，）   并count--
每次排序完成记得将output复制给原有序列，下次的排序在此基础上完成
*/
int * sort::radix_sort(int * original, int n)
{
	int *a = original;
	int *output = new int[n];  //输出序列
	int count[10];  //用于桶排序的计数
	int d = 1;
	//计算最高位的个数 
	for (int i = 0; i < n; i++)
	{
		int c=1;
		int pp = a[i];  //不要修改原始数据
		while (pp / 10)
		{
			pp /= 10;
			c++;
		}
		if (c > d)
			d = c;   //最多有d位
	}
	for (int r = 1; r <= d; r++)  //循环从低位到最高位排序 
	{
		for (int j = 0; j < 10; j++)  //每次将count清零
		{
			count[j] = 0;
		}

		for (int k = 0; k < n; k++)  //记录每个桶中放置的个数
		{
			int ri = pow(10, r - 1);
			int p = (a[k] /ri) % 10;
			count[p]++;
		}

		for (int m = 1; m < 10; m++)   //计算位置
		{
			count[m] += count[m - 1];
		}
		for (int l = n - 1; l >= 0; l--)  //按照计算的位置  输出排序
		{
			int ri1 = pow(10, r - 1);  //为了计算某个位上的值
			int p1 = (a[l] / ri1) % 10;

			output[count[p1] - 1] = a[l];   //关键
			count[p1]--;   //不要忘记计数减一
		}
	
		//每次排序完将输出的数据给原始数据，下次在原始数据的基础上排序
		for (int c = 0; c < n; c++)
		{
			original[c] = output[c];
		}
	}
	return original;
}
/*
桶排序
桶排序的思想就是把区间[0, 1)划分成n个相同大小的子区间，
每一个区间称为桶（bucket）。然后，将n个输入数据分布到各个桶中去。
因为输入数均匀且独立均匀分布在[0, 1)上，
所以一般不会有很多数落在一个桶中的情况。
为得到结果，先对各个桶中的数进行排序，
然后按次序把各个桶中的元素列出来即可。
*/
int * sort::buckt_sort(int * original, int n)
{
	struct list   //定义一个链表结构体
	{
		list *next;
		int value;
		list(int v,list*n):value(v),next(n){}
		list() :value(NULL), next(NULL) {}  //定义构造函数
	};
	int *data = original;  
	list bucket[10], *p,*q;    //要存到10个桶中，范围为1-100  
	int i, j, k, m;

	for (i = 0; i < n; i++)
	{
		list *n = new list;
		(*n).value = data[i];

		m = data[i] / 10;       //确定要存到哪个桶中 data[i]*num/范围
		if(bucket[m].next==NULL)
		bucket[m].next = n;      //桶中第一个节点不存数，只存一个指针
		else
		{
			
			p = &bucket[m];     
			q = bucket[m].next;
			while ((q!=NULL)&&(q->value < data[i]))   //这里是q  不是p  如果是平，那么贵导致第一个值为null  无法比较
			{
				p = p->next;
				q = q->next;
			}
			p->next = n;      //插入链表
			n->next = q;
		}
	}
	int count1 = 0;                //把数据从桶中取出来
	for (int i = 0; i < 10; i++)
	{
		p=bucket[i].next;
		if (p == NULL)     //记得判断是否为空桶 避免错误
			continue;
		while (p!= NULL)
		{
			original[count1++] = p->value;
			p = p->next;
		}
		
	}
	return original;
}
/*
计数排序，是基数排序的内排序  与桶排序一致  但桶排序是一个很大范围的数 排到一个小范围的桶中 内部排序用到了链表
扫描序列A，以A中的每个元素的值为索引，把出现的个数填入C中。此时C[i]可以表示A中值为i的元素的个数。
对于C从头开始累加，使C[i]=C[i]+C[i-1]。这样，C[i]就表示A中值不大于i的元素的个数。
按照统计出的值，输出结果
*/
int * sort::counting_sort(int * original, int n)
{
	int count[100];     //定义桶
	int *output = new int[n];   //定义输出数组
	for (int i = 0; i < 100; i++)  //清空桶
	{
		count[i] = 0;
	}
	for (int i = 0; i < n; i++)  //桶中计数
	{
		count[original[i]]++;
	}

	for (int i = 1; i < 100; i++)  //计算位置
		count[i] += count[i - 1];

	for (int i = n - 1; i >= 0; i--)   //输出
	{
		output[count[original[i]]-1] = original[i];
		count[original[i]]--;
	}
	for (int i = 0; i < n; i++)  //复制给原始数组
	{
		original[i] = output[i];
		
	}
    delete [] output;    //使用c++写程序的时候 new了 要记得delete
	return original;
}



int main()
{
	vector<int> a, *b;
	int *a1,*b1;
	int num = 20;
	sort ss(num);
	a = ss.create_arr();
	a1 = ss.creat_arr1();
	for (int j = 0; j < num; j++)
	{
		if (j % 5 == 0)
			cout << "\n";
		cout << a1[j] << "   ";

	}
	b1 = ss.counting_sort(a1,num);
	cout << "\n 最终结果：\n";
	for (int j = 0; j < num; j++)
	{
		if (j % 5 == 0)
			cout << "\n";
		cout << b1[j] << "   ";

	}
	getchar();

	return 0;
}
```

