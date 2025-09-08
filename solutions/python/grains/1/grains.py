def square(number):
    """
    Calculate the number of grains on a specific square of a chessboard.

    Args:
        number (int): The square number, must be between 1 and 64 (inclusive).

    Returns:
        int: The number of grains on the specified square.

    Raises:
        ValueError: If the square number is not within the valid range (1 to 64).
    """
    # Check if the number is within the valid range for a chessboard.
    if not 1 <= number <= 64:
        raise ValueError("square must be between 1 and 64")

    # The number of grains on any given square 'n' follows the formula 2**(n-1).
    # For example, square 1 has 2**(1-1) = 2**0 = 1 grain.
    # Square 2 has 2**(2-1) = 2**1 = 2 grains.
    return 2 ** (number - 1)


def total():
    """
    Calculate the total number of grains on the entire chessboard.

    Returns:
        int: The total number of grains on all 64 squares.
    """
    # The total number of grains is the sum of a geometric series: 1 + 2 + 4 + ... + 2**63.
    # The sum of this series is calculated as 2**64 - 1.
    # This value is a very large integer.
    return 2 ** 64 - 1
