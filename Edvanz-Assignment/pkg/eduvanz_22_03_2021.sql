--
-- PostgreSQL database dump
--

-- Dumped from database version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)

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

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: meetup; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.meetup (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    age integer NOT NULL,
    date character varying(100) NOT NULL,
    profession character varying(100) NOT NULL,
    locality character varying(100) NOT NULL,
    address character varying(200) NOT NULL,
    create_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.meetup OWNER TO postgres;

--
-- Name: meetup_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meetup_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.meetup_id_seq OWNER TO postgres;

--
-- Name: meetup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.meetup_id_seq OWNED BY public.meetup.id;


--
-- Name: meetup id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup ALTER COLUMN id SET DEFAULT nextval('public.meetup_id_seq'::regclass);


--
-- Data for Name: meetup; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.meetup (id, name, age, date, profession, locality, address, create_at) FROM stdin;
\.


--
-- Name: meetup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meetup_id_seq', 1, false);


--
-- Name: meetup meetup_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup
    ADD CONSTRAINT meetup_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

