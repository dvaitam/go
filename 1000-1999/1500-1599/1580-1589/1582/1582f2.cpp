// #include<iostream>
// #include<vector>
#include <bits/stdc++.h>
using namespace std;
vector<vector<bool> > table(5001, vector<bool>(8192, false));
vector<vector<int> > upd(5001);

void setv(int a, int b) {
    if(!table[a][b]) {
        table[a][b] = true;
        if (a) {
            setv(a-1, b);
        }
        upd[a].push_back(b);
    }
}
void roll(int a) {
    while (!upd[a].empty()) {
        setv(a-1, upd[a].back()^a);
        upd[a].pop_back();
    }
}
int main() {
    int n;
    cin>>n;
    vector<int> a(n);
    for(int i=0; i < n; i++){
        cin>>a[i];
    }
    for(int i =0;i <=5000; i++ ){
        setv(i, 0);
    }
    for(int i=n-1;i >=0; i--) {
        if(a[i] > 0 ){
            roll(a[i]);
        }
    }
    vector<int> out;
    for(int i=0;i < 8192;i++) {
        if (table[0][i]) {
            out.push_back(i);
        }
    }
    cout<<out.size()<<"\n";
    for(auto v:out){
        cout<<v<<" ";
    }
}