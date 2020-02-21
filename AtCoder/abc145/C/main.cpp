#include <bits/stdc++.h>
using namespace std;

int N;
int x[9], y[9];

double calc(int i, int j) {
    double dx = x[i] - x[j];
    double dy = y[i] - y[j];
    return pow(dx * dx + dy * dy, 0.5);
}

int main() {
    cin >> N;
    for (int i=1; i <= N; i++) cin >> x[i] >> y[i];

    double sum = 0.0;
    vector<int> indices(N);
    for (int i=0; i < N; i++) indices[i] = i + 1;

    double step = 0.0;
    do {
        for(int i=0; i < N-1; i++) {
            sum += calc(indices[i], indices[i+1]);
        }
    step++;
    } while (next_permutation(indices.begin(), indices.end()));

    cout << fixed << setprecision(10) << sum / step << endl;

    return 0;
}