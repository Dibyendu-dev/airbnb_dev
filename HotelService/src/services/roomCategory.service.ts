import { CreateRoomCategoryDTO } from "../dto/roomCategory.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import { RoomCategoryRepository } from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";

const roomCategoryRepository = new RoomCategoryRepository   
const hotelRepository = new HotelRepository();

export async function createToomCategoryByService(CreateRoomCategoryDTO:CreateRoomCategoryDTO) {
    const roomCategory = await roomCategoryRepository.create(CreateRoomCategoryDTO)
    return roomCategory;
}

export async function getRoomCategoryByService(id:number){
    const roomCategory = await roomCategoryRepository.findById(id);
    return roomCategory;
}

export async function getAllRoomCategoriesByHotelIdService(hotelId:number) {
    const hotel = await hotelRepository.findById(hotelId)

    if(!hotel){
        throw new NotFoundError(`hotel with ${hotelId} is not found`)
    }

    const room_categories = await roomCategoryRepository.findAllByHotelId(hotelId)

    return room_categories
}

export async function deleteRoomCategoryService(id: number){
    const roomCategory = await roomCategoryRepository.findById(id);

    if(!roomCategory){
        throw new NotFoundError(`Room category with ${id} is not found`)

    }

    await roomCategoryRepository.delete({id})
    return true
}
