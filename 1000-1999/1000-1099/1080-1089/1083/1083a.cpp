#include <iostream>
#include <vector>
#include<map>
using namespace std;
int64_t max(int64_t a, int64_t b) {
    if(a > b) {
        return a;
    }
    return b;
}
int64_t go_down(int64_t root, int64_t parent, vector<int64_t> & w, map<int64_t, int64_t> & l, map<int64_t, vector<int64_t> > & adj,map<int64_t, int64_t> & cache ) {
    if (cache[root*1000000+parent] > 0) {
		return cache[root*1000000+parent] - 1;
	}
    int64_t mx = w[root];
	int64_t fans = mx;
    vector<int64_t> v = adj[root];
	for(int i=0;i < v.size(); i++) {
		if (v[i] == parent ){
			continue;
		}
		int64_t mxc;
        mxc = go_down(v[i], root, w, l, adj, cache);
		//fmt.Println("mxc for", v, root, mxc, mxc-l[v][root]+mx)
		fans = max(fans, mxc-l[v[i]*1000000+root]+mx);
	}
	cache[root*1000000+parent] = fans + 1;

	return fans;
}

int main() {
    int64_t n;
    cin>>n;
    vector<int64_t> w;
    w.push_back(0);
    int64_t  ww;
    for(int i=1; i <= n; i++) {
        cin>>ww;
        w.push_back(ww);
    }
    map<int64_t, int64_t> l, cache;

    map<int64_t, vector<int64_t> > adj;
    for(int i=1; i < n; i++) {
        int64_t u, v, c;
        cin>>u>>v>>c;
        adj[u].push_back(v);
        adj[v].push_back(u);
        l[u*1000000+v] = c;
                l[v*1000000+u] = c;
    }
    int64_t ans = 0;
    for(int i =1; i <= n; i++) {
        ans = max(ans, go_down(i, i,w, l, adj, cache));
    }
    cout<<ans<<"\n";

    return 0;
}
