#include <vector>
#include <iostream>
#include <unordered_map>
#include <stack>

using namespace std;

struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL){};
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
        unordered_map<int, int> hash;
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
    //二叉树的后续递归实现
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
            treenode *t = (treenode *)malloc(sizeof(treenode));
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
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 3, 5, 7, 9};
    vector<int> nums1 = {2, 4, 6};
    double result=solution.find_mid_num(nums,nums1);
    cout<<result;
    cin.get();
    return 0;
}
