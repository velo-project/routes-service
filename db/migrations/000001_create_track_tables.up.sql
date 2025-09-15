CREATE TABLE tbl_track(
    id_track SERIAL,
    fk_user INTEGER NOT NULL,
    tx_initial_location VARCHAR(255) NOT NULL,
    tx_final_location VARCHAR(255) NOT NULL,
    tx_visited_at TIMESTAMP NULL,
    PRIMARY KEY (id_track)
);

CREATE TABLE tbl_location(
    id_location SERIAL,
    fk_track INTEGER NULL,
    tx_lat FLOAT NOT NULL,
    tx_lng FLOAT NOT NULL,
    PRIMARY KEY (id_location),
    FOREIGN KEY (fk_track) REFERENCES tbl_track(id_track)
);
