import logger from "../config/logger.config";
import Hotel from "../db/models/hotel.model";
import { createhotelDTO } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

export async function createhotel(hotelData:createhotelDTO){
    const hotel = await Hotel.create({
        name:hotelData.name,
        address:hotelData.address,
        location:hotelData.location,
        rating:hotelData.rating,
        ratingCount:hotelData.ratingCount

    })
    logger.info(`Hotel created :${hotel.id}`)
    return hotel
}

export async function getHotelById(id:number){
    const hotel = await Hotel.findByPk(id)

    if(!hotel){
        logger.error(`hotel not found ${id}`)
        throw new NotFoundError(`hotel with ${id} is not found`)
    }
    logger.info(`hotel with ${id} is  found`)
    return hotel
}

export async function getAllHotels(){
    const hotel = await Hotel.findAll({
        where:{
            deletedAt:null
        }
    })
    if(!hotel){
        logger.error(`hotel not found `)
        throw new NotFoundError(`No hotels found`)
        
    }
    logger.info(`Hotel found ${Hotel.length}`)
    return hotel
}

export async function softDeleteHotel(id:number){
    const hotel = await Hotel.findByPk(id)

    if(!hotel){
        logger.error(`hotel not found ${id}`)
        throw new NotFoundError(`hotel with ${id} is not found`)
    }
    hotel.deletedAt = new Date()
    await hotel.save()
    logger.info(`Hotel soft deleted: ${hotel.id}`)
    return true
}