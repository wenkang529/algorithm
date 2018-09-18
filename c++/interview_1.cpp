#include <vector>
#include <iostream>
#include <unordered_map>
#include <stack>
#include <string>
#include <map>
#include <queue>
#include<malloc.h>

using namespace std;

struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
typedef int datatype;
struct Treenode
{
    datatype value;
    struct Treenode *left;
    struct Treenode *right;
    Treenode(datatype _value) : value(_value), left(NULL), right(NULL) {}
};
class Solution
{
  public:
    //找出和为目标的两个数
    vector<int> twoSum(vector<int> &nums, int target)
    {
        map<int, int> hash;
        vector<int> result;
        for (int i = 0; i < nums.size(); i++)
        {
            int numtofind = target - nums[i];
            if (hash.find(numtofind) != hash.end())
            {
                result.push_back(hash[numtofind]);
                result.push_back(i);
                return result;
            }
            hash[nums[i]] = i;
        }
    }
    //两个列表累加
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
    {
        int sum, low, high;
        while (l1 && l2)
        {
            sum = (*l1).val + (*l2).val;
            low = sum % 10;
            high = sum / 10;
            ListNode *newlist;
        }
    }
    // 二叉树的遍历 前序遍历 根->左->右
    //中序遍历 左 根 右
    //后序遍历 左 右 根
    //递归方式
    void preoderTraverse(Treenode *head)
    {
        if (head == NULL)
        {
            return;
        }
        cout << head->value << endl;
        preoderTraverse(head->left);
        preoderTraverse(head->right);
    }
    //二叉树的前序遍历非递归
    void N_preoderTraverse(Treenode *head)
    {
        if (head == NULL)
        {
            return;
        }
        stack<Treenode *> s;
        Treenode *p;
        p = head;
        while (p != NULL || !s.empty())
        {
            while (p != NULL)
            {
                cout << p->value << endl; //输出永远只能输出当前的val,不输出左或右的val
                s.push(p);
                p = p->left;
            }
            if (!s.empty())
            {
                p = s.top();
                s.pop();
                p = p->right;
            }
        }
    }
    //二叉树中序遍历递归实现
    void midTraverse(Treenode *head)
    {
        if (head != NULL)
        {
            midTraverse(head->left);
            cout << head->value << endl;
            midTraverse(head->right);
        }
    }
    //二叉树中序遍历非递归实现
    void N_midTraverse(Treenode *head)
    {
        Treenode *p = head;
        stack<Treenode *> s;

        while (p != NULL || !s.empty())
        {
            while (p != NULL)
            {
                s.push(p);
                p = p->left;
            }
            if (!s.empty())
            {
                p = s.top();
                cout << p->value << endl;
                p = p->right;
                s.pop();
            }
        }
    }
    //二叉树后序遍历递归实现
    void postTraverse(Treenode *head)
    {
        if (head != NULL)
        {
            postTraverse(head->left);
            postTraverse(head->right);
            cout << head->value;
        }
    }
    //用于二叉树的后续递归实现
    struct treenode
    {
        Treenode *node;
        bool isFirst;
    };
    //总体思路是先遍历左子树,这时根节点在栈顶,再访问右子树,这个节点右出现在栈顶,所以当第二次访问到这个元素时就可以输出.
    //再定义一个变量存储是否第一次输出,并将这个元素放入栈
    void N_postTraverse(Treenode *head)
    {
        stack<treenode *> s;
        Treenode *p = head;
        treenode *tem;
        //先把左子树遍历一遍 把标志节点放入堆栈
        while (p != NULL || !s.empty())
        {
            // treenode *t = (treenode *)malloc(sizeof(treenode));
            treenode *t=new treenode;
            t->node = p;
            t->isFirst = true;
            s.push(t);
            p = p->left;
        }
        if (!s.empty())
        {
            tem = s.top();
            s.pop();
            if (tem->isFirst == true)
            {
                tem->isFirst = false;
                s.push(tem);
                p = tem->node->right;
            }
            else
            {
                cout << tem->node->value << endl;
                p = NULL;
            }
        }
    }
    //最大子序列
    //采用的是类似递推的方法,设以第i个结尾的最大子序列的和为sum,如果前一个状态的和为负数,则子序列的和为当前的值.否则就累加
    //找出最大的和即可
    int maxsum(int *arr, int len)
    {
        // assert(arr==NULL);
        int sum = arr[0];
        int amax = sum;
        for (int i = 1; i < len; i++)
        {
            if (sum < 0)
            {
                sum = arr[i];
            }
            else
            {
                sum += arr[i];
            }
            if (amax < sum)
                amax = sum;
        }
        return amax;
    }
    //找出两个有序数组的中位数
    //中位数有两中,当为偶数时,中位数可以是两个数的平均,也可以分为上中位数和下中位数
    //一个比较好的方法是转化为求递i个小的数
    //getkth()中的k 范围是1:n  不能是0 否则会出现找不到文件malloc.c
    double find_mid_num(vector<int> arr1, vector<int> arr2)
    {
        int len1 = arr1.size();
        int len2 = arr2.size();
        //因为要求求中位数,所以如果是奇数那么这两个是一样的
        int left = (len1 + len2 + 1) / 2;
        int right = (len1 + len2 + 2) / 2;
        return (getkth(arr1, 0, len1 - 1, arr2, 0, len2 - 1, left) + getkth(arr1, 0, len1 - 1, arr2, 0, len2 - 1, right)) * 0.5;
    }
    int getkth(vector<int> arr1, int start1, int end1, vector<int> arr2, int start2, int end2, int k)
    {
        int len1 = end1 - start1 + 1;
        int len2 = end2 - start2 + 1;
        if (len1 > len2)
        {
            return getkth(arr2, start2, end2, arr1, start1, end1, k);
        }
        if (len1 == 0)
        {
            return arr2[start2 + k - 1];
        }
        if (k == 1)
        {
            return arr1[start1] < arr2[start2] ? arr1[start1] : arr2[start2];
        }
        //i,j 为切到i,j的位置
        int i = start1 - 1 + (len1 < (k / 2) ? len1 : (k / 2));
        int j = start2 - 1 + (len2 < (k / 2) ? len2 : (k / 2));
        if (arr1[i] > arr2[j])
        {
            return getkth(arr1, start1, end1, arr2, j + 1, end2, k - (j - start2 + 1));
        }
        else
        {
            return getkth(arr1, i + 1, end1, arr2, start2, end2, k - (i - start1 + 1));
        }
    }
    //最长不重复子序列
    //解题思路,简单的思路是直接遍历的时候回查是否在子序列中,这样复杂度为O(N2)
    //其实第二遍的遍历可以用空间换时间,用O(1)的方法(数组下标,哈希查找等),这里采用了数组,用一个256长度的数组,存储对应的最后一个的位置.
    // 这样只需遍历一遍就可以
    int lengthOfLongestSubstring(string s)
    {
        vector<int> dict(256, -1);
        int current = 0, maxnum = 0;
        for (int i = 0; i < (s.length()); i++)
        {
            char c = s[i];
            if (dict[int(c)] != -1)
            {
                current = current < (dict[int(c)] + 1) ? (dict[int(c)] + 1) : current;
            }

            dict[int(c)] = i;
            maxnum = maxnum < (i - current + 1) ? (i - current + 1) : maxnum;
        }
        return maxnum;
    }
    //container with most water
    //包含最多的水
    //主要思路是从两测开始i,j 我们只需移动短的那个,因为移动长的那个只会变短
    //If we try to move the pointer at the longer line inwards, we won't gain any increase in area,
    int maxarea(vector<int> arr)
    {
        int i = 0, j = (arr.size() - 1);
        int maxa = 0;
        while (i < j)
        {
            maxa = maxa < (min(arr[i], arr[j]) * (j - i)) ? (min(arr[i], arr[j]) * (j - i)) : maxa;
            if (arr[i] < arr[j])
                i++;
            else
                j--;
        }
        return maxa;
    }
    //合并k个有序的链表
    //重点关注链表的操作
    ListNode *mergeKLists(vector<ListNode *> &lists)
    {
        int len = lists.size();
        vector<ListNode *> j(len, NULL);
        ListNode *result = new ListNode(0);
        ListNode *p = result;
        // result->next = p;
        int index = -1;
        while (lists != j)
        {
            ListNode *new_l = new ListNode(0);
            index = -1;
            int tem_max;
            for (int i = 0; i < len; i++)
            {
                if (lists[i] == NULL)
                    continue;
                if (index == -1 || tem_max > lists[i]->val)
                {
                    index = i;
                    tem_max = lists[index]->val;
                }
            }
            new_l->val = tem_max;
            lists[index] = lists[index]->next;
            p->next = new_l;
            p = p->next;
        }
        return result->next;
    }
    /*
    每k个节点反转一次链表
    中间遇到的问题
    1.注意链表反转的时候那个是头节点 哪个需要与下一个连接
    2.反转后的最后一个节点next必须置为NULL ,否则会因为其本身指向的一个节点而形成的循环.
    3.反转链表只需要每次反转k个就可以,而不是k加一个
    */
    ListNode *reverseKGroup(ListNode *head, int k)
    {
        vector<ListNode *> s(k, NULL); //只有new出来的才需要初始化才分配空间,创建向量已经分配空间了,可以初始化为null
        ListNode a(0);
        ListNode *r1 = &a;
        ListNode *r2 = r1;
        int index;
        if (head == NULL)
            return head;
        while (head != NULL)
        {
            //初始化向量
            for (int i = 0; i < k; i++)
            {
                s[i] = NULL;
            }
            //把节点依次放进向量,如果还未放完遇到了空就break 返回index
            for (int i = 0; i < k; i++)
            {
                if (head != NULL)
                {
                    s[i] = head;
                    head = head->next;
                }
                else
                {
                    index = i - 1;
                    break;
                }
            }
            //判断并反转,注意最后按是否需要反转分别处理
            if (s[k - 1] != NULL)
            {
                for (int i = 1; i < k; i++)
                {
                    s[i]->next = s[i - 1];
                }
                r1->next = s[k - 1];
                r1 = s[0];
            }
            else
            {
                r1->next = s[0];
                r1 = s[index];
            }
            //最后的一个节点 next指向空
            r1->next = NULL;
        }
        return r2->next;
    }
    //返回needle字符串在haystack字符串中的位置
    int strStr(string haystack, string needle)
    {
        if (haystack == "" && needle == "")
            return 0;
        int len_s = haystack.length();
        int len_n = needle.length();
        if (len_s < len_n)
        {
            return -1;
        }
        for (int i = 0; i < len_s - len_n; i++)
        {
            if (haystack.substr(i, len_n) == needle)
                return i;
        }
        return -1;
    }
    //Substring with Concatenation of All Words
    //a包含b列表中的所有字符串的位置
    //还有一个想法,用map和队列的集合,固定长度的队列,每次进一个新的元素就删除一个元素
    vector<int> findSubstring(string s, vector<string> &words)
    {

        vector<int> r1;
        int lens = s.size();
        int len_v = words.size();

        if (len_v == 0 || lens == 0)
            return r1;

        int len_vs = words[0].size();
        int all_len = len_v * len_vs;
        if (lens < all_len)
            return r1;

        vector<int> flag_copy(len_v, 0);
        vector<int> flag = flag_copy;

        vector<int> right(len_v, 1);

        map<string, int> count;
        map<string, int> seen;

        //记录words每个单词的个数
        for (int i = 0; i < len_v; i++)
        {
            if (count.find(words[i]) != count.end())
            {
                count[words[i]]++;
            }
            else
                count[words[i]] = 1;
        }
        //iterate
        for (int i = 0; i < len_vs; i++)
        {
            flag = flag_copy;
            for (int j = i; j < lens - all_len + 1; j += len_vs)
            {
                if (count.find(s.substr(j, len_vs)) == count.end())
                {
                    continue;
                }
                else
                {
                    map<string, int> tmp = count;
                    //计算是否符合 (计算长度为all_len) 其实这样也是有重复计算的
                    for (int m = j; m < j + all_len; m += len_vs)
                    {
                        if (count.find(s.substr(m, len_vs)) == count.end())
                        {
                            j = m;
                            break;
                        }
                        else
                        {
                            tmp[s.substr(m, len_vs)] -= 1;
                        }
                    }
                    int f = 1;
                    for (map<string, int>::iterator iter = tmp.begin(); iter != tmp.end(); iter++)
                    {
                        if (iter->second != 0)
                        {
                            f = 0;
                            break;
                        }
                    }
                    if (f == 1)
                    {
                        r1.push_back(j);
                    }
                }
            }
            if (flag == right)
            {
                r1.push_back(i);
            }
        }
        return r1;
    }
    //二叉数的重建
    /*
    已知二叉树的前序遍历,中序遍历,重建二叉树
    思路:假设已知前序遍历为[1,2,4,7,3,5,6,8],中序遍历为[4,7,2,1,5,3,6,8]
    - 前序遍历第一个为 1  那么1为根节点 
    - 在中序遍历中找到 1 的位置,那么1左边的为左子树,右边的为右子树
    - 1左边有三个值  那么前序遍历中 除了第一个根节点 还有三个值是左子树 剩下的值为右子树 
    - 这样就可以递归的用在左子树和右子树中
    */
    Treenode *tree_rebuild(vector<int> preoder, vector<int> midoder)
    {
        int pre_len = preoder.size();
        int mid_len = midoder.size();

        if (pre_len != mid_len || pre_len == 0)
            return NULL;

        return tree_rebuild_f(preoder, midoder, 0, pre_len - 1, 0, mid_len - 1);
    }
    Treenode *tree_rebuild_f(vector<int> preoder, vector<int> midoder, int pre_begin, int pre_end, int mid_begin, int mid_end)
    {
        //判断是否合法
        if (pre_begin > pre_end)
        {
            return NULL;
        }
        //找出中间节点
        int value = preoder[pre_begin];
        int index = mid_begin;
        while (index <= mid_end && midoder[index] != value)
        {
            index++;
        }

        if (index > mid_end)
        {
            throw "Invalid input";
        }
        //递归调用左节点和由节点
        Treenode *root = new Treenode(value);
        root->left = tree_rebuild_f(preoder, midoder, pre_begin + 1, pre_begin + index - mid_begin, mid_begin, index - 1);
        root->right = tree_rebuild_f(preoder, midoder, pre_begin + index - mid_begin + 1, pre_end, index + 1, mid_end);
        return root;
    }
    //Gavin a collection of distinct integers,return possible permutations.
    //can't iterater element if use queue,can't ust pop to del first element powerful if use vector
    //so use recursive(递归) is well solution.
    vector<vector<int> > permute(vector<int> &nums)
    {
        vector<vector<int> > Q;
        int size = nums.size();
        if (size == 0)
            return Q;
        premutation(nums, Q, 0, size - 1);
    }
    void premutation(vector<int> &nums, vector<vector<int> > Q, int left, int right)
    {
        if (left == right)
            Q.push_back(nums);
        for (int i = left; i <= right; i++)
        {
            swap(nums[left], nums[i]);
            premutation(nums, Q, left + 1, right);
            swap(nums[left], nums[i]);
        }
    }

