#include <iostream>
#include <stack>
#include <queue>
#include <cstdio>
using namespace std;

int st();
int que();

int main() {
  printf("------stack------\n");
  st();
  printf("------queue------\n");
  que();
  return 0;
}

int st() {
  int nums[] = {0, 1, 2, 3, 4};
  stack<int> s;

  // push
  for (auto x: nums) {
    s.push(x);
    printf("pushed %d\n", s.top());
  }

  // pop
  while (!s.empty()) {
    printf("popping %d\n", s.top());
    s.pop();
  }

  return 0;
}

int que() {
  queue<int> que;
  int nums[] = {0, 1, 2, 3, 4};

  // push
  for (auto x: nums) {
    que.push(x);
    printf("pushed %d\n", que.back());
  }

  // pop
  while (!que.empty()) {
    printf("popping %d\n", que.front());
    que.pop();
  } 
  return 0;  
}