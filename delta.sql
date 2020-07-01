CREATE TABLE users (
   id          integer NOT NULL,
   PRIMARY KEY (id)
);
CREATE SEQUENCE users_id_seq START WITH 1;
ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('users_id_seq');
 
CREATE TABLE courses (
   id          integer NOT NULL,
   id_user     integer NOT NULL,
   title       varchar(255),
   description varchar(255),
   PRIMARY KEY (id),
   FOREIGN KEY (id_user) REFERENCES users (id)
);
CREATE SEQUENCE courses_id_seq START WITH 1;
ALTER TABLE courses ALTER COLUMN id SET DEFAULT nextval('courses_id_seq');
 
CREATE TABLE modules (
   id          integer NOT NULL,
   id_course   integer NOT NULL,
   title       varchar(255),
   description varchar(255),
   PRIMARY KEY (id),
   FOREIGN KEY (id_course) REFERENCES courses (id)
);
CREATE SEQUENCE modules_id_seq START WITH 1;
ALTER TABLE modules ALTER COLUMN id SET DEFAULT nextval('modules_id_seq');
 
CREATE TABLE classes (
   id          integer NOT NULL,
   id_module   integer NOT NULL,
   title       varchar(255),
   description varchar(255),
   video_url   varchar(255),
   PRIMARY KEY (id),
   FOREIGN KEY (id_module) REFERENCES modules (id)
);
CREATE SEQUENCE classes_id_seq START WITH 1;
ALTER TABLE classes ALTER COLUMN id SET DEFAULT nextval('classes_id_seq');
 
CREATE TABLE enrollment (
   id          integer NOT NULL,
   id_user     integer NOT NULL,
   id_course      integer NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (id_user ) REFERENCES users (id),
   FOREIGN KEY (id_course ) REFERENCES courses (id)
);
CREATE SEQUENCE enrollment_id_seq START WITH 1;
ALTER TABLE enrollment ALTER COLUMN id SET DEFAULT nextval('enrollment_id_seq');
