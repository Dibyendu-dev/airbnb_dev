import { CreationAttributes } from "sequelize";
import BaseRepository from "./base.repository";
import Room from "../db/models/room.model";


export class RoomRepository extends BaseRepository<Room>{
    constructor(){
        super(Room)
    }

    async findByRoomCategoryIdAndDate(roomCategoryId: number,currentDate: Date){
        return await this.model.findOne({
            where: {
                roomCategoryId,
                dateOfAvailability: currentDate,
                deletedAt: null
            }
        })
    }

    async bulkCreate(rooms: CreationAttributes<Room>[]){
        return await this.model.bulkCreate(rooms)
    }
}