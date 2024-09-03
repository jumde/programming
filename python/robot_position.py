def getCurrentPosition(moves):
    position = [0, 0]
    for move in moves:
        if move[0] == "UP":
            position[0] += move[1]
        elif move[0] == "DOWN":
            position[0] -= move[1]
        elif move[0] == "RIGHT":
            position[1] += move[1]
        elif move[0] == "LEFT":
            position[1] -= move[1]
        else:
            print("Incorrect move")
    return position

print(getCurrentPosition([["UP", 5], ["DOWN", 3], ["RIGHT", 5], ["LEFT", 3]]))
