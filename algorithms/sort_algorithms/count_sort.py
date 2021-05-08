import unittest
from typing import List


class TestCountSort(unittest.TestCase):
    """
    カウントソート(計数ソート)

    時間計算量 O(nk) (最悪時間計算量はO(n^2 * k))
    空間計算量 O(nk)
    不安定ソート
    """
    def sort(self, arr: List[int]) -> List[int]:
        min_num = min(arr)
        max_num = max(arr)

        counts = [0] * (max_num - min_num + 1)
        for elm in arr:
            counts[elm-min_num] += 1

        arr = []
        for i, count in enumerate(counts):
            arr += [i+min_num] * count

        return arr

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
