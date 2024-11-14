--
-- PostgreSQL database dump
--

-- Dumped from database version 14.13 (Ubuntu 14.13-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.13 (Ubuntu 14.13-0ubuntu0.22.04.1)

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
-- Name: category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category (
    id text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    description text
);


ALTER TABLE public.category OWNER TO postgres;

--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    id_category text,
    name text,
    photo_url text,
    price numeric,
    purchase_date timestamp with time zone
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category (id, created_at, updated_at, deleted_at, name, description) FROM stdin;
	2024-11-13 23:56:42.988163+07	2024-11-13 23:56:42.988163+07	\N	eektronik	alat alat memasak
7c042494-abd5-495e-86ab-cf2c039afbbe	2024-11-14 02:52:52.754162+07	2024-11-14 03:13:12.222346+07	2024-11-14 03:16:44.155159+07	afdasfa	alat alat elektronik
68dac652-c973-4ae2-a019-80394809f637	2024-11-14 06:48:17.597953+07	2024-11-14 06:48:17.597953+07	\N	Kuliner	perlengkepan makanan karyawan
468f356e-0901-4235-838b-8fbf05492c9b	2024-11-14 06:55:30.764504+07	2024-11-14 06:55:30.764504+07	2024-11-14 06:58:07.280484+07	new	perlengkepan makanan karyawan
b110cb22-6135-4ab2-bd13-e1263cdb181f	2024-11-14 07:04:35.568969+07	2024-11-14 07:04:35.568969+07	\N	newElektronik	perlengkepan makanan karyawan
da00b7e9-d6c5-4465-8ac1-ad2a22be8485	2024-11-14 02:52:18.824479+07	2024-11-14 08:51:50.89814+07	\N	limbah	Update elektronik category
aaa0b4e3-b3c6-4c3a-9f0a-213a0141e411	2024-11-14 08:52:09.056761+07	2024-11-14 08:52:09.056761+07	\N	newElektronik	perlengkepan makanan karyawan
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, created_at, updated_at, deleted_at, id_category, name, photo_url, price, purchase_date) FROM stdin;
d47eb956-84cf-4e03-ba10-d73439074a08	2024-11-14 08:00:50.530462+07	2024-11-14 09:41:18.751875+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-08-02 22:04:05+07
d4c5698f-8b81-4a9c-a8a8-187315b32994	2024-11-14 03:50:09.921227+07	2024-11-14 03:50:09.921227+07	2024-11-14 07:51:50.243596+07	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	topi	dsfdsf	1000	0001-01-01 07:07:12+07:07:12
ae6b2b9d-8f79-42d0-b2a4-13a22cb1122c	2024-11-14 07:58:24.662588+07	2024-11-14 09:41:43.735465+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-09-02 22:04:05+07
8e6c7974-b4bd-4303-bf94-ae8afeff9d61	2024-11-14 07:56:33.026059+07	2024-11-14 09:42:12.268938+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-12-02 22:04:05+07
73233545-b06b-48a0-a8a7-7a406ef3c07f	2024-11-14 07:50:44.379912+07	2024-11-14 09:42:34.657249+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-11-02 22:04:05+07
eeccbfd2-f11a-47a5-9595-966ed7320367	2024-11-14 08:08:35.740233+07	2024-11-14 09:39:59.537791+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-03-02 22:04:05+07
f7a79a27-58bc-4ef9-aa2f-9cfd2b665374	2024-11-14 03:48:53.061527+07	2024-11-14 09:43:01.703564+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-01-02 22:04:05+07
4fc21fac-098b-4881-b368-9446831f7baf	2024-11-14 08:24:23.218429+07	2024-11-14 09:40:29.266696+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-05-02 22:04:05+07
efc50648-1a4c-4c32-80fc-1173562ec000	2024-11-14 08:02:20.033063+07	2024-11-14 09:40:59.433918+07	\N	da00b7e9-d6c5-4465-8ac1-ad2a22be8485	televisi	image	1233333	2022-04-02 22:04:05+07
\.


--
-- Name: category categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_categories_deleted_at ON public.category USING btree (deleted_at);


--
-- Name: idx_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_items_deleted_at ON public.items USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

