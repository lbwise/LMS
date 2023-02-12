DROP TABLE IF EXISTS modules;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    user_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL
);

CREATE TABLE courses (
    course_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    price MONEY DEFAULT 0,
    created_on TIMESTAMP NOT NULL,
    owner_id uuid REFERENCES users(user_id) ON DELETE SET NULL
);

CREATE TABLE students (
    student_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid REFERENCES users(user_id) ON DELETE SET NULL,
    course_id uuid REFERENCES courses(course_id) ON DELETE SET NULL,
    enrolled_on TIMESTAMP NOT NULL,
    progress INT,
    rating INT
);

CREATE TABLE modules (
    module_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    course_id uuid REFERENCES courses(course_id) ON DELETE SET NULL,
    name VARCHAR(50) NOT NULL,
    content TEXT,
    sequence_id BIGSERIAL
);

