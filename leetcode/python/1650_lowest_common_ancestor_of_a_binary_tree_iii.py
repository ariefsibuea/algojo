"""
Problem          : Lowest Common Ancestor of a Binary Tree III
Topics           : Tree, Binary Tree
Level            : Medium
URL              : https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii
Description      :
Examples         :
                    Example 1:
                    Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
                    Output: 3
                    Explanation: The LCA of nodes 5 and 1 is 3.

                    Example 2:
                    Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
                    Output = 5
                    Explanation: The LCA of nodes 5 and 4 is 5 since a node can be a descendant of itself according to
                    the LCA definition.
"""


class NodeP1650:
    def __init__(self, val) -> None:
        self.val = val
        self.left = None
        self.right = None
        self.parent = None


class Solution:
    def lowest_common_ancestor_iii(self, p: NodeP1650, q: NodeP1650) -> NodeP1650:
        ptr1, ptr2 = p, q

        while ptr1 is not ptr2:
            ptr1 = ptr1.parent if ptr1.parent else q
            ptr2 = ptr2.parent if ptr2.parent else p

        return ptr1


def run_tests():
    construct_tree_p1650()

    inputs = {
        "case_1": [
            case1_p1650()[0],
            case1_p1650()[1],
        ],
        "case_2": [
            case2_p1650()[0],
            case2_p1650()[1],
        ],
    }
    outputs = {
        "case_1": 3,
        "case_2": 5,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.lowest_common_ancestor_iii(input[0], input[1])
        assert result.val == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


root = NodeP1650(3)
node2 = NodeP1650(5)
node3 = NodeP1650(1)
node4 = NodeP1650(6)
node5 = NodeP1650(2)
node6 = NodeP1650(0)
node7 = NodeP1650(8)
node8 = NodeP1650(7)
node9 = NodeP1650(4)


def construct_tree_p1650():
    root.left = node2
    root.right = node3

    node2.left = node4
    node2.right = node5
    node2.parent = root

    node3.left = node6
    node3.right = node7
    node3.parent = root

    node4.parent = node2

    node5.left = node8
    node5.right = node9
    node5.parent = node2

    node6.parent = node3
    node7.parent = node3
    node8.parent = node5
    node9.parent = node5


def case1_p1650() -> list[NodeP1650]:
    return [node2, node3]


def case2_p1650() -> list[NodeP1650]:
    return [node2, node9]


if __name__ == "__main__":
    run_tests()
