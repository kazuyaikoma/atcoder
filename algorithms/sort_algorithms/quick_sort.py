import unittest
from typing import List


class TestMergeSort(unittest.TestCase):
    """
    クイックソート

    時間計算量 O(n log n) (最悪時間計算量はO(n^2))
    空間計算量 O(log n)
    不安定ソート
    """
    def sort(self, arr: List[int]) -> List[int]:
        if len(arr) <= 1:
            return arr

        left, right = [], []
        pivot = arr[0]
        pivot_cnt = 0

        for elm in arr:
            if elm < pivot:
                left.append(elm)
            elif pivot < elm:
                right.append(elm)
            else:
                pivot_cnt += 1

        left = self.sort(left)
        right = self.sort(right)

        return left + [pivot] * pivot_cnt + right

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
