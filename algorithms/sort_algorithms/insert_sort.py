import unittest
from typing import List


class TestInsertSort(unittest.TestCase):
    """
    挿入ソート
    先頭から順に舐めていき、ソート済み配列の良しなな場所に挿入していく

    時間計算量 O(N^2)
    空間計算量 O(1)
    """
    def sort(self, arr: List[int]) -> List[int]:
        for i in range(1, len(arr)):
            elm = arr[i]

            j = i - 1
            while elm < arr[j] and 0 <= j:
                # whileブロックが実行されるということは、ソートされるということなので
                # elmを入れるスペースを開けるために1つずつ位置をずらしていく
                arr[j+1] = arr[j]
                j -= 1

            # while後、min_idx確定時に挿入
            arr[j+1] = elm

        return arr

    def test_sort(self):
        arr = [4, 3, 2, 5, 1]
        arr = self.sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])


if __name__ == '__main__':
    unittest.main()
