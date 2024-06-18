def decode(message_file):
    # read file
    with open(message_file, 'r') as file:
        lines = file.readlines()
    numberWord = []
    for line in lines:
        numberWord.append((int(line.split()[0]), line.split()[1]))
    numberWord.sort(key = lambda x: x[0])
    
    # find number and elemets
    next_line_last_number = 1
    line_words = []
    decoded_message = ''
    
    for ele in numberWord:
        if ele[0] == next_line_last_number:
            line_words.append(ele[1])
            next_line_last_number += len(line_words) + 1  # find next line's last number
    
    decoded_message = ' '.join(line_words)

    return decoded_message

if __name__ == "__main__":
    file = "1.txt"
    print(decode(file))
