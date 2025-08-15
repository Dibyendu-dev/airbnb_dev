export type GetAvailableRoomsDTO = {
    roomsCategoryId : number;
    checkInDate : string;
    checkOutDate: string;
}

export type UpdateBookingIdToRoomsDTO = {
    bookingId: number;
    roomIds: number[];
}