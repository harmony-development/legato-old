import { Socket } from 'socket.io';
import { Events, ILoadMessagesData } from '../types';
import { harmonyServer } from '..';
import { Message } from '../schema/messageSchema';

export default function onLoadMessages(socket: Socket) {
  socket.on(Events.LOAD_MESSAGES, (data: ILoadMessagesData) => {
    harmonyServer.Database.verifyToken(data.token)
      .then(() => {
        if (data.lastmessageid) {
          Message.findOne({ messageid: data.lastmessageid })
            .then(lastmessage => {
              if (lastmessage) {
                Message.find({
                  created_at: {
                    $lt: lastmessage.get('created_at', Date)
                  }
                })
                  .limit(10)
                  .then(messages => {
                    if (messages.length > 0) {
                      harmonyServer
                        .getSocketServer()
                        .emit(Events.LOAD_MESSAGES, messages);
                    }
                  })
                  .catch(err => {
                    console.log(err);
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
                  .emit(
                    Events.LOAD_MESSAGES_ERROR,
                    'Message ID does not exist'
                  );
              }
            })
            .catch(err => {
              console.log(err);
            });
        } else {
          Message.find({
            created_at: {
              $lt: new Date()
            }
          })
            .sort({ created_at: -1 })
            .limit(10)
            .then(messages => {
              harmonyServer
                .getSocketServer()
                .emit(Events.LOAD_MESSAGES, messages.reverse());
            })
            .catch(() => {
              harmonyServer
                .getSocketServer()
                .emit(
                  Events.LOAD_MESSAGES_ERROR,
                  'Unable to get previous messages'
                );
            });
        }
      })
      .catch(() => {
        harmonyServer
          .getSocketServer()
          .emit(Events.INVALIDATE_SESSION, 'Invalid token');
      });
  });
}
