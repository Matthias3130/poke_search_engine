// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func SearchBar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"search-bar\"><input type=\"search\" id=\"search\" pattern=\"[A-Za-z\\s\\-]*\" placeholder=\"Search...\" value=\"\" onkeyup=\"\"> <button type=\"button\" onclick=\"toggleFilter()\"><i class=\"material-icons\" id=\"icon-sort\">filter_list</i></button> <button type=\"button\" onclick=\"toggleOrder()\"><i class=\"material-icons\" id=\"icon-sort\">swap_vert</i></button></div><div id=\"filter\"><div id=\"filter-type\"><h3>Pokemon Type</h3><label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"normal\" value=\"normal\">Normal</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"fire\" value=\"fire\">Fire</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"water\" value=\"water\">Water</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"electric\" value=\"electric\">Electric</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"grass\" value=\"grass\">Grass</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"ice\" value=\"ice\">Ice</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"fighting\" value=\"fighting\">Fighting</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"poison\" value=\"poison\">Poison</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"ground\" value=\"ground\">Ground</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"flying\" value=\"flying\">Flying</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"psychic\" value=\"psychic\">Psychic</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"bug\" value=\"bug\">Bug</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"rock\" value=\"rock\">Rock</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"ghost\" value=\"ghost\">Ghost</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"dragon\" value=\"dragon\">Dragon</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"dark\" value=\"dark\">Dark</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"steel\" value=\"steel\">Steel</label> <label><input type=\"checkbox\" class=\"poke-type checkbox\" name=\"fairy\" value=\"fairy\">Fairy</label></div><div id=\"filter-generation\"><h3>Region</h3><label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"kanto\" value=\"1\">Kanto</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"johto\" value=\"2\">Johto</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"hoenn\" value=\"3\">Hoenn</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"sinnoh\" value=\"4\">Sinnoh</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"unova\" value=\"5\">Unova</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"kalos\" value=\"6\">Kalos</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"alola\" value=\"7\">Alola</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"galar\" value=\"8\">Galar</label> <label><input type=\"checkbox\" class=\"poke-gen checkbox\" name=\"paldea\" value=\"9\">Paldea</label></div><div id=\"filter-group\"><h3>Tags</h3><label><input type=\"checkbox\" class=\"poke-group\" name=\"starters\" id=\"starters\" value=\"starters\">Starters</label> <label><input type=\"checkbox\" class=\"poke-group\" name=\"pseudo-legendary\" id=\"pseudo-legendary\" value=\"pseudo-legendary\">Pseudo-Legendary</label> <label><input type=\"checkbox\" class=\"poke-group\" name=\"sub-legendary\" id=\"sub-legendary\" value=\"sub-legendary\">Sub-Legendary</label> <label><input type=\"checkbox\" class=\"poke-group\" name=\"legendary\" id=\"legendary\" value=\"legendary\">Legendary</label> <label><input type=\"checkbox\" class=\"poke-group\" name=\"mythical\" id=\"mythical\" value=\"mythical\">Mythical</label></div></div><table id=\"table\"><thead><tr><th>ID</th><th>Name</th><th>Image</th><th>Type</th><th>Abilities</th><th>HP</th><th>Atk</th><th>Def</th><th>Sp_Atk</th><th>Sp_Def</th><th>Spd</th><th>Tags</th></tr></thead> <tbody id=\"table-body\"></tbody></table><script>\ndocument.addEventListener(\"DOMContentLoaded\", () => {\n    getAllPokemon();\n});\n\n let isReverse = false;\n let isFilterVisible = false;\n\nconst search = document.getElementById(\"search\");\nsearch.addEventListener('input', function () {\n  console.log('Input value:', search.value); // Log input value\n  findPokemon(search.value);\n});\n\nconst checkboxes = document.querySelectorAll('.checkbox');\ncheckboxes.forEach(checkbox => {\n    checkbox.addEventListener('change', () => {\n        console.log(`${checkbox.value} is ${checkbox.checked ? 'checked' : 'unchecked'}`);\n        const userInput = document.getElementById(\"search\").value;\n        findPokemon(userInput);\n    });\n});\n\nconst pokeGroup = document.querySelectorAll('.poke-group');\npokeGroup.forEach(checkbox => {\n    checkbox.addEventListener('change', () => {\n        console.log(`${checkbox.value} is ${checkbox.checked ? 'checked' : 'unchecked'}`);\n        const userInput = document.getElementById(\"search\").value;\n        findPokemon(userInput);\n    });\n});\n\nfunction getAllPokemon() {\n    const url = \"http://localhost:8008/get_all_pokemon\";\n    fetch(url)\n    .then(response => {\n        if (!response.ok) {\n        throw new Error(`HTTP error! Status: ${response.status}`);\n        }\n        return response.json();\n    })\n    .then(data => {\n        appendPokemon(data)\n        })\n    .catch(error => console.error('Fetch error:', error));\n}\n\nfunction findPokemon(userInput) {\n    validInput(userInput);\n\n    const url = \"http://localhost:8008/find_pokemon\";\n\n    const types = filteredByType();\n    console.log(types);\n    if (types.length === 0) {\n        console.log(\"No types selected.\");\n    }\n\n    const gens = filteredByGen();\n    console.log(gens);\n    if (gens.length === 0) {\n        console.log(\"No gens selected.\");\n    }\n\n    const starter = document.getElementById(\"starters\");\n    const pseudo = document.getElementById(\"pseudo-legendary\");\n    const sub = document.getElementById(\"sub-legendary\");\n    const legend = document.getElementById(\"legendary\");\n    const myth = document.getElementById(\"mythical\");\n\n    const group = [\n        starter ? starter.checked : false,\n        pseudo ? pseudo.checked : false,\n        sub ? sub.checked : false,\n        legend ? legend.checked : false,\n        myth ? myth.checked : false\n    ];\n\n    console.log(group)\n\n    fetch(url, {\n         method: 'POST',  // Use POST to send data\n        headers: {\n            'Content-Type': 'application/json'  // Indicate that we're sending JSON data\n        },\n        body: JSON.stringify(\n            { \n                find: userInput,\n                order: isReverse,\n                types: types,\n                gens: gens,\n                group: group\n            })  // Send inputValue in the request body\n    })\n    .then(response => {\n        if (!response.ok) {\n        throw new Error(`HTTP error! Status: ${response.status}`);\n        }\n        return response.json();\n    })\n    .then(data => {\n        appendPokemon(data)\n    })\n    .catch(error => console.error('Fetch error:', error));\n}\n\nfunction toggleOrder() {\n    const userInput = document.getElementById(\"search\").value;\n    isReverse = !isReverse;\n    findPokemon(userInput);\n}\n\nfunction toggleFilter() {\n    const filter = document.getElementById('filter').classList.toggle('show');\n}\n\nfunction filteredByType() {\n    const types = [];\n    checkboxes.forEach(checkbox => {\n        if (checkbox.checked && !Number.isInteger(parseFloat(checkbox.value))) {\n            types.push(checkbox.value);\n        }\n    });\n    console.log(types);\n    return types\n}\n\nfunction filteredByGen() {\n    const gens = [];\n    checkboxes.forEach(checkbox => {\n        const genNum = parseInt(checkbox.value, 10); \n        if (checkbox.checked && genNum >= 0 && parseInt(checkbox.value)  <= 9) {\n            gens.push(genNum);\n        }\n    });\n    console.log(gens);\n    return gens;\n}\n\nfunction appendPokemon(pokemons) {\n    const tbody = document.getElementById(\"table-body\");\n     tbody.innerHTML = '';\n    for (const pokemon of pokemons) {\n        const row = document.createElement(\"tr\");\n        const types = document.createElement(\"ul\");\n        const abilities = document.createElement(\"ul\");\n        abilities.className = \"abilities\";\n        for (const key in pokemon) {\n            \n            if (key == \"main_t\") {\n                const main_t = document.createElement(\"li\");\n                main_t.className = \"type\";\n                main_t.textContent = capitalizes(String(pokemon[key]));\n                addTypeColor(main_t, main_t.textContent);\n                types.appendChild(main_t);\n            }\n            else if (key == \"sub_t\") {\n                const attribute = document.createElement(\"td\");\n                if (pokemon[key] != \"Null\") {\n                    const sub_t = document.createElement(\"li\");\n                    sub_t.className = \"type\";\n                    sub_t.textContent = capitalizes(String(pokemon[key]));\n                    addTypeColor(sub_t, sub_t.textContent);\n                    types.appendChild(sub_t);\n                }\n                attribute.appendChild(types);\n                row.appendChild(attribute);\n            }\n            else if (key == \"first_ability\") {\n                const first_ability = document.createElement(\"li\");\n                // first_ability.className = \"abilities\";\n                first_ability.textContent = capitalizes(String(pokemon[key]));\n                abilities.appendChild(first_ability);\n            }\n            else if (key == \"second_ability\") {\n                const second_ability = document.createElement(\"li\");\n                // second_ability.className = \"abilities\";\n                if (pokemon[key] != \"Null\") {\n                    second_ability.textContent = capitalizes(String(pokemon[key]));\n                }\n                abilities.appendChild(second_ability);\n            }\n             else if (key == \"hidden_ability\") {\n                const attribute = document.createElement(\"td\");\n                const hidden_ability = document.createElement(\"li\");\n                // hidden_ability.className = \"abilities\";\n                if (pokemon[key] != \"Null\") {\n                    hidden_ability.textContent = capitalizes(String(pokemon[key]));\n                    hidden_ability.style.fontWeight = \"bold\";\n                }\n                abilities.appendChild(hidden_ability);\n                attribute.appendChild(abilities);\n                row.appendChild(attribute);\n            }\n            else {\n                const attribute = document.createElement(\"td\");\n                switch (key) {\n                    case \"img\":\n                        const image = document.createElement(\"img\");\n                        const source = \"assets/images/\" + String(pokemon[\"id\"]) + \".png\";\n                        image.src = source;\n                        image.style.width = \"50px\";  // Corrected width\n                        image.style.height = \"50px\";  // Corrected height\n                        attribute.appendChild(image);\n                        break;\n                    case \"tags\":\n                        const lyst = document.createElement(\"ul\");\n                        const tags = pokemon[\"tags\"];\n                        for (const tag in tags) {\n                            switchTagValues(lyst, tags, tag);\n                        }\n                        attribute.appendChild(lyst);\n                        break;\n                    default:\n                        attribute.textContent = capitalizes(String(pokemon[key]));\n                        break;\n                }\n                row.appendChild(attribute);\n            }\n        }\n        tbody.appendChild(row);\n    }\n}\n\nfunction addTypeColor(li, liType) {\n    typesBgColor = [\n        [\"Normal\", \"#AAAA99\"],\n        [\"Fire\", \"#FF4422\"],\n        [\"Water\", \"#3399FF\"],\n        [\"Electric\", \"#FFCC33\"],\n        [\"Grass\", \"#77CC55\"],\n        [\"Ice\", \"#66CCFF\"],\n        [\"Fighting\", \"#BB5544\"],\n        [\"Poison\", \"#AA5599\"],\n        [\"Ground\", \"#DDBB55\"],\n        [\"Flying\", \"#8899FF\"],\n        [\"Psychic\", \"#FF5599\"],\n        [\"Bug\", \"#AABB22\"],\n        [\"Rock\", \"#BBAA66\"],\n        [\"Ghost\", \"#6666BB\"],\n        [\"Dragon\", \"#7766EE\"],\n        [\"Dark\", \"#775544\"],\n        [\"Steel\", \"#AAAABB\"],\n        [\"Fairy\", \"#EE99EE\"]\n    ];\n    typesBgColor.map( ([type, color]) => {\n        if (liType == type) li.style.backgroundColor = color;\n    });\n}\n\nfunction switchTagValues(lyst, tags, key) {\n    const datapoint = document.createElement(\"li\");\n    datapoint.className = \"tag\";\n    const gen = {\n        1: \"1-Kanto\",\n        2: \"2-Johto\",\n        3: \"3-Hoenn\",\n        4: \"4-Sinnoh\",\n        5: \"5-Unova\",\n        6: \"6-Kalos\",\n        7: \"7-Alola\",\n        8: \"8-Galar\",\n        9: \"9-Paldea\"\n    };\n\n    switch (key) {\n        case \"gens\":\n            datapoint.textContent = gen[tags[key]];\n            lyst.appendChild(datapoint);\n            break;\n        case \"isStarter\":\n            if (tags[key] == true) {\n                datapoint.textContent = \"Starter\";\n                lyst.appendChild(datapoint);\n            }\n            break;\n        case \"isPseudo\":\n            if (tags[key] == true) {\n                datapoint.textContent = \"Pseudo-Legendary\";\n                lyst.appendChild(datapoint);\n            }\n            \n            break;\n        case \"isSubLegend\":\n            if (tags[key] == true) {\n                datapoint.textContent = \"Sub-Legendary\";\n                lyst.appendChild(datapoint);\n            }\n            break;\n        case \"isLegend\":\n            if (tags[key] == true) {\n                datapoint.textContent = \"Legendary\";\n                lyst.appendChild(datapoint);\n            }\n            break;\n        case \"isMyth\":\n            if (tags[key] == true)  {\n                datapoint.textContent =\"Mythical\";\n                lyst.appendChild(datapoint);\n            }\n            break;\n        default:\n            break;\n    }\n}\n\nfunction validInput(userInput) {\n    const pattern = /^[A-Za-z\\s\\-]*$/;\n    if (!pattern.test(userInput)) {\n       console.log(\"Input does not match the pattern.\");\n    //    alert(\"Invalid Input:\\s\" + userInput);\n    }\n}\n\nfunction capitalizes(str) {\n    return str\n        .split(\"-\")\n        .map(segment => segment.charAt(0).toUpperCase() + segment.slice(1).toLowerCase())\n        .join(\"-\");\n}\n\n\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
