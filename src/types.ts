export const Events = {
  USERNAME_UPDATE: 'USERNAME_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT',
  IMAGE: 'IMAGE'
};

export interface IMessage {
  author: string;
  message: string;
  files: File[];
}
