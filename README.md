Credits: https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356 By Bradford Lamson-Scribner

# Database

```sql
$ CREATE SEQUENCE users_id_seq;
```

```sql
$ CREATE TABLE public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    age integer NOT NULL,
    profession character varying(50) COLLATE pg_catalog."default" NOT NULL,
    friendly boolean NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
```