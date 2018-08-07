#include <vector>
#include <iostream>
#include <unordered_map>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
 

class Solution {
public:
//给定一个数组，一个目标数，找出数组中两个相加为目标的数
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int,int> hash;
        vector<int> result;
        for(int i=0;i<nums.size();i++)
        {
            int numtofind=target-nums[i];
            if (hash.find(numtofind)!=hash.end())
            {
                result.push_back(hash[numtofind]);
                result.push_back(i);
                return result;
            }
            hash[nums[i]]=i;
        }
    }
//两个链表相加
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int sum ,low,high;
        while(l1 && l2)
        {
            sum=(*l1).val+(*l2).val;
            low=sum %10;
            high=sum/10;
            ListNode *newlist;
        }
    }
//两个数组求中位数
    vector<int> find_minddle(vector<int> a, vector<int> b ){
        int m= a.size();
        cout << m<<endl;


    }


//最大子序列





};


int main()
{
Solution solution;
vector<int> nums;
nums={2,7,11,15,12};
vector<int> result=solution.find_minddle(nums,nums);
// cout << result << endl;
}
