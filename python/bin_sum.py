def padLeft(val, count):
    for i in range(count):
        val = '0' + val
    return val

def binSum(a, b):
    if len(a) > len(b):
        b = padLeft(b, len(a) - len(b))
    else:
        a = padLeft(a, len(b) - len(a))

    cur_len = len(a) 
    current_val = 0
    reverse_binary_sum = ""

    i = 0
    for i in range(cur_len):
      parsed_index = -(i + 1)
      if i != len(a) and a[parsed_index] == "1":
          current_val += 1
      if i != len(b) and b[parsed_index] == "1":
          current_val += 1

      if current_val % 2 == 1:
          reverse_binary_sum += "1"
      else:
          reverse_binary_sum += "0"
      
      current_val = int(current_val / 2)

    binary_sum = reverse_binary_sum[::-1]

    return binary_sum


print(binSum("0000", "1"))
print(binSum("0000", "10"))
print(binSum("010101", "10"))
print(binSum("010111", "10"))
    

      
