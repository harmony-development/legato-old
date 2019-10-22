import { Schema, model } from 'mongoose';
import { IMessage } from '../types';

export const messageSchema: Schema = new Schema({
  author: {
    unique: false,
    required: true,
    type: String
  },
  message: {
    unique: false,
    required: true,
    type: String
  },
  files: {
    unique: false,
    required: true,
    type: Array<String>()
  }
});

export const Message = model<IMessage>('Message', messageSchema);
