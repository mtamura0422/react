-- migrate:up

create table recipe_materials (
  id integer NOT NULL AUTO_INCREMENT,
  recipe_id integer not null,
  name varchar(15) not null,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);


-- migrate:down
drop table recipe_materials;


