export interface IMessage {
  message: string;
}

export interface IUserData {
  [key: string]: {
    name: string;
  };
}

export interface IImageData {
  file: File;
}

export interface IUsernameUpdate {
  name: string;
}

export interface IConnectData {
  name: string;
}

export const Events = {
  USERNAME_UPDATE: 'USERNAME_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT',
  IMAGE: 'IMAGE'
};
