package db

import "database/sql"

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

func getAll(conn *sql.DB) ([]Pokemon, error) {
	rows, err := conn.Query("SELECT * FROM basic_info;")
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
	return pokemons, nil
}
