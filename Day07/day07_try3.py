class Element:
    def __init__(self, name: str, size: int, is_file: bool, parent, children: {}):
        self.name: str = name
        self.size: int = size
        self.is_file: bool = is_file
        self.parent: Element = parent
        self.children: {} = children

    def add_dir_child(self, child_name):
        self.children[child_name] = Element(child_name, 0, False, self, {})

    def add_file_child(self, child_name, size):
        self.children[child_name] = Element(child_name, size, True, self, {})

    def set_size(self, size: int):
        self.size = size


def get_input_from_file(file_name):
    lines = []

    file = open(file_name, 'r')
    for line in file:
        remove_line_break_length = len(line) - 1
        line = line[:remove_line_break_length]
        lines.append(line.split())

    return lines


def calculate_dir_size(directory: Element):
    global total_size
    size = 0

    for child in directory.children:
        if directory.children[child].is_file:
            size += int(directory.children[child].size)
        else:
            dir_size = calculate_dir_size(directory.children[child])
            size += dir_size

    directory.size = size
    return size


def find_closest_directory(folder: Element):
    global needed_space
    global closest_dir_size
    global closest_difference

    for child in folder.children:
        if not folder.children[child].is_file:
            folder_size = folder.children[child].size
            if folder_size >= needed_space and closest_difference > folder_size - needed_space:
                closest_dir_size = folder.children[child].size
                closest_difference = folder.children[child].size - needed_space
            find_closest_directory(folder.children[child])

    return closest_dir_size


def print_file_tree(folder: Element):
    global indent
    print(indent + '- ' + folder.name + ' (folder)')
    indent += '   '

    for child in folder.children:
        if folder.children[child].is_file:
            print(indent + '- ' + folder.children[child].name + ' (file)')
        else:
            print_file_tree(folder.children[child])

    indent = indent[:-3]


if __name__ == '__main__':
    input_lines = get_input_from_file('day07_input.txt')
    total_size = 0

    current_folder = Element('/', 0, False, None, {})

    for line in input_lines:
        if line[1] == 'ls':
            continue
        elif line[0] == '$' and line[1] == 'cd':
            if line[2] == '..':
                current_folder = current_folder.parent
            else:
                if line[2] == '/':
                    continue
                current_folder = current_folder.children[line[2]]
        elif line[0] == 'dir':
            current_folder.add_dir_child(line[1])
        else:
            current_folder.add_file_child(line[1], line[0])

    while current_folder.name != '/':
        current_folder = current_folder.parent

    calculate_dir_size(current_folder)

    print(current_folder.name)
    print(current_folder.size)
    needed_space = 30_000_000 - (70_000_000 - current_folder.size)
    closest_dir_size = 0
    closest_difference = 70_000_000

    find_closest_directory(current_folder)

    print('Result: ' + str(closest_dir_size))
    print('Difference: ' + str(closest_difference))

    # indent = ''
    # print_file_tree(current_folder)
