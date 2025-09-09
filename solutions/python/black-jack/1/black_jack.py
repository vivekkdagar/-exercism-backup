"""Functions to help play and score a game of blackjack.

How to play blackjack:    https://bicyclecards.com/how-to-play/blackjack/
"Standard" playing cards: https://en.wikipedia.org/wiki/Standard_52-card_deck
"""


def value_of_card(card):
    """Determine the scoring value of a card.
    
    :param card: str - given card.
    :return: int - value of a given card.  See below for values.

    1.  'J', 'Q', or 'K' (otherwise known as "face cards") = 10
    2.  'A' (ace card) = 1
    3.  '2' - '10' = numerical value.
    """
    if card in ['J','Q','K']:
        return 10
    elif card in 'A':
        return 1
    return int(card)


def higher_card(card_one, card_two):
    """Determine which card has a higher value in the hand.

    :param card_one, card_two: str - cards dealt in hand.  See below for values.
    :return: str or tuple - resulting Tuple contains both cards if they are of equal value.

    1.  'J', 'Q', or 'K' (otherwise known as "face cards") = 10
    2.  'A' (ace card) = 1
    3.  '2' - '10' = numerical value.
    """
    v1 = value_of_card(card_one)
    v2 = value_of_card(card_two)
    if v1 > v2:
        return card_one
    elif v2 > v1:
        return card_two
    return card_one, card_two
    
def value_of_ace(card_one, card_two):
    if 'A' in card_one or 'A' in card_two:
        return 1
    v1,v2 = value_of_card(card_one), value_of_card(card_two)
    if v1 + v2 + 11 <= 21:
        return 11
    return 1
    
def is_blackjack(card_one, card_two):
    v1, v2 = value_of_card(card_one), value_of_card(card_two)
    if card_one == 'A':
        v1 = 11
    if card_two == 'A':
        v2 = 11
    return v1 + v2 == 21
    
def can_split_pairs(card_one, card_two):
    return value_of_card(card_one) == value_of_card(card_two)

def can_double_down(card_one, card_two):
    v1, v2 = value_of_card(card_one), value_of_card(card_two)
    return 9 <= v1 + v2 <= 11