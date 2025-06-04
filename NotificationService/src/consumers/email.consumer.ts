import { Job, Worker } from "bullmq";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producers";
import { NotificationDto } from "../dto/notification.dto";


export const setupMailerWorker = ()=>{
    const emailConsumer = new Worker<NotificationDto>(
    MAILER_QUEUE,
    async (job:Job)=>{
        if(job.name !== MAILER_PAYLOAD){
            throw new Error("Invalid job name ")
        }
        const payload =job.data
        console.log(`processing email for: ${JSON.stringify(payload)}`)
    },
    {
        connection:getRedisConnObject()
    }
)

emailConsumer.on("failed",()=>{
    console.error("email processing failed")
});
emailConsumer.on("completed",()=>{
    console.log("email processing completed successfully");
})
}

