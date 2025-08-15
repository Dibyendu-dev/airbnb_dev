import { z } from "zod";

export const getAvailableRoomsSchema = z.object({
    roomsCategoryId: z.number({message: "Rooms category id must be present"}),
    checkInDate: z.string({message: "Check-in date must be present"}),
    checkOutDate: z.string({message: "Check-Out date must be present"}),

})