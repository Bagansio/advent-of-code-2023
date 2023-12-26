

def read_file():
    try:
        with open('input.txt', 'r') as file:
            unparsed_times = file.readline().strip()
            unparsed_distances = file.readline().strip()
            
            times = parse_values(unparsed_times)
            distances = parse_values(unparsed_distances)
            return times, distances       
    except FileNotFoundError:
        print(f"Error: File '{'input.txt'}' not found.")
    except Exception as e:
        print(f"Error: {e}")    
        

def parse_values(line):
    elements = line.split()
    return [int(element) for element in elements[1:]]

def calculate_possiblities(time, distance):
    for time_pushing in range(time):
        time_running = time - time_pushing
        
        if time_running * time_pushing > distance:
            return time_running - time_pushing + 1


def main():
    times, distances = read_file()
    
    result = 1
    
    for time, distance in zip(times,distances):
        possiblities = calculate_possiblities(time, distance)
        
        result *= possiblities
        print(result, possiblities)
    print(times,distances, result)

if __name__ == "__main__":
    main()
