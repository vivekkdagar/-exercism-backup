"""Functions to prevent a nuclear meltdown."""


def is_criticality_balanced(temperature, neutrons_emitted):
    """Verify criticality is balanced.

    :param temperature: int or float - temperature value in kelvin.
    :param neutrons_emitted: int or float - number of neutrons emitted per second.
    :return: bool - is criticality balanced?

    A reactor is said to be critical if it satisfies the following conditions:
    - The temperature is less than 800 K.
    - The number of neutrons emitted per second is greater than 500.
    - The product of temperature and neutrons emitted per second is less than 500000.
    """
    return (
        temperature < 800
        and neutrons_emitted > 500
        and temperature * neutrons_emitted < 500_000
    )


def reactor_efficiency(voltage, current, theoretical_max_power):
    """Assess reactor efficiency zone.

    :param voltage: int or float - voltage value.
    :param current: int or float - current value.
    :param theoretical_max_power: int or float - power that corresponds to a 100% efficiency.
    :return: str - one of ('green', 'orange', 'red', 'black').

    Efficiency bands:
    - 'green'  -> efficiency of 80% or more,
    - 'orange' -> efficiency of at least 60% but less than 80%,
    - 'red'    -> efficiency of at least 30% but less than 60%,
    - 'black'  -> efficiency less than 30%.

    Efficiency = (generated power / theoretical max power) * 100,
    where generated power = voltage * current.
    """
    generated_power = voltage * current
    efficiency = (generated_power / theoretical_max_power) * 100

    if efficiency >= 80:
        return "green"
    if efficiency >= 60:
        return "orange"
    if efficiency >= 30:
        return "red"
    return "black"


def fail_safe(temperature, neutrons_produced_per_second, threshold):
    """Assess and return status code for the reactor.

    :param temperature: int or float - value of the temperature in kelvin.
    :param neutrons_produced_per_second: int or float - neutron flux.
    :param threshold: int or float - threshold for category.
    :return: str - one of ('LOW', 'NORMAL', 'DANGER').

    - 'LOW'    -> power < 90% of threshold
    - 'NORMAL' -> 90% of threshold <= power <= 110% of threshold
    - 'DANGER' -> power outside of the above ranges
    """
    power = temperature * neutrons_produced_per_second

    if power < 0.9 * threshold:
        return "LOW"
    if power <= 1.1 * threshold:
        return "NORMAL"
    return "DANGER"