

def read_file():
    try:
        with open('input.txt', 'r') as file:
            unparsed_times = file.readline().strip()
            unparsed_distances = file.readline().strip()
            
            time = parse_values(unparsed_times)
            distance = parse_values(unparsed_distances)
            return time, distance      
    except FileNotFoundError:
        print(f"Error: File '{'input.txt'}' not found.")
    except Exception as e:
        print(f"Error: {e}")    
        

def parse_values(line):
    elements = line.split()
    return int(''.join(elements[1:]))

def calculate_possiblities(time, distance):
    for time_pushing in range(time):
        time_running = time - time_pushing
        
        if time_running * time_pushing > distance:
            return time_running - time_pushing + 1

def main():
    time, distance = read_file()    

    possiblities = calculate_possiblities(time, distance)
        
    print(possiblities, time, distance)

if __name__ == "__main__":
    main()
