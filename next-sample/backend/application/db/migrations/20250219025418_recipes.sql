-- migrate:up

create table recipes (
  id integer NOT NULL AUTO_INCREMENT,
  title varchar(20) not null,
  content text not null,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);


-- migrate:down
drop table recipes;
