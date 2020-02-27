#include <bits/stdc++.h>
#define all(v) (v).begin(),(v).end()
using namespace std;


void solve(long long N, long long H, std::vector<long long> a, std::vector<long long> b){
    // 振り最強ソード
    long long a_max = *max_element(all(a));
    sort(all(b));
    reverse(all(b));
    long long count = 0;

    for (int i=0; i<N; i++) {
        if (a_max < b[i]) {
            H -= b[i];
            count++;
            if (H <= 0) break;
        }
    }

    while (H > 0) {
        H -= a_max;
        count++;
    }
    cout << count << endl;
}

// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    long long N;
    scanf("%lld",&N);
    long long H;
    scanf("%lld",&H);
    std::vector<long long> a(N);
    std::vector<long long> b(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&a[i]);
        scanf("%lld",&b[i]);
    }
    solve(N, H, std::move(a), std::move(b));
    return 0;
}