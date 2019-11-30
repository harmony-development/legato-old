import { Color, Theme } from '@material-ui/core';

export interface IMessage {
    author: string;
    message: string;
    messageid: string;
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

export interface IUserData {
    [userid: string]: {
        avatar: string;
        username: string;
    };
}

export interface IGetTargetUserData {
    username: string;
    avatar: string;
    userid: string;
}
