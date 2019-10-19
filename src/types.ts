export interface IMessage {
  author: string;
  message: string;
  icon?: string;
  files?: string[];
}

export interface IUserData {
  [key: string]: {
    name: string;
    icon?: string;
  };
}

export interface IProfileUpdate {
  name: string;
  icon?: string;
}

export interface ILoginData {
  name: string;
}

export const Events = {
  PROFILE_UPDATE: 'PROFILE_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT'
};

export type EventData = IUserData | IProfileUpdate | ILoginData | IMessage;
