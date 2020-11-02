#!/usr/bin/python

import getopt, os, sys

def countLine(arg):
    num_words = 0
    num_lines = 0
    num_spaces = 0
    num_chars = 0

    print "Counting in: " + arg
    if not os.path.exists(arg):
        print "File not found"
        exit(1)

    with open(arg) as fp:
        for line in fp:
            words = line.split()
            num_words += len(words)
            num_lines += 1
            num_chars += len(line)
            num_spaces += line.count(" ")
        print "Number of lines: " + str(num_lines)
        print "Number of spaces: " + str(num_spaces)
        print "Number of characters: " + str(num_chars)
        print "Number of words: " + str(num_words)
        
def main(argv):
    try:
        opts, args = getopt.getopt(argv,"hi:",["input="])
    except getopt.GetoptError:
        print 'file.py -i <inputfile>'
        sys.exit(2)

    for opt, arg in opts:
      if opt == '-h':
         print 'file.py -i <inputfile>'
         sys.exit()
      elif opt in ("-i", "--input"):
         countLine(arg)

if __name__ == "__main__":
    main(sys.argv[1:])
