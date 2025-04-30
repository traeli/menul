create table "user"
(
    id         uuid not null
        primary key,
    open_id    varchar(255) default ' '::character varying
        constraint user_open_id_unique
            unique,
    union_id   varchar(255),
    nickname   varchar(255),
    avatar_url varchar(255),
    gender     integer      default 0,
    language   varchar(255),
    country    varchar(255),
    create_at  timestamp,
    update_at  timestamp,
    login_ip   varchar(255),
    phone      varchar(255),
    status     integer      default 0
);

alter table "user"
    owner to directus;

create trigger update_user_updated_at
    before update
    on "user"
    for each row
    execute procedure update_updated_at_column();

