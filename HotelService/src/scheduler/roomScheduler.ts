import * as corn from 'node-cron';
import { RoomRepository } from '../repositories/room.repository';
import { RoomCategoryRepository } from '../repositories/roomCategory.repository';
import { addRoomGenerationJobQueue } from '../producers/roomGeneration.producer';
import { RoomGenerationJob } from '../dto/roomGeneration.dto';
import logger from '../config/logger.config';
import { serverConfig } from '../config';


const roomRepository = new RoomRepository();
const roomCategoryRepository = new RoomCategoryRepository();

let cronJob: corn.ScheduledTask | null = null;

export const startScheduler = (): void => {

    if(cronJob){
        logger.warn('Room schedular is already running');
        return
    }

    // scheduled cronjob to run
    cronJob =  corn.schedule(serverConfig.ROOM_CORN, async ()=> {
        try {
            logger.info(" starting room availibility check");
             await extendRoomAvailability();
             logger.info("room availibility check completed")
        } catch (error) {
            logger.error('error in room availibility scheduler');
        }
    },{
        timezone: 'UTC'
    });

    cronJob.start();
    logger.info(`room availibility scheduler started- running every ${serverConfig.ROOM_CORN}`);
}

export const stopScheduler = ():void =>{
    
    if(cronJob){
        cronJob.stop();
        cronJob = null;
        logger.info("room availibility scheduler stopped.")
    }
}

export const getSchedulerStatus = (): {isRunning : boolean} => {
    return{
        isRunning: cronJob!== null && cronJob.getStatus() === 'scheduled'
    }
}

const extendRoomAvailability = async (): Promise<void>=> {

    try {

        // get all room categories with latest availability dates
        const roomCategoriesWithLatestDates = await roomRepository.findLatestDatesForAllCategories();

        if(roomCategoriesWithLatestDates.length === 0){
            logger.info(`no room categories found with availabilty dates`)
            return;
        }

        logger.info(`found ${roomCategoriesWithLatestDates.length} room categories to extend`)

        // process each category
        for(const categoryData of roomCategoriesWithLatestDates){
            await extendCategoryAvailability(categoryData)
        }

    } catch (error) {
        logger.error(`error extending room availability`,error)
        throw error
    }
}

const extendCategoryAvailability = async (categoryData: {roomsCategoryId: number , latestDate: Date}) : Promise<void> => {
    try {

        const {roomsCategoryId,latestDate} = categoryData;

        // calculate the next date -> one day after latest date
        const nextDate = new Date(latestDate);
        nextDate.setDate(nextDate.getDate()+1);
        
        // check the room category still exsists
        const roomCategory =await roomCategoryRepository.findById(roomsCategoryId)
        if(!roomCategory){
            logger.warn(`room category ${roomsCategoryId} not found`)
            return
        }

        // check if room for next day already exsist
        const exsistingRoom = await roomRepository.findByRoomCategoryIdAndDate(roomsCategoryId,nextDate);
        if(exsistingRoom){
            logger.debug(`room for category ${roomsCategoryId} on ${nextDate.toISOString()} already exsist, skipping`);
            return;
        }

        const endDate = new Date(nextDate);
        endDate.setDate(endDate.getDate() +1);

        // create job for generate room for next date
        const jobData: RoomGenerationJob = {
            roomsCategoryId,
            startDate: nextDate.toISOString(),
            endDate: endDate.toISOString(),
            priceOverride: roomCategory.price,
            bachSize:1
        }

        // add job to queue
        await addRoomGenerationJobQueue(jobData);

        logger.info(` Addrd room generation job for category ${roomsCategoryId} on ${nextDate.toISOString()}`);
    } catch (error) {
            logger.error(`erron in extending room category ${categoryData.roomsCategoryId} availability `)
    }
} 

export const manualExtendAvailability = async ():Promise<void> => {
    logger.info(`manual room availibility triggered`)
    await extendRoomAvailability()
}