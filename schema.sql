create TABLE conditions (
  id integer NOT NULL,
  zip character varying NOT NULL,
  station_id character varying,
	observation_epoch character varying,
	temp_f real,
	temp_c real,
	relative_humidity character varying,
	pressure_mb character varying,
	pressure_in character varying,
	pressure_trend character varying,
	dewpoint_f integer,
	dewpoint_c integer,
	feelslike_f character varying,
	feelslike_c character varying,
	visibility_mi character varying,
	visibility_km character varying,
	icon character varying,
	icon_url character varying,
	json text,
	created_at timestamp without time zone
);

CREATE SEQUENCE conditions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE conditions_id_seq OWNED BY conditions.id;
ALTER TABLE ONLY conditions ALTER COLUMN id SET DEFAULT nextval('conditions_id_seq'::regclass);