--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.13 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: incidents; Type: TABLE; Schema: public; Owner: uday
--

CREATE TABLE public.incidents (
    id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    affected_service text NOT NULL,
    ai_severity text,
    ai_category text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.incidents OWNER TO uday;

--
-- Name: incidents_id_seq; Type: SEQUENCE; Schema: public; Owner: uday
--

CREATE SEQUENCE public.incidents_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.incidents_id_seq OWNER TO uday;

--
-- Name: incidents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: uday
--

ALTER SEQUENCE public.incidents_id_seq OWNED BY public.incidents.id;

--
-- Name: incidents id; Type: DEFAULT; Schema: public; Owner: uday
--

ALTER TABLE ONLY public.incidents ALTER COLUMN id SET DEFAULT nextval('public.incidents_id_seq'::regclass);


--
-- Name: incidents incidents_pkey; Type: CONSTRAINT; Schema: public; Owner: uday
--

ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT incidents_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

