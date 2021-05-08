import unittest
from typing import List
import random


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
        target = random.choice(arr)
        target_cnt = 0

        for elm in arr:
            if elm < target:
                left.append(elm)
            elif target < elm:
                right.append(elm)
            else:
                target_cnt += 1

        left = self.sort(left)
        right = self.sort(right)

        return left + [target] * target_cnt + right

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
