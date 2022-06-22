create table User
(
    uid varchar(20) primary key ,
    name varchar(20) not null,
    password varchar(20) not null
);

create table Code
(
    cid varchar(20) primary key ,
    uid varchar(20) ,
    code text,
    result text,
    time varchar(20),
    type varchar(10),
    foreign key (uid) references User(uid)
);


