create table
    follow (
        id BIGSERIAL not null,
        followed_user_id uuid not null references "user" (id) on delete cascade,
        following_user_id uuid not null references "user" (id) on delete cascade,
        status text not null default 'active',
        created_date timestamptz not null default now(),
        updated_date timestamptz,
        constraint user_cannot_follow_self check (followed_user_id != following_user_id),
        primary key (following_user_id, followed_user_id)
    );

SELECT
    trigger_updated_at ('follow');