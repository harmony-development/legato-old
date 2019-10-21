import mongoose from 'mongoose';

interface IColor {
  light: string;
  dark: string;
  contrastText: string;
}

export interface ITheme {
  primary: IColor;
  secondary: IColor;
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

export const Events = {
  PROFILE_UPDATE: 'PROFILE_UPDATE',
  PROFILE_UPDATE_ERROR: 'PROFILE_UPDATE_ERROR',
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
  token: string;
}
