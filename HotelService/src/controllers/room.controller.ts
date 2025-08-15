import { NextFunction, Request, Response } from "express";
import { getAvailableRoomService, updateBookingIdToRoomsService } from "../services/room.service";
import { StatusCodes } from "http-status-codes";

export async function getAvailableRoomsHandler(req:Request,res:Response, next: NextFunction) {
   
    const rooms = await getAvailableRoomService(req.body);

    res.status(StatusCodes.OK).json({
        message: "Room found successfully",
        data: rooms,
        success: true,
    })
}


export async function updateBookingIdToRoomsHandler(req:Request,res:Response, next: NextFunction){

    const response = await updateBookingIdToRoomsService(req.body);

      res.status(StatusCodes.OK).json({
        message: "Booking id updated to rooms successfully",
        data: response,
        success: true,
    })

}