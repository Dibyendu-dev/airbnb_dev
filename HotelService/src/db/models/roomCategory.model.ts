import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize  from "./sequelize";
import Hotel from "./hotel.model";

export enum RoomType {
    single = 'SINGLE',
    double = 'DOUBLE',
    family = 'FAMILY',
    deluxe = 'DELUXE',
    suite = 'SUITE',
    premium = 'PREMIUM',
}

class RoomCategory extends Model<InferAttributes<RoomCategory>,InferCreationAttributes<RoomCategory>> {
    declare id: CreationOptional<number>;
    declare hotelId: number;
    declare price: number;
    declare roomType: RoomType;
    declare roomCount: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date> | null;
}

RoomCategory.init({
    id: {
        type: 'INTEGER',
        autoIncrement:true,
        primaryKey:true,
    },
    hotelId: {
        type: 'INTEGER',
        allowNull: false,
        references: {
            model: Hotel,
            key: 'id',
        },
    },
    price: {
        type: 'INTEGER',
        allowNull: false,
    },
    roomType: {
        type: 'ENUM',
        values: [...Object.values(RoomType)],
    },
    roomCount: {
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
    
},{
    tableName: 'room_categories',
    sequelize: sequelize,
    underscored: true,
    timestamps: true,
})

export default RoomCategory;