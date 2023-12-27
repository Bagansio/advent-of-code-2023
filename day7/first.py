

def read_file():
    try:
        with open('input.txt', 'r') as file:
            hands = []
            bids = []
            
            for line in file:
                line_parsed = line.strip().split()
                
                hands.append(line_parsed[0])
                bids.append(int(line_parsed[1]))
                
            return hands, bids        
    except FileNotFoundError:
        print(f"Error: File '{'input.txt'}' not found.")
    except Exception as e:
        print(f"Error: {e}")    
        


def main():
    hands, bids = read_file()
    print(hands, bids)

if __name__ == "__main__":
    main()
