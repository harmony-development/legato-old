import { Socket } from 'socket.io';
import { Events, ILoadMessagesData } from '../types';
import { harmonyServer } from '..';
import { Message } from '../schema/messageSchema';

export default function onLoadMessages(socket: Socket) {
  socket.on(Events.LOAD_MESSAGES, (data: ILoadMessagesData) => {
    harmonyServer.Database.verifyToken(data.token)
      .then(() => {
        Message.findOne({ messageid: data.lastmessageid })
          .then(lastmessage => {
            if (lastmessage) {
              Message.find({
                created_at: {
                  $lt: new Date().setDate(lastmessage.get('created_at', Date))
                }
              })
                .limit(10)
                .then(messages => {
                  harmonyServer
                    .getSocketServer()
                    .emit(Events.LOAD_MESSAGES, messages);
                })
                .catch(() => {
                  harmonyServer
                    .getSocketServer()
                    .emit(
                      Events.LOAD_MESSAGES_ERROR,
                      'Unable to get previous messages'
                    );
                });
            } else {
              harmonyServer
                .getSocketServer()
                .emit(Events.LOAD_MESSAGES_ERROR, 'Message ID does not exist');
            }
          })
          .catch(err => {
            console.log(err);
          });
      })
      .catch(() => {
        harmonyServer
          .getSocketServer()
          .emit(Events.INVALIDATE_SESSION, 'Invalid token');
      });
  });
}
