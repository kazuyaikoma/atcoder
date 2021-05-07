import unittest
from typing import List


class TestSelectSort(unittest.TestCase):
    """
    選択ソート
    ループごとに最小のものを選ぶ→先頭と随時交換していく

    時間計算量 O(N^2)
    空間計算量 O(1)
    """
    def sort(self, arr: List[int]) -> List[int]:
        for i, elm in enumerate(arr):
            min_idx = i
            for j in range(i, len(arr)):
                if arr[j] < arr[min_idx]:
                    min_idx = j
            # 空間計算量 O(1)とするため、新たにリストを作らずswap
            arr[i], arr[min_idx] = arr[min_idx], elm

        return arr

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
