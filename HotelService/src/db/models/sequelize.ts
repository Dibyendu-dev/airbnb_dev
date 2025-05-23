import { Sequelize } from "sequelize";


const sequelize = new Sequelize({
    dialect: 'mysql',
    host: "localhost",
    username: "root",
    password: "ddas4548",
    database: "airbnb_dev",
    logging: true
})

export default sequelize