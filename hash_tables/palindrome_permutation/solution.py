def palindrome_permutation(s):
    """
    Given a string, write a function to check if it is a permutation of a
    palindrome. A palindrome is a word or phrase that is the same forwards and
    backwards. A permutation is a rearrangement of letters. The palindrome does
    not need to be limited to just dictionary words.
    """

    # Remove spaces and lowercase
    s = s.replace(' ', '').lower()

    # Count the number of times each character appears in the string
    counts = {}
    for c in s:
        if c not in counts:
            counts[c] = 0
        counts[c] += 1

    # A palindrome can have at most one character that appears an odd number of
    # times. If there are more than one, then the string is not a permutation
    # of a palindrome.
    odd_counts = 0
    for count in counts.values():
        if count % 2 == 1:
            odd_counts += 1
        if odd_counts > 1:
            return False

    return True