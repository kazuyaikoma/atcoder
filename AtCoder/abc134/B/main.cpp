#include <bits/stdc++.h>
using namespace std;


void solve(long long N, long long D){
    int count = 0;
    while (N > 0) {
        N -= 2 * D + 1;
        count++;
    };
    cout << count << endl;
}

// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    long long N;
    scanf("%lld",&N);
    long long D;
    scanf("%lld",&D);
    solve(N, D);
    return 0;
}
