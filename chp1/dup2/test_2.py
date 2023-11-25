import random
import string
import time

print("Hello World")
print("Hello World")
print("Hello World")
print("Hello World")
print("Hello World")
print("Hello World")
print("Hello World")


class Person:
    def __init__(self, name):
        self.name = name

    def printName(self):
        print(self.name)

    def generate_random_text(self, num_lines, line_length):
        # Generate random text consisting of letters and digits
        random.seed(time.time_ns())
        result_lines = [
            ''.join(random.choices(string.ascii_letters +
                    string.digits, k=line_length))
            for _ in range(num_lines)
        ]
        return '\n'.join(result_lines)

    def write_random_text_to_file(self, file_name, num_lines, line_length):
        random_text = self.generate_random_text(num_lines, line_length)

        with open(file_name, 'w') as file:
            file.write(random_text)

    def getName(self):
        return self.name


persons = [
    Person("Harold"),
    Person("Perla"),
    Person("Finn"),
    Person("Poe"),
    Person("Chewy"),
]

for person in persons:
    person.write_random_text_to_file(f"{person.getName()}.txt", 676, 3)
