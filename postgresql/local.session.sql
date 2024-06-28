CREATE TABLE public.achievement (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    img VARCHAR(100) NOT NULL,
    info VARCHAR(100) NOT NULL,
    user_id INTEGER NOT NULL,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.user (id) ON DELETE CASCADE
);

CREATE TABLE public.user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

INSERT INTO "user" (name, email, password) VALUES ('admin', 'admin@admin.com', 'admin'); -- 1
INSERT INTO "user" (name, email, password) VALUES ('notadmin', 'notadmin@notadmin.com', 'notadmin'); -- 2

INSERT INTO "achievement" (title, img, info, user_id) VALUES ('achievement1',  'achievement1.png',  'infoURL',  1); --  1
INSERT INTO "achievement" (title, img, info, user_id) VALUES ('achievement1',  'achievement1.png',  'infoURL',  2); --  2