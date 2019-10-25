"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
var messageSchema_1 = require("../schema/messageSchema");
function onLoadMessages(socket) {
    socket.on(types_1.Events.LOAD_MESSAGES, function (data) {
        __1.harmonyServer.Database.verifyToken(data.token)
            .then(function () {
            if (data.lastmessageid) {
                messageSchema_1.Message.findOne({ messageid: data.lastmessageid })
                    .then(function (lastmessage) {
                    if (lastmessage) {
                        messageSchema_1.Message.find({
                            created_at: {
                                $lt: lastmessage.get('created_at', Date)
                            }
                        })
                            .limit(10)
                            .then(function (messages) {
                            if (messages.length > 0) {
                                __1.harmonyServer
                                    .getSocketServer()
                                    .emit(types_1.Events.LOAD_MESSAGES, messages);
                            }
                        })
                            .catch(function (err) {
                            console.log(err);
                            __1.harmonyServer
                                .getSocketServer()
                                .emit(types_1.Events.LOAD_MESSAGES_ERROR, 'Unable to get previous messages');
                        });
                    }
                    else {
                        __1.harmonyServer
                            .getSocketServer()
                            .emit(types_1.Events.LOAD_MESSAGES_ERROR, 'Message ID does not exist');
                    }
                })
                    .catch(function (err) {
                    console.log(err);
                });
            }
            else {
                messageSchema_1.Message.find({
                    created_at: {
                        $lt: new Date()
                    }
                })
                    .sort({ created_at: -1 })
                    .limit(10)
                    .then(function (messages) {
                    __1.harmonyServer
                        .getSocketServer()
                        .emit(types_1.Events.LOAD_MESSAGES, messages.reverse());
                })
                    .catch(function () {
                    __1.harmonyServer
                        .getSocketServer()
                        .emit(types_1.Events.LOAD_MESSAGES_ERROR, 'Unable to get previous messages');
                });
            }
        })
            .catch(function () {
            __1.harmonyServer
                .getSocketServer()
                .emit(types_1.Events.INVALIDATE_SESSION, 'Invalid token');
        });
    });
}
exports.default = onLoadMessages;
