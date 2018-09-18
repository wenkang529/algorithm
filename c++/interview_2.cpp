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
