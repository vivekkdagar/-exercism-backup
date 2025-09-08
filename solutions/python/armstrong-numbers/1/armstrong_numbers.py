def is_armstrong_number(n):
    """
    Checks if a number is an Armstrong number without using an intermediate list.
    """
    # Handle non-positive numbers as they aren't Armstrong numbers by definition

    # Convert the number to a string to easily get the number of digits
    # This is a key optimization and a very Pythonic way to solve this
    str_n = str(n)
    num_digits = len(str_n)

    sum_of_powers = 0
    temp_n = n

    while temp_n > 0:
        digit = temp_n % 10
        sum_of_powers += digit ** num_digits
        temp_n //= 10

    return sum_of_powers == n
