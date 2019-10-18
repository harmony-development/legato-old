export interface IMessage {
  author: string;
  message: string;
  files?: File[];
}

export interface IUserData {
  [key: string]: {
    name: string;
  };
}

export interface IUsernameUpdate {
  name: string;
}

export interface ILoginData {
  name: string;
}

export const Events = {
  USERNAME_UPDATE: 'USERNAME_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT'
};

export type EventData = IUserData | IUsernameUpdate | ILoginData | IMessage;
