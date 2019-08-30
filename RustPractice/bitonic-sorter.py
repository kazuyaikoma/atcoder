# リストxの要素をupで指定された向きにソート 
def sort(x, up):
  if len(x) <= 1:
    return x
  else:
    mid_point = len(x) // 2
    first = sort(x[:mid_point], True)
    second = sort(x[mid_point:], False)

    x1 = first + second
    return _sub_sort(x1, up)

# bitonic列の前半と後半を、upの向きに交換しながら再帰的にソート
def _sub_sort(x, up):
  if len(x) == 1:
    return x
  else:
    # up方向への大小の交換
    _compare_and_swap(x, up)

    mid_point = len(x) // 2
    first = _sub_sort(x[:mid_point], up)
    second = _sub_sort(x[mid_point:], up)
    
    return first + second

# 要素数nのbitonic列の要素をn/2ごとに比較 → up向きに交換
def _compare_and_swap(x, up):
  mid_point = len(x) // 2
  for i in range(mid_point):
    if (x[i] > x[mid_point + i]) == up:
      x[i], x[mid_point + i] = x[mid_point + i], x[i]
