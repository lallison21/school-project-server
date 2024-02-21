CREATE TABLE IF NOT EXISTS role_list
(
    id serial PRIMARY KEY not null unique,
    role_name varchar(255) not null unique,
    access_level int not null
);

CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY not null unique,
    login varchar(255) not null unique,
    password varchar(255) not null,
    photo varchar(255),
    role_id int references role_list(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS student_info
(
    id serial PRIMARY KEY not null unique,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    patronymic varchar(255),
    user_id int references users(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS teacher_info
(
    id serial PRIMARY KEY not null unique,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    patronymic varchar(255),
    user_id int references users(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS class_info
(
    id serial PRIMARY KEY not null unique,
    class_number int not null,
    class_subgroup varchar(255) not null,
    teacher_id int references teacher_info on delete cascade not null
);

CREATE TABLE IF NOT EXISTS classroom_info
(
    id serial PRIMARY KEY not null unique,
    classroom_number int not null,
    floor_number int not null,
    responsible_teacher_id int references teacher_info(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS lesson_info
(
    id serial PRIMARY KEY not null unique,
    lesson_name varchar(255) not null unique,
    main_teacher_id int references teacher_info(id) on delete cascade not null,
    main_classroom_id int references classroom_info(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS student_in_class
(
    id serial PRIMARY KEY not null unique,
    student_id int references student_info(id) on delete cascade not null,
    class_id int references class_info(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS lessons_schedule
(
    id serial PRIMARY KEY not null unique,
    lesson_id int references lesson_info(id) on delete cascade not null,
    datetime_start date not null,
    teacher_id int references teacher_info(id) on delete cascade not null,
    class_id int references class_info(id) on delete cascade not null,
    classroom_id int references classroom_info(id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS journal
(
    id serial PRIMARY KEY not null unique,
    lessons_schedule_id int references lessons_schedule(id) on delete cascade not null,
    student_id int references student_info(id) on delete cascade not null,
    assessment int,
    visiting bool default false not null
);

INSERT INTO role_list(role_name, access_level) VALUES('Администратор', 1)
