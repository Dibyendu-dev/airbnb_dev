import { createhotelDTO } from "../dto/hotel.dto";
import { HotelRepository } from "../repositories/hotel.repository";

const hotelRepository = new HotelRepository();

export async function createHotelService(hotelData:createhotelDTO){
    const hotel = await hotelRepository.create(hotelData)
    return hotel
}

export async function getHotelByIdService(id:number){
    const hotel = await hotelRepository.findById(id)
    return hotel
}

export async function getAllHotelsService(){
    const hotel = hotelRepository.findAll()
    return hotel
}

export async function deleteHotelService(id:number){
    const response = hotelRepository.softDelete(id)
    return response
}
