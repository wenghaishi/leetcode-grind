#  Find the Index of the First Occurrence in a String
def lc(needle, haystack):
    needle_length = len(needle)
    haystack_length = len(haystack)
    for i, char in enumerate(haystack):
        starting_index = i
        ending_index = i + needle_length

        if ending_index <= haystack_length:
            # Extract substring and compare with needle
            if haystack[starting_index:ending_index] == needle:
                return starting_index
    return -1



print(lc("efg", "avcdefgh"))