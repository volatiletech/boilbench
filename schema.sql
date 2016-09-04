CREATE TABLE airports (
    id integer NOT NULL,
    size integer
);

ALTER TABLE airports OWNER TO postgres;
CREATE SEQUENCE airports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE airports_id_seq OWNER TO postgres;
ALTER SEQUENCE airports_id_seq OWNED BY airports.id;

CREATE TABLE hangars (
    id integer NOT NULL,
    name text NOT NULL
);

ALTER TABLE hangars OWNER TO postgres;
CREATE SEQUENCE hangars_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE hangars_id_seq OWNER TO postgres;
ALTER SEQUENCE hangars_id_seq OWNED BY hangars.id;

CREATE TABLE jets (
    id integer NOT NULL,
    pilot_id integer NOT NULL,
    airport_id integer NOT NULL,
    name text NOT NULL,
    color text,
    uuid text NOT NULL,
    identifier text NOT NULL,
    cargo bytea NOT NULL,
    manifest bytea NOT NULL
);

ALTER TABLE jets OWNER TO postgres;
CREATE SEQUENCE jets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE jets_id_seq OWNER TO postgres;
ALTER SEQUENCE jets_id_seq OWNED BY jets.id;

CREATE TABLE languages (
    id integer NOT NULL,
    language text NOT NULL
);

ALTER TABLE languages OWNER TO postgres;
CREATE SEQUENCE languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE languages_id_seq OWNER TO postgres;
ALTER SEQUENCE languages_id_seq OWNED BY languages.id;

CREATE TABLE licenses (
    id integer NOT NULL,
    pilot_id integer
);

ALTER TABLE licenses OWNER TO postgres;
CREATE SEQUENCE licenses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE licenses_id_seq OWNER TO postgres;
ALTER SEQUENCE licenses_id_seq OWNED BY licenses.id;

CREATE TABLE pilot_languages (
    pilot_id integer NOT NULL,
    language_id integer NOT NULL
);

ALTER TABLE pilot_languages OWNER TO postgres;
CREATE TABLE pilots (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE pilots OWNER TO postgres;
CREATE SEQUENCE pilots_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE pilots_id_seq OWNER TO postgres;
ALTER SEQUENCE pilots_id_seq OWNED BY pilots.id;

ALTER TABLE ONLY airports ALTER COLUMN id SET DEFAULT nextval('airports_id_seq'::regclass);
ALTER TABLE ONLY hangars ALTER COLUMN id SET DEFAULT nextval('hangars_id_seq'::regclass);
ALTER TABLE ONLY jets ALTER COLUMN id SET DEFAULT nextval('jets_id_seq'::regclass);
ALTER TABLE ONLY languages ALTER COLUMN id SET DEFAULT nextval('languages_id_seq'::regclass);
ALTER TABLE ONLY licenses ALTER COLUMN id SET DEFAULT nextval('licenses_id_seq'::regclass);
ALTER TABLE ONLY pilots ALTER COLUMN id SET DEFAULT nextval('pilots_id_seq'::regclass);

ALTER TABLE ONLY airports ADD CONSTRAINT airports_pkey PRIMARY KEY (id);
ALTER TABLE ONLY hangars ADD CONSTRAINT hangars_pkey PRIMARY KEY (id);
ALTER TABLE ONLY jets ADD CONSTRAINT jets_pkey PRIMARY KEY (id);
ALTER TABLE ONLY languages ADD CONSTRAINT languages_pkey PRIMARY KEY (id);
ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_pkey PRIMARY KEY (id);
ALTER TABLE ONLY pilot_languages ADD CONSTRAINT pilot_languages_pkey PRIMARY KEY (pilot_id, language_id);
ALTER TABLE ONLY pilots ADD CONSTRAINT pilots_pkey PRIMARY KEY (id);

CREATE INDEX idx_pilot_languages ON languages USING btree (language);

ALTER TABLE ONLY jets ADD CONSTRAINT jets_airport_id_airports_id_foreign FOREIGN KEY (airport_id) REFERENCES airports(id);
ALTER TABLE ONLY jets ADD CONSTRAINT jets_pilot_id_pilots_id_foreign FOREIGN KEY (pilot_id) REFERENCES pilots(id);
ALTER TABLE ONLY pilot_languages ADD CONSTRAINT languages_fkey FOREIGN KEY (language_id) REFERENCES languages(id);
ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_pilot_id_pilots_id_foreign FOREIGN KEY (pilot_id) REFERENCES pilots(id);
ALTER TABLE ONLY pilot_languages ADD CONSTRAINT pilots_fkey FOREIGN KEY (pilot_id) REFERENCES pilots(id);

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;

