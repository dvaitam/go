#include<iostream>
using namespace std;
int main(){
    int a, n;
    long long ans = 0;
    cin>>a>>n;
    int primes[674579];
    int largest_square[a+n];
    primes[0] = 2;
	int next_prime_index = 1;
    largest_square[1] = 1;
	if (2 < a+n ){
		largest_square[2] = 1;
	}
    for(int i = 3; i < a+n; i++) {
		int prime = 1;
		for(int j = 0; j < next_prime_index; j++ ){
			if( i%primes[j] == 0 ){
				int rem = i / primes[j];
				if( rem%primes[j] == 0) {
					int sq = primes[j] * primes[j];
					largest_square[i] = largest_square[i/sq] * sq;
				} else {
					largest_square[i] = largest_square[i/primes[j]];
				}
				prime = false;
				break;
			}
			if( primes[j]*primes[j] > i ){
				break;
			}
		}
		if (prime == 1) {
			largest_square[i] = 1;
			primes[next_prime_index] = i;
			next_prime_index++;
			
		}
		if( i >= a) {
			ans += i / largest_square[i];
		}
		
	}
	for(int i = 1; i < 3; i++ ){
		if (i >= a && i < a+n) {
			ans += i / largest_square[i];
		}
	}
    cout<<ans<<"\n";

return 0;
}