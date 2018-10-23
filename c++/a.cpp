
#include <iostream>
#include <cstdio>
#include <vector>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int arr[n][4];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < 4; j++)
        {
            cin >> arr[i][j];
        }
    }
    for (int j = 0; j < n; j++)
    {
        int nn[4]={arr[j][0],arr[j][1],arr[j][2],arr[j][3]};
        cout << find(nn);
        int n;
    }
    cin.get();
    cin.get();
    return 0;
}

int find(int arr[])
{
    vector<int> k;
    int old[3]={arr[0],arr[1],arr[2]};
    int j=3;
    int a;
    while (j < arr[3]+1)
    {
        a = old[0] + old[1] + old[2];
        vector<int> k;
        while (a)
        {
            k.push_back(a % 10);
            a = a / 10;
        }
        int len=k.size();
        if (len+j >= arr[3])
            return k[arr[3]-j-1];
        if (len>2)
        {
            old[0]=k[2];
            old[1]=k[1];
            old[2]=k[0];
        }
        
        else if (len==2)
        {
            old[0]=old[2];
            old[1]=k[1];
            old[2]=k[0];
        }
        else if (len==1)
        {
            old[0]=old[1];
            old[1]=old[2];
            old[2]=k[0];
        }
        
}
}
