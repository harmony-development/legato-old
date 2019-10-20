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
  LOGIN_ERROR: 'LOGIN_ERROR',
  REGISTER: 'REGISTER',
  REGISTER_ERROR: 'REGISTER_ERROR',
  DISCONNECT: 'DISCONNECT'
};

export interface ILoginDetails {
  email: string;
  password: string;
}

export interface IRegisterDetails {
  email: string;
  username: string;
  password: string;
}
