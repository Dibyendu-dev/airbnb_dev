import { Job,Worker } from "bullmq";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue";
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer";
import { generateRooms } from "../services/roomGeneration.service";
import logger from "../config/logger.config";
import { getRedisConnObject } from "../config/redis.config";


export const setupRoomGenerationWorker = ()=>{

    const roomGenerationConsumer = new Worker<RoomGenerationJob>(
        ROOM_GENERATION_QUEUE,
        async (job: Job) =>{

            if(job.name !== ROOM_GENERATION_PAYLOAD){
                throw new Error(`Invalid job name`)
            }

            const payload = job.data;
            console.log(`processing room generation for: ${JSON.stringify(payload)}`);

            await generateRooms(payload)

            logger.info(`Room generation completed for  ${JSON.stringify(payload)}`)
        },
        {
            connection: getRedisConnObject()
        }
    )
    roomGenerationConsumer.on("failed",()=> {
        console.error("Room generation processing failed");
    });

    roomGenerationConsumer.on("completed",()=>{
        console.log(`room generation completed successfully`)
    })

}