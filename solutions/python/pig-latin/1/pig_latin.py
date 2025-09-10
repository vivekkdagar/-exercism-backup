def translate(text):
    if " " in text:
        text = text.split()
        return " ".join(translate(word) for word in text)
        
    vowels = ['a', 'e', 'i', 'o', 'u']
    flag = False
    
    if text.startswith(('a','e','i','o','u',"xr","yt")):
        return text + "ay"
    if text.startswith('y'):
        return text[1:] + 'yay'
    if text.startswith("qu"):
        return text[2:] + "quay"
        
    while text[0] not in vowels:
        if text[0] == 'q' or text[0] == 'y':
            flag = True
            break
        x = text[0]
        text = text[1:]
        text += x
    if text[0] == 'q':
        if text[1] == 'u':
            text = text[2:]
            text += "qu"
        else:
            text = text[1:] + "q"
    text += "ay"
    return text
    
    