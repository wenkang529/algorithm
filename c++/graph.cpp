#include <iostream>
#include <vector>
#include <queue>
using namespace std;

/*
save graph:connection matrix

*/

/*
traverse the graph
two methods:
1.DFS 
2.WFS
*/

//DFS
class Graph
{
    typedef int graph_data;
    static const int MAX_VERTEX = 100;
    struct Matrix_graph
    {
        graph_data vertex[MAX_VERTEX];    //maxnum of vertex
        int edge[MAX_VERTEX][MAX_VERTEX]; //info of each edge
        int vertex_num, edge_num;
        int visit[MAX_VERTEX];
    };

    void DFS(Matrix_graph G, int i)
    {
        G.visit[i] = true;
        cout << G.vertex[i];
        for (int j = 0; j < G.vertex_num; j++)
        {
            if (G.edge[i][j] == 1 && !G.visit[j])
            {
                DFS(G, j);
            }
        }
    }
    void DFs_travel(Matrix_graph G)
    {
        for (int i = 0; i < G.vertex_num; i++)
        {
            G.visit[i] = false;
        }
        for (int i = 0; i < G.vertex_num; i++)
        {
            if (!G.visit[i])
            {
                DFS(G, i);
            }
        }
    }
//BFS
/*
use queue to realize BFS
two-level loop
for 0->num
    Q push
    cout
    for 0->num
        Q pop 
        cout
        Q push
*/
    void BFS_travel(Matrix_graph G, int i)
    {
        queue<int> Q;
        for (int i = 0; i < G.vertex_num; i++)
        {
            G.visit[i] = false;
        }
        for (int i = 0; i < G.vertex_num; i++)
        {
            if (!G.visit[i])
            {
                G.visit[i] = true;
                cout << G.vertex[i];
                Q.push(G.vertex[i]);
                while (!Q.empty())
                {
                    int tmp = Q.front();
                    Q.pop();
                    for (int j = 0; j < G.vertex_num; j++)
                    {
                        if (G.visit[j] && G.edge[tmp][j] == 1)
                        {
                            Q.push(G.vertex[i]);
                            G.visit[j] = true;
                            cout<<G.vertex[j];

                        }
                    }
                }
            }
        }
    }
};