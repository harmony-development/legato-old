import { Schema, model } from 'mongoose';
import { IMessage } from '../types';
import randomstring from 'crypto-random-string';

export const messageSchema: Schema = new Schema(
  {
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
    },
    messageid: {
      unique: true,
      required: true,
      type: String
    }
  },
  { timestamps: { createdAt: 'created_at' } }
);

messageSchema.pre<IMessage>('validate', function(next) {
  //if (!this.isModified('messageid')) return next();
  this.messageid = randomstring({ length: 30 });
  next();
});

export const Message = model<IMessage>('Message', messageSchema);
