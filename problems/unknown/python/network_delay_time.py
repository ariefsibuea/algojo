from typing import List


class NetworkDelayTime:
    def solve_by_bellman_ford(self, times: List[List[int]], n: int, k: int) -> int:
        """Finds the time it takes for all nodes to receive a signal using Bellman-Ford algorithm.

        This solution uses the Bellman-Ford algorithm to find the shortest paths from a source node
        to all other nodes in a weighted directed graph. The algorithm relaxes all edges for n-1
        iterations, where n is the number of nodes.

        Args:
            times: List of edges [u,v,w] where u is the source node, v is the target node, and w is the time it takes
                    for a signal to travel from u to v.
            n: The number of nodes in the network (labeled from 1 to n).
            k: The starting node to send the signal from.

        Returns:
            The time it takes for all nodes to receive the signal, or -1 if it's impossible.

        Time Complexity:
            O(E * V): Where E is the number of edges and V is the number of vertices.

        Space Complexity:
            O(V): For the distances array.
        """
        distances = [float("inf")] * (n + 1)
        distances[k] = 0
        distances[0] = 0  # we don't use this index since the node start from 1

        for _ in range(n - 1):
            for u, v, w in times:
                if distances[u] != float("inf") and distances[u] + w < distances[v]:
                    distances[v] = distances[u] + w

        max_distance = max(distances[1:])

        return max_distance if max_distance < float("inf") else -1


if __name__ == "__main__":
    inputs = {
        "case_1": [[[2, 1, 1], [2, 3, 1], [3, 4, 1]], 4, 2],
        "case_2": [[[1, 2, 1]], 2, 1],
        "case_3": [[[1, 2, 1]], 2, 2],
    }

    outputs = {
        "case_1": 2,
        "case_2": 1,
        "case_3": -1,
    }

    solution = NetworkDelayTime()

    for case, input in inputs.items():
        result = solution.solve_by_bellman_ford(input[0], input[1], input[2])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
