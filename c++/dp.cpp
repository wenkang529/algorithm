#include <iostream>
#include <vector>
using namespace std;

//dynamic programing
//key is to find state and state transition equation(状态转移方程)
//
/*
LIS longest increasing subsequence
state:d[i] :lis which end with i. return the max[]
state transition equation: d[p]-->d[p+1]:
    if arr[p+1] is bigger than arr[0--i],so d[p+1]=max(d[0-p](arr[p+1]>arr[m])+1,1)
    so if arr[p+1] bigger than some elment in arry before. max(d[those elment])+1,if no element
    d[p+1]=1
remeber to save each d[p]
*/
int lengthOfLIS(vector<int> &nums) // reference (alias:another name)
{
    int n = nums.size();
    if (n <= 1)
    {
        return n;
    }
    int d[n];
    d[0] = 1;
    int maxr = 0;
    for (int i = 0; i < n; i++)
    {
        int current = nums[i];
        int maxp = 0;
        for (int j = i - 1; j >= 0; j--)
        {
            if (current > nums[j])
            {
                maxp = maxp < d[j] + 1 ? d[j] : maxp;
            }
        }
        d[i] = maxp + 1;
        maxr = maxr < d[i] ? d[i] : maxr;
    }
    return maxr;
}
//valid parenthess "()"
//key: d[i] ending with s[i] valid parenthess,so if s[i]="(" d[i]=0
//     if s[i]=")" then if s[i-1]=="(" or s[i-d[i-1]-1]=="("
int longestValidParentheses(string s)
{
    int len = s.size();
    if (len < 2) //place judge before int len[0]
        return 0;
    int d[len];
    d[0] = 0;
    int max = 0;

    for (int i = 1; i < len; i++)
    {
        if (s[i] == '(')
        {
            d[i] = 0;
        }
        else //if s[i]==')'
        {
            if (s[i - 1] == '(')
                d[i] = i > 2 ? d[i - 2] + 2 : 2;
            else if ((i - d[i - 1] - 1 >= 0) && (s[i - d[i - 1] - 1] == '('))
                d[i] = i - d[i - 1] - 2 > 0 ? d[i - d[i - 1] - 2] + d[i - 1] + 2 : d[i - 1] + 2;
            else
                d[i] = 0;
        }
        max = max < d[i] ? d[i] : max;
    }

    return max;
}
//paliindrome 回文
//most important is to find the correct equation of state transition,
//first i had missed the target,and all i did is wrong
//key:d[i][j]
//init: d[i][i]=true    d[i][i+1]=s[i]==s[i+1]?true:false
//dp i(n-->0) j(i-->n)  d[i][j]=(s[i]==s[j]&&d[i+1][j-1])

string longestPalindrome(string s)
{
    int len = s.size();
    if (len < 2)
        return s;
    bool d[len][len];
    //init
    for (int i = 0; i < len; i++)
    {
        d[i][i] = true;
        if (i < len - 1)
            d[i][i + 1] = s[i] == s[i + 1] ? true : false;
    }
    //dp
    for (int i = len - 3; i >= 0; i--)
    {
        for (int j = i + 2; j < len; j++)
        {
            d[i][j] = (s[i] == s[j] && d[i + 1][j - 1]);
        }
    }
    //return the longest palindrome
    int maxnum = 0;
    int m = 0, n = 0;
    for (int i = 0; i < len; i++)
    {
        for (int j = i; j < len; j++)
        {
            if (d[i][j] && j - i + 1 > maxnum)
            {
                maxnum = j - i + 1;
                m = i;
                n = j;
            }
        }
    }
    // cout<< maxnum<<endl;
    return s.substr(m, n - m + 1); //string.substr(begin,size)
}
/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false

*/

bool isMatch(string s, string p)
{
    int lens = s.length();
    int lenp = p.length();
    bool d[lens + 1][lenp + 1];
    d[0][0] = true;

    //初始化行 p为空的时候 s不为空 则必为F
    for (int i = 1; i < lens + 1; i++)
    {
        d[i][0] = false;
    }
    //初始化列 s为空 只有p为 X* 才有可能为T 且为  d[0][j-2]&&p[j]=='*'
    for (int j = 1; j < lenp + 1; j++)
    {
        d[0][j] = j > 1 && p[j - 1] == '*' && d[0][j - 2];
    }
    //遍历的顺序为 外层为s  内层为p
    for(int i=1;i<lens+1;i++){
        for(int j=1;j<lenp+1;j++){
            if(p[j-1]=='*'){
                d[i][j]=d[i][j-2]||((p[j-2]==s[i-1]||p[j-2]=='.')&&d[i-1][j]);  //匹配0个||多个   0个即为d[i][j-2] 多个即为满足后面的两个条件
            }else{
                d[i][j]=(p[j - 1] == '.' || s[i - 1] == p[j - 1]) && d[i - 1][j - 1];  //没有*的时候  一个对一个的匹配
            }
        }
    }

    return d[lens][lenp];
}

int main()
{
    // vector<int> arr = {1, 3, 6, 7, 9, 4, 10, 5, 6};
    bool result = isMatch("ab",".*");
    cout << "result:" << result << endl;
    cin.get();
    return 0;
}
