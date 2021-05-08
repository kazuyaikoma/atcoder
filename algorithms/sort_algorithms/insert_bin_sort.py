import unittest
from typing import List


class TestInsertSort(unittest.TestCase):
    """
    挿入ソート(二分探索バージョン)
    先頭から順に舐めていき、ソート済み配列の良しなな場所に挿入していく
    挿入場所の探索に二分探索を使うと高速化が図れる。

    時間計算量 O(N^2)
    空間計算量 O(1)
    安定ソート
    """

    # 挿入ソートにおける挿入位置を二分探索し、挿入箇所のidx+1を返す
    def search(self, arr: List[int], low: int, high: int, target: int) -> int:
        if low == high:
            if target < arr[low]:
                return low
            else:
                # arr[high] <= target; ということなので、targetのidxを返すことになる
                # return low + 1; と等価だが、highの方がイメージに合っている
                # これが返される場合、配列のソート順は変わらない(targetのidxを返すので)
                return high + 1
        # lowがhighを越したのであれば、その時点でlowを返す
        elif high < low:
            return low

        mid = (low + high) // 2
        if arr[mid] < target:
            return self.search(arr, mid+1, high, target)
        elif target < arr[mid]:
            return self.search(arr, low, mid-1, target)
        else:
            return mid

    def sort(self, arr: List[int]) -> List[int]:
        for i, elm in enumerate(arr):
            if i == 0:
                continue
            idx = self.search(arr, 0, i-1, elm)
            arr = arr[:idx] + [elm] + arr[idx:i] + arr[i+1:]
        return arr

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
