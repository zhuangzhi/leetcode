import unittest
import le1
class TestPromeMethods(unittest.TestCase):
    def setUp(self):
        self.solution = le1.Solution()

    def tearDown(self):
        del(self.solution)
        pass
    def testTwoSum(self):
        units = [
        ([2, 7, 11, 15], 9, [0, 1]),
		([3, 2, 4], 6, [1,2]),
        ]
        for _, unit in enumerate(units):
            self.assertEqual(unit[2], self.solution.twoSum(unit[0], unit[1]))
        for _, unit in enumerate(units):
            self.assertEqual(unit[2], self.solution.twoSumDoublePointer(unit[0], unit[1]))
if __name__ == '__main__':
    unittest.main()
