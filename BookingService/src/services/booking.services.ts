import { CreateBookingDTO } from "../dto/booking.dto";
import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotencyKey,
  getIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import {
  BadRequestError,
  InternalServerError,
  NotFoundError,
} from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";

import prismaClient from "../prisma/client";
import { redlock } from "../config/redis.config";
import { serverConfig } from "../config";
import { getAvailableRooms, updateBookingIdToRooms } from "../api/hotel.api";

type AvailableRoom = {
  id: number;
  roomCategoryId: number;
  dateOfAvailability: Date;
};

export async function createBookingService(createBookingDTO: CreateBookingDTO) {
  const ttl = serverConfig.LOCK_TTL;
  const bookingResource = `hotel:${createBookingDTO.hotelId}`; //todo: modify the lock to use available rooms

  // check the available rooms from hotel service

  const availableRooms = await getAvailableRooms(
    createBookingDTO.roomsCategoryId,
    createBookingDTO.checkInDate,
    createBookingDTO.checkOutDate
  );

  const checkInDate = new Date(createBookingDTO.checkInDate);
  const checkOutDate = new Date(createBookingDTO.checkOutDate);

  const totalNights = Math.ceil(
    (checkOutDate.getTime() - checkInDate.getTime()) / (1000 * 60 * 60 * 24)
  );

  if (availableRooms.length === 0 || availableRooms.length < totalNights) {
    throw new BadRequestError("no rooms available for given dates");
  }

  try {
    await redlock.acquire([bookingResource], ttl); 
    const booking = await createBooking({
      userId: createBookingDTO.userId,
      hotelId: createBookingDTO.hotelId,
      totalGuests: createBookingDTO.totalGuests,
      bookingAmount: createBookingDTO.bookingAmount,
      checkInDate: new Date(createBookingDTO.checkInDate),
      checkOutDate: new Date(createBookingDTO.checkOutDate),
      roomsCategoryId: createBookingDTO.roomsCategoryId,
    });

    const idempotencyKey = generateIdempotencyKey();

    await createIdempotencyKey(idempotencyKey, booking.id);

    //  update rooms
    await updateBookingIdToRooms(
      booking.id,
      availableRooms.data.map((room: AvailableRoom) => room.id)
    );

    return {
      bookingId: booking.id,
      idempotencyKey: idempotencyKey,
    };
  } catch (error) {
    throw new InternalServerError(
      "faild to accuire lock the booking resourses "
    );
  }
}

// Todo: explore the function for potential issues and improvements
export async function confirmBookingService(idempotencyKey: string) {
  return await prismaClient.$transaction(async (tx) => {
    const idempotencyKeyData = await getIdempotencyKeyWithLock(
      tx,
      idempotencyKey
    );

    if (!idempotencyKeyData || !idempotencyKeyData.bookingId) {
      throw new NotFoundError("Idempotency key not found");
    }

    if (idempotencyKeyData.finalized) {
      throw new BadRequestError("Idempotency key already finalized");
    }

    const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
    await finalizeIdempotencyKey(tx, idempotencyKey);  // todo: mark the room as null, if booking fails

    return booking;
  });
}
