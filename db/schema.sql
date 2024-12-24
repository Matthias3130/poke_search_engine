

DROP TABLE IF EXISTS basic_info;

CREATE TABLE basic_info (
    id INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    img TEXT NOT NULL,
    main_t TEXT NOT NULL,
    sub_t TEXT,
    first_ability TEXT NOT NULL,
    second_ability TEXT,
    hidden_ability TEXT NOT NULL,
    hp INTEGER NOT NULL,
    atk INTEGER NOT NULL,
    def INTEGER NOT NULL,
    sp_atk INTEGER NOT NULL,
    sp_def INTEGER NOT NULL,
    spd INTEGER NOT NULL
);

.mode csv

.import ./db/poke_database.csv basic_info

SELECT * FROM basic_info;
