import { NextFunction, Request, Response } from "express";

import { createHotelService, deleteHotelService, getAllHotelsService, getHotelByIdService} from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";

export const createHotelHandler = async (req: Request, res: Response, next: NextFunction) => {
    // call the service layer
    const hotelResponse = await createHotelService(req.body)

    //send the response
    res.status(StatusCodes.CREATED).json({
        message:"Hotel created successfully",
        data:hotelResponse,
        success:true
    })
}

export const getHotelByIdHandler =async (req: Request, res: Response, next: NextFunction)=>{
    const hotelResponse = await getHotelByIdService(Number(req.params.id))

    res.status(StatusCodes.OK).json({
        message:"Hotel found successfully",
        data:hotelResponse,
        success:true
    })

}

export const getHotelsHandler = async(req: Request, res: Response, next: NextFunction)=>{
    const hotelsResponse = await getAllHotelsService()
     res.status(StatusCodes.OK).json({
        message:"Hotel found successfully",
        data:hotelsResponse,
        success:true
     })
}


export async function deleteHotelHandler(req: Request, res: Response, next: NextFunction){
    const hotelResponse = await deleteHotelService(Number(req.params.id))

    res.status(StatusCodes.OK).json({
        message:"Hotel deleted successfully",
        data:hotelResponse,
        success:true
    })
}