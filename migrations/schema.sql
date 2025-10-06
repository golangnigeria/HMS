--
-- PostgreSQL database dump
--

\restrict Y7WJwYicTWDTeaAI82x5CBbFNuMxydq6YSuBG9tsM2o68eGA6bVlHRWnpTSi8BS

-- Dumped from database version 17.5 (84bec44)
-- Dumped by pg_dump version 17.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: doctors; Type: TABLE; Schema: public; Owner: neondb_owner
--

CREATE TABLE public.doctors (
    id uuid NOT NULL,
    full_name character varying(255),
    email character varying(255),
    phone_number character varying(20) NOT NULL,
    specialization character varying(255) NOT NULL,
    license_number character varying(255),
    clinic_name character varying(255) NOT NULL,
    clinic_location character varying(255) NOT NULL,
    bank_account_number character varying(255) NOT NULL,
    bank_name character varying(255) NOT NULL,
    password_hash character varying(60),
    rating numeric(3,2) DEFAULT '0'::numeric NOT NULL,
    created_at timestamp without time zone DEFAULT '2025-10-06 16:35:14.534249'::timestamp without time zone NOT NULL,
    updated_at timestamp without time zone DEFAULT '2025-10-06 16:35:14.534249'::timestamp without time zone NOT NULL
);


ALTER TABLE public.doctors OWNER TO neondb_owner;

--
-- Name: patients; Type: TABLE; Schema: public; Owner: neondb_owner
--

CREATE TABLE public.patients (
    id uuid NOT NULL,
    full_name character varying(255),
    email character varying(255),
    phone_number character varying(20) NOT NULL,
    dob date NOT NULL,
    gender character varying(255) DEFAULT 'unspecified'::character varying NOT NULL,
    address text NOT NULL,
    password_hash character varying(60),
    created_at timestamp without time zone DEFAULT '2025-10-06 16:33:44.130518'::timestamp without time zone NOT NULL,
    updated_at timestamp without time zone DEFAULT '2025-10-06 16:33:44.130518'::timestamp without time zone NOT NULL
);


ALTER TABLE public.patients OWNER TO neondb_owner;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: neondb_owner
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO neondb_owner;

--
-- Name: doctors doctors_pkey; Type: CONSTRAINT; Schema: public; Owner: neondb_owner
--

ALTER TABLE ONLY public.doctors
    ADD CONSTRAINT doctors_pkey PRIMARY KEY (id);


--
-- Name: patients patients_pkey; Type: CONSTRAINT; Schema: public; Owner: neondb_owner
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT patients_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: neondb_owner
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: neondb_owner
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: DEFAULT PRIVILEGES FOR SEQUENCES; Type: DEFAULT ACL; Schema: public; Owner: cloud_admin
--

ALTER DEFAULT PRIVILEGES FOR ROLE cloud_admin IN SCHEMA public GRANT ALL ON SEQUENCES TO neon_superuser WITH GRANT OPTION;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: public; Owner: cloud_admin
--

ALTER DEFAULT PRIVILEGES FOR ROLE cloud_admin IN SCHEMA public GRANT ALL ON TABLES TO neon_superuser WITH GRANT OPTION;


--
-- PostgreSQL database dump complete
--

\unrestrict Y7WJwYicTWDTeaAI82x5CBbFNuMxydq6YSuBG9tsM2o68eGA6bVlHRWnpTSi8BS

