export const Events = {
  PROFILE_UPDATE: 'PROFILE_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT'
};

export interface IMessage {
  author: string;
  message: string;
  avatar: string;
  files: string[];
}
