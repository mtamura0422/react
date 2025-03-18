-- migrate:up
ALTER TABLE recipe_materials ADD INDEX idx_recipe_materials_1(recipe_id)


-- migrate:down
ALTER TABLE recipe_materials DROP INDEX idx_recipe_materials_1
