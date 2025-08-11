import { CreationAttributes } from "sequelize";
import logger from "../config/logger.config";
import RoomCategory from "../db/models/roomCategory.model";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { RoomCategoryRepository } from "../repositories/roomCategory.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import Room from "../db/models/room.model";
import { RoomRepository } from "../repositories/room.repository";

const roomCategoryRepository = new RoomCategoryRepository();
const roomRepository = new RoomRepository();

export async function generateRooms(jobData: RoomGenerationJob) {
  let totalRoomsCreated = 0;
  let totalDatesProcessed = 0;

  const roomCategory = await roomCategoryRepository.findById(
    jobData.roomCategoryId
  );

  if (!roomCategory) {
    logger.error(
      `Room Category with id : ${jobData.roomCategoryId} is not found`
    );
    throw new NotFoundError(
      `Room Category with id : ${jobData.roomCategoryId} is not found`
    );
  }

  const startDate = new Date(jobData.startDate);
  const endDate = new Date(jobData.endDate);

  if (startDate >= endDate) {
    logger.error(`start date must be before end date`);
    throw new BadRequestError(`start date must be before end date`);
  }

  if (startDate < new Date()) {
    logger.error(`start date must be in the future`);
    throw new BadRequestError(`start date must be in the future`);
  }

  const totalDays = Math.ceil(
    (endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24)
  );

  logger.info(`Generating rooms for ${totalDays} days`)

  const batchSize = jobData.bachSize || 100;

  const currentDate = new Date(startDate)

  while(currentDate < endDate){
    const batchEndDate = new Date(currentDate);
    batchEndDate.setDate(batchEndDate.getDate() + batchSize)

    if(batchEndDate > endDate){
        batchEndDate.setTime(endDate.getTime())
    }

    const batchResult = await processDateBatch(roomCategory,currentDate,batchEndDate,jobData.priceOverride)

    totalRoomsCreated +=batchResult.roomsCreated
    totalDatesProcessed +=batchResult.datesProcessed

    currentDate.setTime(batchEndDate.getTime())
  }

  return{
    totalRoomsCreated,
    totalDatesProcessed,
  }
}

export async function processDateBatch(
  roomCategory: RoomCategory,
  startDate: Date,
  endDate: Date,
  priceOverride?: number
) {
  let roomsCreated = 0;
  let datesProcessed = 0;

  const roomsToCreate: CreationAttributes<Room>[] = [];

  const currentDate = new Date(startDate);

  // select * from room_category where id = ? and date_of_availiability between ? and  ?
  // TODO: better query to get the rooms
  while (currentDate <= endDate) {
    const exsistingRoom = await roomRepository.findByRoomCategoryIdAndDate(
      roomCategory.id,
      currentDate
    );

    logger.info(
      `Exsisting room: ${JSON.stringify(exsistingRoom)} :${currentDate}`
    );

    if (!exsistingRoom) {
      const roomPayload = {
        hotelId: roomCategory.hotelId,
        roomCategoryId: roomCategory.id,
        dateOfAvailability: new Date(currentDate),
        price: priceOverride || roomCategory.price,
        createdAt: new Date(),
        updatedAt: new Date(),
        deletedAt: null,
      };
      console.log(`Rooms payload: ${JSON.stringify(roomPayload)}`);
      roomsToCreate.push(roomPayload);
    }

    currentDate.setDate(currentDate.getDate() + 1);
    datesProcessed++;
  }

  console.log(`Rooms to create: ${JSON.stringify(roomsToCreate)}`);

  if (roomsToCreate.length > 0) {
    logger.info(`creating rooms ${roomsToCreate.length} rooms`);

    await roomRepository.bulkCreate(roomsToCreate);
    roomsCreated += roomsToCreate.length;
  }

  return {
    roomsCreated,
    datesProcessed,
  };
}
