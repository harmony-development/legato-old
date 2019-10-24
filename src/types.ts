import mongoose from 'mongoose';

export interface ITheme {
  primary: object;
  secondary: object;
  type: 'dark' | 'light';
}

export interface IUser extends mongoose.Document {
  email?: string;
  password?: string;
  username?: string;
  userid?: string;
  avatar?: string;
  theme: ITheme;
}

export interface IMessage extends mongoose.Document {
  author: string;
  message: string;
  files: string[];
  messageid: string;
}

export const Events = {
  PROFILE_UPDATE: 'PROFILE_UPDATE',
  PROFILE_UPDATE_ERROR: 'PROFILE_UPDATE_ERROR',
  GET_USER_DATA: 'GET_USER_DATA',
  GET_USER_DATA_ERROR: 'GET_USER_DATA_ERROR',
  LOAD_MESSAGES: 'LOAD_MESSAGES',
  LOAD_MESSAGES_ERROR: 'LOAD_MESSAGES_ERROR',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  LOGIN_ERROR: 'LOGIN_ERROR',
  REGISTER: 'REGISTER',
  REGISTER_ERROR: 'REGISTER_ERROR',
  INVALIDATE_SESSION: 'INVALIDATE_SESSION'
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

export interface IMessageData {
  token: string;
  message: string;
  files: [];
}

export interface IToken {
  userid: string;
}

export interface IUserUpdateData {
  username?: string;
  avatar?: string;
  theme: ITheme;
  token: string;
}

export interface IGetUserData {
  token: string;
}

export interface ILoadMessagesData {
  token: string;
  lastmessageid: string;
}
