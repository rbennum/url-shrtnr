/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function(knex) {
    return knex.schema.createTable('link_mappers', (table) => {
        table.increments('id').primary();
        table.string('url', 400).notNullable();
        table.string('short_tag', 5).notNullable();
    });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function(knex) {
    return knex.schema.dropTableIfExists('link_mappers');
};
