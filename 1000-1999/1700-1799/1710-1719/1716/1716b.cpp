#include<iostream>
using namespace std;
int main(){
    int T;
    cin>>T;
    for(int t=0;t<T;t++) {
        int n;
        cin>>n;
        cout<<n<<"\n";
		for(int i = 1; i <= n; i++){
			cout<<i;
			if(i != n) {
				cout<<" ";
			}
		}
        cout<<"\n";
		for(int k = n - 1; k >= 1; k-- ){
			for (int i = 1; i <= n; i++ ){
				if( i != k ){
					cout<<i<<" ";
				}
			}
			cout<<k<<"\n";
		}
    }

}