    void createBinaryTree(Treenode *p)
    {
        char c;
        cin >> c;

        if (c == '#') //如果到了叶子节点，接下来的左、右子树分别赋值为0
        {
            p = NULL;
        }
        else
        {
            p = new Treenode(0);
            //          p = (Node *)malloc(sizeof(Node));  //两种方式都可以,malloc需要加入头文件 #include<malloc.h>
            p->value = c;
            createBinaryTree(p->left);  //递归创建左子树
            createBinaryTree(p->right); //递归创建右子树
        }
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {2,4,6};
    ListNode *a = new ListNode(1);
    ListNode *b = new ListNode(2);
    ListNode *c = new ListNode(3);
    ListNode *d = new ListNode(4);
    ListNode *e = new ListNode(5);
    ListNode *f = new ListNode(6);
    ListNode *g = new ListNode(7);
    a->next = b;
    b->next = c;
    c->next = d;
    d->next = e;
    e->next = f;
    vector<string> pa = {"bar", "foo"};
    vector<int> preoder = {10, 6, 4, 8, 14, 12, 16};
    vector<int> midoder = {4, 6, 8, 10, 12, 14, 16};
    Treenode *result = solution.tree_rebuild(preoder, midoder);

    solution.postTraverse(result);
    std::cin.get();
    return 0;
}
