# pg tut

# starting the server
    docker-compose up

## connect to the psql client in another shell
    docker-compose exec database bash

## start the psql tool to connect to the generic postgres db...
    psql -U unicorn_user postgres 

    ....> magical_password

# Postgres Tutorial

Online can be found here: https://www.postgresqltutorial.com/

# Golang code samples:

Can be found in the ```cmd``` folder

# create a sample db from the Postgres tutorials
    postgres=# CREATE DATABASE dvdrental;

# create an ADMIN ROLE
    CREATE ROLE dbadmin WITH CREATEDB CREATEROLE LOGIN  PASSWORD 'bob';

# create a USER ROLE
    CREATE ROLE dbuser WITH LOGIN;

# create a USER 
    CREATE ROLE some_user IN ROLE dbuser WITH LOGIN PASSWORD 'bob';

# list users in psql

    dvdrental=# \du
                                         List of roles
      Role name   |                         Attributes                         | Member of
    --------------+------------------------------------------------------------+-----------
     dbadmin      | Create role, Create DB                                     | {}
     dbuser       |                                                            | {}
     some_user    |                                                            | {dbuser}
     unicorn_user | Superuser, Create role, Create DB, Replication, Bypass RLS | {}

# Set the default search path to 'public' schema
This saves having to prefix eveything if you're only going to use the default public schema

    SET search_path TO public;

# Grant access to all tables to a user
    dvdrental=# grant select,insert,update on all tables in schema "public" to some_user;

# create a table in the public schema of dvdrental
    dvdrental=# create table public.hello (
    message character varying(25) NOT NULL,
    last_update timestamp with time zone DEFAULT now() NOT NULL
    );
    CREATE TABLE

# create a sequence
    CREATE SEQUENCE seq_mytable_id;

# create a table using that sequence
    CREATE TABLE Blah (
        id int nextval("seq_mytable_id"),
        name character varying(15)
    )

# add constraint
    ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.address(address_id) ON UPDATE CASCADE ON DELETE RESTRICT;


# interval arithmetic
    select date_trunc('month', now()) + interval '1 month - 1d'


# create a partition
    CREATE TABLE public.transactions (
        transaction_id character varying(64),
        transaction_time timestampz,
        summary_id integer NOT NULL,
        amount numeric(10, 2),
        loc_id integer NOT NULL
    ) PARTITION BY RANGE(transaction_time);

    CREATE OR REPLACE FUNCTION public.partionFor(forMonthContaining date) RETURNS void
    AS $$ 
        CREATE TABLE public.transaction_yyyy_mm PARTITION OF public.transactions FOR VALUES 
        FROM (SELECT date_trunc('month', forDate) ) TO (SELECT date_trunc('month', forDate) + interval '1 month - 1 day');
    $$ LANGUAGE sql IMMUTABLE;
  

