#include <iostream>

using namespace std;

class a{
    public:

    void p(){
        cout<<'a';
    };
    void pp(){
        cout<<"bb";
    };
};

int main()
{
    a aa;
    aa.p();
    cin.get();
}