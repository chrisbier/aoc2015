#!/usr/bin/python3
# Day 09 part 1

inputfile = open('input1', 'r')
input_list = inputfile.read()

counter = 0

for line in input_list:
    if line == "(":
        counter += 1
    if line == ")":
        counter -= 1
    print(line, counter)

print(counter)
