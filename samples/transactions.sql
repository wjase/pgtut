--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0
-- Dumped by pg_dump version 13.0

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
-- Name: seq_summaries_id; Type: SEQUENCE; Schema: public; Owner: unicorn_user
--

CREATE SEQUENCE public.seq_summaries_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_summaries_id OWNER TO unicorn_user;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: summaries; Type: TABLE; Schema: public; Owner: unicorn_user
--

CREATE TABLE public.summaries (
    id integer DEFAULT nextval('public.seq_summaries_id'::regclass) NOT NULL,
    amount1_sum numeric(10,2),
    amount2_sum numeric(10,2),
    amount3_sum numeric(10,2),
    amount4_sum numeric(10,2),
    amount_count integer,
    loc_id integer NOT NULL,
    date date
);


ALTER TABLE public.summaries OWNER TO unicorn_user;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: unicorn_user
--

CREATE TABLE public.transactions (
    transaction_id character varying(64) NOT NULL,
    transaction_time timestamp with time zone,
    transaction_date date,
    amount1 numeric(10,2),
    amount2 numeric(10,2),
    amount3 numeric(10,2),
    amount4 numeric(10,2),
    loc_id integer NOT NULL
);


ALTER TABLE public.transactions OWNER TO unicorn_user;

--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: unicorn_user
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id);


--
-- PostgreSQL database dump complete
--

