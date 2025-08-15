import { z } from "zod"

export const createBookingSchema = z.object({
    userId: z.number({message: " user id must be present"}),
    hotelId: z.number({message: " hotel id must be present"}),
    totalGuests:z.number({message: " total guests must be present"}).min(1,{message: " total guests must be greter than 1"}),
    bookingAmount:z.number({message: " booking ammount must be present"}).min(1,{message: " booking ammount must be greter than 1"}),
    checkInDate:z.string({message: "Check-in date must be present"}),
    checkOutDate:z.string({message: "Check-Out date must be present"}),
    roomsCategoryId:z.number({message: "Rooms-category-id must be present"}),
})