import numpy as np


def sum_array(array):
    result = 0
    for value in array:
        result += int(value)
    return result


def get_input_from_file(file_name):
    elves_calories_array = []
    input_array = []

    file = open(file_name, 'r')
    for line in file:
        if line == '\n':
            elves_calories_array.append(sum_array(input_array))
            input_array = []
        else:
            input_array.append(line)

    return elves_calories_array


if __name__ == '__main__':
    calories_array = get_input_from_file('Day01/day01_input.txt')
    # print('Most calories carried by an Elf: ' + str(np.amax(calories_array)))

    n1 = np.amax(calories_array)
    calories_array.remove(n1)

    n2 = np.amax(calories_array)
    calories_array.remove(n2)

    n3 = np.amax(calories_array)
    sum_top3_calories = n1 + n2 + n3
    print('Calories carried by top 3 elves: ' + str(sum_top3_calories))

