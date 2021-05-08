import unittest
from typing import List


class TestMergeSort(unittest.TestCase):
    """
    マージソート

    時間計算量 O(n log n)
    空間計算量 O(n)
    安定ソート
    """
    def sort(self, arr: List[int]) -> List[int]:
        if len(arr) <= 1:
            return arr

        mid = len(arr) // 2
        left_arr = arr[:mid]
        right_arr = arr[mid:]

        left_arr = self.sort(left_arr)
        right_arr = self.sort(right_arr)

        return self.merge(left_arr, right_arr)

    def merge(self, left_arr: List[int], right_arr: List[int]) -> List[int]:
        merged = []
        left, right = 0, 0

        while right < len(right_arr) and left < len(left_arr):
            if left_arr[left] < right_arr[right]:
                merged.append(left_arr[left])
                left += 1
            else:
                merged.append(right_arr[right])
                right += 1

        merged.extend(left_arr[left:])
        merged.extend(right_arr[right:])

        return merged

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
