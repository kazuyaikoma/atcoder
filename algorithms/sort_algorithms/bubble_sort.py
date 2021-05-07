import unittest
from typing import List


class TestBubbleSort(unittest.TestCase):
    """
    バブルソート

    時間計算量 O(N^2)
    空間計算量 O(1)
    安定ソート
    """
    def sort(self, arr: List[int]) -> List[int]:
        swapped = True
        # swapが起きている間はひたすらswap処理のループを回す
        while swapped:
            swapped = False
            for i in range(0, len(arr)-1):
                if arr[i+1] < arr[i]:
                    arr[i+1], arr[i] = arr[i], arr[i+1]
                    # もしswapが起きたらswappedフラグをTrueに切り替え
                    swapped = True

        return arr

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
