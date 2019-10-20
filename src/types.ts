import mongoose from 'mongoose';

export interface IUser extends mongoose.Document {
  userid?: string;
  email?: string;
  password?: string;
  username?: string;
}

export const Events = {
  PROFILE_UPDATE: 'PROFILE_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT'
};
