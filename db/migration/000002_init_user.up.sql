create table
    "user" (
        id uuid not null default uuid_generate_v1mc (),
        -- By applying our custom collation we can simply mark this column as `unique` and Postgres will enforce
        -- case-insensitive uniqueness for us, and lookups over `username` will be case-insensitive by default.
        --
        -- Note that this collation doesn't support the `LIKE`/`ILIKE` operators so if you want to do searches
        -- over `username` you will want a separate index with the default collation:
        --
        -- create index on "user" (username collate "ucs_basic");
        --
        -- select * from "user" where (username collate "ucs_basic") ilike ($1 || '%')
        username text collate "case_insensitive" not null,
        email text collate "case_insensitive" not null,
        bio text not null default '',
        password_hash text not null,
        status text not null default 'active',
        -- If you want to be really pedantic you can add a trigger that enforces this column will never change,
        -- but that seems like overkill for something that's relatively easy to enforce in code-review.
        created_date timestamptz not null default now(),
        updated_date timestamptz,
        constraint pk_user primary key (id)
    );

-- create indexing
create unique index idx_user_email on "user" (email);

-- -- And applying our `updated_at` trigger is as easy as this.
select
    trigger_updated_at ('"user"');