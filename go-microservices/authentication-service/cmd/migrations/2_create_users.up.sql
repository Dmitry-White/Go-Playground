--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE users (
    id integer DEFAULT nextval('user_id_seq' :: regclass) NOT NULL,
    email character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    password character varying(60),
    user_active integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);