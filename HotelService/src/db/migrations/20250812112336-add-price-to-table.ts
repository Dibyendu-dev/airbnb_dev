import { QueryInterface } from "sequelize";



module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
        alter table rooms add column price int not null default 0;
      `)
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
          alter table rooms drop column price;
      `)
    
  }
};
