import { Request, Response } from "express";

import { StatusCodes } from "http-status-codes";
import { addRoomGenerationJobQueue } from "../producers/roomGeneration.producer";

export async function generateRoomHandler(req:Request,res:Response){

    await addRoomGenerationJobQueue(req.body)

    res.status(StatusCodes.OK).json({
        message: "Rooms generation job added to queue",
        success:true,
        data:{},
    })
}