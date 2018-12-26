def check(s1, s2):
  s1_backspace = 0
  s2_backspace = 0
  s1_ind = len(s1)-1
  s2_ind = len(s2)-1
  while (s1_ind >= 0 or s2_ind >= 0):
    s1_char = s1[s1_ind] if s1_ind >= 0 else ""
    s2_char = s2[s2_ind] if s2_ind >= 0 else ""
    # print("{}, {}".format(s1_char, s2_char))
    if s1_char == "#":
      s1_backspace += 1
    if s2_char == "#":
      s2_backspace += 1
    if s1_backspace > 0 or s2_backspace > 0:
      if s1_backspace > 0:
        s1_ind -= 1
        move = 1 if s1_char != "#" else 0
        s1_backspace -= move
      if s2_backspace > 0:
        s2_ind -= 1
        move = 0 if s2_char == "#" else 1
        s2_backspace -= move
    else:
      if s1_char != s2_char:
        return False
      else:
        s1_ind -= 1
        s2_ind -= 1
  return s1_ind < 0 and s2_ind < 0


print(True == check("a", "a"))
print(True == check("", ""))
print(True == check("ab", "ab"))
print(True == check("ab#", "a"))
print(True == check("ab#", "ab#"))
print(True == check("ab#", "ac#"))
print(True == check("a#", ""))
print(True == check("", "a#"))
print(True == check("abc#d##", "a"))
print(True == check("a", "abc#d##"))
print(True == check("abc#d##", "a"))
print(True == check("", "bc#d##"))
print(True == check("#", ""))
print(True == check("##", ""))
print(True == check("##", "#"))
print(True == check("abcdd##def", "##abcdeera###fhj##"))

import string
import random

def oracle(n):
  size = 10
  base = "ABCDEF" #string.ascii_uppercase
  for i in range(n):
    chars = base + "#"*(int(len(base)/5))
    # print(chars)
    s1 = ''.join(random.choice(chars) for _ in range(size))
    s2 = ''.join(random.choice(chars) for _ in range(size+2))
    assert(check(s1, s2) == check(s2, s1))
    ch = check(s1, s2)
    num = count_backspace(s1)
    # print(s1, num)
    if ch and num < size - num:
      return s1, s2
  return None

def count_backspace(s):
  count = 0
  for i in range(len(s)):
    if s[i] == "#":
      count += 1
  return count

ret = None
counter = 1
while ret == None:
  num = 100000
  print("testing {} on iteration {}".format(num, counter))
  counter += 1
  ret = oracle(num)
print(ret)