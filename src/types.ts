import { Color, Theme } from '@material-ui/core';

export interface IMessage {
  username: string;
  message: string;
  avatar: string;
  files: string[];
}

export interface IProfileUpdate {
  username?: string;
  avatar?: string;
  theme?: {
    primary: Color;
    secondary: Color;
    type: 'dark' | 'light';
  };
  token: string;
}

export interface IGetUserData {
  username: string;
  avatar?: string;
  theme?: {
    primary: Color;
    secondary: Color;
    type: 'dark' | 'light';
  };
}
