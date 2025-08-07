import { RoomType } from "../db/models/roomCategory.model";

export type CreateRoomCategoryDTO ={
    hotelId: number;
    price: number;
    roomType: RoomType;
    roomCount: number;
}