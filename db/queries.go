package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Pokemon struct {
	ID            int    `json:"id"`                       // Primary Key
	Name          string `json:"name"`                     // Not Null
	Img           string `json:"img"`                      // Not Null
	MainType      string `json:"main_t"`                   // Not Null
	SubType       string `json:"sub_t,omitempty"`          // Nullable
	FirstAbility  string `json:"first_ability"`            // Not Null
	SecondAbility string `json:"second_ability,omitempty"` // Nullable
	HiddenAbility string `json:"hidden_ability"`           // Not Null
	HP            int    `json:"hp"`                       // Not Null
	Atk           int    `json:"atk"`                      // Not Null
	Def           int    `json:"def"`                      // Not Null
	SpAtk         int    `json:"sp_atk"`                   // Not Null
	SpDef         int    `json:"sp_def"`                   // Not Null
	Spd           int    `json:"spd"`                      // Not Null
}

func GetAllPokemon(conn *sql.DB) ([]Pokemon, error) {
	var query = `
	SELECT id, name, img, main_t, sub_t, first_ability, second_ability, hidden_ability, hp, atk, def, sp_atk, sp_def, spd
	FROM basic_info;`

	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []Pokemon
	for rows.Next() {
		var pokemon Pokemon
		err = rows.Scan(
			&pokemon.ID,
			&pokemon.Name,
			&pokemon.Img,
			&pokemon.MainType,
			&pokemon.SubType,
			&pokemon.FirstAbility,
			&pokemon.SecondAbility,
			&pokemon.HiddenAbility,
			&pokemon.HP,
			&pokemon.Atk,
			&pokemon.Def,
			&pokemon.SpAtk,
			&pokemon.SpDef,
			&pokemon.Spd,
		)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pokemon)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func FindPokemon(conn *sql.DB, txtInput string, isReverse bool, types []string, gens []int, group []bool) ([]Pokemon, error) {
	var query string
	var args []interface{}

	query = `
	SELECT id, name, img, main_t, sub_t, first_ability, second_ability, hidden_ability, hp, atk, def, sp_atk, sp_def, spd 
	FROM basic_info 
	WHERE name LIKE ?
	`

	args = append(args, txtInput+"%")

	if len(types) > 0 {
		placeholders := make([]string, len(types))
		for i := range types {
			placeholders[i] = "?"
		}
		typesPlaceholder := strings.Join(placeholders, ",")

		query += fmt.Sprintf(` AND (main_t IN (%s) OR sub_t IN (%s))`, typesPlaceholder, typesPlaceholder)

		args = append(args, interfaceSlice(types)...)
		args = append(args, interfaceSlice(types)...)
	}

	if len(gens) > 0 {
		placeholders := make([]string, len(gens))
		for i := range gens {
			placeholders[i] = "?"
		}
		gensPlaceholder := strings.Join(placeholders, ",")
		query += fmt.Sprintf(` AND gen IN (%s)`, gensPlaceholder)
		args = append(args, intSliceToInterface(gens)...)
	}

	var tmpQuery = ``
	for i, value := range group {
		groupLabel := []string{"isStarter", "isPseudo", "isSubLegend", "isLegend", "isMyth"}
		if i < len(groupLabel) && value {
			tmpQuery += fmt.Sprintf(`(%s) = 1 OR`, groupLabel[i])
		}
	}
	if len(tmpQuery) > 0 {
		tmpQuery = tmpQuery[:len(tmpQuery)-3]
		query += `AND (` + tmpQuery + `)`
	}

	if isReverse {
		query += ` ORDER BY id DESC`
	}

	query += ";"

	log.Println("Executing Query:", query)
	log.Println("With Args:", args)

	// Execute the query
	rows, err := conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []Pokemon
	for rows.Next() {
		var pokemon Pokemon
		err = rows.Scan(
			&pokemon.ID,
			&pokemon.Name,
			&pokemon.Img,
			&pokemon.MainType,
			&pokemon.SubType,
			&pokemon.FirstAbility,
			&pokemon.SecondAbility,
			&pokemon.HiddenAbility,
			&pokemon.HP,
			&pokemon.Atk,
			&pokemon.Def,
			&pokemon.SpAtk,
			&pokemon.SpDef,
			&pokemon.Spd,
		)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pokemon)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func interfaceSlice(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func intSliceToInterface(slice []int) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
