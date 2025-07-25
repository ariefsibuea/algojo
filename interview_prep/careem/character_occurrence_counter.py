class CharacterCounter:
    def __init__(self):
        self.cache = {}

    def linear_scan(self, text: str, char: str) -> int:
        """Returns number of occurrence of specified character in a given string.

        Traverse the entire string character by character, incrementing a counter each time the 'char' was found.

        Args:
            text (str): The string to search in.
            char (str): The character to count occurrences of.

        Returns:
            int: Number of occurrence of 'char' in 'text'.

        Raises:
            ValueError: If 'text' or 'char' is empty.
            ValueError: If 'char' length is not 1.

        Time Complexity:
            O(n): Traverse all characters in the specified 'text'.

        Space Complexity:
            O(1): Only need to define 'counter' as additional variable.
        """
        if not text or not char:
            raise ValueError("text and char should be exist")
        elif len(char) != 1:
            raise ValueError("char length must be 1")

        counter = 0
        for c in text:
            if c == char:
                counter += 1

        return counter

    def caching_layer(self, text: str, char: str) -> int:
        """Returns number of occurrence of specified character in a given string.

        Initialize a frequenct map, traverse the entire string character by character, store the character as the key in
        the 'freq_map' with its count as value. In the final, store the 'freq_map' into 'cache'.

        Args:
            text (str): The string to search in.
            char (str): The character to count occurrences of.

        Returns:
            int: Number of occurrence of 'char' in 'text'.

        Raises:
            ValueError: If 'text' or 'char' is empty.
            ValueError: If 'char' length is not 1.

        Time Complexity:
            O(n): Traverse all characters in the specified 'text'.

        Space Complexity:
            O( m x k): m is the number of text and k is the number of characters.
        """
        if not text or not char:
            raise ValueError("text and char should be exist")
        elif len(char) != 1:
            raise ValueError("char length must be 1")

        if text not in self.cache:
            freq_map = {}
            for c in text:
                freq_map[c] = freq_map.get(c, 0) + 1
            self.cache[text] = freq_map

        return self.cache[text].get(char, 0)


if __name__ == "__main__":
    inputs = {
        "case_1": ["hello world", "o"],
        "case_2": ["programming", "g"],
        "case_3": ["test", "x"],
    }

    outputs = {
        "case_1": 2,
        "case_2": 2,
        "case_3": 0,
    }

    solution = CharacterCounter()

    for case, input in inputs.items():
        result = solution.caching_layer(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
