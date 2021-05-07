import unittest
from typing import List


class TestSelectSort(unittest.TestCase):
    # 選択ソート
    def sort(self, arr: List[int]) -> List[int]:
        for i in range(0, len(arr)-1):
            # そのループで最小値をとるindex
            min_idx = i
            for j in range(i, len(arr)):
                if arr[j] < arr[min_idx]:
                    min_idx = j

            arr[i], arr[min_idx] = arr[min_idx], arr[i]

        return arr

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
