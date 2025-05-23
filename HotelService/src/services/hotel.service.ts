import { createhotelDTO } from "../dto/hotel.dto";
import { createhotel, getAllHotels, getHotelById, softDeleteHotel } from "../repositories/hotel.repository";

export async function createHotelService(hotelData:createhotelDTO){
    const hotel = await createhotel(hotelData)
    return hotel
}

export async function getHotelByIdService(id:number){
    const hotel = await getHotelById(id)
    return hotel
}

export async function getAllHotelsService(){
    const hotel = getAllHotels()
    return hotel
}

export async function deleteHotelService(id:number){
    const response = softDeleteHotel(id)
    return response
}
