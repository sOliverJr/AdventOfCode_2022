class Element:
    def __init__(self, name: str, size: int, is_file: bool, parent, children: {}):
        self.name: str = name
        self.size: int = size
        self.is_file: bool = is_file
        self.parent: Element = parent
        self.children: {} = children

    def add_child(self, child_name):
        self.children[child_name] = Element(child_name, 0, False, self, {})

    def set_size(self, size: int):
        self.size = size


current_folder = Element('/', 0, False, None, {})
current_folder.set_size(50)

current_folder.add_child('subfolder')

current_folder = current_folder.children['subfolder']
current_folder.set_size(25)

print(current_folder.name + ' ' + str(current_folder.size))

current_folder = current_folder.parent
print(current_folder.name + ' ' + str(current_folder.size))
