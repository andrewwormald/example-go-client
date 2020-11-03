create table users (
    id bigint not null auto_increment,
    email varchar(255),


    unique(email),
    primary key (id)
);
