CREATE TABLE airports (
    id INT NOT NULL DEFAULT unique_rowid(),
    size integer,
    CONSTRAINT airports_pkey PRIMARY KEY (id)
);

CREATE TABLE hangars (
    id INT NOT NULL DEFAULT unique_rowid(),
    name text NOT NULL,
    CONSTRAINT hangars_pkey PRIMARY KEY (id)
);

CREATE TABLE jets (
    id INT NOT NULL DEFAULT unique_rowid(),
    pilot_id integer NOT NULL,
    airport_id integer NOT NULL,
    name text NOT NULL,
    color text,
    uuid text NOT NULL,
    identifier text NOT NULL,
    cargo bytea NOT NULL,
    manifest bytea NOT NULL,
    CONSTRAINT jets_pkey PRIMARY KEY (id)
);

CREATE TABLE languages (
    id INT NOT NULL DEFAULT unique_rowid(),
    language text NOT NULL,
    CONSTRAINT languages_pkey PRIMARY KEY (id)
);


CREATE TABLE licenses (
    id INT NOT NULL DEFAULT unique_rowid(),
    pilot_id integer,
    CONSTRAINT licenses_pkey PRIMARY KEY (id)
);

CREATE TABLE pilot_languages (
    pilot_id INT NOT NULL,
    language_id INT NOT NULL,
    CONSTRAINT pilot_languages_pkey PRIMARY KEY (pilot_id, language_id)
);

CREATE TABLE pilots (
    id INT NOT NULL DEFAULT unique_rowid(),
    name text NOT NULL,
    CONSTRAINT pilots_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_pilot_languages ON languages USING btree (language);

ALTER TABLE ONLY jets ADD CONSTRAINT jets_airport_id_airports_id_foreign FOREIGN KEY (airport_id) REFERENCES airports(id);
ALTER TABLE ONLY jets ADD CONSTRAINT jets_pilot_id_pilots_id_foreign FOREIGN KEY (pilot_id) REFERENCES pilots(id);
ALTER TABLE ONLY pilot_languages ADD CONSTRAINT languages_fkey FOREIGN KEY (language_id) REFERENCES languages(id);
ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_pilot_id_pilots_id_foreign FOREIGN KEY (pilot_id) REFERENCES pilots(id);
ALTER TABLE ONLY pilot_languages ADD CONSTRAINT pilots_fkey FOREIGN KEY (pilot_id) REFERENCES pilots(id);
