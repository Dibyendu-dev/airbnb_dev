import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize  from "./sequelize";
import Hotel from "./hotel.model";
import RoomCategory from "./roomCategory.model"

class Room extends Model <InferAttributes<Room>,InferCreationAttributes<Room>>{
    declare id: CreationOptional<number>;
    declare hotelId: number;
    declare  roomCategoryId: number;
    declare dateOfAvailability: Date;
    declare price: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date> | null;
    declare bookingId?: number | null; 
}

Room.init({
    id: {
        type: 'INTEGER',
        autoIncrement: true,
        primaryKey: true,
    },
     hotelId: {
                type: 'INTEGER',
                allowNull: false,
                references:{
                model : Hotel,
                key: 'id',
                },
             },
     roomCategoryId: {
        type: 'INTEGER',
        allowNull: false,
             references:{
                model : RoomCategory,
                key: 'id',
                },
     },
     dateOfAvailability: {
        type: 'DATE',
        allowNull: false,
     },
     price: {
        type: 'INTEGER',
        allowNull: false,
     },
     createdAt: {
        type: 'DATE',
        defaultValue: new Date(),
     },
     updatedAt: {
        type: 'DATE',
        defaultValue: new Date(),
     },
     deletedAt: {
        type: 'DATE',
        defaultValue: new Date(),
     },
     bookingId: {
        type: 'INTEGER',
        defaultValue: null,
     }
},{
 tableName: 'rooms',
 sequelize: sequelize,
 underscored: true,
 timestamps: true,
})

export default Room;