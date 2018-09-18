#include <iostream>
#include<malloc.h>
using namespace std;

typedef struct BinaryTreeNode
{
    char data;
    struct BinaryTreeNode *Left;
    struct BinaryTreeNode *Right;
} *Nodes,Node;

void createBinaryTree(Nodes & p)
{
    char c;
    cin>>c;

        if (c == '#') //如果到了叶子节点，接下来的左、右子树分别赋值为0
        {
            p = NULL;
        }
        else
        {
            p = new (Node);
        
            p->data = c;
            createBinaryTree(p->Left);  //递归创建左子树
            createBinaryTree(p->Right); //递归创建右子树
        }
}

//124###3##
int main()
{
    Node *a;
    createBinaryTree(a);
    cin.get();
    return 0;
}
