import json
import sys

def process_data(data):
    person = json.loads(data)
    results = []
    for i in range(3):
        obpersonj = {
            "id": person["id"] + str(i),
            "result1":"1",
            "result2":"2",
            "result3":"3",
            "result4":"4",
            "result5":"5",
            
        }
        results.append(obpersonj)
    return json.dumps(results)

if __name__ == "__main__":
    data = sys.argv[1]
    results = process_data(data)
    print(results)
