# A simple script that extract the wanted information about a pokemon through pokeapi.co

import requests
from pprint import pprint

def extract_info(data):
    main_t = "Null"
    sub_t = "Null"
    first_ability = "Null"
    second_ability = "Null"
    hidden_ability = "Null"
    
    for i, typ in enumerate(data["types"]):
        if main_t == "Null":
            main_t = data["types"][i]["type"]["name"]
        else:
            sub_t = data["types"][i]["type"]["name"]
    
    for i, ability in enumerate(data["abilities"]):
        if data["abilities"][i]["is_hidden"] == True:
            hidden_ability = data["abilities"][i]["ability"]["name"]
        else:
            if first_ability == "Null":
                first_ability = data["abilities"][i]["ability"]["name"]
            else:
                second_ability = data["abilities"][i]["ability"]["name"]
            
    poke_stats = {
        "id": data["id"],
        "name": data["name"],
        "img": data["sprites"]["other"]["official-artwork"]["front_default"],
        "main_t": main_t,
        "sub_t": sub_t,
        "first_ability":first_ability,
        "second_ability": second_ability,
        "hidden_ability": hidden_ability,
        "hp": data["stats"][0]["base_stat"],
        "atk": data["stats"][1]["base_stat"],
        "def": data["stats"][2]["base_stat"],
        "sp_atk": data["stats"][3]["base_stat"],
        "sp_def": data["stats"][4]["base_stat"],
        "spd": data["stats"][5]["base_stat"]
    }
    return poke_stats
    
    
def main():
    LAST_POKE = 1025
    f = open("./db/poke_database.csv", "w")
    
    for i in range(1,LAST_POKE):
        progress = (i / LAST_POKE) * 100
        print(f"Progress: {progress:.2f}%")
    
        url = "https://pokeapi.co/api/v2/pokemon/" + str(i)
        response = requests.get(url)
        
        if response.status_code == 200:
            data = response.json()
            poke_stats = extract_info(data)
            result = ",".join(f"{value}" for value in poke_stats.values())
            # pprint(f"{result}")
            f.write(result + "\n")
        else:
            print("Failed to fetch data:", response.status_code)
            
    f.close()
    
if __name__ == "__main__":
    main()