# A simple script that extract the wanted information about a pokemon through pokeapi.co

import requests

poke_gen = {
    1: [1,151],
    2: [152,251],
    3: [252,386],
    4: [387,493],
    5: [494,649],
    6: [650,721],
    7: [722,809],
    8: [810,905],
    9: [906,1025]
}

def find_starters():
    gen_num = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    starters = []

    for gen in gen_num:
        start, _ = poke_gen[gen]
        starters.extend(range(start, start + 9)) # victini is added warning
        
    return starters

poke_groups = {
    "starters": [
        25, 26, 172, #  Pikachu evoline
        133, 134, 135, 136, 196, 197, 470, 471, 700], # Eevee evo-line
    "pseudo-legendary": [
        149, 148, 147,  # Dragonite -> Dragonair -> Dratini
        248, 247, 246,  # Tyranitar -> Pupitar -> Larvitar
        373, 372, 371,  # Salamence -> Shelgon -> Bagon
        376, 375, 374,  # Metagross -> Metang -> Beldum
        445, 444, 443,  # Garchomp -> Gabite -> Gible
        635, 634, 633,  # Hydreigon -> Zweilous -> Deino
        706, 705, 704,  # Goodra -> Sliggoo -> Goomy
        784, 783, 782,  # Kommo-o -> Hakamo-o -> Jangmo-o
        887, 886, 885,  # Dragapult -> Drakloak -> Dreepy
        998, 997, 996  # Baxcalibur -> Arctibax -> Frigibax]
    ],
    "sub-legendary": [144, 145, 146, 243, 244, 245, 377, 378, 379, 380,
                      381, 480, 481, 482, 485, 486, 488, 638, 639, 640, 
                      641, 642, 645, 772, 773, 785, 786, 787, 788, 891, 
                      892, 894, 895, 896, 897, 905, 1001, 1002, 1003, 
                      1004, 1014, 1015, 1016, 793, 794, 795, 796, 797, 
                      798, 799, 803, 804, 805, 806], 
    "legendary": [150, 249, 250, 382, 383, 384, 483, 484, 487, 643, 644, 646, 716, 717, 718, 790, 791, 789, 792, 800, 888, 889, 890, 898, 1007, 1008, 1024],
    "mythical": [151, 251, 385, 386, 489, 490, 491, 492, 493, 494, 647, 648, 649, 719, 720, 721, 801, 802, 807, 808, 809, 893, 1025]
}

def find_generation(id):
     for gen, (start, end) in poke_gen.items():
        if id in range(start, end + 1):
            return gen

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
        "spd": data["stats"][5]["base_stat"],
        "gen": find_generation(data["id"]),
        "isStarter": int(data["id"] in poke_groups["starters"]),
        "isPseudo": int(data["id"] in poke_groups["pseudo-legendary"]),
        "isSubLegend": int(data["id"] in poke_groups["sub-legendary"]),
        "isLegend": int(data["id"] in poke_groups["legendary"]),
        "isMyth": int(data["id"] in poke_groups["mythical"])
    }
    return poke_stats
    
def downloadImg(url, num):
    response = requests.get(url)
    if response.status_code == 200:
        with open("assets/images/" + str(num)+'.png', 'wb') as file:
                file.write(response.content)

def createTables():
    LAST_POKE = 1025
    f = open("./db/poke_database.csv", "w")
    
    for i in range(1,LAST_POKE+1):
        progress = (i / LAST_POKE) * 100
        print(f"Progress: {progress:.2f}%")
    
        url = "https://pokeapi.co/api/v2/pokemon/" + str(i)
        response = requests.get(url)
        
        if response.status_code == 200:
            data = response.json()
            poke_stats = extract_info(data)
            downloadImg(poke_stats["img"], i)
            result = ",".join(f"{value}" for value in poke_stats.values())
            f.write(result + "\n")
        else:
            print("Failed to fetch data:", response.status_code)
            
    f.close()
      
def main():
    poke_groups["starters"].extend(find_starters())
    # print(poke_groups)
    # print(int(1 in poke_groups["starters"]))
    createTables()
    
if __name__ == "__main__":
    main()