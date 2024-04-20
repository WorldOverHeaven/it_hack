CREATE TABLE IF NOT EXISTS public."task"
(
    task_id serial PRIMARY KEY ,
    title text NOT NULL ,
    completion_status boolean
    );
