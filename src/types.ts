export interface IMessage {
  username: string;
  message: string;
  avatar: string;
  files: string[];
}

export interface IProfileUpdate {
  username?: string;
  avatar?: string;
  token: string;
}
