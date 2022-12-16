my_legend = {
    'X': 'Rock',
    'Y': 'Paper',
    'Z': 'Scissors'
}

opponent_legend = {
    'A': 'Rock',
    'B': 'Paper',
    'C': 'Scissors'
}

play_to_point_dict = {
    'X': 1,
    'Y': 2,
    'Z': 3
}

win_combinations_dict = {
    'A': 'Y',
    'B': 'Z',
    'C': 'X'
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


def test_play_outcome(play_dict):
    global win_combinations_dict, my_legend, opponent_legend

    # Test if play ist Draw, then if it is a win and assign points
    if my_legend[play_dict['my_play']] == opponent_legend[play_dict['opponent_play']]:
        return 3
    elif win_combinations_dict[play_dict['opponent_play']] == play_dict['my_play']:
        return 6
    else:
        return 0


def calculate_score(play_string):
    global play_to_point_dict
    score = 0
    play_dict = make_play_readable(play_string)

    score += int(play_to_point_dict[play_dict['my_play']])
    score += test_play_outcome(play_dict)

    return score


if __name__ == '__main__':
    total_score = 0

    plays_array = get_input_from_file('day02_input.txt')
    for play in plays_array:
        total_score += calculate_score(play)

    print('Total score: ' + str(total_score))
