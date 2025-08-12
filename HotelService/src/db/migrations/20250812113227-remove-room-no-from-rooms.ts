import { QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
        alter table rooms drop column room_no;
      `)
  },

  async down (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
        alter table rooms add column room_no int not null;
      `)
  }
};
