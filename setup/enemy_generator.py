"""Generate a SQL file full of enemies to populate the database with"""

import random
import sys


def main():
    """Where the magic happens"""
    target = open("enemies.sql", "w")

    for i in range(int(sys.argv[1])):
        target.write("insert into enemies (latitude, longitude) values ('" + str(random.randint(-179, 179)) + "." + str(random.randint(0, 100000)) + "', '" + str(random.randint(-179, 179)) + "." + str(random.randint(0, 100000)) + "');\n")

    target.close()
main()  # Call the function so the program runs
