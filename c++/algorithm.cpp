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
};


int main()
{
Solution solution;
vector<int> nums;
nums={2,7,11,15};
vector<int> result=solution.twoSum(nums,9);
for(int i=0;i<result.size();i++)
{
    cout << result[i] << endl;
}
}
