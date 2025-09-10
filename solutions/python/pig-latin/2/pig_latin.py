def translate(text):
    if " " in text:
        text = text.split()
        return " ".join(translate(word) for word in text)

    moved = False
    
    vowels = {'a', 'e', 'i', 'o', 'u'}

    if text.startswith(("xr", "yt")):
         return text + "ay"

    while text[0] not in vowels:
        if text[0] in 'qy':
            break
        x = text[0]
        text = text[1:] + x
        moved = True

    if text[0] == 'q':
        if text[1] == 'u':
            text = text[2:] + "qu"
        else:
            text = text[1:] + "q"
    if text[0] == 'y' and not moved:
        return text[1:] + 'yay'
    text += "ay"
    return text