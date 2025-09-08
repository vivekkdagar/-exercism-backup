def isTriangle(sides):
    if sides[0] + sides[1] < sides[2]:
        return False
    if sides[1] + sides[2] < sides[0]:
        return False
    if sides[0] + sides[2] < sides[1]:
        return False
    return True

def equilateral(sides):
    if len(list(set(sides))) == 1 and sides[0] != 0:
        return True
    return False and isTriangle(sides)

def isosceles(sides):
    return len(set(sides)) < 3 and isTriangle(sides)

def scalene(sides):
    return len(set(sides)) == len(sides) and isTriangle(sides)
