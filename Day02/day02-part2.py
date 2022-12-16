my_legend = {
    'X': 0,     # loose
    'Y': 3,     # draw
    'Z': 6      # win
}

draw_dict = {
    'A': 1,
    'B': 2,
    'C': 3
}

win_dict = {
    'A': 2,
    'B': 3,
    'C': 1
}

loose_dict = {
    'A': 3,
    'B': 1,
    'C': 2
}


def get_input_from_file(file_name):
    plays = []

    file = open(file_name, 'r')
    for line in file:
        plays.append(line)

    return plays


def make_play_readable(play_string):
    return {
        'opponent_play': str(play_string[0]),
        'my_play': str(play_string[2])
    }


def calculate_score(play_string):
    global win_dict, loose_dict, draw_dict
    play_dict = make_play_readable(play_string)

    if play_dict['my_play'] == 'X':                         # Loose
        score = loose_dict[play_dict['opponent_play']]
    elif play_dict['my_play'] == 'Y':                       # Draw
        score = 3 + draw_dict[play_dict['opponent_play']]
    else:                                                   # Win
        score = 6 + win_dict[play_dict['opponent_play']]

    return score


if __name__ == '__main__':
    total_score = 0

    plays_array = get_input_from_file('day02_input.txt')
    for play in plays_array:
        total_score += calculate_score(play)

    print('Total score: ' + str(total_score))
