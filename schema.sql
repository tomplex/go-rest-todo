CREATE SEQUENCE public.todo_id_seq MINVALUE 1000;

CREATE TABLE public.todo(
	id INT PRIMARY KEY DEFAULT nextval('todo_id_seq'),
	name TEXT,
	completed BOOLEAN DEFAULT FALSE,
	due_date TIMESTAMP,
	entered_date TIMESTAMP DEFAULT now(),
	completed_date TIMESTAMP
);


