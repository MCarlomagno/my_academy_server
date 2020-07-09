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


ALTER TABLE courses ADD COLUMN thumbnail_image VARCHAR(255);
ALTER TABLE classes ADD COLUMN thumbnail_image VARCHAR(255);

ALTER TABLE users ADD email VARCHAR(255);

ALTER TABLE users ADD name VARCHAR(255);

ALTER TABLE users ADD surname VARCHAR(255);

ALTER TABLE users ADD image_url VARCHAR(255);


CREATE EXTENSION pgcrypto;

ALTER TABLE users ADD COLUMN password TEXT;

/* SAMPLE QUERIES */
INSERT INTO users (nombre, email, password) VALUES (
    'Jon',
  'jon@doe.com',
  crypt('jon', gen_salt('bf'))
);

SELECT id FROM users WHERE email = 'jon@doe.com' AND password = crypt('jon', password);
/* SAMPLE QUERIES */

ALTER TABLE users ALTER COLUMN password SET NOT NULL;

