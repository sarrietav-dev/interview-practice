def print_tree(arr):
    """
    Prints a tree in a human readable format
    The array must be in BFS order
    """

    # The index of the first element in the next row
    next_row = 1

    # The index of the current element
    curr_index = 0

    # The number of elements in the array
    size = len(arr)

    # While there are still elements to print
    while curr_index < size:
        # Print the current element
        print(arr[curr_index], end=' ')

        # If we have reached the end of the current row
        if curr_index == next_row - 1:
            # Print a newline
            print()
            # Update the index of the first element in the next row
            next_row = next_row * 2 + 1

        # Update the current index
        curr_index += 1


tree = [1,2,2,3,None,None,3,4,None,None,4]

print_tree(tree)

