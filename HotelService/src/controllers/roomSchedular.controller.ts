import { Request,Response } from "express";
import { getSchedulerStatus, manualExtendAvailability, startScheduler, stopScheduler } from "../scheduler/roomScheduler";
import { StatusCodes } from "http-status-codes";
import logger from "../config/logger.config";



export async function startSchedulerHandler(req:Request,res:Response){

    try {
         startScheduler()
         res.status(StatusCodes.OK).json({
            message: "room availibility schedular started successfully",
            success: true,
            data: {
                status: "started"
            }
         })
    } catch (error) {
        logger.error(`eroor in startng scheduler:`,error)
        res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({
            message: " failed to start room availibility schedular ",
            success: false,
            error: error instanceof Error ? error.message : "unknown error"
         })
    }
}

export async function stopSchedulerHandler(req:Request,res:Response){

    try {
         stopScheduler()
         res.status(StatusCodes.OK).json({
            message: "room availibility schedular stopped successfully",
            success: true,
            data: {
                status: "stopped"
            }
         })
    } catch (error) {
        logger.error(`eroor in stopping scheduler:`,error)
        res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({
            message: " failed to stop room availibility schedular ",
            success: false,
            error: error instanceof Error ? error.message : "unknown error"
         })
    }
}

export async function getSchedulerHandler(req:Request,res:Response){

    try {
         const status = getSchedulerStatus()
         res.status(StatusCodes.OK).json({
            message: " schedular status recived successfully",
            success: true,
            data:  status
            
         })
    } catch (error) {
        logger.error(`eroor in getting scheduler status:`,error)
        res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({
            message: " failed to get schedular status ",
            success: false,
            error: error instanceof Error ? error.message : "unknown error"
         })
    }
}


export async function manualExtendAvailabilityHandler(req:Request,res:Response){

    try {
         await manualExtendAvailability()
         res.status(StatusCodes.OK).json({
            message: " manual room availibility schedular completed successfully",
            success: true,
            data: {
                action: "manual_extention_completed"
            }
         })
    } catch (error) {
        logger.error(`Error in manual room availability extension:`,error)
        res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({
            message: " failed to perform manual room availability extension ",
            success: false,
            error: error instanceof Error ? error.message : "unknown error"
         })
    }
}