def get_input_from_file(file_name):
    backpacks = []

    file = open(file_name, 'r')
    for line in file:
        remove_line_break_lenght = len(line) - 1
        line = line[:remove_line_break_lenght]
        backpacks.append(line)

    return backpacks


def sort_backpack_array(backpack_array):
    sorted_backpack_array = []

    for backpack in backpack_array:
        half_amount_items_in_backpack = int(len(backpack) / 2)
        sorted_backpack = [backpack[:half_amount_items_in_backpack], backpack[half_amount_items_in_backpack:]]
        sorted_backpack_array.append(sorted_backpack)

    return sorted_backpack_array


def find_duplicates_in_backpack(sorted_backpack):
    duplicates_array = []
    pouch_1 = list(sorted_backpack[0])

    for element in pouch_1:
        if element in sorted_backpack[1] and element not in duplicates_array:
            duplicates_array.append(element)

    return duplicates_array


def panic():
    print('Help :(')


def assign_priority_to_duplicates_in_backpack(duplicates_in_backpack):
    priority = 0
    for duplicate_item in duplicates_in_backpack:
        if duplicate_item.islower():
            priority += ord(duplicate_item) - 96
        elif duplicate_item.isupper():
            priority += ord(duplicate_item) - 38
        else:
            panic()

    return priority


def find_group_badge(group_backpacks):
    group_badges = []
    backpack_one_array = list(group_backpacks[0])

    for element in backpack_one_array:
        if element in group_backpacks[1] and element in group_backpacks[2] and element not in group_badges:
            group_badges.append(element)

    return group_badges


if __name__ == '__main__':
    backpack_array = get_input_from_file('day03_input.txt')
    summ_of_priorities = 0

    # Teil 1
    # sorted_backpack_array = sort_backpack_array(backpack_array)
    # duplicates_array = []
    # for backpack in sorted_backpack_array:
    #     duplicates_array.append(find_duplicates_in_backpack(backpack))
    #
    # for duplicates_in_backpack in duplicates_array:
    #     summ_of_priorities += assign_priority_to_duplicates_in_backpack(duplicates_in_backpack)

    # Teil 2
    badges_array = []
    group_backpacks = []
    for backpack in backpack_array:
        group_backpacks.append(backpack)
        if len(group_backpacks) == 3:
            group_badges = find_group_badge(group_backpacks)
            badges_array.append(group_badges)
            group_backpacks = []

    for badges in badges_array:
        summ_of_priorities += assign_priority_to_duplicates_in_backpack(badges)

    print('Summ of priorities of duplicate items: ' + str(summ_of_priorities))
